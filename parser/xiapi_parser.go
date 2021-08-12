// Code generated from /Users/xiaoqiang.tian/Opensource/Mu/src/main/antlr4/mu/Xiapi.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Xiapi

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 39, 137,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 31, 10, 3, 12, 3, 14, 3, 34, 11, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 42, 10, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 51, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	7, 6, 58, 10, 6, 12, 6, 14, 6, 61, 11, 6, 3, 6, 3, 6, 5, 6, 65, 10, 6,
	3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 75, 10, 8, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 5, 11, 91, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 114, 10, 11, 12, 11, 14, 11,
	117, 11, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 135, 10, 13,
	3, 13, 2, 3, 20, 14, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 2, 8,
	3, 2, 15, 17, 3, 2, 13, 14, 3, 2, 9, 12, 3, 2, 7, 8, 3, 2, 34, 35, 3, 2,
	25, 26, 2, 149, 2, 26, 3, 2, 2, 2, 4, 32, 3, 2, 2, 2, 6, 41, 3, 2, 2, 2,
	8, 50, 3, 2, 2, 2, 10, 52, 3, 2, 2, 2, 12, 66, 3, 2, 2, 2, 14, 74, 3, 2,
	2, 2, 16, 76, 3, 2, 2, 2, 18, 80, 3, 2, 2, 2, 20, 90, 3, 2, 2, 2, 22, 118,
	3, 2, 2, 2, 24, 134, 3, 2, 2, 2, 26, 27, 5, 4, 3, 2, 27, 28, 7, 2, 2, 3,
	28, 3, 3, 2, 2, 2, 29, 31, 5, 6, 4, 2, 30, 29, 3, 2, 2, 2, 31, 34, 3, 2,
	2, 2, 32, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 5, 3, 2, 2, 2, 34, 32,
	3, 2, 2, 2, 35, 42, 5, 8, 5, 2, 36, 42, 5, 10, 6, 2, 37, 42, 5, 16, 9,
	2, 38, 42, 5, 18, 10, 2, 39, 40, 7, 39, 2, 2, 40, 42, 8, 4, 1, 2, 41, 35,
	3, 2, 2, 2, 41, 36, 3, 2, 2, 2, 41, 37, 3, 2, 2, 2, 41, 38, 3, 2, 2, 2,
	41, 39, 3, 2, 2, 2, 42, 7, 3, 2, 2, 2, 43, 44, 7, 33, 2, 2, 44, 45, 7,
	20, 2, 2, 45, 51, 5, 20, 11, 2, 46, 47, 5, 22, 12, 2, 47, 48, 7, 20, 2,
	2, 48, 49, 5, 20, 11, 2, 49, 51, 3, 2, 2, 2, 50, 43, 3, 2, 2, 2, 50, 46,
	3, 2, 2, 2, 51, 9, 3, 2, 2, 2, 52, 53, 7, 28, 2, 2, 53, 59, 5, 12, 7, 2,
	54, 55, 7, 29, 2, 2, 55, 56, 7, 28, 2, 2, 56, 58, 5, 12, 7, 2, 57, 54,
	3, 2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2,
	60, 64, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 63, 7, 29, 2, 2, 63, 65, 5,
	14, 8, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 11, 3, 2, 2, 2, 66,
	67, 5, 20, 11, 2, 67, 68, 5, 14, 8, 2, 68, 13, 3, 2, 2, 2, 69, 70, 7, 23,
	2, 2, 70, 71, 5, 4, 3, 2, 71, 72, 7, 24, 2, 2, 72, 75, 3, 2, 2, 2, 73,
	75, 5, 6, 4, 2, 74, 69, 3, 2, 2, 2, 74, 73, 3, 2, 2, 2, 75, 15, 3, 2, 2,
	2, 76, 77, 7, 30, 2, 2, 77, 78, 5, 20, 11, 2, 78, 79, 5, 14, 8, 2, 79,
	17, 3, 2, 2, 2, 80, 81, 7, 31, 2, 2, 81, 82, 5, 20, 11, 2, 82, 19, 3, 2,
	2, 2, 83, 84, 8, 11, 1, 2, 84, 85, 7, 14, 2, 2, 85, 91, 5, 20, 11, 12,
	86, 87, 7, 19, 2, 2, 87, 91, 5, 20, 11, 11, 88, 91, 5, 22, 12, 2, 89, 91,
	5, 24, 13, 2, 90, 83, 3, 2, 2, 2, 90, 86, 3, 2, 2, 2, 90, 88, 3, 2, 2,
	2, 90, 89, 3, 2, 2, 2, 91, 115, 3, 2, 2, 2, 92, 93, 12, 13, 2, 2, 93, 94,
	7, 18, 2, 2, 94, 114, 5, 20, 11, 14, 95, 96, 12, 10, 2, 2, 96, 97, 9, 2,
	2, 2, 97, 114, 5, 20, 11, 11, 98, 99, 12, 9, 2, 2, 99, 100, 9, 3, 2, 2,
	100, 114, 5, 20, 11, 10, 101, 102, 12, 8, 2, 2, 102, 103, 9, 4, 2, 2, 103,
	114, 5, 20, 11, 9, 104, 105, 12, 7, 2, 2, 105, 106, 9, 5, 2, 2, 106, 114,
	5, 20, 11, 8, 107, 108, 12, 6, 2, 2, 108, 109, 7, 6, 2, 2, 109, 114, 5,
	20, 11, 7, 110, 111, 12, 5, 2, 2, 111, 112, 7, 5, 2, 2, 112, 114, 5, 20,
	11, 6, 113, 92, 3, 2, 2, 2, 113, 95, 3, 2, 2, 2, 113, 98, 3, 2, 2, 2, 113,
	101, 3, 2, 2, 2, 113, 104, 3, 2, 2, 2, 113, 107, 3, 2, 2, 2, 113, 110,
	3, 2, 2, 2, 114, 117, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 116, 3, 2,
	2, 2, 116, 21, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 118, 119, 7, 33, 2, 2,
	119, 120, 7, 3, 2, 2, 120, 121, 5, 20, 11, 2, 121, 122, 7, 4, 2, 2, 122,
	23, 3, 2, 2, 2, 123, 124, 7, 21, 2, 2, 124, 125, 5, 20, 11, 2, 125, 126,
	7, 22, 2, 2, 126, 135, 3, 2, 2, 2, 127, 135, 9, 6, 2, 2, 128, 135, 9, 7,
	2, 2, 129, 135, 7, 33, 2, 2, 130, 135, 7, 36, 2, 2, 131, 135, 7, 27, 2,
	2, 132, 133, 7, 32, 2, 2, 133, 135, 5, 20, 11, 2, 134, 123, 3, 2, 2, 2,
	134, 127, 3, 2, 2, 2, 134, 128, 3, 2, 2, 2, 134, 129, 3, 2, 2, 2, 134,
	130, 3, 2, 2, 2, 134, 131, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 135, 25, 3,
	2, 2, 2, 12, 32, 41, 50, 59, 64, 74, 90, 113, 115, 134,
}
var literalNames = []string{
	"", "'['", "']'", "'||'", "'&&'", "'=='", "'!='", "'\u5168\u5FC3\u6295\u5165'",
	"'<'", "'>='", "'<='", "'\u6536\u996D'", "'\u51FA\u996D'", "'*'", "'/'",
	"'\u867E\u7C73'", "'^'", "'!'", "'\u987A\u52BF\u5E94\u53D8'", "'('", "')'",
	"'{'", "'}'", "'true'", "'false'", "'nil'", "'\u9886\u5934\u867E'", "'\u4FDD\u6301\u8C26\u900A'",
	"'\u5206\u79D2\u5FC5\u4E89'", "'\u6D77\u804A'", "'\u7528\u6237\u81F3\u4E0A'",
}
var symbolicNames = []string{
	"", "L", "R", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS",
	"MINUS", "XIAPILT", "DIV", "MOD", "POW", "NOT", "ASSIGN", "OPAR", "CPAR",
	"OBRACE", "CBRACE", "TRUE", "FALSE", "NIL", "IF", "ELSE", "WHILE", "LOG",
	"LEN", "ID", "INT", "FLOAT", "STRING", "COMMENT", "SPACE", "OTHER",
}

