package data

import (
	"flow-blueprint/branch"
	"flow-blueprint/tools"
	"testing"
)

func TestNewDataMap(t *testing.T) {
	// node:A->B,B取: data2.data2key2
	data := &Map{
		globalData: AnyMap{
			"data1": AnyMap{
				"data1key1": "value1",
			},
			"data2": AnyMap{
				"data2key2": "value1",
			},
			"list": []AnyMap{
				{
					"listkey1": "value1",
				},
				{
					"listkey2": "value2",
				},
				{
					"listkey3": "value3",
				},
			},
		},
	}

	data.Get(nil, branch.Key(tools.Md5("A")), []string{"data2", "data2key2"})
}
