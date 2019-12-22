package parser

import (
	"decode/src/ast"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ZipVisitor struct {
	BasezipVisitor
}

func (v *ZipVisitor) Visit(tree antlr.ParseTree) interface{} {
	res := tree.Accept(v)
	return res
}

func (v *ZipVisitor) VisitTimes(ctx *TimesContext) interface{} {
	res, _ := strconv.ParseInt(ctx.GetText(), 10, 8)
	return int(res)
}

func (v *ZipVisitor) VisitContent(ctx *ContentContext) interface{} {
	return ctx.GetText()
}

func (v *ZipVisitor) VisitSingal(ctx *SingalContext) interface{} {
	unit := &ast.Unit{}
	unit.Cur = ctx.GetText()
	unit.Ty = ast.SingalType
	return unit
}

func (v *ZipVisitor) VisitTimesSignal(ctx *TimesSignalContext) interface{} {
	unit := &ast.Unit{}
	unit.Times = ctx.Times().Accept(v).(int)
	unit.Mine = ctx.Unit().Accept(v).(*ast.Unit)
	unit.Ty = ast.TimesType
	return unit
}

func (v *ZipVisitor) VisitSingalfComb(ctx *SingalfCombContext) interface{} {
	unit := &ast.Unit{}
	unit.Cur = ctx.Content().Accept(v).(string)
	unit.Other = ctx.Unit().Accept(v).(*ast.Unit)
	unit.Ty = ast.SingalCombType
	return unit
}

func (v *ZipVisitor) VisitTimesComb(ctx *TimesCombContext) interface{} {
	unit := &ast.Unit{}
	unit.Times = ctx.Times().Accept(v).(int)
	unit.Mine = ctx.Unit(0).Accept(v).(*ast.Unit)
	unit.Other = ctx.Unit(1).Accept(v).(*ast.Unit)
	unit.Ty = ast.TimesCombType
	return unit
}