var ruleNames = []string{
	"parse", "block", "stat", "assignment", "if_stat", "condition_block", "stat_block",
	"while_stat", "log", "expr", "arrExpr", "atom",
}

type XiapiParser struct {
	*antlr.BaseParser
}

// NewXiapiParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *XiapiParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewXiapiParser(input antlr.TokenStream) *XiapiParser {
	this := new(XiapiParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Xiapi.g4"

	return this
}

// XiapiParser tokens.
const (
	XiapiParserEOF     = antlr.TokenEOF
	XiapiParserL       = 1
	XiapiParserR       = 2
	XiapiParserOR      = 3
	XiapiParserAND     = 4
	XiapiParserEQ      = 5
	XiapiParserNEQ     = 6
	XiapiParserGT      = 7
	XiapiParserLT      = 8
	XiapiParserGTEQ    = 9
	XiapiParserLTEQ    = 10
	XiapiParserPLUS    = 11
	XiapiParserMINUS   = 12
	XiapiParserXIAPILT = 13
	XiapiParserDIV     = 14
	XiapiParserMOD     = 15
	XiapiParserPOW     = 16
	XiapiParserNOT     = 17
	XiapiParserASSIGN  = 18
	XiapiParserOPAR    = 19
	XiapiParserCPAR    = 20
	XiapiParserOBRACE  = 21
	XiapiParserCBRACE  = 22
	XiapiParserTRUE    = 23
	XiapiParserFALSE   = 24
	XiapiParserNIL     = 25
	XiapiParserIF      = 26
	XiapiParserELSE    = 27
	XiapiParserWHILE   = 28
	XiapiParserLOG     = 29
	XiapiParserLEN     = 30
	XiapiParserID      = 31
	XiapiParserINT     = 32
	XiapiParserFLOAT   = 33
	XiapiParserSTRING  = 34
	XiapiParserCOMMENT = 35
	XiapiParserSPACE   = 36
	XiapiParserOTHER   = 37
)

