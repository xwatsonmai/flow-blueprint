package consts

type PushEndDataType int

const (
	PushEndDataTypeEndNode PushEndDataType = iota
	PushEndDataTypeCloseNode
)

type LogicFlowNodeType string

const (
	LogicFlowRootNode     LogicFlowNodeType = "root"
	LogicFlowEndNode      LogicFlowNodeType = "end"
	LogicFlowBusinessNode LogicFlowNodeType = "business"
	LogicFlowSwitchNode   LogicFlowNodeType = "switch"
)

type LogicFlowBusinessNodeType string

const (
	ApiLogicFlowNode       LogicFlowBusinessNodeType = "api"
	ModelLogicFlowNode     LogicFlowBusinessNodeType = "Module"
	FuncLogicFlowNode      LogicFlowBusinessNodeType = "func"
	ApiNodeTypeSiApiConfig LogicFlowBusinessNodeType = "si"
)

type SwitchNodeCheckRule string

const (
	SwitchNodeCheckRuleRange SwitchNodeCheckRule = "range"
	SwitchNodeCheckRuleOne   SwitchNodeCheckRule = "one"
	SwitchNodeCheckRuleAll   SwitchNodeCheckRule = "all"
)

type DataType string

const (
	DataTypeString   DataType = "string"
	DataTypeNumber   DataType = "number"
	DataTypeBool     DataType = "bool"
	DataTypeMetadata DataType = "metadata" // 元数据，包含了字符串、数字、bool，表示改元素为非结构化数据，可能是一个字符串，也可能是一个数字，也可能是一个bool
	DataTypeArray    DataType = "array"
	DataTypeObject   DataType = "object"
)

const (
	NumberNotSet = -1001 // 表示数据类型为number，但是没有获取到数据
)

type InputDataFromType string

const (
	InputFromGlobal InputDataFromType = "global"
	InputFromNode   InputDataFromType = "node"
)

const (
	SuccessCode           = 0
	NoAuthCode            = 1005
	LogicErrCode          = 1000
	DBErrorCode           = 1001
	InternalServerErrCode = 1500
	NotFoundCode          = 1004
)

type NodeKey string
type NodeKeys []NodeKey
type NodeResultKey string

const (
	RootNodeKey NodeKey = "root"
	EndNodeKey  NodeKey = "end"
)
