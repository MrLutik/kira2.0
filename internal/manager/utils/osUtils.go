package utils

import (
	"io"
	"os"
	"os/user"
)

func (h *HelperManager) GetCurrentOSUser() *user.User {
	//geting curent user home folder even if it runned by sudo
	sudoUser := os.Getenv("SUDO_USER")

	if sudoUser != "" {
		usr, err := user.Lookup(sudoUser)
		if err != nil {
			panic(err)
		}
		return usr
	} else {
		// Fallback to the current user if not running via sudo
		usr, err := user.Current()
		if err != nil {
			panic(err)
		}
		return usr
	}
}

func (h *HelperManager) copyFile(src, dst string) error {
	// Open source file for reading
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create destination file for writing
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy the contents from srcFile to dstFile
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
