package assertions

import (
	"testing"
)

type TestPredicate[A any] func(A) bool

type Assertion struct {
	test *testing.T
}

// AssertValue[A comparable] is a convenient type for test case assertions
// on a certain value of type A
type AssertValue[A comparable] struct {
	Assertion
	test *testing.T
	val  A
}

type PointerAssertion[A any] struct {
	Assertion
	test    *testing.T
	pointer *A
}

func AssertPointer[A any](tst *testing.T, pointer *A) PointerAssertion[A] {
	return PointerAssertion[A]{test: tst, pointer: pointer}
}

// AssertThat[A comparable] creates an AssertValue holding a concrete value of type A
func AssertThat[A comparable](tst *testing.T, value A) AssertValue[A] {
	return AssertValue[A]{test: tst, val: value}
}

// (it AssertValue[A]) EqualsTo(otherVal A) compares that the value held by 'it' is
// equal to the parameter otherVal, if not the test fails
func (it AssertValue[A]) EqualsTo(otherVal A) {
	if it.val == otherVal {
		return
	}
	it.test.Errorf("Test failed %v is not equal to %v", it.val, otherVal)
}

func (it PointerAssertion[A]) NotNil() {
	if it.pointer == nil {
		it.test.Errorf("Test failed, value is null")
	}
}

func (it PointerAssertion[A]) IsNil() {
	if it.pointer != nil {
		it.test.Errorf("Test failed, value is present")
	}
}

// (it AssertValue[A]) Satisfies(predicate Predicate[A]) AssertValue[A]
// evaluates that value held by 'it' satisfies the predicate parameter
func (it AssertValue[A]) Satisfies(predicate TestPredicate[A]) AssertValue[A] {
	if predicate(it.val) {
		return it
	}
	it.test.Errorf("%v does not satisfy the given condition", it.val)
	return it
}

type TestCondition struct {
	test *testing.T
	cond bool
}

func Assert(tst *testing.T) Assertion {
	return Assertion{test: tst}
}

func (it Assertion) NoError(err error) {
	if err != nil {
		it.test.Errorf("Expected no error but got %v", err)
	}
}

func (it Assertion) Condition(condition bool) TestCondition {
	return TestCondition{test: it.test, cond: condition}
}

func (it TestCondition) IsTrue(messageOtherwise string) {
	if it.cond {
		return
	}
	it.test.Errorf(messageOtherwise)
}

func (it TestCondition) IsTrueV() {
	it.IsTrue("Failure expected a condition to be true")
}

func (it TestCondition) IsFalse(messageOtherwise string) {
	if !it.cond {
		return
	}
	it.test.Errorf(messageOtherwise)
}

func (it TestCondition) IsFalseV() {
	it.IsFalse("Failure expected a condition to be false")
}
