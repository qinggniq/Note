// Code generated from /Users/qinggniq/IdeaProjects/Cacl/src/zip.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // zip

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasezipVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasezipVisitor) VisitSingal(ctx *SingalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasezipVisitor) VisitTimesSignal(ctx *TimesSignalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasezipVisitor) VisitSingalfComb(ctx *SingalfCombContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasezipVisitor) VisitTimesComb(ctx *TimesCombContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasezipVisitor) VisitTimes(ctx *TimesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasezipVisitor) VisitContent(ctx *ContentContext) interface{} {
	return v.VisitChildren(ctx)
}
