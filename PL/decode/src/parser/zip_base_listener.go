// Code generated from /Users/qinggniq/IdeaProjects/Cacl/src/zip.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // zip

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasezipListener is a complete listener for a parse tree produced by zipParser.
type BasezipListener struct{}

var _ zipListener = &BasezipListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasezipListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasezipListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasezipListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasezipListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSingal is called when production singal is entered.
func (s *BasezipListener) EnterSingal(ctx *SingalContext) {}

// ExitSingal is called when production singal is exited.
func (s *BasezipListener) ExitSingal(ctx *SingalContext) {}

// EnterTimesSignal is called when production timesSignal is entered.
func (s *BasezipListener) EnterTimesSignal(ctx *TimesSignalContext) {}

// ExitTimesSignal is called when production timesSignal is exited.
func (s *BasezipListener) ExitTimesSignal(ctx *TimesSignalContext) {}

// EnterSingalfComb is called when production singalfComb is entered.
func (s *BasezipListener) EnterSingalfComb(ctx *SingalfCombContext) {}

// ExitSingalfComb is called when production singalfComb is exited.
func (s *BasezipListener) ExitSingalfComb(ctx *SingalfCombContext) {}

// EnterTimesComb is called when production timesComb is entered.
func (s *BasezipListener) EnterTimesComb(ctx *TimesCombContext) {}

// ExitTimesComb is called when production timesComb is exited.
func (s *BasezipListener) ExitTimesComb(ctx *TimesCombContext) {}

// EnterTimes is called when production times is entered.
func (s *BasezipListener) EnterTimes(ctx *TimesContext) {}

// ExitTimes is called when production times is exited.
func (s *BasezipListener) ExitTimes(ctx *TimesContext) {}

// EnterContent is called when production content is entered.
func (s *BasezipListener) EnterContent(ctx *ContentContext) {}

// ExitContent is called when production content is exited.
func (s *BasezipListener) ExitContent(ctx *ContentContext) {}
