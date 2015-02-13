package gogit

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Git struct {
	Dir string
}

type Log struct {
	Revision string `json:"revision"`
	Message  string `json:"message"`
}

func (git *Git) Branches() (branches []string, err error) {
	stdout, stderr, err := git.command("branch")
	if err != nil {
		return branches, fmt.Errorf("%s", stderr.String())
	}
	scanner := bufio.NewScanner(strings.NewReader(stdout.String()))
	for scanner.Scan() {
		branchSlice := strings.SplitN(scanner.Text(), " ", 2)
		if len(branchSlice) > 1 {
			branches = append(branches, branchSlice[1])
		}
	}
	return branches, nil
}

func (git *Git) RemoteBranches(remote string) (branches []string, err error) {
	stdout, stderr, err := git.command("branch", "-r")
	if err != nil {
		return branches, fmt.Errorf("%s", stderr.String())
	}
	scanner := bufio.NewScanner(strings.NewReader(stdout.String()))
	for scanner.Scan() {
		branchSlice := strings.SplitN(scanner.Text(), remote+"/", 2)
		if len(branchSlice) > 1 {
			if strings.Contains(branchSlice[1], "HEAD") {
				continue
			}
			branches = append(branches, branchSlice[1])
		}
	}
	return branches, nil
}

func (git *Git) ActiveBranch() (string, error) {
	stdout, stderr, err := git.command("branch")
	if err != nil {
		return "", fmt.Errorf("%s", stderr.String())
	}
	scanner := bufio.NewScanner(strings.NewReader(stdout.String()))
	for scanner.Scan() {
		branchSlice := strings.SplitN(scanner.Text(), " ", 2)
		if len(branchSlice) > 1 && branchSlice[0] == "*" {
			return branchSlice[1], nil
		}
	}
	return "", fmt.Errorf("detect active branch failed")
}

func (git *Git) Pull(remote string) error {
	_, stderr, err := git.command("pull", remote)
	if err != nil {
		return fmt.Errorf("%s", stderr.String())
	}
	return nil
}

func (git *Git) Fetch(remote string, args ...string) error {
	newArgs := []string{"fetch", remote}
	for _, arg := range args {
		newArgs = append(newArgs, arg)
	}
	_, stderr, err := git.command(newArgs...)
	if err != nil {
		return fmt.Errorf("%s", stderr.String())
	}
	return nil
}

func (git *Git) SimpleLog(path string, limit int) (logs []*Log, err error) {
	args := []string{"--no-pager", "log", path, "-n", fmt.Sprintf("%d", limit), "--oneline"}
	stdout, stderr, err := git.command(args...)
	if err != nil {
		return logs, fmt.Errorf("%s", stderr.String())
	}
	scanner := bufio.NewScanner(strings.NewReader(stdout.String()))
	for scanner.Scan() {
		logSlice := strings.SplitN(scanner.Text(), " ", 2)
		if len(logSlice) > 1 {
			log := &Log{logSlice[0], logSlice[1]}
			logs = append(logs, log)
		}
	}
	return
}

func (git *Git) RevisionLog(revision string) (string, error) {
	args := []string{"--no-pager", "log", "-1", revision, "--oneline"}
	stdout, stderr, err := git.command(args...)
	if err != nil {
		return "", fmt.Errorf("%s", stderr.String())
	}
	logSlice := strings.SplitN(stdout.String(), " ", 2)
	if len(logSlice) > 1 {
		return logSlice[1], nil
	}
	return "", fmt.Errorf("Unexpected log format: %s", stdout.String())
}

func (git *Git) Checkout(branch string) error {
	_, stderr, err := git.command("checkout", branch)
	if err != nil {
		return fmt.Errorf("%s", stderr.String())
	}
	return nil
}

func (git *Git) Export(destination string) error {
	return nil
}

func (git *Git) Command(args ...string) (string, error) {
	stdout, stderr, err := git.command(args...)
	if err != nil {
		return "", fmt.Errorf("%s", stderr.String())
	}
	return stdout.String(), nil
}

func (git *Git) command(args ...string) (stdout bytes.Buffer, stderr bytes.Buffer, err error) {
	command := exec.Command("git", args...)
	command.Stdout = &stdout
	command.Stderr = &stderr
	command.Dir = git.Dir
	err = command.Run()
	return
}
