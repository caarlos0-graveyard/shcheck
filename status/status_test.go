package status

import (
	"testing"
)

func TestAllStatuses(t *testing.T) {
	var file = "/path/to/my/file.sh"
	Ignore(file)
	Success(file)
	Fail(file)
}
