package branch

import "flow-blueprint/consts"

type Key string

type Data [][]consts.NodeKey

type Map map[Key]Data
