package define

import "errors"

var (
	// ErrNoRootNode 无ROOT节点错误
	ErrNoRootNode = errors.New("缺少ROOT节点")
	// ErrNoEndNode 缺少END节点错误
	ErrNoEndNode = errors.New("缺少END节点")
	// ErrInputDataRequired 必填参数为空
	ErrInputDataRequired = errors.New("必填参数为空")

	ErrNodeNotFound = errors.New("节点未找到")
)
