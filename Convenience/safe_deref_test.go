package convenience

import (
	assertions "github.com/rafaelbfs/GoConvenience/Assertions"
	"testing"
)

func TestSafeNvl(t *testing.T) {
	var str = "NOT NULL STR"
	var alternative = "ALT STR"
	var strPtr = &str
	var nilPtr *string = nil

	maker := func() *string { return &alternative }

	assertions.Assert(t).Condition(Nvl(strPtr).NotNil()).IsTrueV()
	assertions.Assert(t).Condition(Nvl(nilPtr).IsNil()).IsTrueV()
	assertions.AssertThat(t, Nvl(strPtr).Or(alternative)).EqualsTo(str)
	assertions.AssertThat(t, *Nvl(nilPtr).OrCall(maker)).EqualsTo(alternative)
}