// XiapiParser rules.
const (
	XiapiParserRULE_parse           = 0
	XiapiParserRULE_block           = 1
	XiapiParserRULE_stat            = 2
	XiapiParserRULE_assignment      = 3
	XiapiParserRULE_if_stat         = 4
	XiapiParserRULE_condition_block = 5
	XiapiParserRULE_stat_block      = 6
	XiapiParserRULE_while_stat      = 7
	XiapiParserRULE_log             = 8
	XiapiParserRULE_expr            = 9
	XiapiParserRULE_arrExpr         = 10
	XiapiParserRULE_atom            = 11
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XiapiParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XiapiParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(XiapiParserEOF, 0)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitParse(s)
	}
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XiapiParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, XiapiParserRULE_parse)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Block()
	}
	{
		p.SetState(25)
		p.Match(XiapiParserEOF)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XiapiParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XiapiParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *BlockContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XiapiParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, XiapiParserRULE_block)
	var _la int

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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(XiapiParserIF-26))|(1<<(XiapiParserWHILE-26))|(1<<(XiapiParserLOG-26))|(1<<(XiapiParserID-26))|(1<<(XiapiParserOTHER-26)))) != 0 {
		{
			p.SetState(27)
			p.Stat()
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XiapiParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XiapiParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatContext) If_stat() IIf_statContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_statContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_statContext)
}

func (s *StatContext) While_stat() IWhile_statContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhile_statContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhile_statContext)
}

func (s *StatContext) Log() ILogContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogContext)
}

func (s *StatContext) OTHER() antlr.TerminalNode {
	return s.GetToken(XiapiParserOTHER, 0)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitStat(s)
	}
}

func (s *StatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XiapiParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, XiapiParserRULE_stat)

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

	p.SetState(39)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case XiapiParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Assignment()
		}

	case XiapiParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.If_stat()
		}

	case XiapiParserWHILE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(35)
			p.While_stat()
		}

	case XiapiParserLOG:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(36)
			p.Log()
		}

	case XiapiParserOTHER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(37)
			p.Match(XiapiParserOTHER)
		}
		panic("unkonwn")

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XiapiParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XiapiParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) CopyFrom(ctx *AssignmentContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdAssignContext struct {
	*AssignmentContext
}

func NewIdAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdAssignContext {
	var p = new(IdAssignContext)

	p.AssignmentContext = NewEmptyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignmentContext))

	return p
}

func (s *IdAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdAssignContext) ID() antlr.TerminalNode {
	return s.GetToken(XiapiParserID, 0)
}

func (s *IdAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(XiapiParserASSIGN, 0)
}

func (s *IdAssignContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IdAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterIdAssign(s)
	}
}

func (s *IdAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitIdAssign(s)
	}
}

func (s *IdAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitIdAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrAssignContext struct {
	*AssignmentContext
	arr IArrExprContext
}

func NewArrAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrAssignContext {
	var p = new(ArrAssignContext)

	p.AssignmentContext = NewEmptyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignmentContext))

	return p
}

func (s *ArrAssignContext) GetArr() IArrExprContext { return s.arr }

func (s *ArrAssignContext) SetArr(v IArrExprContext) { s.arr = v }

func (s *ArrAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(XiapiParserASSIGN, 0)
}

func (s *ArrAssignContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrAssignContext) ArrExpr() IArrExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrExprContext)
}

func (s *ArrAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterArrAssign(s)
	}
}

