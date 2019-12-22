// Code generated from /Users/qinggniq/IdeaProjects/Cacl/src/zip.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // zip

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 6, 36, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 24, 10, 2, 3, 
	3, 3, 3, 3, 3, 5, 3, 29, 10, 3, 3, 4, 6, 4, 32, 10, 4, 13, 4, 14, 4, 33, 
	3, 4, 2, 2, 5, 2, 4, 6, 2, 2, 2, 37, 2, 23, 3, 2, 2, 2, 4, 28, 3, 2, 2, 
	2, 6, 31, 3, 2, 2, 2, 8, 24, 5, 6, 4, 2, 9, 10, 5, 4, 3, 2, 10, 11, 7, 
	3, 2, 2, 11, 12, 5, 2, 2, 2, 12, 13, 7, 4, 2, 2, 13, 24, 3, 2, 2, 2, 14, 
	15, 5, 6, 4, 2, 15, 16, 5, 2, 2, 2, 16, 24, 3, 2, 2, 2, 17, 18, 5, 4, 3, 
	2, 18, 19, 7, 3, 2, 2, 19, 20, 5, 2, 2, 2, 20, 21, 7, 4, 2, 2, 21, 22, 
	5, 2, 2, 2, 22, 24, 3, 2, 2, 2, 23, 8, 3, 2, 2, 2, 23, 9, 3, 2, 2, 2, 23, 
	14, 3, 2, 2, 2, 23, 17, 3, 2, 2, 2, 24, 3, 3, 2, 2, 2, 25, 29, 7, 6, 2, 
	2, 26, 27, 7, 6, 2, 2, 27, 29, 5, 4, 3, 2, 28, 25, 3, 2, 2, 2, 28, 26, 
	3, 2, 2, 2, 29, 5, 3, 2, 2, 2, 30, 32, 7, 5, 2, 2, 31, 30, 3, 2, 2, 2, 
	32, 33, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 7, 3, 2, 
	2, 2, 5, 23, 28, 33,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "Char", "Digit",
}

var ruleNames = []string{
	"unit", "times", "content",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type zipParser struct {
	*antlr.BaseParser
}

func NewzipParser(input antlr.TokenStream) *zipParser {
	this := new(zipParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "zip.g4"

	return this
}

// zipParser tokens.
const (
	zipParserEOF = antlr.TokenEOF
	zipParserT__0 = 1
	zipParserT__1 = 2
	zipParserChar = 3
	zipParserDigit = 4
)

// zipParser rules.
const (
	zipParserRULE_unit = 0
	zipParserRULE_times = 1
	zipParserRULE_content = 2
)

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = zipParserRULE_unit
	return p
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = zipParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) CopyFrom(ctx *UnitContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type SingalContext struct {
	*UnitContext
}

func NewSingalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingalContext {
	var p = new(SingalContext)

	p.UnitContext = NewEmptyUnitContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnitContext))

	return p
}

func (s *SingalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingalContext) Content() IContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContentContext)
}


func (s *SingalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(zipListener); ok {
		listenerT.EnterSingal(s)
	}
}

func (s *SingalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(zipListener); ok {
		listenerT.ExitSingal(s)
	}
}

func (s *SingalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case zipVisitor:
		return t.VisitSingal(s)

	default:
		return t.VisitChildren(s)
	}
}


type TimesCombContext struct {
	*UnitContext
}

func NewTimesCombContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimesCombContext {
	var p = new(TimesCombContext)

	p.UnitContext = NewEmptyUnitContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnitContext))

	return p
}

func (s *TimesCombContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimesCombContext) Times() ITimesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimesContext)
}

func (s *TimesCombContext) AllUnit() []IUnitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnitContext)(nil)).Elem())
	var tst = make([]IUnitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnitContext)
		}
	}

	return tst
}

func (s *TimesCombContext) Unit(i int) IUnitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}


func (s *TimesCombContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(zipListener); ok {
		listenerT.EnterTimesComb(s)
	}
}

func (s *TimesCombContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(zipListener); ok {
		listenerT.ExitTimesComb(s)
	}
}

func (s *TimesCombContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case zipVisitor:
		return t.VisitTimesComb(s)

	default:
		return t.VisitChildren(s)
	}
}


type SingalfCombContext struct {
	*UnitContext
}

func NewSingalfCombContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingalfCombContext {
	var p = new(SingalfCombContext)

	p.UnitContext = NewEmptyUnitContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnitContext))

	return p
}

