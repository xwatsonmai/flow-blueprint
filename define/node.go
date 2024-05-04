package define

import (
	"flow-blueprint/consts"
	"flow-blueprint/runTime"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"sync"
)

type LoopType int

const (
	LoopTypeAll LoopType = iota + 1
	LoopTypeItem
)

type NodeStatus struct {
	sync.RWMutex
	status consts.NodeStatus
}

func NewNodeStatus() *NodeStatus {
	return &NodeStatus{
		status: consts.NodeStatusInit,
	}
}

func (n *NodeStatus) GetStatus() consts.NodeStatus {
	n.RLock()
	defer n.RUnlock()
	return n.status
}

func (n *NodeStatus) SetStatus(status consts.NodeStatus) {
	n.Lock()
	defer n.Unlock()
	n.status = status
}

type INode interface {
	GetNodeKey() consts.NodeKey
	GetNodeName() string
	GetNodeDataKey() string
	GetNodeType() consts.LogicFlowNodeType
	// GetNextNodeKeys 获取下级节点的key
	GetNextNodeKeys() consts.NodeKeys
	Status() *NodeStatus
	Do(runTime.RunTime, string, map[string]any) (NodeResultData, error)
	ParentNodeFinish(node INode)
	GetBranchKey() BranchKey
	BranchPath() [][]consts.NodeKey
}

type IEndNode interface {
	GetResult(*gjson.Json) map[string]any
}

type NodeResultData struct {
	Data             map[DataSetPath]*gvar.Var
	ReplaceData      bool
	NodeType         consts.LogicFlowNodeType
	NodeKey          consts.NodeKey
	SwitchNodeResult SwitchNodeResult
	ResultSetSlice   bool
}

type NodeDebuggerData struct {
	NodeInputOutPutData
	NodeRunStatus consts.NodeHandlerResult
	Error         error
}

//func (n NodeDebuggerData) MarshalJSON() ([]byte, error) {
//	if n.Error == nil {
//		n.Error = fmt.Errorf("")
//	}
//	fmt.Println("in MarshalJSON")
//	resultData := struct {
//		NodeInputOutPutData
//		NodeRunStatus consts.NodeHandlerResult
//		Error         string
//	}{
//		NodeInputOutPutData: n.NodeInputOutPutData,
//		NodeRunStatus:       n.NodeRunStatus,
//		Error:               n.Error.Error(),
//	}
//	// 对resultData进行json序列化
//	return gjson.Marshal(resultData)
//}

type NodeInputOutPutData struct {
	InputData any
	OutData   any
}

type SwitchNodeResult struct {
	IsFilter      bool
	FilterIndexes []int
	FilterPath    string
}

type NodeSign struct {
	Status consts.NoticeChildrenStatus
	Key    string
}
type ISwitchParser interface {
	Parse(data map[string]any) (bool, error)
}

type IBusinessNodeParser interface {
	Parser(runTime.RunTime, map[string]interface{}) (*gvar.Var, error)
}

type ISwitchNodeParser interface {
	Parser(runTime.RunTime, string, map[string]interface{}) (bool, string, error)
}

type NodeReadyStatus int

const (
	// NodeReadyStatusWait 该节点正处于等待状态，尚有未执行完毕的父级节点
	NodeReadyStatusWait NodeReadyStatus = iota
	// NodeReadyStatusReady 该节点已经准备就绪，可以执行
	NodeReadyStatusReady
	// NodeReadyStatusClose 该节点不执行，直接关闭
	NodeReadyStatusClose
)
