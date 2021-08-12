package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io"
	"strconv"
	"strings"
	"xiapi-lang/parser"
)

type Value struct {
	Val interface{}
}

func (v *Value) AsString() string {
	return fmt.Sprintf("%v", v.Val)
}

func (v *Value) IsInt() bool {
	_, ok := v.Val.(int)
	return ok
}

func (v *Value) AsInt() int {
	return v.Val.(int)
}

func (v *Value) AsBool() bool {
	return v.Val.(bool)
}

type XiapiLangVisitor struct {
	parser.BaseXiapiVisitor
	Mem     map[string]*Value
	Console io.StringWriter
}

func (v *XiapiLangVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.ParseContext:
		return v.VisitParse(val)
	case *parser.BlockContext:
		return v.VisitBlock(val)
	case *parser.StatContext:
		return v.VisitStat(val)
	case *parser.IdAssignContext:
		return v.VisitIdAssign(val)
	case *parser.ArrAssignContext:
		return v.VisitArrAssign(val)
	case *parser.If_statContext:
		return v.VisitIf_stat(val)
	case *parser.While_statContext:
		return v.VisitWhile_stat(val)
	case *parser.LogContext:
		return v.VisitLog(val)
	case *parser.NotExprContext:
		return v.VisitNotExpr(val)
	case *parser.UnaryMinusExprContext:
		return v.VisitUnaryMinusExpr(val)
	case *parser.LenAtomContext:
		return v.VisitLenAtom(val)
	case *parser.OrExprContext:
		return v.VisitOrExpr(val)
	case *parser.XiapiltiplicationExprContext:
		return v.VisitXiapiltiplicationExpr(val)
	case *parser.AdditiveExprContext:
		return v.VisitAdditiveExpr(val)
	case *parser.RelationalExprContext:
		return v.VisitRelationalExpr(val)
	case *parser.EqualityExprContext:
		return v.VisitEqualityExpr(val)
	case *parser.AndExprContext:
		return v.VisitAndExpr(val)
	case *parser.ArrExprContext:
		return v.VisitArrExpr(val)
	case *parser.ParExprContext:
		return v.VisitParExpr(val)
	case *parser.NumberAtomContext:
		return v.VisitNumberAtom(val)
	case *parser.BooleanAtomContext:
		return v.VisitBooleanAtom(val)
	case *parser.IdAtomContext:
		return v.VisitIdAtom(val)
	case *parser.StringAtomContext:
		return v.VisitStringAtom(val)
	case *parser.NilAtomContext:
		return v.VisitNilAtom(val)
	case *parser.AtomExprContext:
		return v.VisitAtomExpr(val)
	case *parser.Stat_blockContext:
		return v.VisitStat_block(val)
	case *parser.LrExprContext:
		return v.VisitLrExpr(val)
	default:
		panic(fmt.Sprintf("Unknown context %T", tree))
	}
	return &Value{}
}

func (v *XiapiLangVisitor) VisitStat_block(ctx *parser.Stat_blockContext) interface{} {
	if ctx.Block() != nil {
		v.Visit(ctx.Block())
	}
	if ctx.Stat() != nil {
		v.Visit(ctx.Stat())
	}
	return &Value{}
}

func (v *XiapiLangVisitor) VisitLrExpr(ctx *parser.LrExprContext) interface{} {
	return v.Visit(ctx.ArrExpr())
}

func (v *XiapiLangVisitor) VisitParse(ctx *parser.ParseContext) interface{} {
	return v.Visit(ctx.Block())
}

func (v *XiapiLangVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	for _, s := range ctx.AllStat() {
		v.Visit(s)
	}
	return &Value{}
}

func (v *XiapiLangVisitor) VisitStat(ctx *parser.StatContext) interface{} {
	if ctx.Assignment() != nil {
		v.Visit(ctx.Assignment())
	}
	if ctx.If_stat() != nil {
		v.Visit(ctx.If_stat())
	}
	if ctx.While_stat() != nil {
		v.Visit(ctx.While_stat())
	}
	if ctx.Log() != nil {
		v.Visit(ctx.Log())
	}
	if ctx.OTHER() != nil {
		v.Visit(ctx.OTHER())
	}
	return &Value{}
}

func (v *XiapiLangVisitor) VisitIdAssign(ctx *parser.IdAssignContext) interface{} {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr()).(*Value)
	v.Mem[id] = value
	return value
}

func (v *XiapiLangVisitor) VisitArrAssign(ctx *parser.ArrAssignContext) interface{} {
	arrName := ctx.ArrExpr().GetName().GetText()
	value := v.Visit(ctx.Expr()).(*Value)

	arrValue := v.Mem[arrName]

	str := arrValue.AsString()
	idx := v.Visit(ctx.ArrExpr().GetIdx()).(*Value).AsInt()

	replace := value.AsString()

	newValue := &Value{
		Val: str[:idx] + replace + str[idx+1:],
	}
	v.Mem[arrName] = newValue
	return newValue
}

func (v *XiapiLangVisitor) VisitIf_stat(ctx *parser.If_statContext) interface{} {
	blocks := ctx.AllCondition_block()
	evaluatedBlock := false
	for _, b := range blocks {
		evaluated := v.Visit(b.GetExp()).(*Value)
		if evaluated.AsBool() {
			evaluatedBlock = true
			v.Visit(b.GetStatb())
			break
		}
	}
	if !evaluatedBlock && ctx.GetStatb() != nil {
		v.Visit(ctx.GetStatb())
	}
	return &Value{}
}

func (v *XiapiLangVisitor) VisitWhile_stat(ctx *parser.While_statContext) interface{} {
	value := v.Visit(ctx.Expr()).(*Value)
	for value.AsBool() {
		v.Visit(ctx.Stat_block())

		value = v.Visit(ctx.Expr()).(*Value)
	}
	return &Value{}
}

