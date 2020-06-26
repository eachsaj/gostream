package lexer

import (
	"strings"
	"testing"
)

func TestLexerTokenizeMin(t *testing.T) {
	q := "select min(Level) from LogEvent.time(10 sec) where Level > 2"
	lexer := New(strings.NewReader(q))

	var test = []struct {
		token   Token
		literal string
	}{
		{SELECT, "select"},
		{MIN, "min"},
		{LPAREN, "("},
		{IDENTIFIER, "Level"},
		{RPAREN, ")"},
		{FROM, "from"},
		{IDENTIFIER, "LogEvent"},
		{DOT, "."},
		{TIME, "time"},
		{LPAREN, "("},
		{IDENTIFIER, "10"},
		{SEC, "sec"},
		{RPAREN, ")"},
		{WHERE, "where"},
		{IDENTIFIER, "Level"},
		{LARGER, ">"},
		{IDENTIFIER, "2"},
	}

	for _, tt := range test {
		token, literal := lexer.TokenizeIgnoreWhiteSpace()
		if token != tt.token || literal != tt.literal {
			t.Error(token, literal)
		}
	}
}

func TestLexerTokenizeFloat(t *testing.T) {
	q := "select count(*) from LogEvent.time(10 sec) where Level > 2.5"
	lexer := New(strings.NewReader(q))

	var test = []struct {
		token   Token
		literal string
	}{
		{SELECT, "select"},
		{COUNT, "count"},
		{LPAREN, "("},
		{ASTERISK, "*"},
		{RPAREN, ")"},
		{FROM, "from"},
		{IDENTIFIER, "LogEvent"},
		{DOT, "."},
		{TIME, "time"},
		{LPAREN, "("},
		{IDENTIFIER, "10"},
		{SEC, "sec"},
		{RPAREN, ")"},
		{WHERE, "where"},
		{IDENTIFIER, "Level"},
		{LARGER, ">"},
		{IDENTIFIER, "2"},
		{DOT, "."},
		{IDENTIFIER, "5"},
	}

	for _, tt := range test {
		token, literal := lexer.TokenizeIgnoreWhiteSpace()
		if token != tt.token || literal != tt.literal {
			t.Error(token, literal)
		}
	}
}

func TestLexerTokenize(t *testing.T) {
	q := "select Value, count(*), avg(Value), sum(Value) from MyEvent.time(10 sec) where Value > 97"
	lexer := New(strings.NewReader(q))

	var test = []struct {
		token   Token
		literal string
	}{
		{SELECT, "select"},
		{WHITESPACE, " "},
		{IDENTIFIER, "Value"},
		{COMMA, ","},
		{WHITESPACE, " "},
		{COUNT, "count"},
		{LPAREN, "("},
		{ASTERISK, "*"},
		{RPAREN, ")"},
		{COMMA, ","},
		{WHITESPACE, " "},
		{AVG, "avg"},
		{LPAREN, "("},
		{IDENTIFIER, "Value"},
		{RPAREN, ")"},
		{COMMA, ","},
		{WHITESPACE, " "},
		{SUM, "sum"},
		{LPAREN, "("},
		{IDENTIFIER, "Value"},
		{RPAREN, ")"},
		{WHITESPACE, " "},
		{FROM, "from"},
		{WHITESPACE, " "},
		{IDENTIFIER, "MyEvent"},
		{DOT, "."},
		{TIME, "time"},
		{LPAREN, "("},
		{IDENTIFIER, "10"},
		{WHITESPACE, " "},
		{SEC, "sec"},
		{RPAREN, ")"},
		{WHITESPACE, " "},
		{WHERE, "where"},
		{WHITESPACE, " "},
		{IDENTIFIER, "Value"},
		{WHITESPACE, " "},
		{LARGER, ">"},
		{WHITESPACE, " "},
		{IDENTIFIER, "97"},
	}

	for _, tt := range test {
		token, literal := lexer.Tokenize()
		if token != tt.token || literal != tt.literal {
			t.Error(token, literal)
		}
	}
}

func TestLexerTokenizeIgnoreSpaceTimeWindow(t *testing.T) {
	q := "select count(*) from LogEvent.time(10 sec) where Level > 2"
	lexer := New(strings.NewReader(q))

	var test = []struct {
		token   Token
		literal string
	}{
		{SELECT, "select"},
		{COUNT, "count"},
		{LPAREN, "("},
		{ASTERISK, "*"},
		{RPAREN, ")"},
		{FROM, "from"},
		{IDENTIFIER, "LogEvent"},
		{DOT, "."},
		{TIME, "time"},
		{LPAREN, "("},
		{IDENTIFIER, "10"},
		{SEC, "sec"},
		{RPAREN, ")"},
		{WHERE, "where"},
		{IDENTIFIER, "Level"},
		{LARGER, ">"},
		{IDENTIFIER, "2"},
	}

	for _, tt := range test {
		token, literal := lexer.TokenizeIgnoreWhiteSpace()
		if token != tt.token || literal != tt.literal {
			t.Error(token, literal)
		}
	}
}

func TestLexerTokenizeIgnoreSpace(t *testing.T) {
	q := "select Value, count(*), avg(Value), sum(Value) from MyEvent.time(10 sec) where Value > 97"
	lexer := New(strings.NewReader(q))

	var test = []struct {
		token   Token
		literal string
	}{
		{SELECT, "select"},
		{IDENTIFIER, "Value"},
		{COMMA, ","},
		{COUNT, "count"},
		{LPAREN, "("},
		{ASTERISK, "*"},
		{RPAREN, ")"},
		{COMMA, ","},
		{AVG, "avg"},
		{LPAREN, "("},
		{IDENTIFIER, "Value"},
		{RPAREN, ")"},
		{COMMA, ","},
		{SUM, "sum"},
		{LPAREN, "("},
		{IDENTIFIER, "Value"},
		{RPAREN, ")"},
		{FROM, "from"},
		{IDENTIFIER, "MyEvent"},
		{DOT, "."},
		{TIME, "time"},
		{LPAREN, "("},
		{IDENTIFIER, "10"},
		{SEC, "sec"},
		{RPAREN, ")"},
		{WHERE, "where"},
		{IDENTIFIER, "Value"},
		{LARGER, ">"},
		{IDENTIFIER, "97"},
	}

	for _, tt := range test {
		token, literal := lexer.TokenizeIgnoreWhiteSpace()
		if token != tt.token || literal != tt.literal {
			t.Error(token, literal)
		}
	}
}
