package gogit

import (
	"testing"
)

func TestBranches(t *testing.T) {
	git := Git{"/Users/tuzki/htdocs/go/src/gogit"}
	branches, err := git.Branches()
	if err != nil {
		t.Fatal(err)
	}
	if len(branches) < 1 {
		t.Fatal("there is no branches")
	}
}
