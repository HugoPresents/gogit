package gogit

import (
	"fmt"
	"testing"
)

var path string
var git *Git

func init() {
	git = &Git{"/Users/tuzki/htdocs/go/src/gogit"}
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
	fmt.Printf("remote %s branches: %v\n", "origin", remoteBranches)
}
