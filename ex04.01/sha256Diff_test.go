// test file for the sha256Diff package
//
//

package sha256Diff

import (
	"crypto/sha256"
	"testing"
)

func TestSha265Diff(t *testing.T) {
	a := sha256.Sum256([]byte{})
	b := sha256.Sum256([]byte{0x01})
	if Diff(a, a) != 0 {
		t.Errorf("\n%x\n%x\n diff should be 0, and it is not\n", a, a)
	}
	if diff := Diff(a, b); diff != 139 {
		t.Errorf("\n%x\n%x\n diff should be 139, and it is %d\n", a, b, diff)
	}
}
