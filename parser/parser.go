package parser

import (
	"flow-blueprint/branch"
	"flow-blueprint/consts"
	"flow-blueprint/define"
	"flow-blueprint/runTime"
)

type Parser struct {
	nodeMap define.IMap
	query   map[string]any
	errChan chan error
	branch  branch.Map
}

func (p *Parser) Parse(ctx runTime.RunTime) (map[string]any, error, any) {
	rootNode := p.nodeMap.Get(consts.RootNodeKey)
	if rootNode == nil {
		return nil, define.ErrNoRootNode, nil
	}
	endNode := p.nodeMap.Get(consts.EndNodeKey)
	if endNode == nil {
		return nil, define.ErrNoEndNode, nil
	}

	for _, key := range rootNode.GetNextNodeKeys() {
		if ctx.GetGo() == nil {
			ctx.ExecuteGo(p.next, consts.RootNodeKey, ctx, key)
		}
	}

}

func (p *Parser) next(ctx runTime.RunTime, branchKey string, parentNode define.INode, nextNodeKey consts.NodeKey) {
	nextNode := p.nodeMap.Get(nextNodeKey)
	if nextNode == nil {
		p.errChan <- define.ErrNodeNotFound
		return
	}
	// 通知子节点，父节点已经完成
	nextNode.ParentNodeFinish(parentNode)

	// 查询子节点的状态，查看是否能执行
	if status := nextNode.Status().GetStatus(); status != consts.NodeStatusReady {
		return
	}

	nextNode.Status().SetStatus(consts.NodeStatusRunning)

	// 执行子节点
	res, err := nextNode.Do(ctx, branchKey, p.query)
	if err != nil {
		p.errChan <- err
		return
	}

}
