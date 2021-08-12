// Code generated from /Users/xiaoqiang.tian/Opensource/Mu/src/main/antlr4/mu/Xiapi.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Xiapi

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by XiapiParser.
type XiapiVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by XiapiParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by XiapiParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by XiapiParser#stat.
	VisitStat(ctx *StatContext) interface{}

	// Visit a parse tree produced by XiapiParser#idAssign.
	VisitIdAssign(ctx *IdAssignContext) interface{}

	// Visit a parse tree produced by XiapiParser#arrAssign.
	VisitArrAssign(ctx *ArrAssignContext) interface{}

	// Visit a parse tree produced by XiapiParser#if_stat.
	VisitIf_stat(ctx *If_statContext) interface{}

	// Visit a parse tree produced by XiapiParser#condition_block.
	VisitCondition_block(ctx *Condition_blockContext) interface{}

	// Visit a parse tree produced by XiapiParser#stat_block.
	VisitStat_block(ctx *Stat_blockContext) interface{}

	// Visit a parse tree produced by XiapiParser#while_stat.
	VisitWhile_stat(ctx *While_statContext) interface{}

	// Visit a parse tree produced by XiapiParser#log.
	VisitLog(ctx *LogContext) interface{}

	// Visit a parse tree produced by XiapiParser#notExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by XiapiParser#unaryMinusExpr.
	VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{}

	// Visit a parse tree produced by XiapiParser#lrExpr.
	VisitLrExpr(ctx *LrExprContext) interface{}

	// Visit a parse tree produced by XiapiParser#atomExpr.
	VisitAtomExpr(ctx *AtomExprContext) interface{}

	// Visit a parse tree produced by XiapiParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by XiapiParser#xiapiltiplicationExpr.
	VisitXiapiltiplicationExpr(ctx *XiapiltiplicationExprContext) interface{}

	// Visit a parse tree produced by XiapiParser#additiveExpr.
	VisitAdditiveExpr(ctx *AdditiveExprContext) interface{}

	// Visit a parse tree produced by XiapiParser#powExpr.
	VisitPowExpr(ctx *PowExprContext) interface{}

	// Visit a parse tree produced by XiapiParser#relationalExpr.
	VisitRelationalExpr(ctx *RelationalExprContext) interface{}

	// Visit a parse tree produced by XiapiParser#equalityExpr.
	VisitEqualityExpr(ctx *EqualityExprContext) interface{}

	// Visit a parse tree produced by XiapiParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by XiapiParser#arrExpr.
	VisitArrExpr(ctx *ArrExprContext) interface{}

	// Visit a parse tree produced by XiapiParser#parExpr.
	VisitParExpr(ctx *ParExprContext) interface{}

	// Visit a parse tree produced by XiapiParser#numberAtom.
	VisitNumberAtom(ctx *NumberAtomContext) interface{}

	// Visit a parse tree produced by XiapiParser#booleanAtom.
	VisitBooleanAtom(ctx *BooleanAtomContext) interface{}

	// Visit a parse tree produced by XiapiParser#idAtom.
	VisitIdAtom(ctx *IdAtomContext) interface{}

	// Visit a parse tree produced by XiapiParser#stringAtom.
	VisitStringAtom(ctx *StringAtomContext) interface{}

	// Visit a parse tree produced by XiapiParser#nilAtom.
	VisitNilAtom(ctx *NilAtomContext) interface{}

	// Visit a parse tree produced by XiapiParser#lenAtom.
	VisitLenAtom(ctx *LenAtomContext) interface{}
}
