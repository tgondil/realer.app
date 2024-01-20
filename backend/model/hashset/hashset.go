package hashset

import (
	"fmt"
	"github.com/valyala/bytebufferpool"
)

type HashSet[T comparable] map[T]struct{}

func NewWithValues[T comparable](values ...T) *HashSet[T] {
	defaultLen := 0
	switch l := len(values); l {
	case 0:
		break
	default:
		defaultLen = int(float64(l) * 1.25)
	}
	s := make(HashSet[T], defaultLen)
	for _, v := range values {
		s[v] = struct{}{}
	}
	return &s
}

func New[T comparable](length ...int) *HashSet[T] {
	defaultLen := 0
	switch l := len(length); l {
	case 0:
		break
	default:
		defaultLen = length[0]
	}
	s := make(HashSet[T], defaultLen)
	return &s
}

func (s *HashSet[T]) Add(v T) {
	if s == nil {
		*s = make(HashSet[T])
	}
	(*s)[v] = struct{}{}
}

func (s *HashSet[T]) Remove(v T) {
	if s == nil {
		return
	}
	delete(*s, v)
}

func (s *HashSet[T]) Contains(v T) bool {
	if s == nil {
		return false
	}
	_, ok := (*s)[v]
	return ok
}

func (s *HashSet[T]) ToSlice() []T {
	if s == nil {
		return nil
	}
	value := make([]T, 0, len(*s))
	for key := range *s {
		value = append(value, key)
	}
	return value
}

func (s *HashSet[T]) Join(separator string, stringer ...func(T) string) string {
	if s == nil {
		return ""
	}
	defaultStringer := func(t T) string {
		return fmt.Sprint(t)
	}
	if len(stringer) > 0 {
		defaultStringer = stringer[0]
	}
	buf := new(bytebufferpool.ByteBuffer)
	firstEntry := true
	for k := range *s {
		if firstEntry {
			firstEntry = false
		} else {
			_, _ = buf.WriteString(separator)
		}
		_, _ = buf.WriteString(defaultStringer(k))
	}
	value := buf.String()
	return value
}

func (s *HashSet[T]) Len() int {
	if s == nil {
		return 0
	}
	return len(*s)
}

func (s *HashSet[T]) IsEmpty() bool {
	if s == nil {
		return true
	}
	return len(*s) == 0
}

func (s *HashSet[T]) Clear() {
	if s == nil {
		return
	}
	*s = make(HashSet[T])
}

func (s *HashSet[T]) ForEach(f func(T)) {
	if s == nil {
		return
	}
	for key := range *s {
		f(key)
	}
}

func (s *HashSet[T]) Grow(len int) {
	if s == nil {
		return
	}
	if len > 0 {
		newM := make(HashSet[T], len)
		for key := range *s {
			newM[key] = struct{}{}
		}
		*s = newM
	}
}