func (s *ArrAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitArrAssign(s)
	}
}

func (s *ArrAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitArrAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XiapiParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, XiapiParserRULE_assignment)

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

	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIdAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Match(XiapiParserID)
		}
		{
			p.SetState(42)
			p.Match(XiapiParserASSIGN)
		}
		{
			p.SetState(43)
			p.expr(0)
		}

	case 2:
		localctx = NewArrAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)

			var _x = p.ArrExpr()

			localctx.(*ArrAssignContext).arr = _x
		}
		{
			p.SetState(45)
			p.Match(XiapiParserASSIGN)
		}
		{
			p.SetState(46)
			p.expr(0)
		}

	}

	return localctx
}

// IIf_statContext is an interface to support dynamic dispatch.
type IIf_statContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStatb returns the statb rule contexts.
	GetStatb() IStat_blockContext

	// SetStatb sets the statb rule contexts.
	SetStatb(IStat_blockContext)

	// IsIf_statContext differentiates from other interfaces.
	IsIf_statContext()
}

type If_statContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	statb  IStat_blockContext
}

func NewEmptyIf_statContext() *If_statContext {
	var p = new(If_statContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XiapiParserRULE_if_stat
	return p
}

func (*If_statContext) IsIf_statContext() {}

func NewIf_statContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statContext {
	var p = new(If_statContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XiapiParserRULE_if_stat

	return p
}

func (s *If_statContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statContext) GetStatb() IStat_blockContext { return s.statb }

func (s *If_statContext) SetStatb(v IStat_blockContext) { s.statb = v }

func (s *If_statContext) AllIF() []antlr.TerminalNode {
	return s.GetTokens(XiapiParserIF)
}

func (s *If_statContext) IF(i int) antlr.TerminalNode {
	return s.GetToken(XiapiParserIF, i)
}

func (s *If_statContext) AllCondition_block() []ICondition_blockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICondition_blockContext)(nil)).Elem())
	var tst = make([]ICondition_blockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICondition_blockContext)
		}
	}

	return tst
}

func (s *If_statContext) Condition_block(i int) ICondition_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondition_blockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICondition_blockContext)
}

func (s *If_statContext) AllELSE() []antlr.TerminalNode {
	return s.GetTokens(XiapiParserELSE)
}

func (s *If_statContext) ELSE(i int) antlr.TerminalNode {
	return s.GetToken(XiapiParserELSE, i)
}

func (s *If_statContext) Stat_block() IStat_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStat_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStat_blockContext)
}

func (s *If_statContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_statContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterIf_stat(s)
	}
}

func (s *If_statContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitIf_stat(s)
	}
}

func (s *If_statContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitIf_stat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XiapiParser) If_stat() (localctx IIf_statContext) {
	localctx = NewIf_statContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, XiapiParserRULE_if_stat)

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
	{
		p.SetState(50)
		p.Match(XiapiParserIF)
	}
	{
		p.SetState(51)
		p.Condition_block()
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(52)
				p.Match(XiapiParserELSE)
			}
			{
				p.SetState(53)
				p.Match(XiapiParserIF)
			}
			{
				p.SetState(54)
				p.Condition_block()
			}

		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(60)
			p.Match(XiapiParserELSE)
		}
		{
			p.SetState(61)

			var _x = p.Stat_block()

			localctx.(*If_statContext).statb = _x
		}

	}

	return localctx
}

// ICondition_blockContext is an interface to support dynamic dispatch.
type ICondition_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExp returns the exp rule contexts.
	GetExp() IExprContext

	// GetStatb returns the statb rule contexts.
	GetStatb() IStat_blockContext

	// SetExp sets the exp rule contexts.
	SetExp(IExprContext)

	// SetStatb sets the statb rule contexts.
	SetStatb(IStat_blockContext)

	// IsCondition_blockContext differentiates from other interfaces.
	IsCondition_blockContext()
}

type Condition_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	exp    IExprContext
	statb  IStat_blockContext
}

func NewEmptyCondition_blockContext() *Condition_blockContext {
	var p = new(Condition_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XiapiParserRULE_condition_block
	return p
}

func (*Condition_blockContext) IsCondition_blockContext() {}

func NewCondition_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Condition_blockContext {
	var p = new(Condition_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XiapiParserRULE_condition_block

	return p
}

func (s *Condition_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Condition_blockContext) GetExp() IExprContext { return s.exp }

func (s *Condition_blockContext) GetStatb() IStat_blockContext { return s.statb }

func (s *Condition_blockContext) SetExp(v IExprContext) { s.exp = v }

func (s *Condition_blockContext) SetStatb(v IStat_blockContext) { s.statb = v }

func (s *Condition_blockContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Condition_blockContext) Stat_block() IStat_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStat_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStat_blockContext)
}

