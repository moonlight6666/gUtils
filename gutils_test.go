package gUtils

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

//func TestReverse(t *testing.T) {
//	var r = []int{3, 2, 1, 4, 5}
//	want, _ := json.Marshal([]int{5, 4, 1, 2, 3})
//	Reverse(r)
//	got, _ := json.Marshal(r)
//	//t.Logf("got:%s", got)
//	//t.Logf("want:%s", want)
//
//	assert.Equal(t, string(got), string(want))
//}

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

func TestMd5(t *testing.T) {
	s := Md5String("172380068220000009908a96d8fbf694e1911a2c6e1d91a198.CQGAME")
	assert.Equal(t, s, "94ad8936057de9133c1d60dd9e6370de")
}

func TestSHA1(t *testing.T) {
	s := SHA1("94ad8936057de9133c1d60dd9e6370de")
	assert.Equal(t, s, "ecf9cf8ee545da387d81f3eb897a9a0b40614af0")
}

type testStruct struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Level int    `json:"level,omitempty"`
	From  string `json:"from,omitempty"`
}

func TestMakeSignSource(t *testing.T) {
	v := testStruct{
		Name:  "xiaoming",
		Age:   99,
		Level: 10,
		From:  "china",
	}
	source, err := MakeSignSource(v, "json", false, ' ', '#')
	assert.Equal(t, err, nil)
	assert.Equal(t, source, "99#china#10#xiaoming")

	source, err = MakeSignSource(v, "json", true, '=', '#')
	assert.Equal(t, err, nil)
	assert.Equal(t, source, "age=99#from=china#level=10#name=xiaoming")

	source, err = MakeSignSource(v, "json", true, '=', '&')
	assert.Equal(t, err, nil)
	assert.Equal(t, source, "age=99&from=china&level=10&name=xiaoming")
}

func TestMakeSignSourceForm(t *testing.T) {
	v := testStruct{
		Name:  "xiaoming",
		Age:   99,
		Level: 10,
		From:  "china",
	}
	source, err := MakeSignSourceForm(v, "json")
	assert.Equal(t, err, nil)
	assert.Equal(t, source, "age=99&from=china&level=10&name=xiaoming")
}

func TestUniq(t *testing.T) {
	r := Uniq([]int{9, 1, 1, 2, 9})
	assert.Equal(t, r, []int{9, 1, 2})

	r = Uniq([]int{})
	assert.Equal(t, r, []int{})

	r2 := Uniq([]string{"qq", "wx", "qq", "qq"})
	assert.Equal(t, r2, []string{"qq", "wx"})

	r3 := Uniq([]interface{}{1, "2", "a", "2"})
	assert.Equal(t, r3, []interface{}{1, "2", "a"})
}

func TestFilter(t *testing.T) {
	r := Filter([]int{1, 2, 3, 4, 5}, func(v int, _ int) bool {
		return v%2 == 0
	})
	assert.Equal(t, r, []int{2, 4})

	r2 := Filter([]string{"", "1", "3", ""}, func(v string, _ int) bool {
		return v != ""
	})
	assert.Equal(t, r2, []string{"1", "3"})
}

func TestMap(t *testing.T) {
	r := Map([]int{1, 2, 3}, func(item int, index int) string {
		return fmt.Sprintf("s%d", item)
	})

	assert.Equal(t, r, []string{"s1", "s2", "s3"})
}

func TestGroupBy(t *testing.T) {
	r := GroupBy([]int{0, 1, 2, 3}, func(i int) int {
		return i % 2
	})
	assert.Equal(t, r, map[int][]int{0: []int{0, 2}, 1: []int{1, 3}})
}

func TestReverse(t *testing.T) {
	assert.Equal(t, Reverse([]int{2, 3, 4}), []int{4, 3, 2})
	assert.Equal(t, Reverse([]int{2, 3, 4, 1}), []int{1, 4, 3, 2})
	assert.Equal(t, Reverse([]int{}), []int{})
	assert.Equal(t, Reverse([]int{1}), []int{1})

	assert.Equal(t, Reverse([]string{"a", "c"}), []string{"c", "a"})
}
