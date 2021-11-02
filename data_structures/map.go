package data_structures

// A Map is a an associative array which
// means values are stored and retrieved
// along with keys which identify them

type Map struct {
	items []*MapNode
}

type MapNode struct {
	key   interface{}
	value interface{}
}

// Instantiates a new Map
func NewMap() *Map {
	return &Map{}
}

// Sets an item at key: k to value: v
// If the key already exists, its value
// will be over-written
func (m *Map) Set(k interface{}, v interface{}) {
	_, node := m.findMapNode(k)
	if node != nil {
		node.value = v
		return
	}

	m.items = append(m.items, &MapNode{k, v})
}

// Retrieves an item from the map located
// at key: k
func (m *Map) Get(k interface{}) interface{} {
	_, node := m.findMapNode(k)
	if node == nil {
		return nil
	}

	return node.value
}

// Deletes an item from the map located at
// key: k
func (m *Map) Delete(k interface{}) {
	i, node := m.findMapNode(k)
	if node == nil {
		return
	}

	m.items[i] = m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
}

// returns the index and item belonging to key: k
// if no key is found, the index will be -1
// and node will be nil
func (m *Map) findMapNode(k interface{}) (int, *MapNode) {
	for i, item := range m.items {
		if item.key == k {
			return i, item
		}
	}
	return -1, nil
}