func (s *Condition_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Condition_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Condition_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterCondition_block(s)
	}
}

func (s *Condition_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitCondition_block(s)
	}
}

func (s *Condition_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitCondition_block(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XiapiParser) Condition_block() (localctx ICondition_blockContext) {
	localctx = NewCondition_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, XiapiParserRULE_condition_block)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)

		var _x = p.expr(0)

		localctx.(*Condition_blockContext).exp = _x
	}
	{
		p.SetState(65)

		var _x = p.Stat_block()

		localctx.(*Condition_blockContext).statb = _x
	}

	return localctx
}

// IStat_blockContext is an interface to support dynamic dispatch.
type IStat_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStat_blockContext differentiates from other interfaces.
	IsStat_blockContext()
}

type Stat_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStat_blockContext() *Stat_blockContext {
	var p = new(Stat_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XiapiParserRULE_stat_block
	return p
}

func (*Stat_blockContext) IsStat_blockContext() {}

func NewStat_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stat_blockContext {
	var p = new(Stat_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XiapiParserRULE_stat_block

	return p
}

func (s *Stat_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Stat_blockContext) OBRACE() antlr.TerminalNode {
	return s.GetToken(XiapiParserOBRACE, 0)
}

func (s *Stat_blockContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Stat_blockContext) CBRACE() antlr.TerminalNode {
	return s.GetToken(XiapiParserCBRACE, 0)
}

func (s *Stat_blockContext) Stat() IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *Stat_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stat_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Stat_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterStat_block(s)
	}
}

func (s *Stat_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitStat_block(s)
	}
}

func (s *Stat_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitStat_block(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XiapiParser) Stat_block() (localctx IStat_blockContext) {
	localctx = NewStat_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, XiapiParserRULE_stat_block)

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

	p.SetState(72)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case XiapiParserOBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.Match(XiapiParserOBRACE)
		}
		{
			p.SetState(68)
			p.Block()
		}
		{
			p.SetState(69)
			p.Match(XiapiParserCBRACE)
		}

	case XiapiParserIF, XiapiParserWHILE, XiapiParserLOG, XiapiParserID, XiapiParserOTHER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Stat()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IWhile_statContext is an interface to support dynamic dispatch.
type IWhile_statContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhile_statContext differentiates from other interfaces.
	IsWhile_statContext()
}

type While_statContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_statContext() *While_statContext {
	var p = new(While_statContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XiapiParserRULE_while_stat
	return p
}

func (*While_statContext) IsWhile_statContext() {}

func NewWhile_statContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_statContext {
	var p = new(While_statContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XiapiParserRULE_while_stat

	return p
}

func (s *While_statContext) GetParser() antlr.Parser { return s.parser }

func (s *While_statContext) WHILE() antlr.TerminalNode {
	return s.GetToken(XiapiParserWHILE, 0)
}

func (s *While_statContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *While_statContext) Stat_block() IStat_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStat_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStat_blockContext)
}

func (s *While_statContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_statContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_statContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterWhile_stat(s)
	}
}

func (s *While_statContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitWhile_stat(s)
	}
}

func (s *While_statContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitWhile_stat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XiapiParser) While_stat() (localctx IWhile_statContext) {
	localctx = NewWhile_statContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, XiapiParserRULE_while_stat)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(XiapiParserWHILE)
	}
	{
		p.SetState(75)
		p.expr(0)
	}
	{
		p.SetState(76)
		p.Stat_block()
	}

	return localctx
}

// ILogContext is an interface to support dynamic dispatch.
type ILogContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogContext differentiates from other interfaces.
	IsLogContext()
}

type LogContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogContext() *LogContext {
	var p = new(LogContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XiapiParserRULE_log
	return p
}

func (*LogContext) IsLogContext() {}

func NewLogContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogContext {
	var p = new(LogContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XiapiParserRULE_log

	return p
}

func (s *LogContext) GetParser() antlr.Parser { return s.parser }

func (s *LogContext) LOG() antlr.TerminalNode {
	return s.GetToken(XiapiParserLOG, 0)
}

func (s *LogContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LogContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterLog(s)
	}
}

func (s *LogContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitLog(s)
	}
}

