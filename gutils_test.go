package gUtils

import (
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	var r = []int{3, 2, 1, 4, 5}
	want, _ := json.Marshal([]int{5, 4, 1, 2, 3})
	Reverse(r)
	got, _ := json.Marshal(r)
	//t.Logf("got:%s", got)
	//t.Logf("want:%s", want)

	assert.Equal(t, string(got), string(want))
}

func TestSlice(t *testing.T) {
	var r = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t.Logf("r:%+v", r)
	t.Logf(" r[:]:%+v", r[:])
	t.Logf(" r[1:]:%+v", r[1:])
	t.Logf(" r[0:3]:%+v", r[0:3])
	t.Logf(" r[1:3]:%+v", r[1:3])
	t.Logf(" r[2:3]:%+v", r[2:3])
	t.Logf(" r[:3]:%+v", r[:3])
	t.Logf(" r[:len(r)]:%+v", r[:len(r)])
	//t.Logf(" r[:11]:%+v", r[:11])
	//t.Logf(" r[:100]:%+v", r[:100])
	t.Logf(" r[:10]:%+v", r[:10])
}

func TestIsInArray(t *testing.T) {
	var intList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var strList = []string{"1", "2", "3"}
	assert.Equal(t, IsInArray(1, intList), true, "IsInArray(1, intList)")
	assert.Equal(t, IsInArray(100, intList), false, "IsInArray(100, intList)")
	assert.Equal(t, IsInArray("1", strList), true, "IsInArray(\"1\", strList)")
	assert.Equal(t, IsInArray("100", strList), false, "IsInArray(\"2\", strList)")
}

func TestMin(t *testing.T) {
	var intList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, Min(intList...), 1)
}

func TestMax(t *testing.T) {
	var intList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, Max(intList...), 10)
}

func TestSum(t *testing.T) {
	var intList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, Sum(intList...), 55)
}

func TestCalcWeight(t *testing.T) {
	weights := []*Weight{
		{Value: "1", Weight: 0},
		{Value: "2", Weight: 0},
		{Value: "3", Weight: 3},
	}

	w, err := CalcWeight(weights)
	assert.Equal(t, err, nil)
	assert.Equal(t, w.Value, "3")

	weights = []*Weight{
		{Value: "1", Weight: 0},
		{Value: "2", Weight: 30},
		{Value: "3", Weight: 0},
	}

	w, err = CalcWeight(weights)
	assert.Equal(t, err, nil)
	assert.Equal(t, w.Value, "2")
}

func TestIsChinese(t *testing.T) {
	assert.Equal(t, IsChinese("你好"), true)
	assert.Equal(t, IsChinese("你好1"), false)
	assert.Equal(t, IsChinese("你好a"), false)
	assert.Equal(t, IsChinese("你 好"), false)
	assert.Equal(t, IsChinese(""), true)
}

func TestRandomString(t *testing.T) {
	t.Logf("TestRandomString:%s", RandomString(10))
}

func BenchmarkRandomString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RandomString(16)
	}
}

// go test -bench=.

func TestCheckStructEmpty(t *testing.T) {
	var s struct {
		Name string
		Age  int
	}
	err := CheckStructEmpty(s)
	t.Logf("err:%s", err)
	assert.Equal(t, err.Error(), "filed empty:Name")

	err = CheckStructEmpty(&s)
	t.Logf("err:%s", err)
	assert.Equal(t, err.Error(), "filed empty:Name")

	s.Name = "xiaoming"
	err = CheckStructEmpty(s)
	t.Logf("err:%s", err)
	assert.Equal(t, err.Error(), "filed empty:Age")

	err = CheckStructEmpty(s, "Age")
	t.Logf("err:%s", err)
	assert.Equal(t, err, nil)

	s.Age = 1
	err = CheckStructEmpty(s)
	t.Logf("err:%s", err)
	assert.Equal(t, err, nil)
}

func TestIfThen(t *testing.T) {

	assert.Equal(t, IfThen(true, 1, 0), 1)
	assert.Equal(t, IfThen(false, 1, 0), 0)

	age := 20
	assert.Equal(t, IfThen(age > 10, "old", "young"), "old")

	age = 1
	assert.Equal(t, IfThen(age > 10, "old", "young"), "young")
}
