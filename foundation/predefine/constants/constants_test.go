package constants_test

import (
	"makeng/foundation/predefine/constants"
	"testing"
)

// TestMAKE_NG_VERSION checks if the MAKE_NG_VERSION constant is correct
func TestMAKE_NG_VERSION(t *testing.T) {
	expected := "0.0.1"
	if constants.MAKE_NG_VERSION != expected {
		t.Errorf("MAKE_NG_VERSION = %v; want %v", constants.MAKE_NG_VERSION, expected)
	}
}
