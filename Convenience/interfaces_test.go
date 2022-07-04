package convenience

import (
	a "github.com/rafaelbfs/GoConvenience/Assertions"
	"math"
	"testing"
)

func TestFunctor(t *testing.T) {
	var f Function[int8, int16] = func(i8 int8) int16 {
		return 2 * int16(i8)
	}
	var plus2 LoopAction[*int16] = func(i16 *int16) {
		v := (*i16) + 2
		*i16 = v
	}
	var chars List[int8] = []int8{122, 125, 127}
	result := FMapF(f)(chars).Loop(plus2)

	greaterThan100 := func(n int16) bool { return n > 100 }

	for i, i8 := range chars {
		res := int16(i8)*2 + 2
		a.AssertThat(t, res).Satisfies(greaterThan100).EqualsTo(result[i])
		a.Assert(t).Condition((result[i] % 2) == 1).IsFalseV()
	}
}

func TestReducing(t *testing.T) {
	var squared = func(acc *map[int8]int, item *int8) *map[int8]int {
		(*acc)[*item] = int(*item) * int(*item)
		return acc
	}

	chars := []int8{12, 16, 20}
	aux := make(map[int8]int)
	result := MkReducer[int8, map[int8]int](squared, &aux)(chars)
	t.Log(result)
}

func TestComposing(t *testing.T) {
	chars := []int8{72, 32, 127}
	var times2 Function[int8, int16] = func(i8 int8) int16 {
		return 2 * int16(i8)
	}
	var root = func(i16 int16) float64 {
		return math.Sqrt(float64(i16))
	}
	result := FMapF(Compose(root, times2))(chars)

	var isRound = func(n float64) bool {
		floor := math.Floor(n)
		return (n - floor) == 0
	}
	onlyRoundNrs := result.Filter(isRound)
	t.Logf("Before filter =%v, after=%v", result, onlyRoundNrs)

	for _, n := range onlyRoundNrs {
		a.AssertThat(t, n).Satisfies(isRound)
	}
}
