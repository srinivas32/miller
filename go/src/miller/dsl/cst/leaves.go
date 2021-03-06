// ================================================================
// CST build/execute for AST leaf nodes
// ================================================================

package cst

import (
	"errors"
	"math"

	"miller/dsl"
	"miller/lib"
	"miller/types"
)

// ----------------------------------------------------------------
func (this *RootNode) BuildLeafNode(
	astNode *dsl.ASTNode,
) (IEvaluable, error) {
	lib.InternalCodingErrorIf(astNode.Children != nil)
	sval := string(astNode.Token.Lit)

	switch astNode.Type {

	case dsl.NodeTypeDirectFieldValue:
		return this.BuildDirectFieldRvalueNode(sval), nil
		break
	case dsl.NodeTypeFullSrec:
		return this.BuildFullSrecRvalueNode(sval), nil
		break

	case dsl.NodeTypeDirectOosvarValue:
		return this.BuildDirectOosvarRvalueNode(sval), nil
		break
	case dsl.NodeTypeFullOosvar:
		return this.BuildFullOosvarRvalueNode(sval), nil
		break

	case dsl.NodeTypeLocalVariable:
		return this.BuildLocalVariableNode(sval), nil
		break

	case dsl.NodeTypeStringLiteral:
		return this.BuildStringLiteralNode(sval), nil
		break
	case dsl.NodeTypeIntLiteral:
		return this.BuildIntLiteralNode(sval), nil
		break
	case dsl.NodeTypeFloatLiteral:
		return this.BuildFloatLiteralNode(sval), nil
		break
	case dsl.NodeTypeBoolLiteral:
		return this.BuildBoolLiteralNode(sval), nil
		break
	case dsl.NodeTypeContextVariable:
		return this.BuildContextVariableNode(astNode)
		break
	case dsl.NodeTypeConstant:
		return this.BuildConstantNode(astNode)
		break

	case dsl.NodeTypePanic:
		return this.BuildPanicNode(astNode)
		break
	}

	return nil, errors.New(
		"CST BuildLeafNode: unhandled AST node " + string(astNode.Type),
	)
}

// ----------------------------------------------------------------
type DirectFieldRvalueNode struct {
	fieldName string
}

func (this *RootNode) BuildDirectFieldRvalueNode(fieldName string) *DirectFieldRvalueNode {
	return &DirectFieldRvalueNode{
		fieldName: fieldName,
	}
}
func (this *DirectFieldRvalueNode) Evaluate(state *State) types.Mlrval {
	value := state.Inrec.Get(&this.fieldName)
	if value == nil {
		return types.MlrvalFromAbsent()
	} else {
		return *value
	}
}

// ----------------------------------------------------------------
type FullSrecRvalueNode struct {
}

func (this *RootNode) BuildFullSrecRvalueNode(fieldName string) *FullSrecRvalueNode {
	return &FullSrecRvalueNode{}
}
func (this *FullSrecRvalueNode) Evaluate(state *State) types.Mlrval {
	return types.MlrvalFromMap(state.Inrec)
}

// ----------------------------------------------------------------
type DirectOosvarRvalueNode struct {
	variableName string
}

func (this *RootNode) BuildDirectOosvarRvalueNode(variableName string) *DirectOosvarRvalueNode {
	return &DirectOosvarRvalueNode{
		variableName: variableName,
	}
}
func (this *DirectOosvarRvalueNode) Evaluate(state *State) types.Mlrval {
	value := state.Oosvars.Get(&this.variableName)
	if value == nil {
		return types.MlrvalFromAbsent()
	} else {
		return *value
	}
}

// ----------------------------------------------------------------
type FullOosvarRvalueNode struct {
}

func (this *RootNode) BuildFullOosvarRvalueNode(fieldName string) *FullOosvarRvalueNode {
	return &FullOosvarRvalueNode{}
}
func (this *FullOosvarRvalueNode) Evaluate(state *State) types.Mlrval {
	return types.MlrvalFromMap(state.Oosvars)
}

// ----------------------------------------------------------------
type LocalVariableNode struct {
	variableName string
}

