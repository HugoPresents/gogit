package gogit

import (
	"fmt"
	"testing"
)

var path string
var git *Git

func init() {
	git = &Git{"/Users/rabbit/htdocs/tff/front"}
}

func TestBranches(t *testing.T) {
	branches, err := git.Branches()
	if err != nil {
		t.Fatal(err)
	}
	if len(branches) < 1 {
		t.Fatal("there is no branches")
	}
	fmt.Printf("local branches: %v\n", branches)
}

func TestDeleteBranch(t *testing.T) {
	err := git.DeleteBranch("none_branch")
	if err == nil {
		t.Fatal("error is nil")
	}
}

func TestActiveBranch(t *testing.T) {
	branch, err := git.ActiveBranch()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("active branch: %s\n", branch)
}

func TestRemoteBranches(t *testing.T) {
	remoteBranches, err := git.RemoteBranches("origin")
	if err != nil {
		t.Fatal(err)
	}
	if len(remoteBranches) < 1 {
		t.Fatal("there is no remote branches")
	}
	//fmt.Printf("remote %s branches: %v\n", "origin", remoteBranches)
}

func TestFetch(t *testing.T) {
	err := git.Fetch("office")
	if err == nil {
		t.Fatal("error is nil")
	}
}

func TestCommand(t *testing.T) {
	_, err := git.Command("clean", "-xxx")
	if err == nil {
		t.Fatal("error is nil")
	}
}
