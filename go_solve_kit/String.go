package go_solve_kit

import (
	"sort"
	"strconv"
	"strings"
)

type String string
type StringArray []String

func (s String) ValueOf() string {
	return string(s)
}

func (s String) ToInt() Int {
	val, _ := strconv.Atoi(s.ValueOf())
	return Int(val)
}

func (s String) Split(sep string) StringArray {
	seps := strings.Split(s.ValueOf(), sep)
	var output StringArray
	for _, v := range seps {
		output = append(output, String(v))
	}
	return output
}

func (s String) Length() Int {
	return Int(len(s.ValueOf()))
}

func (s String) Trim(cutset string) String {
	return String(strings.Trim(s.ValueOf(), cutset))
}

func (s String) TrimSpace() String {
	return String(strings.TrimSpace(s.ValueOf()))
}

func (s String) TrimLeft(prefix string) String {
	return String(strings.TrimPrefix(s.ValueOf(), prefix))
}

func (s String) TrimRight(suffix string) String {
	return String(strings.TrimSuffix(s.ValueOf(), suffix))
}

func (s String) IndexOf(v string) Int {
	return Int(strings.Index(s.ValueOf(), v))
}

func (s String) Contains(v string) bool {
	return strings.Contains(s.ValueOf(), v)
}

func (s String) ToLower() String {
	return String(strings.ToLower(s.ValueOf()))
}

func (s String) ToUpper() String {
	return String(strings.ToUpper(s.ValueOf()))
}

func (array StringArray) Length() Int {
	return Int(len(array))
}

func (array StringArray) Map(lambda func(s String, i int) interface{}) TypeArray {
	var output TypeArray
	for i, v := range array {
		output = append(output, Type{lambda(v, i)})
	}
	return output
}

func (array StringArray) ForEach(lambda func(s String, i int)) {
	for i, v := range array {
		lambda(v, i)
	}
}

func (array StringArray) Filter(lambda func(s String, i int) bool) StringArray {
	var output StringArray
	for i, v := range array {
		if lambda(v, i) {
			output = append(output, v)
		}
	}
	return output
}

func (array StringArray) ToTypeArray() TypeArray {
	var output TypeArray
	for _, v := range array {
		output = append(output, Type{v.ValueOf()})
	}
	return output
}

func (array StringArray) ToIntArray() IntArray {
	var output IntArray
	for _, v := range array {
		output = append(output, v.ToInt())
	}
	return output
}

func (array StringArray) Join(sep string) String {
	tmpArray := make([]string, 0)
	for _, v := range array {
		tmpArray = append(tmpArray, v.ValueOf())
	}
	return String(strings.Join(tmpArray, sep))
}

func (array StringArray) Fill(v String) StringArray {
	return NewArray(len(array)).Map(func(_ Type, _ int) interface{} {
		return v
	}).ToStringArray()
}

func (array StringArray) Every(lambda func(v String, i int) bool) bool {
	for i, val := range array {
		if !lambda(val, i) {
			return false
		}
	}
	return true
}

func (array StringArray) Some(lambda func(v String, i int) bool) bool {
	for i, val := range array {
		if lambda(val, i) {
			return true
		}
	}
	return false
}

func (array StringArray) Contains(v string) bool {
	return array.Some(func(i String, _ int) bool {
		return i.ValueOf() == v
	})
}

func (array StringArray) FindIndex(lambda func(v String, i int) bool) Int {
	for i, val := range array {
		if lambda(val, i) {
			return Int(i)
		}
	}
	return -1
}

func (array StringArray) IndexOf(v string) Int {
	return array.FindIndex(func(i String, _ int) bool {
		return i.ValueOf() == v
	})
}

func (array *StringArray) Append(v string) {
	*array = append(*array, String(v))
}

func (array *StringArray) Pop() String {
	output := (*array)[array.Length()-1]
	*array = (*array)[:array.Length()-1]
	return output
}

func (array *StringArray) Enqueue(v String) {
	*array = append(StringArray{String(v)}, (*array)...)
}

func (array *StringArray) Dequeue() String {
	output := (*array)[0]
	*array = (*array)[1:]
	return output
}

func (array StringArray) First() String {
	return array[0]
}

func (array StringArray) Last() String {
	return array[array.Length()-1]
}

func (array *StringArray) Remove(v string) {
	i := array.IndexOf(v)
	*array = append((*array)[:i], (*array)[i+1:]...)
}

func (array StringArray) Sort() {
	sort.SliceStable(array, func(i, j int) bool {
		return array[i] < array[j]
	})
}

func (array StringArray) SortBy(lambda func(x, y String) bool) {
	sort.SliceStable(array, func(i, j int) bool {
		return lambda(array[i], array[j])
	})
}

func (array StringArray) Copy() StringArray {
	return append(make(StringArray, 0), array...)
}

func (array StringArray) Slice(start, end int) StringArray {
	if end <= 0 {
		end = array.Length().ValueOf() + end
	}

	return append(make(StringArray, 0), array[start:end]...)
}
