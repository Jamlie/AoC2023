package utils

type Array[T any] []T

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



type Coordinates2D struct {
	X int
	Y int
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
