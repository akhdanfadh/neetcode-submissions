// these are without balancing
// avg: O(log n), worst: O(n)

type TreeMap struct {
	key, val int
	exists bool
	left, right *TreeMap
}

func NewTreeMap() *TreeMap {
	return &TreeMap{}
}

func (tm *TreeMap) Insert(key, val int) {
	if !tm.exists { // we are in "root"
		tm.key, tm.val, tm.exists = key, val, true
		tm.left = &TreeMap{}
		tm.right = &TreeMap{}
		return
	}

	if key > tm.key {
		tm.right.Insert(key, val)
	} else if key < tm.key {
		tm.left.Insert(key, val)
	} else {
		tm.val = val
	}
}

func (tm *TreeMap) Get(key int) int {
	if !tm.exists { return -1 }

	if key > tm.key {
		return tm.right.Get(key)
	} else if key < tm.key {
		return tm.left.Get(key)
	} else {
		return tm.val
	}
}

func (tm *TreeMap) GetMin() int {
	if !tm.exists { return -1 }
	if !tm.left.exists { return tm.val }
	return tm.left.GetMin()
}

func (tm *TreeMap) GetMax() int {
	if !tm.exists { return -1 }
	if !tm.right.exists { return tm.val }
	return tm.right.GetMax()
}

func (tm *TreeMap) Remove(key int) {
	if !tm.exists { return }

	if key > tm.key {
		tm.right.Remove(key)
	} else if key < tm.key {
		tm.left.Remove(key)
	} else { // this node is removed

		// case 1: two children
		if tm.left.exists && tm.right.exists {
			// move min node from right children here
			min := tm.right.getMinNode()
			tm.key, tm.val = min.key, min.val
			tm.right.Remove(min.key)
			return
		}

		// case 2: 0/1 children
		if tm.left.exists {
			*tm = *tm.left // replace
		} else if tm.right.exists {
			*tm = *tm.right
		} else {
			*tm = TreeMap{}
		}
	}
}

func (tm *TreeMap) GetInorderKeys() []int {
	if !tm.exists { return nil } // base case
	keys := tm.left.GetInorderKeys() // give left part
	keys = append(keys, tm.key) // append this node
	keys = append(keys, tm.right.GetInorderKeys()...) // give right part
	return keys
}

func (tm *TreeMap) getMinNode() *TreeMap {
	if !tm.left.exists { return tm }
	return tm.left.getMinNode()
}
