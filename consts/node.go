package consts

type NodeStatus int

const (
	NodeStatusInit NodeStatus = iota
	NodeStatusWait            // 等待上级节点执行完成
	NodeStatusRunning
	NodeStatusFinish
	NodeStatusClose
	// NodeStatusReady 该节点已经准备好，可以执行
	NodeStatusReady
)

type NodeHandlerResult int

const (
	NodeHandlerResultSuccess NodeHandlerResult = iota
	NodeHandlerResultFail
	NodeHandlerResultBeClose
)

func (s NodeHandlerResult) String() string {
	switch s {
	case NodeHandlerResultSuccess:
		return "success"
	case NodeHandlerResultFail:
		return "fail"
	case NodeHandlerResultBeClose:
		return "be close"
	default:
		return "unknown"
	}
}

type NoticeChildrenStatus int

const (
	// NoticeChildrenStatusSuccess 通知子节点该父级节点执行成功
	NoticeChildrenStatusSuccess NoticeChildrenStatus = iota
	// NoticeChildrenStatusFail 通知子节点该父级节点执行失败
	NoticeChildrenStatusFail
	// NoticeChildrenStatusClose 通知子节点转为关闭状态
	NoticeChildrenStatusClose
)
