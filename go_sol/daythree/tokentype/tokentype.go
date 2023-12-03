package tokentype

type TokenType int

const (
	PartNumber TokenType = iota
	WholeNumber
	Dot
)