func (s *LogContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitLog(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XiapiParser) Log() (localctx ILogContext) {
	localctx = NewLogContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, XiapiParserRULE_log)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(XiapiParserLOG)
	}
	{
		p.SetState(79)
		p.expr(0)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XiapiParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XiapiParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NotExprContext struct {
	*ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(XiapiParserNOT, 0)
}

func (s *NotExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryMinusExprContext struct {
	*ExprContext
}

func NewUnaryMinusExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExprContext {
	var p = new(UnaryMinusExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *UnaryMinusExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(XiapiParserMINUS, 0)
}

func (s *UnaryMinusExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryMinusExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterUnaryMinusExpr(s)
	}
}

func (s *UnaryMinusExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitUnaryMinusExpr(s)
	}
}

func (s *UnaryMinusExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitUnaryMinusExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LrExprContext struct {
	*ExprContext
}

func NewLrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LrExprContext {
	var p = new(LrExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LrExprContext) ArrExpr() IArrExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrExprContext)
}

func (s *LrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterLrExpr(s)
	}
}

func (s *LrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitLrExpr(s)
	}
}

func (s *LrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitLrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtomExprContext struct {
	*ExprContext
}

func NewAtomExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomExprContext {
	var p = new(AtomExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AtomExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomExprContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *AtomExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterAtomExpr(s)
	}
}

func (s *AtomExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitAtomExpr(s)
	}
}

func (s *AtomExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitAtomExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	*ExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OrExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(XiapiParserOR, 0)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type XiapiltiplicationExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewXiapiltiplicationExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *XiapiltiplicationExprContext {
	var p = new(XiapiltiplicationExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *XiapiltiplicationExprContext) GetOp() antlr.Token { return s.op }

func (s *XiapiltiplicationExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *XiapiltiplicationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XiapiltiplicationExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *XiapiltiplicationExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *XiapiltiplicationExprContext) XIAPILT() antlr.TerminalNode {
	return s.GetToken(XiapiParserXIAPILT, 0)
}

func (s *XiapiltiplicationExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(XiapiParserDIV, 0)
}

func (s *XiapiltiplicationExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(XiapiParserMOD, 0)
}

func (s *XiapiltiplicationExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterXiapiltiplicationExpr(s)
	}
}

func (s *XiapiltiplicationExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitXiapiltiplicationExpr(s)
	}
}

