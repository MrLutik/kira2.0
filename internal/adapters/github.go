package adapters

import (
	"context"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/logging"
)

var log = logging.Log

const (
	envGithubTokenVariableName = "GITHUB_TOKEN"
	kiraGit                    = "KiraCore"
	sekaiRepo                  = "sekai"
	interxRepo                 = "interx"
)

func DownloadBinaries(ctx context.Context, cfg *config.KiraConfig, sekaiDebFileName, interxDebFileName string) {
	repositories := repositories{}

	repositories.Set(kiraGit, sekaiRepo, cfg.SekaiVersion)
	repositories.Set(kiraGit, interxRepo, cfg.InterxVersion)
	log.Infof("Getting repositories: %+v", repositories.Get())

	token, exists := os.LookupEnv(envGithubTokenVariableName)
	if !exists {
		log.Fatalf("'%s' variable is not set", envGithubTokenVariableName)
	}

	repositories = fetch(repositories, token)

	gitHubAdapter := newGitHubAdapter(token)

	gitHubAdapter.downloadBinaryFromRepo(ctx, kiraGit, sekaiRepo, sekaiDebFileName, cfg.SekaiVersion)
	gitHubAdapter.downloadBinaryFromRepo(ctx, kiraGit, interxRepo, interxDebFileName, cfg.InterxVersion)
}

// gitHubAdapter is a struct to hold the GitHub client
type gitHubAdapter struct {
	client *github.Client
}
type repository struct {
	Owner   string
	Repo    string
	Version string
}

type repositories struct {
	repos []repository
}

// Add a new Repository to Repositories, version can be = ""
func (r *repositories) Set(owner, repo, version string) {
	newRepo := repository{Owner: owner, Repo: repo, Version: version}
	r.repos = append(r.repos, newRepo)
}

func (r *repositories) Get() []repository {
	return r.repos
}

func fetch(r repositories, accessToken string) repositories {
	adapter := newGitHubAdapter(accessToken)

	var wg sync.WaitGroup
	results := make(chan repository)

	for _, repo := range r.repos {
		wg.Add(1)
		go func(owner, repo string) {
			defer wg.Done()

			latestRelease, err := adapter.getLatestRelease(owner, repo)
			if err != nil {
				log.Errorf("Error fetching latest release for %s/%s: %s\n", owner, repo, err)
				return
			}

			results <- repository{Owner: owner, Repo: repo, Version: *latestRelease.TagName}
		}(repo.Owner, repo.Repo)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var updatedRepos []repository
	for result := range results {
		updatedRepos = append(updatedRepos, result)
	}

	return repositories{repos: updatedRepos}
}

// newGitHubAdapter initializes a new GitHubAdapter instance
func newGitHubAdapter(accessToken string) *gitHubAdapter {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return &gitHubAdapter{
		client: github.NewClient(tc),
	}
}

// getLatestRelease fetches the latest release from the specified repository
func (gh *gitHubAdapter) getLatestRelease(owner, repo string) (*github.RepositoryRelease, error) {
	release, _, err := gh.client.Repositories.GetLatestRelease(context.Background(), owner, repo)
	if err != nil {
		return nil, err
	}
	return release, nil
}

// downloadBinaryFromRepo downloads a binary file from a GitHub repository.
// ctx: The context for the operation.
// owner: The owner of the GitHub repository.
// repo: The name of the GitHub repository.
// binaryName: The name of the binary file to download.
// tag: The tag or version of the release containing the binary file.
func (gh *gitHubAdapter) downloadBinaryFromRepo(ctx context.Context, owner, repo, binaryName, tag string) {
	var release *github.RepositoryRelease
	var err error
	log.Printf("downloading %s from %s/%s, tag:%s\n", binaryName, owner, repo, tag)
	switch tag {
	case "latest":
		release, _, err = gh.client.Repositories.GetLatestRelease(ctx, owner, repo)
		if err != nil {
			log.Fatalf("Error fetching latest release: %v", err)
		}
	default:
		release, _, err = gh.client.Repositories.GetReleaseByTag(ctx, owner, repo, tag)
		if err != nil {
			log.Fatalf("Error fetching latest release: %v", err)
		}
	}

	var asset *github.ReleaseAsset
	for _, a := range release.Assets {
		if *a.Name == binaryName {
			asset = &a
			break
		}
	}

	if asset == nil {
		log.Fatalf("Binary not found in the latest release: %s", binaryName)
	}

	resp, err := http.Get(*asset.BrowserDownloadURL)
	if err != nil {
		log.Fatalf("Error downloading binary: %v", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(binaryName)
	if err != nil {
		log.Fatalf("Error creating binary file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("Error writing binary to file: %v", err)
	}

	log.Infof("Binary file downloaded successfully")
}
