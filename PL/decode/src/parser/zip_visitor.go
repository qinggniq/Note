// Code generated from /Users/qinggniq/IdeaProjects/Cacl/src/zip.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // zip

import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by zipParser.
type zipVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by zipParser#singal.
	VisitSingal(ctx *SingalContext) interface{}

	// Visit a parse tree produced by zipParser#timesSignal.
	VisitTimesSignal(ctx *TimesSignalContext) interface{}

	// Visit a parse tree produced by zipParser#singalfComb.
	VisitSingalfComb(ctx *SingalfCombContext) interface{}

	// Visit a parse tree produced by zipParser#timesComb.
	VisitTimesComb(ctx *TimesCombContext) interface{}

	// Visit a parse tree produced by zipParser#times.
	VisitTimes(ctx *TimesContext) interface{}

	// Visit a parse tree produced by zipParser#content.
	VisitContent(ctx *ContentContext) interface{}

}