func (s *XiapiltiplicationExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitXiapiltiplicationExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AdditiveExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewAdditiveExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExprContext {
	var p = new(AdditiveExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AdditiveExprContext) GetOp() antlr.Token { return s.op }

func (s *AdditiveExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AdditiveExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AdditiveExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AdditiveExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(XiapiParserPLUS, 0)
}

func (s *AdditiveExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(XiapiParserMINUS, 0)
}

func (s *AdditiveExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitAdditiveExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PowExprContext struct {
	*ExprContext
}

func NewPowExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowExprContext {
	var p = new(PowExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PowExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PowExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PowExprContext) POW() antlr.TerminalNode {
	return s.GetToken(XiapiParserPOW, 0)
}

func (s *PowExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterPowExpr(s)
	}
}

func (s *PowExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitPowExpr(s)
	}
}

func (s *PowExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitPowExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelationalExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewRelationalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalExprContext {
	var p = new(RelationalExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RelationalExprContext) GetOp() antlr.Token { return s.op }

func (s *RelationalExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RelationalExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RelationalExprContext) LTEQ() antlr.TerminalNode {
	return s.GetToken(XiapiParserLTEQ, 0)
}

func (s *RelationalExprContext) GTEQ() antlr.TerminalNode {
	return s.GetToken(XiapiParserGTEQ, 0)
}

func (s *RelationalExprContext) LT() antlr.TerminalNode {
	return s.GetToken(XiapiParserLT, 0)
}

func (s *RelationalExprContext) GT() antlr.TerminalNode {
	return s.GetToken(XiapiParserGT, 0)
}

func (s *RelationalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterRelationalExpr(s)
	}
}

func (s *RelationalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitRelationalExpr(s)
	}
}

func (s *RelationalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitRelationalExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualityExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewEqualityExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExprContext {
	var p = new(EqualityExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EqualityExprContext) GetOp() antlr.Token { return s.op }

func (s *EqualityExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualityExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EqualityExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqualityExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(XiapiParserEQ, 0)
}

func (s *EqualityExprContext) NEQ() antlr.TerminalNode {
	return s.GetToken(XiapiParserNEQ, 0)
}

func (s *EqualityExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterEqualityExpr(s)
	}
}

func (s *EqualityExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitEqualityExpr(s)
	}
}

func (s *EqualityExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitEqualityExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExprContext struct {
	*ExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AndExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(XiapiParserAND, 0)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XiapiParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *XiapiParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, XiapiParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUnaryMinusExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(82)
			p.Match(XiapiParserMINUS)
		}
		{
			p.SetState(83)
			p.expr(10)
		}

	case 2:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(84)
			p.Match(XiapiParserNOT)
		}
		{
			p.SetState(85)
			p.expr(9)
		}

	case 3:
		localctx = NewLrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(86)
			p.ArrExpr()
		}

	case 4:
		localctx = NewAtomExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(87)
			p.Atom()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, XiapiParserRULE_expr)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(91)
					p.Match(XiapiParserPOW)
				}
				{
					p.SetState(92)
					p.expr(12)
				}

			case 2:
				localctx = NewXiapiltiplicationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, XiapiParserRULE_expr)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(94)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*XiapiltiplicationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<XiapiParserXIAPILT)|(1<<XiapiParserDIV)|(1<<XiapiParserMOD))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*XiapiltiplicationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(95)
					p.expr(9)
				}

			case 3:
				localctx = NewAdditiveExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, XiapiParserRULE_expr)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(97)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AdditiveExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == XiapiParserPLUS || _la == XiapiParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AdditiveExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(98)
					p.expr(8)
				}

			case 4:
				localctx = NewRelationalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, XiapiParserRULE_expr)
				p.SetState(99)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(100)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<XiapiParserGT)|(1<<XiapiParserLT)|(1<<XiapiParserGTEQ)|(1<<XiapiParserLTEQ))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(101)
					p.expr(7)
				}

			case 5:
				localctx = NewEqualityExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, XiapiParserRULE_expr)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(103)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualityExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == XiapiParserEQ || _la == XiapiParserNEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualityExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(104)
					p.expr(6)
				}

			case 6:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, XiapiParserRULE_expr)
				p.SetState(105)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(106)
					p.Match(XiapiParserAND)
				}
				{
					p.SetState(107)
					p.expr(5)
				}

			case 7:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, XiapiParserRULE_expr)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(109)
					p.Match(XiapiParserOR)
				}
				{
					p.SetState(110)
					p.expr(4)
				}

			}

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IArrExprContext is an interface to support dynamic dispatch.
type IArrExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetIdx returns the idx rule contexts.
	GetIdx() IExprContext

	// SetIdx sets the idx rule contexts.
	SetIdx(IExprContext)

	// IsArrExprContext differentiates from other interfaces.
	IsArrExprContext()
}

type ArrExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	idx    IExprContext
}

func NewEmptyArrExprContext() *ArrExprContext {
	var p = new(ArrExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XiapiParserRULE_arrExpr
	return p
}

func (*ArrExprContext) IsArrExprContext() {}

func NewArrExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrExprContext {
	var p = new(ArrExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XiapiParserRULE_arrExpr

	return p
}

func (s *ArrExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrExprContext) GetName() antlr.Token { return s.name }

func (s *ArrExprContext) SetName(v antlr.Token) { s.name = v }

func (s *ArrExprContext) GetIdx() IExprContext { return s.idx }

func (s *ArrExprContext) SetIdx(v IExprContext) { s.idx = v }

func (s *ArrExprContext) L() antlr.TerminalNode {
	return s.GetToken(XiapiParserL, 0)
}

func (s *ArrExprContext) R() antlr.TerminalNode {
	return s.GetToken(XiapiParserR, 0)
}

func (s *ArrExprContext) ID() antlr.TerminalNode {
	return s.GetToken(XiapiParserID, 0)
}

func (s *ArrExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterArrExpr(s)
	}
}

func (s *ArrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitArrExpr(s)
	}
}

func (s *ArrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitArrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XiapiParser) ArrExpr() (localctx IArrExprContext) {
	localctx = NewArrExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, XiapiParserRULE_arrExpr)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)

		var _m = p.Match(XiapiParserID)

		localctx.(*ArrExprContext).name = _m
	}
	{
		p.SetState(117)
		p.Match(XiapiParserL)
	}
	{
		p.SetState(118)

		var _x = p.expr(0)

		localctx.(*ArrExprContext).idx = _x
	}
	{
		p.SetState(119)
		p.Match(XiapiParserR)
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XiapiParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XiapiParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) CopyFrom(ctx *AtomContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParExprContext struct {
	*AtomContext
}

func NewParExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParExprContext {
	var p = new(ParExprContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *ParExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExprContext) OPAR() antlr.TerminalNode {
	return s.GetToken(XiapiParserOPAR, 0)
}

func (s *ParExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParExprContext) CPAR() antlr.TerminalNode {
	return s.GetToken(XiapiParserCPAR, 0)
}

func (s *ParExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterParExpr(s)
	}
}

func (s *ParExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitParExpr(s)
	}
}

