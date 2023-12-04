package utils

import (
	"fmt"
	"log"
	"strconv"
)

type Array[T comparable] []T

func (a Array[T]) ForEach(f func(T)) {
	for _, v := range a {
		f(v)
	}
}

func (a Array[T]) Map(f func(T) T) Array[T] {
	result := Array[T]{}
	for _, v := range a {
		result = append(result, f(v))
	}
	return result
}

func (a Array[T]) Filter(f func(T) bool) Array[T] {
	result := Array[T]{}
	for _, v := range a {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func (a Array[T]) Reduce(f func(T, T) T, initial *T) T {
	result := a[0]
	for _, v := range a[1:] {
		result = f(result, v)
	}
	return result
}

func (a Array[T]) Contains(item T) bool {
	for _, v := range a {
		if v == item {
			return true
		}
	}
	return false
}

func (a Array[T]) Fill(item T) Array[T] {
	for i := range a {
		a[i] = item
	}
	return a
}


type PartType int

const (
	PartOne PartType = iota
	PartTwo
)

type Coordinates2D struct {
	X int
	Y int
}

func (c *Coordinates2D) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

func (c Coordinates2D) Add(other Coordinates2D) Coordinates2D {
	return Coordinates2D{c.X + other.X, c.Y + other.Y}
}

func (c Coordinates2D) Sub(other Coordinates2D) Coordinates2D {
	return Coordinates2D{c.X - other.X, c.Y - other.Y}
}

func (c Coordinates2D) Mul(other Coordinates2D) Coordinates2D {
	return Coordinates2D{c.X * other.X, c.Y * other.Y}
}

func (c Coordinates2D) Div(other Coordinates2D) Coordinates2D {
	return Coordinates2D{c.X / other.X, c.Y / other.Y}
}

func IsSpace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func IsLetter(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func IsDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

func IsSome(c rune, chars []rune) bool {
	for _, ch := range chars {
		if c == ch {
			return true
		}
	}
	return false
}

func Range(start, end int) []int {
	result := []int{}
	if start > end {
		log.Fatalf("Start %d is greater than end %d", start, end)
		return result
	}
	for i := start; i < end; i++ {
		result = append(result, i)
	}
	return result
}

func StrArrayToIntArray(strArray []string) Array[int] {
	result := Array[int]{}
	for _, str := range strArray {
		if str == "" {
			continue
		}
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, num)
	}
	return result
}
