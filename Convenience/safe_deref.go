/*
 * Distributed under BSD 3-Clause License
 *
 * Copyright (c) 2022, Rafael Barros Felix, github.com/rafaelbfs
 * All rights reserved.
 */

package convenience

import "log"

type Nullable[A any] struct {
	val *A
}

type Error struct {
	e error
}

type Maker[A any] func() *A

type Maybe[R any] struct {
	err    error
	result R
}

func Nvl[A any](a *A) Nullable[A] {
	return Nullable[A]{val: a}
}

func Empty[A any]() Nullable[A] {
	return Nullable[A]{val: nil}
}

func (nvl Nullable[A]) IsNil() bool {
	return nvl.val == nil
}

func (nvl Nullable[A]) NotNil() bool {
	return nvl.val != nil
}

func (nvl Nullable[A]) DoIfPresent(action LoopAction[*A]) {
	if nvl.IsNil() {
		return
	}
	action(nvl.val)
}

func MapNvl[A any, B any](function Function[A, *B]) Function[Nullable[A], Nullable[B]] {
	return func(nvlA Nullable[A]) Nullable[B] {
		if nvlA.IsNil() {
			return Empty[B]()
		}
		return Nvl(function(*nvlA.val))
	}
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

func WrapError(err error) Error {
	return Error{e: err}
}

func (err Error) AndHandle(handler func(err error)) {
	if err.e == nil {
		return
	}
	handler(err.e)
}

func (err Error) AndPanic() {
	panic(err)
}

func LogAndDisregard(err Error) {
	log.Printf("Error %v will be disregarded", err)
}

func Try[R any](result R, err error) Maybe[R] {
	return Maybe[R]{err: err, result: result}
}

func (m Maybe[R]) ResultOrPanic() R {
	if m.err != nil {
		panic(m.err)
	}
	return m.result
}

func (m Maybe[R]) IsSuccessful() bool {
	return m.err == nil
}

func (m Maybe[R]) HasError() bool {
	return m.err != nil
}

// (m Maybe[R]) HandleErr(handler func(err error) R) R allows you to pass a function
// which can return a default value for the result in case of error
func (m Maybe[R]) HandleErr(handler func(err error) R) R {
	if m.err != nil {
		return handler(m.err)
	}
	return m.result
}
