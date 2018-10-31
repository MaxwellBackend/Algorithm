package bevtree

/*
	并行节点(Parallel Node)：
	无论子节点返回值是什么都会遍历所有子节点
*/
type BTParallel struct {
	*BTNode
}

func NewBtParallel(parent IBTNode) *BTParallel {
	return &BTParallel{
		BTNode: NewBTNode(parent),
	}
}

func (par *BTParallel) Execute() bool {
	for _, child := range par.Children {
		if !child.PreCondition() {
			continue
		}

		child.Execute()
	}

	return true
}
