package docker

import (
	"context"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/cosign"
	"github.com/mrlutik/kira2.0/internal/errors"
)

const DockerImagePubKey = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE/IrzBQYeMwvKa44/DF/HB7XDpnE+
f+mU9F/Qbfq25bBWV2+NlYMJv3KvKHNtu3Jknt6yizZjUV4b8WGfKBzFYw==
-----END PUBLIC KEY-----`

func VerifyingDockerImage(ctx context.Context, dockerManager *DockerManager, cfg *config.KiraConfig) {
	err := dockerManager.VerifyDockerInstallation(ctx)
	errors.HandleErr("Docker is not available", err)

	dockerBaseImageName := cfg.DockerImageName + ":" + cfg.DockerImageVersion
	err = dockerManager.PullImage(ctx, dockerBaseImageName)
	errors.HandleErr("Pulling image", err)

	checkBool, err := cosign.VerifyImageSignature(ctx, dockerBaseImageName, DockerImagePubKey)
	errors.HandleErr("Verifying image signature", err)

	log.Infoln("Verified:", checkBool)
}
