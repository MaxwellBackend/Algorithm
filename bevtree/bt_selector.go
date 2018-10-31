package bevtree

/*
	选择节点(Selector)：
	顺序执行子节点，只要碰到一个子节点返回true，则停止继续执行，并返回true，否则返回false，类似于程序中的逻辑或。
*/
type BTSelector struct {
	*BTNode
}

func NewBTSelector(parent *BTNode) *BTSelector {
	return &BTSelector{
		BTNode: NewBTNode(parent),
	}
}

func (sel *BTSelector) Execute() bool {
	for _, child := range sel.Children {
		if child.PreCondition() && child.Execute() {
			return true
		}
	}
	return false
}
