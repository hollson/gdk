package gdk

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

const maxRecursive = 100

// FindProjectAbs find absolute path that have go.mod in parent folder.
func FindProjectAbs() (string, error) {
	p, err := filepath.Abs(".")
	if err != nil {
		return "", fmt.Errorf(
			"can't find go.mod in parent ancestor: cannot find absolute path of '.'",
		)
	}

	return findProjectAbs(p, 0)
}

func findProjectAbs(currentPath string, recursive int) (string, error) {
	// Prevent too much recursive that can create deadlock.
	if recursive > maxRecursive {
		return "", fmt.Errorf(
			"can't find go.mod in parent ancestor: '%s' nested more than %d level",
			currentPath,
			maxRecursive,
		)
	}

	// Check if current path is already project root.
	if currentPath == "/" {
		return "", fmt.Errorf("can't find go.mod in parent ancestor")
	}

	files, err := ioutil.ReadDir(currentPath)
	if err != nil {
		return "", fmt.Errorf(
			"can't find go.mod in parent ancestor: cannot read in '%s'",
			currentPath,
		)
	}

	for _, f := range files {
		if f.Name() == "go.mod" {
			return currentPath, nil
		}
	}

	newPath, err := filepath.Abs(currentPath + "/../")
	if err != nil {
		return "", fmt.Errorf(
			"can't find go.mod in parent ancestor: cannot find absolute path of %s",
			currentPath+"/../",
		)
	}

	if currentPath == newPath {
		return "", fmt.Errorf("can't find go.mod in parent ancestor: stuck in %s", newPath)
	}

	return findProjectAbs(newPath, recursive+1)
}
