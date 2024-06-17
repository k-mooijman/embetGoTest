package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

func Log() {
	// Clones the given repository, creating the remote, the local branches
	// and fetching the objects, everything in memory:
	Info("git clone https://github.com/k-mooijman/embetGoTest.git")
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/k-mooijman/embetGoTest.git",
	})
	CheckIfError(err)

	// Gets the HEAD history from HEAD, just like this command:
	Info("git log")

	// ... retrieves the branch pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)

	// ... retrieves the commit history
	since := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	until := time.Date(2024, 7, 30, 0, 0, 0, 0, time.UTC)
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &since, Until: &until})
	CheckIfError(err)

	// ... just iterates over the commits, printing it
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)

		return nil
	})
	CheckIfError(err)

}

func Stat() {
	repoPath := "/home/kasper/development/kasper/projects/cleanProjectGo" // Replace with your repository path

	r, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatalf("Failed to open repository: %v", err)
	}

	w, err := r.Worktree()
	if err != nil {
		log.Fatalf("Failed to get worktree: %v", err)
	}

	status, err := w.Status()
	if err != nil {
		log.Fatalf("Failed to get status: %v", err)
	}

	fmt.Printf("Is the repository clean? %v\n", status.IsClean())
	fmt.Println(status.String())

	fmt.Printf("####################################### ")

	StatGI()

}

func StatGI() {

	// Paths to the local.gitignore and global.gitignore
	localIgnorePath := ".gitignore"
	globalIgnorePath := os.Getenv("HOME") + "/.gitignore_global"

	// Read the.gitignore files
	localIgnorePatterns, _ := readGitIgnore(localIgnorePath)
	globalIgnorePatterns, _ := readGitIgnore(globalIgnorePath)

	// Example path to check
	pathToCheck := "/home/kasper/development/kasper/projects/cleanProjectGo"

	// Check if the path should be ignored
	isIgnored := shouldIgnore(pathToCheck, append(localIgnorePatterns, globalIgnorePatterns...))

	fmt.Printf("\n \nShould ignore '%s'? %v\n", pathToCheck, isIgnored)
}

// Function to read and return the content of a.gitignore file
func readGitIgnore(filePath string) ([]string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}

// Function to check if a path should be ignored based on the ignore patterns
func shouldIgnore(path string, patterns []string) bool {
	for _, pattern := range patterns {
		if strings.Contains(path, pattern) {
			return true
		}
	}
	return false
}
