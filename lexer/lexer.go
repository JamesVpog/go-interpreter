package lexer

import (
	"github.com/JamesVpog/go-interpreter/token"
)
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}


func (l *Lexer) NextToken() token.Token {
	// return the token of the current character under examination (l.ch)
	var tok token.Token
	
	
	l.skipWhitespace()
	// any new bytes are converted into a token for AST 
	switch l.ch {
		case '=':
			if l.peekChar() == '=' { // we are doing eq check
				ch := l.ch
				l.readChar()
				literal := string(ch) + string(ch)
				tok = token.Token{Type: token.EQ, Literal: literal}	
			} else {	// just one assignment
				tok = newToken(token.ASSIGN, l.ch)
			}
		case '+':
			tok = newToken(token.PLUS, l.ch)
		case '-':
			tok = newToken(token.MINUS, l.ch)
		case '/':
			tok = newToken(token.SLASH, l.ch)
		case '*':
			tok = newToken(token.ASTERISK, l.ch)
		case '!':		
			if l.peekChar() == '=' { // we are doing not_eq check
				ch := l.ch
				l.readChar()
				literal := string(ch) + string(l.ch)
				tok = token.Token{Type: token.NOT_EQ, Literal: literal}	
			} else {	// just excited!
				tok = newToken(token.BANG, l.ch)
			}
		case '<':
			tok = newToken(token.LT, l.ch)
		case '>':
			tok = newToken(token.GT, l.ch)
		case '(':
			tok = newToken(token.LPAREN, l.ch)
		case ')':
			tok = newToken(token.RPAREN, l.ch)
		case '{':
			tok = newToken(token.LBRACE, l.ch)
		case '}':
			tok = newToken(token.RBRACE, l.ch)
		case ';':
			tok = newToken(token.SEMICOLON, l.ch)
		case ',':
			tok = newToken(token.COMMA, l.ch)
		case 0: // eof char, the 0 byte is same as EOF
			tok.Literal = ""
			tok.Type = token.EOF
		// adding support for keywords and identifiers
		default:
			if isLetter(l.ch) {
				tok.Literal = l.readIdentifier()
				tok.Type = token.LookupIdent(tok.Literal)
				return tok // early return to not advance to far
			} else if isDigit(l.ch) {
				tok.Type = token.INT
				tok.Literal = l.readNumber()
				return tok
			} else {
				tok = newToken(token.ILLEGAL, l.ch)
			}
	}
	l.readChar()
	return tok 
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
// go to next char in string and advance our position pointers
func (l *Lexer) readChar() {	
	// TODO: support unicode/emojis... right now reads each token into one byte
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// look ahead for 2 char tokens 
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// read the identifier (names of things in our program) for all its chars
// and advnace the lexer's position until non letter char
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9';
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}