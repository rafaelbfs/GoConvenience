package assertions

import (
	"os"
	"testing"
)

func TestAssertion_NoError(t *testing.T) {
	file, err := os.Open("dummy.txt")
	Assert(t).NoError(err)
	AssertPointer(t, file).NotNil()
}
