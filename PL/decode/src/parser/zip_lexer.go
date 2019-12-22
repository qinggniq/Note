// Code generated from /Users/qinggniq/IdeaProjects/Cacl/src/zip.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 6, 19, 8, 
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 3, 3, 
	3, 3, 4, 3, 4, 3, 5, 3, 5, 2, 2, 6, 3, 3, 5, 4, 7, 5, 9, 6, 3, 2, 4, 4, 
	2, 67, 92, 99, 124, 3, 2, 50, 59, 2, 18, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 3, 11, 3, 2, 2, 2, 5, 13, 3, 
	2, 2, 2, 7, 15, 3, 2, 2, 2, 9, 17, 3, 2, 2, 2, 11, 12, 7, 93, 2, 2, 12, 
	4, 3, 2, 2, 2, 13, 14, 7, 95, 2, 2, 14, 6, 3, 2, 2, 2, 15, 16, 9, 2, 2, 
	2, 16, 8, 3, 2, 2, 2, 17, 18, 9, 3, 2, 2, 18, 10, 3, 2, 2, 2, 3, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "", "", "Char", "Digit",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "Char", "Digit",
}

type zipLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewzipLexer(input antlr.CharStream) *zipLexer {

	l := new(zipLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "zip.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// zipLexer tokens.
const (
	zipLexerT__0 = 1
	zipLexerT__1 = 2
	zipLexerChar = 3
	zipLexerDigit = 4
)

