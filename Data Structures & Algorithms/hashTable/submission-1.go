type HashTable struct {
	size, capacity int
	table []*hashNode
}

// linked list to handle collisions
type hashNode struct {
	key, val int
	next *hashNode
}

func NewHashTable(capacity int) *HashTable {
	return &HashTable{
		size: 0, 
		capacity: capacity,
		table: make([]*hashNode, capacity),
	}
}

func newNode(key, val int) *hashNode {
	return &hashNode{
		key: key,
		val: val,
		next: nil,
	}
}

func (ht *HashTable) hash(key int) int {
	return key % ht.capacity
}

func (ht *HashTable) Insert(key, value int) {
	index := ht.hash(key)
	node := ht.table[index]
	if node == nil {
		// if no node in this table's index
		ht.table[index] = newNode(key, value)
	} else {
		// check if key exists in the linked list
		for cur := node; cur != nil; cur = cur.next {
			if key == cur.key {
				cur.val = value // replaced
				return
			}
		}
		// prepend new node to linked list if not exist
		newNode := newNode(key, value)
		newNode.next = ht.table[index]
		ht.table[index] = newNode
	}
	ht.size++

	// load factor = size / capacity
	if float32(ht.size)/float32(ht.capacity) >= 0.5 {
		ht.Resize()
	}
}

func (ht *HashTable) Get(key int) int {
	index := ht.hash(key)
	node := ht.table[index]
	for node != nil {
		if key == node.key {
			return node.val
		}
		node = node.next
	}
	return -1
}

func (ht *HashTable) Remove(key int) bool {
	index := ht.hash(key)
	node := ht.table[index]
	var prev *hashNode
	for node != nil {
		if key == node.key {
			if prev != nil {
				prev.next = node.next
			} else {
				ht.table[index] = node.next
			}
			ht.size--
			return true
		}
		prev, node = node, node.next
	}
	return false
}

func (ht *HashTable) GetSize() int {
	return ht.size
}

func (ht *HashTable) GetCapacity() int {
	return ht.capacity
}

func (ht *HashTable) Resize() {
	newCap := ht.capacity * 2
	newTable := make([]*hashNode, newCap)
	// since cap increases, we rehash all node
	for _, oldNode := range ht.table {
		for oldNode != nil {
			newIndex := oldNode.key % newCap
			if newTable[newIndex] == nil {
				// new linked list
				newTable[newIndex] = newNode(oldNode.key, oldNode.val)
			} else {
				// add to beginning of linked list, order doesn't matter
				newNode := newNode(oldNode.key, oldNode.val)
				newNode.next = newTable[newIndex]
				newTable[newIndex] = newNode
			}
			oldNode = oldNode.next
		}
	}
	ht.capacity = newCap
	ht.table = newTable
}
