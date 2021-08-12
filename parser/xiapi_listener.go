// Code generated from /Users/xiaoqiang.tian/Opensource/Mu/src/main/antlr4/mu/Xiapi.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Xiapi

import "github.com/antlr/antlr4/runtime/Go/antlr"

// XiapiListener is a complete listener for a parse tree produced by XiapiParser.
type XiapiListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterIdAssign is called when entering the idAssign production.
	EnterIdAssign(c *IdAssignContext)

	// EnterArrAssign is called when entering the arrAssign production.
	EnterArrAssign(c *ArrAssignContext)

	// EnterIf_stat is called when entering the if_stat production.
	EnterIf_stat(c *If_statContext)

	// EnterCondition_block is called when entering the condition_block production.
	EnterCondition_block(c *Condition_blockContext)

	// EnterStat_block is called when entering the stat_block production.
	EnterStat_block(c *Stat_blockContext)

	// EnterWhile_stat is called when entering the while_stat production.
	EnterWhile_stat(c *While_statContext)

	// EnterLog is called when entering the log production.
	EnterLog(c *LogContext)

	// EnterNotExpr is called when entering the notExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterUnaryMinusExpr is called when entering the unaryMinusExpr production.
	EnterUnaryMinusExpr(c *UnaryMinusExprContext)

	// EnterLrExpr is called when entering the lrExpr production.
	EnterLrExpr(c *LrExprContext)

	// EnterAtomExpr is called when entering the atomExpr production.
	EnterAtomExpr(c *AtomExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterXiapiltiplicationExpr is called when entering the xiapiltiplicationExpr production.
	EnterXiapiltiplicationExpr(c *XiapiltiplicationExprContext)

	// EnterAdditiveExpr is called when entering the additiveExpr production.
	EnterAdditiveExpr(c *AdditiveExprContext)

	// EnterPowExpr is called when entering the powExpr production.
	EnterPowExpr(c *PowExprContext)

	// EnterRelationalExpr is called when entering the relationalExpr production.
	EnterRelationalExpr(c *RelationalExprContext)

	// EnterEqualityExpr is called when entering the equalityExpr production.
	EnterEqualityExpr(c *EqualityExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterArrExpr is called when entering the arrExpr production.
	EnterArrExpr(c *ArrExprContext)

	// EnterParExpr is called when entering the parExpr production.
	EnterParExpr(c *ParExprContext)

	// EnterNumberAtom is called when entering the numberAtom production.
	EnterNumberAtom(c *NumberAtomContext)

	// EnterBooleanAtom is called when entering the booleanAtom production.
	EnterBooleanAtom(c *BooleanAtomContext)

	// EnterIdAtom is called when entering the idAtom production.
	EnterIdAtom(c *IdAtomContext)

	// EnterStringAtom is called when entering the stringAtom production.
	EnterStringAtom(c *StringAtomContext)

	// EnterNilAtom is called when entering the nilAtom production.
	EnterNilAtom(c *NilAtomContext)

	// EnterLenAtom is called when entering the lenAtom production.
	EnterLenAtom(c *LenAtomContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitIdAssign is called when exiting the idAssign production.
	ExitIdAssign(c *IdAssignContext)

	// ExitArrAssign is called when exiting the arrAssign production.
	ExitArrAssign(c *ArrAssignContext)

	// ExitIf_stat is called when exiting the if_stat production.
	ExitIf_stat(c *If_statContext)

	// ExitCondition_block is called when exiting the condition_block production.
	ExitCondition_block(c *Condition_blockContext)

	// ExitStat_block is called when exiting the stat_block production.
	ExitStat_block(c *Stat_blockContext)

	// ExitWhile_stat is called when exiting the while_stat production.
	ExitWhile_stat(c *While_statContext)

	// ExitLog is called when exiting the log production.
	ExitLog(c *LogContext)

	// ExitNotExpr is called when exiting the notExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitUnaryMinusExpr is called when exiting the unaryMinusExpr production.
	ExitUnaryMinusExpr(c *UnaryMinusExprContext)

	// ExitLrExpr is called when exiting the lrExpr production.
	ExitLrExpr(c *LrExprContext)

	// ExitAtomExpr is called when exiting the atomExpr production.
	ExitAtomExpr(c *AtomExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitXiapiltiplicationExpr is called when exiting the xiapiltiplicationExpr production.
	ExitXiapiltiplicationExpr(c *XiapiltiplicationExprContext)

	// ExitAdditiveExpr is called when exiting the additiveExpr production.
	ExitAdditiveExpr(c *AdditiveExprContext)

	// ExitPowExpr is called when exiting the powExpr production.
	ExitPowExpr(c *PowExprContext)

	// ExitRelationalExpr is called when exiting the relationalExpr production.
	ExitRelationalExpr(c *RelationalExprContext)

	// ExitEqualityExpr is called when exiting the equalityExpr production.
	ExitEqualityExpr(c *EqualityExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitArrExpr is called when exiting the arrExpr production.
	ExitArrExpr(c *ArrExprContext)

	// ExitParExpr is called when exiting the parExpr production.
	ExitParExpr(c *ParExprContext)

	// ExitNumberAtom is called when exiting the numberAtom production.
	ExitNumberAtom(c *NumberAtomContext)

	// ExitBooleanAtom is called when exiting the booleanAtom production.
	ExitBooleanAtom(c *BooleanAtomContext)

	// ExitIdAtom is called when exiting the idAtom production.
	ExitIdAtom(c *IdAtomContext)

	// ExitStringAtom is called when exiting the stringAtom production.
	ExitStringAtom(c *StringAtomContext)

	// ExitNilAtom is called when exiting the nilAtom production.
	ExitNilAtom(c *NilAtomContext)

	// ExitLenAtom is called when exiting the lenAtom production.
	ExitLenAtom(c *LenAtomContext)
}
