package concepts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHashTable(t *testing.T) {
	size := 5

	hashTable := NewHashTable(size)

	assert.Equal(t, &HashTable{
		size:   size,
		values: make([]interface{}, size),
	}, hashTable)
}

func TestInsertOneItem(t *testing.T) {
	hm := NewHashTable(5)

	hm.Insert("testkey", 3)

	assert.Equal(t, []interface{}{
		nil, nil, 3, nil, nil,
	}, hm.values)
}

func TestInsertOneMultipleItems(t *testing.T) {
	hm := NewHashTable(100)

	hm.Insert("Dad", 38)
	hm.Insert("Mom", 37)
	hm.Insert("Son", 15)
	hm.Insert("Daughter", 12)
	hm.Insert("Dog", 4)

	assert.Equal(t, 38, hm.values[65])
	assert.Equal(t, 37, hm.values[97])
	assert.Equal(t, 15, hm.values[4])
	assert.Equal(t, 12, hm.values[20])
	assert.Equal(t, 4, hm.values[82])
}

func TestGet(t *testing.T) {
	hm := NewHashTable(100)

	hm.Insert("Dad", 38)
	hm.Insert("Mom", 37)
	hm.Insert("Son", 15)
	hm.Insert("Daughter", 12)
	hm.Insert("Dog", 4)

	assert.Equal(t, 38, hm.Get("Dad"))
	assert.Equal(t, 37, hm.Get("Mom"))
	assert.Equal(t, 15, hm.Get("Son"))
	assert.Equal(t, 12, hm.Get("Daughter"))
	assert.Equal(t, 4, hm.Get("Dog"))
}

func TestGetValueDoesNotExist(t *testing.T) {
	hm := NewHashTable(10)

	assert.Equal(t, nil, hm.Get("nope"))
}

func TestDelete(t *testing.T) {
	hm := NewHashTable(100)

	hm.Insert("Dad", 38)
	hm.Insert("Mom", 37)
	hm.Insert("Son", 15)
	hm.Insert("Daughter", 12)
	hm.Insert("Dog", 4)

	assert.Equal(t, 38, hm.Get("Dad"))
	assert.Equal(t, 37, hm.Get("Mom"))
	assert.Equal(t, 15, hm.Get("Son"))
	assert.Equal(t, 12, hm.Get("Daughter"))
	assert.Equal(t, 4, hm.Get("Dog"))

	hm.Delete("Dad")
	hm.Delete("Mom")
	hm.Delete("Son")
	hm.Delete("Daughter")
	hm.Delete("Dog")

	assert.Equal(t, nil, hm.Get("Dad"))
	assert.Equal(t, nil, hm.Get("Mom"))
	assert.Equal(t, nil, hm.Get("Son"))
	assert.Equal(t, nil, hm.Get("Daughter"))
	assert.Equal(t, nil, hm.Get("Dog"))
}
