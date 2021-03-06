package git

import (
	"log"
	"os/exec"
	"strings"

	"github.com/denis-tingajkin/go-header/messages"
)

//Git represents API for git
type Git struct {
	projectDir string
}

//New creates new Git instance
func New(projectDir string) *Git {
	return &Git{projectDir: projectDir}
}

//Author returns author of file
func (g *Git) Author(path string) string {
	out, err := g.do("log", `--pretty=format:%an - %ae`, "--diff-filter=A", "--", path, "--e")
	if err != nil {
		log.Printf("can't get author of file: %v. %v", path, messages.ErrorMsg(err))
		return ""
	}
	authors := strings.Split(out, "\n")
	return authors[0]
}

//DiffFiles returns list of changed files
func (g *Git) DiffFiles(branch string) []string {
	out, err := g.do("diff", "--name-only", branch)
	if err != nil {
		log.Printf("can't get diff: %v", messages.ErrorMsg(err))
		return nil
	}
	return strings.Split(out, "\n")
}

//OnlyNewFiles returns list of new files
func (g *Git) OnlyNewFiles(branch string) []string {
	out, err := g.do("diff", "--name-only", "--diff-filter=A", branch)
	if err != nil {
		log.Printf("can't get diff: %v", messages.ErrorMsg(err))
		return nil
	}
	return strings.Split(out, "\n")
}

func (g *Git) do(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = g.projectDir
	out, err := cmd.Output()
	return string(out), err
}
