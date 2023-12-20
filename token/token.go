package token

type TokenType string

type Token Struct {
    Type TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // Identifiers + literals
    IDENT = "IDENT" // Add, foobar, x, y
    INT = "INT" //12345

    // OPERATORS
    ASSIGN = "="
    PLUS = "+"


    // Delimiters
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