func (s *SingalfCombContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingalfCombContext) Content() IContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContentContext)
}

func (s *SingalfCombContext) Unit() IUnitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}


func (s *SingalfCombContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(zipListener); ok {
		listenerT.EnterSingalfComb(s)
	}
}

func (s *SingalfCombContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(zipListener); ok {
		listenerT.ExitSingalfComb(s)
	}
}

func (s *SingalfCombContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case zipVisitor:
		return t.VisitSingalfComb(s)

	default:
		return t.VisitChildren(s)
	}
}


type TimesSignalContext struct {
	*UnitContext
}

func NewTimesSignalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimesSignalContext {
	var p = new(TimesSignalContext)

	p.UnitContext = NewEmptyUnitContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnitContext))

	return p
}

func (s *TimesSignalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimesSignalContext) Times() ITimesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimesContext)
}

func (s *TimesSignalContext) Unit() IUnitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}


func (s *TimesSignalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(zipListener); ok {
		listenerT.EnterTimesSignal(s)
	}
}

func (s *TimesSignalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(zipListener); ok {
		listenerT.ExitTimesSignal(s)
	}
}

func (s *TimesSignalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case zipVisitor:
		return t.VisitTimesSignal(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *zipParser) Unit() (localctx IUnitContext) {
	localctx = NewUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, zipParserRULE_unit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(6)
			p.Content()
		}


	case 2:
		localctx = NewTimesSignalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(7)
			p.Times()
		}
		{
			p.SetState(8)
			p.Match(zipParserT__0)
		}
		{
			p.SetState(9)
			p.Unit()
		}
		{
			p.SetState(10)
			p.Match(zipParserT__1)
		}


	case 3:
		localctx = NewSingalfCombContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(12)
			p.Content()
		}
		{
			p.SetState(13)
			p.Unit()
		}


	case 4:
		localctx = NewTimesCombContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(15)
			p.Times()
		}
		{
			p.SetState(16)
			p.Match(zipParserT__0)
		}
		{
			p.SetState(17)
			p.Unit()
		}
		{
			p.SetState(18)
			p.Match(zipParserT__1)
		}
		{
			p.SetState(19)
			p.Unit()
		}

	}


	return localctx
}


// ITimesContext is an interface to support dynamic dispatch.
type ITimesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimesContext differentiates from other interfaces.
	IsTimesContext()
}

type TimesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimesContext() *TimesContext {
	var p = new(TimesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = zipParserRULE_times
	return p
}

func (*TimesContext) IsTimesContext() {}

func NewTimesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimesContext {
	var p = new(TimesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = zipParserRULE_times

	return p
}

func (s *TimesContext) GetParser() antlr.Parser { return s.parser }

func (s *TimesContext) Digit() antlr.TerminalNode {
	return s.GetToken(zipParserDigit, 0)
}

func (s *TimesContext) Times() ITimesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimesContext)
}

func (s *TimesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TimesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(zipListener); ok {
		listenerT.EnterTimes(s)
	}
}

func (s *TimesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(zipListener); ok {
		listenerT.ExitTimes(s)
	}
}

func (s *TimesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case zipVisitor:
		return t.VisitTimes(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *zipParser) Times() (localctx ITimesContext) {
	localctx = NewTimesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, zipParserRULE_times)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.Match(zipParserDigit)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.Match(zipParserDigit)
		}
		{
			p.SetState(25)
			p.Times()
		}

	}


	return localctx
}


// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContentContext differentiates from other interfaces.
	IsContentContext()
}

type ContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentContext() *ContentContext {
	var p = new(ContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = zipParserRULE_content
	return p
}

func (*ContentContext) IsContentContext() {}

func NewContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentContext {
	var p = new(ContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = zipParserRULE_content

	return p
}

func (s *ContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentContext) AllChar() []antlr.TerminalNode {
	return s.GetTokens(zipParserChar)
}

func (s *ContentContext) Char(i int) antlr.TerminalNode {
	return s.GetToken(zipParserChar, i)
}

func (s *ContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(zipListener); ok {
		listenerT.EnterContent(s)
	}
}

func (s *ContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(zipListener); ok {
		listenerT.ExitContent(s)
	}
}

func (s *ContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case zipVisitor:
		return t.VisitContent(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *zipParser) Content() (localctx IContentContext) {
	localctx = NewContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, zipParserRULE_content)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(28)
					p.Match(zipParserChar)
				}




		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}



	return localctx
}


