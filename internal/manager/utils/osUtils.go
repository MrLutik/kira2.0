package utils

import (
	"io"
	"os"
)

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
