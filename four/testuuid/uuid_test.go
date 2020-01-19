package testuuid

import (
	"testing"
)

func TestGetUUID(t *testing.T) {
	luuid := GetUUID()
	if len(luuid) < 31 {
		t.Errorf("uuid is OK")
	}

}

func TestLdkAdd(t *testing.T) {
	var (
		a        = 7
		b        = 5
		expected = 11
	)
	actual := LdkAdd(a, b)
	if actual != expected {
		t.Errorf("LdkAdd(%d,%d) = %d; expected %d", a, b, actual, expected)
	}
}
