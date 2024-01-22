package token

// トークンの種類を識別するためのタイプ　整数、識別子（変数名）、キーワード（let,fn）、記号
type TokenType string

// トークンはタイプとリテラル（実際の値）で構成する
type Token struct {
	Type    TokenType
	Literal string
}

// トークンタイプの定数を定義
const (
	//エラー用
	ILLEGAL = "ILLEGAL"
	//終端用
	EOF = "EOF"

	//識別子　＋　リテラル
	IDENT  = "IDENT //add, foobar, x,y, ..."
	INT    = "INT" //123456
	STRING = "STRING"

	//演算子
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	//デリミタ
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	//キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	MACRO    = "MACRO"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"macro":  MACRO,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
