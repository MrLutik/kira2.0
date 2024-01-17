package docker

import "errors"

var (
	ErrPackageInstallationFailed = errors.New("package installation failed")
	ErrFileNotFoundInTarBase     = errors.New("file not found in tar archive")
	ErrStderrNotEmpty            = errors.New("stderr is not empty")
)
