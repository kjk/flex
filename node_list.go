package flex

// YGNodeList is a list of nodes
type YGNodeList struct {
	items []*YGNode
}

// YGNodeListNew creates  a new node
func YGNodeListNew(initialCapacity int) *YGNodeList {
	return &YGNodeList{}
}

// YGNodeListFree frees a node
func YGNodeListFree(list *YGNodeList) {
	// nothing to do
}

// YGNodeListCount returns list of nodes
func YGNodeListCount(list *YGNodeList) int {
	if list != nil {
		return len(list.items)
	}
	return 0
}

// YGNodeListAdd adds a node
func YGNodeListAdd(listp **YGNodeList, node *YGNode) {
	if *listp == nil {
		*listp = YGNodeListNew(4)
	}
	YGNodeListInsert(listp, node, YGNodeListCount(*listp))
}

// YGNodeListInsert insertes a node
func YGNodeListInsert(listp **YGNodeList, node *YGNode, i int) {
	if *listp == nil {
		*listp = YGNodeListNew(4)
	}
	list := *listp
	a := list.items
	// https://github.com/golang/go/wiki/SliceTricks
	list.items = append(a[:i], append([]*YGNode{node}, a[i:]...)...)
}

// YGNodeListRemove removes a node from list
func YGNodeListRemove(list *YGNodeList, index int) *YGNode {
	removed := list.items[index]
	list.items[index] = nil

	n := len(list.items)
	for i := index; i < n-1; i++ {
		list.items[i] = list.items[i+1]
		list.items[i+1] = nil
	}

	list.items = list.items[:n-1]
	return removed
}

// YGNodeListDelete deletes a node
func YGNodeListDelete(list *YGNodeList, node *YGNode) *YGNode {
	n := len(list.items)
	for i := 0; i < n; i++ {
		if list.items[i] == node {
			return YGNodeListRemove(list, i)
		}
	}

	return nil
}

// YGNodeListGet retruns a node at a given position
func YGNodeListGet(list *YGNodeList, index int) *YGNode {
	if YGNodeListCount(list) > 0 {
		return list.items[index]
	}

	return nil
}
