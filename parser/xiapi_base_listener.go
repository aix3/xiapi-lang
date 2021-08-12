// Code generated from /Users/xiaoqiang.tian/Opensource/Mu/src/main/antlr4/mu/Xiapi.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Xiapi

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseXiapiListener is a complete listener for a parse tree produced by XiapiParser.
type BaseXiapiListener struct{}

var _ XiapiListener = &BaseXiapiListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseXiapiListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseXiapiListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseXiapiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseXiapiListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseXiapiListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseXiapiListener) ExitParse(ctx *ParseContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseXiapiListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseXiapiListener) ExitBlock(ctx *BlockContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseXiapiListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseXiapiListener) ExitStat(ctx *StatContext) {}

// EnterIdAssign is called when production idAssign is entered.
func (s *BaseXiapiListener) EnterIdAssign(ctx *IdAssignContext) {}

// ExitIdAssign is called when production idAssign is exited.
func (s *BaseXiapiListener) ExitIdAssign(ctx *IdAssignContext) {}

// EnterArrAssign is called when production arrAssign is entered.
func (s *BaseXiapiListener) EnterArrAssign(ctx *ArrAssignContext) {}

// ExitArrAssign is called when production arrAssign is exited.
func (s *BaseXiapiListener) ExitArrAssign(ctx *ArrAssignContext) {}

// EnterIf_stat is called when production if_stat is entered.
func (s *BaseXiapiListener) EnterIf_stat(ctx *If_statContext) {}

// ExitIf_stat is called when production if_stat is exited.
func (s *BaseXiapiListener) ExitIf_stat(ctx *If_statContext) {}

// EnterCondition_block is called when production condition_block is entered.
func (s *BaseXiapiListener) EnterCondition_block(ctx *Condition_blockContext) {}

// ExitCondition_block is called when production condition_block is exited.
func (s *BaseXiapiListener) ExitCondition_block(ctx *Condition_blockContext) {}

// EnterStat_block is called when production stat_block is entered.
func (s *BaseXiapiListener) EnterStat_block(ctx *Stat_blockContext) {}

// ExitStat_block is called when production stat_block is exited.
func (s *BaseXiapiListener) ExitStat_block(ctx *Stat_blockContext) {}

// EnterWhile_stat is called when production while_stat is entered.
func (s *BaseXiapiListener) EnterWhile_stat(ctx *While_statContext) {}

// ExitWhile_stat is called when production while_stat is exited.
func (s *BaseXiapiListener) ExitWhile_stat(ctx *While_statContext) {}

// EnterLog is called when production log is entered.
func (s *BaseXiapiListener) EnterLog(ctx *LogContext) {}

// ExitLog is called when production log is exited.
func (s *BaseXiapiListener) ExitLog(ctx *LogContext) {}

