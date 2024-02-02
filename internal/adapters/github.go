package adapters

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/logging"
)

type (
	// GitHubAdapter is a struct to hold the GitHub client
	GitHubAdapter struct {
		client *github.Client
		log    *logging.Logger
	}

	BinaryNotFoundError struct {
		BinaryName string
	}
)

const (
	kiraGit    = "KiraCore"
	sekaiRepo  = "sekai"
	interxRepo = "interx"
)

func (g *GitHubAdapter) MustDownloadBinaries(ctx context.Context, cfg *config.KiraConfig) {
	err := g.downloadBinaryFromRepo(ctx, kiraGit, sekaiRepo, cfg.SekaiDebFileName, cfg.SekaiVersion)
	if err != nil {
		g.log.Fatalf("Cannot download '%s/%s' from '%s', error: %s", cfg.SekaiDebFileName, cfg.SekaiVersion, sekaiRepo, err)
	}

	err = g.downloadBinaryFromRepo(ctx, kiraGit, interxRepo, cfg.InterxDebFileName, cfg.InterxVersion)
	if err != nil {
		g.log.Fatalf("Cannot download '%s/%s' from '%s', error: %s", cfg.InterxDebFileName, cfg.InterxVersion, interxRepo, err)
	}
}

// NewGitHubAdapter initializes a new GitHubAdapter instance
func NewGitHubAdapter(ctx context.Context, accessToken string, log *logging.Logger) *GitHubAdapter {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)

	return &GitHubAdapter{
		client: github.NewClient(tc),
		log:    log,
	}
}

// getLatestRelease fetches the latest release from the specified repository
func (g *GitHubAdapter) getLatestRelease(ctx context.Context, owner, repo string) (*github.RepositoryRelease, error) {
	release, _, err := g.client.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		return nil, err
	}
	return release, nil
}

// downloadBinaryFromRepo downloads a binary file from a GitHub repository.
// - ctx: The context for the operation.
// - owner: The owner of the GitHub repository.
// - repo: The name of the GitHub repository.
// - binaryName: The name of the binary file to download.
// tag: The tag or version of the release containing the binary file.
func (g *GitHubAdapter) downloadBinaryFromRepo(ctx context.Context, owner, repo, binaryName, tag string) error {
	g.log.Infof("Downloading '%s' from '%s/%s', tag: %s", binaryName, owner, repo, tag)

	var (
		release *github.RepositoryRelease
		err     error
	)
	switch tag {
	case "latest":
		release, _, err = g.client.Repositories.GetLatestRelease(ctx, owner, repo)
		if err != nil {
			return fmt.Errorf("fetching latest release error: %w", err)
		}
	default:
		release, _, err = g.client.Repositories.GetReleaseByTag(ctx, owner, repo, tag)
		if err != nil {
			return fmt.Errorf("fetching '%s' tag release error: %w", tag, err)
		}
	}

	var asset *github.ReleaseAsset
	for _, a := range release.Assets {
		if *a.Name == binaryName {
			assetCopy := a     // Create a copy of 'a'
			asset = &assetCopy // Assign the address of the copy
			break
		}
	}

	if asset == nil {
		return &BinaryNotFoundError{BinaryName: binaryName}
	}

	// Create a request with context
	req, err := http.NewRequestWithContext(ctx, "GET", *asset.BrowserDownloadURL, nil)
	if err != nil {
		return fmt.Errorf("creating request error: %w", err)
	}

	// Create an HTTP client and do the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("downloading binary error: %w", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(binaryName)
	if err != nil {
		return fmt.Errorf("creating binary file error: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("writing binary to file error: %w", err)
	}

	g.log.Infof("Binary file downloaded successfully")
	return nil
}

func (e *BinaryNotFoundError) Error() string {
	return fmt.Sprintf("binary not found in the latest release: %s", e.BinaryName)
}
