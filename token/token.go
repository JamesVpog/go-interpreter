package token

type TokenType string // special strings for our tokens

type Token struct {
	Type TokenType
	Literal string 
}

// all the token type strings
const (
	ILLEGAL = "ILLEGAL"
	EOF		= "EOF"
	
	// Identifiers and literals
	IDENT = "IDENT" // names of vars/functions: , add, foobar, x, y ...
	INT   = "INT"
	
	// Operators 
	ASSIGN 	= "="
	PLUS 	= "+"
	
	// Delimiters/special characters
	COMMA 	  = ","
	SEMICOLON = ";"
	
	RPAREN = ")"
	LPAREN = "("
	RBRACE = "{"
	LBRACE = "{"
	
	// Keywords
	FUNCTION = "FUNCTION"
	LET 	 = "LET"
)