package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const staticServingDir = "www"

func isFileInDirectory(filePath, directoryPath string) (bool, error) {
	relPath, err := filepath.Rel(directoryPath, filePath)
	if err != nil {
		return false, err
	}
	return !strings.HasPrefix(relPath, ".."+string(filepath.Separator)), nil
}

func resolveRoute(route string) response {
	if route == "/" {
		route = "/index.html"
	}

	joined := path.Join(staticServingDir, route)
	is, err := isFileInDirectory(joined, staticServingDir)
	if !is {
		log.Printf("blocked malicious file access to %s\n", joined)
		return response{status: http.StatusNotFound}
	}
	content, err := os.ReadFile(joined)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return response{status: http.StatusNotFound}
		}
		return response{status: http.StatusInternalServerError}
	}

	return response{status: http.StatusOK, content: string(content)}
}
