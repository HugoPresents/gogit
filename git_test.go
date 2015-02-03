package gogit

import (
	"testing"
)

var path string

func init() {
	path = "/Users/tuzki/htdocs/go/src/gogit"
}

func TestBranches(t *testing.T) {
	git := Git{path}
	branches, err := git.Branches()
	if err != nil {
		t.Fatal(err)
	}
	if len(branches) < 1 {
		t.Fatal("there is no branches")
	}
}

func TestRemoteBranches(t *testing.T) {
	git := Git{path}
	remoteBranches, err := git.RemoteBranches("origin")
	if err != nil {
		t.Fatal(err)
	}
	if len(remoteBranches) < 1 {
		t.Fatal("there is no remote branches")
	}
}