// EnterNotExpr is called when production notExpr is entered.
func (s *BaseXiapiListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production notExpr is exited.
func (s *BaseXiapiListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterUnaryMinusExpr is called when production unaryMinusExpr is entered.
func (s *BaseXiapiListener) EnterUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// ExitUnaryMinusExpr is called when production unaryMinusExpr is exited.
func (s *BaseXiapiListener) ExitUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// EnterLrExpr is called when production lrExpr is entered.
func (s *BaseXiapiListener) EnterLrExpr(ctx *LrExprContext) {}

// ExitLrExpr is called when production lrExpr is exited.
func (s *BaseXiapiListener) ExitLrExpr(ctx *LrExprContext) {}

// EnterAtomExpr is called when production atomExpr is entered.
func (s *BaseXiapiListener) EnterAtomExpr(ctx *AtomExprContext) {}

// ExitAtomExpr is called when production atomExpr is exited.
func (s *BaseXiapiListener) ExitAtomExpr(ctx *AtomExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseXiapiListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseXiapiListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterXiapiltiplicationExpr is called when production xiapiltiplicationExpr is entered.
func (s *BaseXiapiListener) EnterXiapiltiplicationExpr(ctx *XiapiltiplicationExprContext) {}

// ExitXiapiltiplicationExpr is called when production xiapiltiplicationExpr is exited.
func (s *BaseXiapiListener) ExitXiapiltiplicationExpr(ctx *XiapiltiplicationExprContext) {}

// EnterAdditiveExpr is called when production additiveExpr is entered.
func (s *BaseXiapiListener) EnterAdditiveExpr(ctx *AdditiveExprContext) {}

// ExitAdditiveExpr is called when production additiveExpr is exited.
func (s *BaseXiapiListener) ExitAdditiveExpr(ctx *AdditiveExprContext) {}

// EnterPowExpr is called when production powExpr is entered.
func (s *BaseXiapiListener) EnterPowExpr(ctx *PowExprContext) {}

// ExitPowExpr is called when production powExpr is exited.
func (s *BaseXiapiListener) ExitPowExpr(ctx *PowExprContext) {}

// EnterRelationalExpr is called when production relationalExpr is entered.
func (s *BaseXiapiListener) EnterRelationalExpr(ctx *RelationalExprContext) {}

// ExitRelationalExpr is called when production relationalExpr is exited.
func (s *BaseXiapiListener) ExitRelationalExpr(ctx *RelationalExprContext) {}

// EnterEqualityExpr is called when production equalityExpr is entered.
func (s *BaseXiapiListener) EnterEqualityExpr(ctx *EqualityExprContext) {}

// ExitEqualityExpr is called when production equalityExpr is exited.
func (s *BaseXiapiListener) ExitEqualityExpr(ctx *EqualityExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseXiapiListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseXiapiListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterArrExpr is called when production arrExpr is entered.
func (s *BaseXiapiListener) EnterArrExpr(ctx *ArrExprContext) {}

// ExitArrExpr is called when production arrExpr is exited.
func (s *BaseXiapiListener) ExitArrExpr(ctx *ArrExprContext) {}

// EnterParExpr is called when production parExpr is entered.
func (s *BaseXiapiListener) EnterParExpr(ctx *ParExprContext) {}

// ExitParExpr is called when production parExpr is exited.
func (s *BaseXiapiListener) ExitParExpr(ctx *ParExprContext) {}

// EnterNumberAtom is called when production numberAtom is entered.
func (s *BaseXiapiListener) EnterNumberAtom(ctx *NumberAtomContext) {}

// ExitNumberAtom is called when production numberAtom is exited.
func (s *BaseXiapiListener) ExitNumberAtom(ctx *NumberAtomContext) {}

// EnterBooleanAtom is called when production booleanAtom is entered.
func (s *BaseXiapiListener) EnterBooleanAtom(ctx *BooleanAtomContext) {}

// ExitBooleanAtom is called when production booleanAtom is exited.
func (s *BaseXiapiListener) ExitBooleanAtom(ctx *BooleanAtomContext) {}

// EnterIdAtom is called when production idAtom is entered.
func (s *BaseXiapiListener) EnterIdAtom(ctx *IdAtomContext) {}

// ExitIdAtom is called when production idAtom is exited.
func (s *BaseXiapiListener) ExitIdAtom(ctx *IdAtomContext) {}

// EnterStringAtom is called when production stringAtom is entered.
func (s *BaseXiapiListener) EnterStringAtom(ctx *StringAtomContext) {}

// ExitStringAtom is called when production stringAtom is exited.
func (s *BaseXiapiListener) ExitStringAtom(ctx *StringAtomContext) {}

// EnterNilAtom is called when production nilAtom is entered.
func (s *BaseXiapiListener) EnterNilAtom(ctx *NilAtomContext) {}

// ExitNilAtom is called when production nilAtom is exited.
func (s *BaseXiapiListener) ExitNilAtom(ctx *NilAtomContext) {}

// EnterLenAtom is called when production lenAtom is entered.
func (s *BaseXiapiListener) EnterLenAtom(ctx *LenAtomContext) {}

// ExitLenAtom is called when production lenAtom is exited.
func (s *BaseXiapiListener) ExitLenAtom(ctx *LenAtomContext) {}
