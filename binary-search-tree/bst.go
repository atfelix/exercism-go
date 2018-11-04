package binarysearchtree

type SearchTreeData struct {
	data int
	left, right *SearchTreeData
}

func Bst(data int) *SearchTreeData {
	return &SearchTreeData{data: data}
}

func (tree *SearchTreeData) Insert(data int) {
	var parent *SearchTreeData

	for tree != nil {
		parent = tree
		if data <= tree.data {
			tree = tree.left
		} else {
			tree = tree.right
		}
	}

	if data <= parent.data {
		parent.left = Bst(data)
	} else {
		parent.right = Bst(data)
	}
}

func (tree *SearchTreeData) MapString(f func(int) string) []string {
	slice := []string{}

	for _, element := range tree.preOrderTraversal() {
		slice = append(slice, f(element))
	}

	return slice
}

func (tree *SearchTreeData) MapInt(f func(int) int) []int {
	slice := []int{}

	for _, element := range tree.preOrderTraversal() {
		slice = append(slice, f(element))
	}

	return slice
}

func (tree *SearchTreeData) preOrderTraversal() []int {
	slice := []int{}
	if tree == nil {
		return slice
	}
	slice = append(slice, tree.left.preOrderTraversal()...)
	slice = append(slice, tree.data)
	slice = append(slice, tree.right.preOrderTraversal()...)
	return slice
}
