package bevtree

/*
	行为树节点抽象接口
 */
type IBTNode interface {
	SetParent(node IBTNode) // 设置父节点
	AddNode(node IBTNode)   // 添加子节点
	PreCondition() bool     // 准入条件判断
	Execute() bool          // 执行逻辑
}

/*
	行为树基础节点实现
 */
type BTNode struct {
	Parent   IBTNode
	Children []IBTNode
}

func NewBTNode(parent IBTNode) *BTNode {
	return &BTNode{
		Parent:   parent,
		Children: make([]IBTNode, 0),
	}
}

func (bt *BTNode) SetParent(node IBTNode) {
	bt.Parent = node
}

func (bt *BTNode) AddNode(node IBTNode) {
	if node != nil {
		node.SetParent(bt)
	}
	bt.Children = append(bt.Children, node)
}


func (bt *BTNode) PreCondition() bool {
	return true
}

func (bt *BTNode) Execute() bool {
	return true
}
