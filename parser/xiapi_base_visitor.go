// Code generated from /Users/xiaoqiang.tian/Opensource/Mu/src/main/antlr4/mu/Xiapi.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Xiapi

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseXiapiVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseXiapiVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitStat(ctx *StatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitIdAssign(ctx *IdAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitArrAssign(ctx *ArrAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitIf_stat(ctx *If_statContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitCondition_block(ctx *Condition_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitStat_block(ctx *Stat_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitWhile_stat(ctx *While_statContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitLog(ctx *LogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitLrExpr(ctx *LrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitAtomExpr(ctx *AtomExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitXiapiltiplicationExpr(ctx *XiapiltiplicationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitAdditiveExpr(ctx *AdditiveExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitPowExpr(ctx *PowExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitRelationalExpr(ctx *RelationalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitEqualityExpr(ctx *EqualityExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitArrExpr(ctx *ArrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitParExpr(ctx *ParExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitNumberAtom(ctx *NumberAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitBooleanAtom(ctx *BooleanAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitIdAtom(ctx *IdAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitStringAtom(ctx *StringAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitNilAtom(ctx *NilAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXiapiVisitor) VisitLenAtom(ctx *LenAtomContext) interface{} {
	return v.VisitChildren(ctx)
}
