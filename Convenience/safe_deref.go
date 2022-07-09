package convenience

type Nullable[A any] struct {
	val *A
}

type Maker[A any] func() *A

func Nvl[A any](a *A) Nullable[A] {
	return Nullable[A]{val: a}
}

func (nvl Nullable[A]) IsNil() bool {
	return nvl.val == nil
}

func (nvl Nullable[A]) NotNil() bool {
	return nvl.val != nil
}

func (nvl Nullable[A]) Or(a A) A {
	if nvl.IsNil() {
		return a
	}
	return *nvl.val
}

func (nvl Nullable[A]) OrCall(mk func() *A) *A {
	if nvl.IsNil() {
		return mk()
	}
	return nvl.val
}