func (s *ParExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitParExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanAtomContext struct {
	*AtomContext
}

func NewBooleanAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanAtomContext {
	var p = new(BooleanAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *BooleanAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanAtomContext) TRUE() antlr.TerminalNode {
	return s.GetToken(XiapiParserTRUE, 0)
}

func (s *BooleanAtomContext) FALSE() antlr.TerminalNode {
	return s.GetToken(XiapiParserFALSE, 0)
}

func (s *BooleanAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterBooleanAtom(s)
	}
}

func (s *BooleanAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitBooleanAtom(s)
	}
}

func (s *BooleanAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitBooleanAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdAtomContext struct {
	*AtomContext
}

func NewIdAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdAtomContext {
	var p = new(IdAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *IdAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdAtomContext) ID() antlr.TerminalNode {
	return s.GetToken(XiapiParserID, 0)
}

func (s *IdAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterIdAtom(s)
	}
}

func (s *IdAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitIdAtom(s)
	}
}

func (s *IdAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitIdAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type LenAtomContext struct {
	*AtomContext
}

func NewLenAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LenAtomContext {
	var p = new(LenAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *LenAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LenAtomContext) LEN() antlr.TerminalNode {
	return s.GetToken(XiapiParserLEN, 0)
}

func (s *LenAtomContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LenAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterLenAtom(s)
	}
}

func (s *LenAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitLenAtom(s)
	}
}

func (s *LenAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitLenAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringAtomContext struct {
	*AtomContext
}

func NewStringAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringAtomContext {
	var p = new(StringAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *StringAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringAtomContext) STRING() antlr.TerminalNode {
	return s.GetToken(XiapiParserSTRING, 0)
}

func (s *StringAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterStringAtom(s)
	}
}

func (s *StringAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitStringAtom(s)
	}
}

func (s *StringAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitStringAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilAtomContext struct {
	*AtomContext
}

func NewNilAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilAtomContext {
	var p = new(NilAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *NilAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilAtomContext) NIL() antlr.TerminalNode {
	return s.GetToken(XiapiParserNIL, 0)
}

func (s *NilAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterNilAtom(s)
	}
}

func (s *NilAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitNilAtom(s)
	}
}

func (s *NilAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitNilAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberAtomContext struct {
	*AtomContext
}

func NewNumberAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberAtomContext {
	var p = new(NumberAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *NumberAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberAtomContext) INT() antlr.TerminalNode {
	return s.GetToken(XiapiParserINT, 0)
}

func (s *NumberAtomContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(XiapiParserFLOAT, 0)
}

func (s *NumberAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.EnterNumberAtom(s)
	}
}

func (s *NumberAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XiapiListener); ok {
		listenerT.ExitNumberAtom(s)
	}
}

func (s *NumberAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case XiapiVisitor:
		return t.VisitNumberAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *XiapiParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, XiapiParserRULE_atom)
	var _la int

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

	p.SetState(132)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case XiapiParserOPAR:
		localctx = NewParExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Match(XiapiParserOPAR)
		}
		{
			p.SetState(122)
			p.expr(0)
		}
		{
			p.SetState(123)
			p.Match(XiapiParserCPAR)
		}

	case XiapiParserINT, XiapiParserFLOAT:
		localctx = NewNumberAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			_la = p.GetTokenStream().LA(1)

			if !(_la == XiapiParserINT || _la == XiapiParserFLOAT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case XiapiParserTRUE, XiapiParserFALSE:
		localctx = NewBooleanAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(126)
			_la = p.GetTokenStream().LA(1)

			if !(_la == XiapiParserTRUE || _la == XiapiParserFALSE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case XiapiParserID:
		localctx = NewIdAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(127)
			p.Match(XiapiParserID)
		}

	case XiapiParserSTRING:
		localctx = NewStringAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(128)
			p.Match(XiapiParserSTRING)
		}

	case XiapiParserNIL:
		localctx = NewNilAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(129)
			p.Match(XiapiParserNIL)
		}

	case XiapiParserLEN:
		localctx = NewLenAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(130)
			p.Match(XiapiParserLEN)
		}
		{
			p.SetState(131)
			p.expr(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *XiapiParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *XiapiParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
