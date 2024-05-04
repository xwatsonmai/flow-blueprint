package data

import (
	"flow-blueprint/branch"
	"flow-blueprint/runTime"
)

// Map 蓝图数据结构
// 用于存储蓝图数据
// globalData 为全局数据，蓝图所有产生的数据会存储在这里
// dataMap 节点可访问的数据，用于映射节点数据，实现分支数据隔离访问。
// 分支key的生成规则是，该分支上的所有节点的key的拼接做md5
type Map struct {
	// 全局数据
	globalData AnyMap
	dataMap    map[branch.Key]interface{}
}

//func (m *Map) SetData(ctx runTime.RunTime, node define.INode, data AnyMap) {
//	m.globalData = data
//}

func (m *Map) Get(ctx runTime.RunTime, branchKey branch.Key, getPath []string) {



}

//func NewDataMap() *Map {
//	return &Map{
//		globalData: make(map[string]interface{}),
//		dataMap:    make(map[BranchKey]interface{}),
//	}
//}
