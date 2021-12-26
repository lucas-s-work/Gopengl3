package util

import (
	"io"
	"os"
	"path"
)

func RelativePath(relpath string) string {
	return path.Join(os.Getenv("root_file_path"), relpath)
}

func ReadFile(loc string) (string, error) {
	// Use relative path to ensure loaded relative to importing package
	f, err := os.Open(RelativePath(loc))
	if err != nil {
		return "", err
	}

	d, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(d), nil
}
