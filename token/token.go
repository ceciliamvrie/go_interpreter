// token/token.go

package token

type TokenType string
type Token struct {
	Type	TokenType
	Literal	string
}

const (
	ILLEGAL = "ILLEGAL" // unknown tokens
	EOF = "EOF" // end of file
	
	// IDENTIFIERS AND LITERALS
	IDENT = "IDENT" // variable names, function names
	INT = "INT" // integers

	// OPERATORS
	ASSIGN = "="
	PLUS = "+"
	
	// DELIMITERS
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// KEYWORDS
	FUNCTION = "FUNCTION"
	LET = "LET"
)
