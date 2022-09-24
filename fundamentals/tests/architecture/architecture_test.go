package architecture

import (
	"runtime"
	"testing"
)

func TestDependent(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH != "amd64" {
		t.Logf("current architecture: %v", runtime.GOARCH)
		t.Skip("It doesn't work on amd64 architecture")
	}

	t.Fail()
}
