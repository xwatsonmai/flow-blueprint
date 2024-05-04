package define

import (
	"flow-blueprint/consts"
)

type IMap interface {
	Get(key consts.NodeKey) INode
	GetNodeParser(key consts.NodeKey) INode
	GetEndNode() IEndNode
	Len() int
	Range(runFunc rangeFunc)
}

//	type Map struct {
//		data map[consts.NodeKey]define.INode
//		lock *sync.RWMutex
//	}
//
//	func NewMap(data map[consts.NodeKey]define.INode) *Map {
//		return &Map{
//			data: data,
//			lock: new(sync.RWMutex),
//		}
//	}
//
//	func (m *Map) Get(key consts.NodeKey) define.INode {
//		m.lock.RLock()
//		defer m.lock.RUnlock()
//		if val, ok := m.data[key]; ok {
//			return val
//		}
//		return nil
//	}
//
//	func (m *Map) GetNodeParser(key consts.NodeKey) define.INode {
//		m.lock.Lock()
//		defer m.lock.Unlock()
//		if _, ok := m.data[key]; !ok {
//			return nil
//		}
//		node := m.data[key]
//		// 检查node的状态
//		if node.GetStatus() == nil {
//			// 未进行过初始
//			node.Init()
//		} else {
//			return node
//		}
//		// 根据node的类型，进行不同的处理
//		switch node.GetNodeType() {
//		case consts.LogicFlowBusinessNode:
//			businessNode := NewBusinessNode(node.(*Node))
//			m.data[key] = businessNode
//			return businessNode
//		case consts.LogicFlowSwitchNode:
//			switchNode := NewSwitchNode(node.(*Node))
//			m.data[key] = switchNode
//			return switchNode
//		}
//		return nil
//	}
//
//	func (m *Map) GetEndNode() define.IEndNode {
//		m.lock.Lock()
//		defer m.lock.Unlock()
//		endNode := m.data[consts.EndNodeKey]
//		if endNode == nil {
//			return nil
//		}
//		return NewEndNode(endNode.(*Node))
//	}
//
//	func (m *Map) Len() int {
//		return len(m.data)
//	}
type rangeFunc func(nodeKey consts.NodeKey, nodeItem INode) bool

//
//// Range 循环遍历
//func (m *Map) Range(runFunc rangeFunc) {
//	for key, node := range m.data {
//		isNext := runFunc(key, node)
//		if !isNext {
//			break
//		}
//	}
//}
