package convenience

import (
	assertions "github.com/rafaelbfs/GoConvenience/Assertions"
	"strconv"
	s "strings"
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

func errorAtoi(err error) bool {
	msg := err.Error()
	return s.Contains(msg, "invalid syntax") &&
		s.Contains(msg, "strconv.Atoi")
}

func strMap(str string) *int {
	legalNr, err := strconv.Atoi(str)
	if err == nil {
		return &legalNr
	}
	return nil
}

func TestMapNvl(t *testing.T) {
	var str = "123"
	var notNr = "NO NUMBER"
	var f = MapNvl(strMap)
	var opt = f(Nvl(&str))
	var notNrOpt = f(Nvl(&notNr))
	assertions.AssertThat(t, opt.Or(321)).EqualsTo(123)
	assertions.AssertThat(t, notNrOpt.Or(321)).EqualsTo(321)

	_, err := strconv.Atoi(notNr)
	assertions.Assert(t).ThatError(err).Matches(errorAtoi)
}
