package sliceutils

import (
	"fmt"
	"github.com/valyala/bytebufferpool"
)

var pool bytebufferpool.Pool

func SliceGroupBy[Key comparable, SliceType any](array []SliceType, key func(SliceType) Key) map[Key][]SliceType {
	res := make(map[Key][]SliceType)
	for _, val := range array {
		k := key(val)
		existingSlice, alreadyExists := res[k]
		if alreadyExists {
			res[k] = append(existingSlice, val)
		} else {
			newV := make([]SliceType, 1, 4)
			newV[0] = val
			res[k] = newV
		}
	}
	return res
}

func SliceMapTo[from any, to any](array []from, mapper func(*from) to) (value []to) {
	value = make([]to, 0, cap(array))
	for _, v := range array {
		value = append(value, mapper(&v))
	}
	return value
}

func SliceJoin[T any](values []T, separator string, stringer ...func(T) string) string {
	switch l := len(values); l {
	case 0:
		return ""
	case 1:
		return fmt.Sprint(values[0])
	default:
		toStr := func(t T) string {
			return fmt.Sprint(t)
		}
		if len(stringer) > 0 {
			toStr = stringer[0]
		}
		buf := pool.Get()
		_, _ = buf.WriteString(toStr(values[0]))
		for _, s := range values[1:] {
			_, _ = buf.WriteString(separator)
			_, _ = buf.WriteString(toStr(s))
		}
		value := buf.String()
		pool.Put(buf)
		return value
	}
}

func KeysJoin[T comparable, T2 any](values map[T]T2, separator string) string {
	buf := pool.Get()
	firstEntry := true
	for k := range values {
		if firstEntry {
			firstEntry = false
		} else {
			_, _ = buf.WriteString(separator)
		}
		_, _ = buf.WriteString(fmt.Sprint(k))
	}
	value := buf.String()
	pool.Put(buf)
	return value
}

func ValuesJoin[T comparable, T2 any](values map[T]T2, separator string) string {
	buf := pool.Get()
	firstEntry := true
	for _, v := range values {
		if firstEntry {
			firstEntry = false
		} else {
			_, _ = buf.WriteString(separator)
		}
		_, _ = buf.WriteString(fmt.Sprint(v))
	}
	value := buf.String()
	pool.Put(buf)
	return value
}

func MapKeysToSlice[T comparable, T2 any](values map[T]T2) []T {
	res := make([]T, 0, len(values))
	for k := range values {
		res = append(res, k)
	}
	return res
}

func MapValuesToSlice[T comparable, T2 any](values map[T]T2) []T2 {
	res := make([]T2, 0, len(values))
	for _, v := range values {
		res = append(res, v)
	}
	return res
}

func SliceWhere[T any](values []T, condition func(T) bool) []T {
	res := make([]T, 0, len(values))
	for _, v := range values {
		if condition(v) {
			res = append(res, v)
		}
	}
	return res
}
