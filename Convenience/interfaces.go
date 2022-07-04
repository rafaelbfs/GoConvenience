package convenience

type Function[I any, R any] func(I) R

// Predicate[A any] represents a function that when evaluated on a value of type A,
// will return either true or false
type Predicate[A any] func(A) bool
type Combinator[I any, R any] func(R, I) R
type LoopAction[I any] func(I)

type List[A any] []A

func Not[A any](predicate Predicate[A]) Predicate[A] {
	return func(v A) bool { return !predicate(v) }
}

func Compose[A any, B any, C any](f Function[B, C], g Function[A, B]) Function[A, C] {
	return func(a A) C { return f(g(a)) }
}

func Chain[A any, B any, C any](f Function[A, B], g Function[B, C]) Function[A, C] {
	return func(a A) C { return g(f(a)) }
}

func (l List[A]) Loop(action LoopAction[*A]) List[A] {
	for i, _ := range l {
		action(&l[i])
	}
	return l
}

func (l List[A]) Filter(pred Predicate[A]) List[A] {
	var r = make([]A, 0, len(l))
	for _, v := range l {
		if pred(v) {
			r = append(r, v)
		}
	}
	return r
}

func FMap[A any, B any](function Function[A, B], a List[A]) List[B] {
	b := make([]B, len(a))
	for i, e := range a {
		b[i] = function(e)
	}
	return b
}

func FMapF[A any, B any](function Function[A, B]) Function[List[A], List[B]] {
	return func(a List[A]) List[B] {
		return FMap(function, a)
	}
}

func reduce[A any, B any](combiner func(*B, *A) *B, initial *B, a List[A]) *B {
	var acc = initial
	for i, _ := range a {
		acc = combiner(acc, &a[i])
	}
	return acc
}

func MkReducer[A any, B any](combiner func(*B, *A) *B, initial *B) Function[List[A], *B] {
	return func(a List[A]) *B {
		return reduce(combiner, initial, a)
	}
}
