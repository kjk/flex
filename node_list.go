package flex

type YGNodeList struct {
	items []YGNodeRef
}

type YGNodeListRef *YGNodeList

func YGNodeListNew(initialCapacity uint32) YGNodeListRef {
	return &YGNodeList{}
}

func YGNodeListFree(list YGNodeListRef) {
	// nothing to do
}

func YGNodeListCount(list YGNodeListRef) uint32 {
	if list != nil {
		return uint32(len(list.items))
	}
	return 0
}

func YGNodeListAdd(listp *YGNodeListRef, node YGNodeRef) {
	if *listp == nil {
		*listp = YGNodeListNew(4)
	}
	YGNodeListInsert(listp, node, YGNodeListCount(*listp))
}

// https://github.com/golang/go/wiki/SliceTricks
func YGNodeListInsert(listp *YGNodeListRef, node YGNodeRef, index uint32) {
	if *listp == nil {
		*listp = YGNodeListNew(4)
	}
	list := *listp
	i := int(index)
	a := list.items
	list.items = append(a[:i], append([]YGNodeRef{node}, a[i:]...)...)
}

func YGNodeListRemove(list YGNodeListRef, index uint32) YGNodeRef {
	removed := list.items[index]
	list.items[index] = nil

	n := len(list.items)
	for i := int(index); i < n-1; i++ {
		list.items[i] = list.items[i+1]
		list.items[i+1] = nil
	}

	list.items = list.items[:n-1]
	return removed
}

func YGNodeListDelete(list YGNodeListRef, node YGNodeRef) YGNodeRef {
	n := len(list.items)
	for i := 0; i < n; i++ {
		if list.items[i] == node {
			return YGNodeListRemove(list, uint32(i))
		}
	}

	return nil
}

func YGNodeListGet(list YGNodeListRef, index uint32) YGNodeRef {
	if YGNodeListCount(list) > 0 {
		return list.items[index]
	}

	return nil
}
