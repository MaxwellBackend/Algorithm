package bevtree

/*
	顺序节点(Sequence)：
	顺序执行子节点，只要碰到一个子节点返回false，则停止继续执行，并返回false，否则返回true，类似于程序中的逻辑与。
*/
type BTSequence struct {
	*BTNode
}

func NewBtSequence(parent IBTNode) *BTSequence {
	return &BTSequence{
		BTNode: NewBTNode(parent),
	}
}

func (sel *BTSequence) Execute() bool {
	for _, child := range sel.Children {
		if !child.PreCondition() {
			return false
		}

		if !child.Execute() {
			return false
		}
	}

	return true
}
