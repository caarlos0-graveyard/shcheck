package sh

import (
	"testing"

	"github.com/tj/assert"
)

func TestSuccess(t *testing.T) {
	var assert = assert.New(t)
	for _, checker := range Checkers(Options{}) {
		assert.NoError(checker.Check("testdata/success.sh"))
	}
}

func TestShellcheckError(t *testing.T) {
	var assert = assert.New(t)
	var checker = &shellcheck{}
	assert.Error(checker.Check("testdata/shellcheck.sh"))
}

func TestShfmtError(t *testing.T) {
	var assert = assert.New(t)
	var checker = &shfmt{}
	assert.Error(checker.Check("testdata/malformat.sh"))
}