func (v *XiapiLangVisitor) VisitLog(ctx *parser.LogContext) interface{} {
	value := v.Visit(ctx.Expr()).(*Value)
	v.Console.WriteString(fmt.Sprintf("%v\n", value.Val))
	return value
}

func (v *XiapiLangVisitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	value := v.Visit(ctx.Expr()).(*Value)
	return &Value{Val: !value.AsBool()}
}

func (v *XiapiLangVisitor) VisitUnaryMinusExpr(ctx *parser.UnaryMinusExprContext) interface{} {
	value := v.Visit(ctx.Expr()).(*Value)
	return &Value{Val: -value.AsInt()}
}

func (v *XiapiLangVisitor) VisitOrExpr(ctx *parser.OrExprContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(*Value)
	right := v.Visit(ctx.Expr(1)).(*Value)
	return &Value{Val: left.AsBool() || right.AsBool()}
}

func (v *XiapiLangVisitor) VisitXiapiltiplicationExpr(ctx *parser.XiapiltiplicationExprContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(*Value)
	right := v.Visit(ctx.Expr(1)).(*Value)

	switch ctx.GetOp().GetTokenType() {
	case parser.XiapiParserXIAPILT:
		return &Value{left.AsInt() * right.AsInt()}
	case parser.XiapiParserDIV:
		return &Value{left.AsInt() / right.AsInt()}
	case parser.XiapiParserMOD:
		return &Value{left.AsInt() % right.AsInt()}
	default:
		panic("unknown op")
	}
}

func (v *XiapiLangVisitor) VisitAtomExpr(ctx *parser.AtomExprContext) interface{} {
	if ctx.Atom() != nil {
		return v.Visit(ctx.Atom()).(*Value)
	}
	return &Value{}
}

func (v *XiapiLangVisitor) VisitAdditiveExpr(ctx *parser.AdditiveExprContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(*Value)
	right := v.Visit(ctx.Expr(1)).(*Value)

	switch ctx.GetOp().GetTokenType() {
	case parser.XiapiParserPLUS:
		if left.IsInt() && right.IsInt() {
			return &Value{left.AsInt() + right.AsInt()}
		}
		return &Value{left.AsString() + right.AsString()}
	case parser.XiapiParserMINUS:
		return &Value{left.AsInt() - right.AsInt()}
	default:
		panic("unknown op")
	}
}

func (v *XiapiLangVisitor) VisitRelationalExpr(ctx *parser.RelationalExprContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(*Value)
	right := v.Visit(ctx.Expr(1)).(*Value)

	switch ctx.GetOp().GetTokenType() {
	case parser.XiapiLexerLT:
		return &Value{left.AsInt() < right.AsInt()}
	case parser.XiapiParserLTEQ:
		return &Value{left.AsInt() <= right.AsInt()}
	case parser.XiapiParserGT:
		return &Value{left.AsInt() > right.AsInt()}
	case parser.XiapiParserGTEQ:
		return &Value{left.AsInt() >= right.AsInt()}
	default:
		panic("unknown op")
	}
}

func (v *XiapiLangVisitor) VisitEqualityExpr(ctx *parser.EqualityExprContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(*Value)
	right := v.Visit(ctx.Expr(1)).(*Value)

	switch ctx.GetOp().GetTokenType() {
	case parser.XiapiParserEQ:
		return &Value{left.Val == right.Val}
	case parser.XiapiParserNEQ:
		return &Value{left.Val != right.Val}
	default:
		panic("unknown op")
	}
}

func (v *XiapiLangVisitor) VisitAndExpr(ctx *parser.AndExprContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(*Value)
	right := v.Visit(ctx.Expr(1)).(*Value)
	return &Value{Val: left.AsBool() && right.AsBool()}
}

func (v *XiapiLangVisitor) VisitArrExpr(ctx *parser.ArrExprContext) interface{} {
	idx := v.Visit(ctx.GetIdx()).(*Value)
	id := ctx.GetName().GetText()
	arrValue := v.Mem[id]
	if arrValue == nil {
		panic(id + " is nil")
	}
	str := arrValue.AsString()
	return &Value{fmt.Sprintf("%c", str[idx.AsInt()])}
}

func (v *XiapiLangVisitor) VisitParExpr(ctx *parser.ParExprContext) interface{} {
	return v.Visit(ctx.Expr()).(*Value)
}

func (v *XiapiLangVisitor) VisitNumberAtom(ctx *parser.NumberAtomContext) interface{} {
	number, _ := strconv.Atoi(ctx.GetText())
	return &Value{number}
}

func (v *XiapiLangVisitor) VisitBooleanAtom(ctx *parser.BooleanAtomContext) interface{} {
	b, _ := strconv.ParseBool(ctx.GetText())
	return &Value{b}
}

func (v *XiapiLangVisitor) VisitIdAtom(ctx *parser.IdAtomContext) interface{} {
	id := ctx.GetText()
	value := v.Mem[id]
	if value == nil {
		panic(id + " is nil")
	}
	return value
}

func (v *XiapiLangVisitor) VisitStringAtom(ctx *parser.StringAtomContext) interface{} {
	text := ctx.GetText()
	text = strings.TrimSuffix(text, "\"")
	text = strings.TrimPrefix(text, "\"")
	return &Value{text}
}

func (v *XiapiLangVisitor) VisitNilAtom(ctx *parser.NilAtomContext) interface{} {
	return &Value{nil}
}

func (v *XiapiLangVisitor) VisitLenAtom(ctx *parser.LenAtomContext) interface{} {
	value := v.Visit(ctx.Expr()).(*Value)
	return &Value{len(value.AsString())}
}