func (this *RootNode) BuildLocalVariableNode(variableName string) *LocalVariableNode {
	return &LocalVariableNode{
		variableName: variableName,
	}
}
func (this *LocalVariableNode) Evaluate(state *State) types.Mlrval {
	value := state.stack.ReadVariable(this.variableName)
	//state.stack.Dump()
	if value == nil {
		return types.MlrvalFromAbsent()
	} else {
		return *value
	}
}

// ----------------------------------------------------------------
type StringLiteralNode struct {
	literal types.Mlrval
}

func (this *RootNode) BuildStringLiteralNode(literal string) *StringLiteralNode {
	return &StringLiteralNode{
		literal: types.MlrvalFromString(literal),
	}
}
func (this *StringLiteralNode) Evaluate(state *State) types.Mlrval {
	return this.literal
}

// ----------------------------------------------------------------
type IntLiteralNode struct {
	literal types.Mlrval
}

func (this *RootNode) BuildIntLiteralNode(literal string) *IntLiteralNode {
	return &IntLiteralNode{
		literal: types.MlrvalFromInt64String(literal),
	}
}
func (this *IntLiteralNode) Evaluate(state *State) types.Mlrval {
	return this.literal
}

// ----------------------------------------------------------------
type FloatLiteralNode struct {
	literal types.Mlrval
}

func (this *RootNode) BuildFloatLiteralNode(literal string) *FloatLiteralNode {
	return &FloatLiteralNode{
		literal: types.MlrvalFromFloat64String(literal),
	}
}
func (this *FloatLiteralNode) Evaluate(state *State) types.Mlrval {
	return this.literal
}

// ----------------------------------------------------------------
type BoolLiteralNode struct {
	literal types.Mlrval
}

func (this *RootNode) BuildBoolLiteralNode(literal string) *BoolLiteralNode {
	return &BoolLiteralNode{
		literal: types.MlrvalFromBoolString(literal),
	}
}
func (this *BoolLiteralNode) Evaluate(state *State) types.Mlrval {
	return this.literal
}

// ================================================================
func (this *RootNode) BuildContextVariableNode(astNode *dsl.ASTNode) (IEvaluable, error) {
	lib.InternalCodingErrorIf(astNode.Token == nil)
	sval := string(astNode.Token.Lit)

	switch sval {

	case "FILENAME":
		return this.BuildFILENAMENode(), nil
		break
	case "FILENUM":
		return this.BuildFILENUMNode(), nil
		break

	case "NF":
		return this.BuildNFNode(), nil
		break
	case "NR":
		return this.BuildNRNode(), nil
		break
	case "FNR":
		return this.BuildFNRNode(), nil
		break

	case "IRS":
		return this.BuildIRSNode(), nil
		break
	case "IFS":
		return this.BuildIFSNode(), nil
		break
	case "IPS":
		return this.BuildIPSNode(), nil
		break

	case "ORS":
		return this.BuildORSNode(), nil
		break
	case "OFS":
		return this.BuildOFSNode(), nil
		break
	case "OPS":
		return this.BuildOPSNode(), nil
		break

	}

	return nil, errors.New(
		"CST BuildContextVariableNode: unhandled context variable " + sval,
	)
}

// ----------------------------------------------------------------
type FILENAMENode struct {
}

func (this *RootNode) BuildFILENAMENode() *FILENAMENode {
	return &FILENAMENode{}
}
func (this *FILENAMENode) Evaluate(state *State) types.Mlrval {
	return types.MlrvalFromString(state.Context.FILENAME)
}

// ----------------------------------------------------------------
type FILENUMNode struct {
}

func (this *RootNode) BuildFILENUMNode() *FILENUMNode {
	return &FILENUMNode{}
}
func (this *FILENUMNode) Evaluate(state *State) types.Mlrval {
	return types.MlrvalFromInt64(state.Context.FILENUM)
}

// ----------------------------------------------------------------
type NFNode struct {
}

func (this *RootNode) BuildNFNode() *NFNode {
	return &NFNode{}
}
func (this *NFNode) Evaluate(state *State) types.Mlrval {
	return types.MlrvalFromInt64(state.Inrec.FieldCount)
}

