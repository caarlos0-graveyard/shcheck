package status

import (
	"testing"

	"time"
)

func TestAllStatuses(t *testing.T) {
	var file = "/path/to/my/file.sh"
	Info(file)
	time.Sleep(1 * time.Second)
	Success(file)
	Info(file)
	time.Sleep(1 * time.Second)
	Fail(file)
}