// ----------------------------------------------------------------
type NRNode struct {
}

func (this *RootNode) BuildNRNode() *NRNode {
	return &NRNode{}
}
func (this *NRNode) Evaluate(state *State) types.Mlrval {
	return types.MlrvalFromInt64(state.Context.NR)
}

// ----------------------------------------------------------------
type FNRNode struct {
}

func (this *RootNode) BuildFNRNode() *FNRNode {
	return &FNRNode{}
}
func (this *FNRNode) Evaluate(state *State) types.Mlrval {
	return types.MlrvalFromInt64(state.Context.FNR)
}

// ----------------------------------------------------------------
type IRSNode struct {
}

func (this *RootNode) BuildIRSNode() *IRSNode {
	return &IRSNode{}
}
func (this *IRSNode) Evaluate(state *State) types.Mlrval {
	return types.MlrvalFromString(state.Context.IRS)
}

// ----------------------------------------------------------------
type IFSNode struct {
}

func (this *RootNode) BuildIFSNode() *IFSNode {
	return &IFSNode{}
}
func (this *IFSNode) Evaluate(state *State) types.Mlrval {
	return types.MlrvalFromString(state.Context.IFS)
}

// ----------------------------------------------------------------
type IPSNode struct {
}

func (this *RootNode) BuildIPSNode() *IPSNode {
	return &IPSNode{}
}
func (this *IPSNode) Evaluate(state *State) types.Mlrval {
	return types.MlrvalFromString(state.Context.IPS)
}

// ----------------------------------------------------------------
type ORSNode struct {
}

func (this *RootNode) BuildORSNode() *ORSNode {
	return &ORSNode{}
}
func (this *ORSNode) Evaluate(state *State) types.Mlrval {
	return types.MlrvalFromString(state.Context.ORS)
}

// ----------------------------------------------------------------
type OFSNode struct {
}

func (this *RootNode) BuildOFSNode() *OFSNode {
	return &OFSNode{}
}
func (this *OFSNode) Evaluate(state *State) types.Mlrval {
	return types.MlrvalFromString(state.Context.OFS)
}

// ----------------------------------------------------------------
type OPSNode struct {
}

func (this *RootNode) BuildOPSNode() *OPSNode {
	return &OPSNode{}
}
func (this *OPSNode) Evaluate(state *State) types.Mlrval {
	return types.MlrvalFromString(state.Context.OPS)
}

// ================================================================
func (this *RootNode) BuildConstantNode(astNode *dsl.ASTNode) (IEvaluable, error) {
	lib.InternalCodingErrorIf(astNode.Token == nil)
	sval := string(astNode.Token.Lit)

	switch sval {

	case "M_PI":
		return this.BuildMathPINode(), nil
		break
	case "M_E":
		return this.BuildMathENode(), nil
		break

	}

	return nil, errors.New(
		"CST BuildContextVariableNode: unhandled context variable " + sval,
	)
}

// ----------------------------------------------------------------
type MathPINode struct {
}

func (this *RootNode) BuildMathPINode() *MathPINode {
	return &MathPINode{}
}
func (this *MathPINode) Evaluate(state *State) types.Mlrval {
	return types.MlrvalFromFloat64(math.Pi)
}

// ----------------------------------------------------------------
type MathENode struct {
}

func (this *RootNode) BuildMathENode() *MathENode {
	return &MathENode{}
}
func (this *MathENode) Evaluate(state *State) types.Mlrval {
	return types.MlrvalFromFloat64(math.E)
}

// ----------------------------------------------------------------
// The panic token is a special token which causes a panic when evaluated.
// This is for testing that AND/OR short-circuiting is implemented correctly:
// output = input1 || panic should NOT panic the process when input1 is true.

type PanicNode struct {
}

func (this *RootNode) BuildPanicNode(astNode *dsl.ASTNode) (*PanicNode, error) {
	return &PanicNode{}, nil
}
func (this *PanicNode) Evaluate(state *State) types.Mlrval {
	lib.InternalCodingErrorPanic("Panic token was evaluated, not short-circuited.")
	return types.MlrvalFromError() // not reached
}
