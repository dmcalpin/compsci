package data_structures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type HashTableSuite struct {
	suite.Suite
}

func (s *HashTableSuite) TestNewHashTable() {
	size := 5

	hashTable := NewHashTable(size)

	s.Equal(&HashTable{
		size:   size,
		values: make([]interface{}, size),
	}, hashTable)
}

func (s *HashTableSuite) TestInsertOneItem() {
	hm := NewHashTable(5)

	hm.Insert("testkey", 3)

	s.Equal([]interface{}{
		nil, nil, 3, nil, nil,
	}, hm.values)
}

func (s *HashTableSuite) TestInsertOneMultipleItems() {
	hm := NewHashTable(100)

	hm.Insert("Dad", 38)
	hm.Insert("Mom", 37)
	hm.Insert("Son", 15)
	hm.Insert("Daughter", 12)
	hm.Insert("Dog", 4)

	s.Equal(38, hm.values[65])
	s.Equal(37, hm.values[97])
	s.Equal(15, hm.values[4])
	s.Equal(12, hm.values[20])
	s.Equal(4, hm.values[82])
}

func (s *HashTableSuite) TestGet() {
	hm := NewHashTable(100)

	hm.Insert("Dad", 38)
	hm.Insert("Mom", 37)
	hm.Insert("Son", 15)
	hm.Insert("Daughter", 12)
	hm.Insert("Dog", 4)

	s.Equal(38, hm.Get("Dad"))
	s.Equal(37, hm.Get("Mom"))
	s.Equal(15, hm.Get("Son"))
	s.Equal(12, hm.Get("Daughter"))
	s.Equal(4, hm.Get("Dog"))
}

func (s *HashTableSuite) TestGetValueDoesNotExist() {
	hm := NewHashTable(10)

	s.Equal(nil, hm.Get("nope"))
}

func (s *HashTableSuite) TestDelete() {
	hm := NewHashTable(100)

	hm.Insert("Dad", 38)
	hm.Insert("Mom", 37)
	hm.Insert("Son", 15)
	hm.Insert("Daughter", 12)
	hm.Insert("Dog", 4)

	s.Equal(38, hm.Get("Dad"))
	s.Equal(37, hm.Get("Mom"))
	s.Equal(15, hm.Get("Son"))
	s.Equal(12, hm.Get("Daughter"))
	s.Equal(4, hm.Get("Dog"))

	hm.Delete("Dad")
	hm.Delete("Mom")
	hm.Delete("Son")
	hm.Delete("Daughter")
	hm.Delete("Dog")

	s.Equal(nil, hm.Get("Dad"))
	s.Equal(nil, hm.Get("Mom"))
	s.Equal(nil, hm.Get("Son"))
	s.Equal(nil, hm.Get("Daughter"))
	s.Equal(nil, hm.Get("Dog"))
}

func BenchmarkHashTableInsert(b *testing.B) {
	keys := []string{
		"aaa", "bbb", "ccc", "ddd", "eee", "fff",
		"ggg", "hhh", "iii", "jjj", "kkk", "lll",
		"mmm", "nnn", "ooo", "ppp", "qqq", "rrr",
		"sss", "ttt", "uuu", "vvv", "www", "xxx",
		"yyy", "zzz",
	}
	hm := NewHashTable(100)
	for i := 0; i < b.N; i++ {
		hm.Insert(keys[i%26], i)
	}
}

func BenchmarkHashTableGet(b *testing.B) {
	keys := []string{
		"aaa", "bbb", "ccc", "ddd", "eee", "fff",
		"ggg", "hhh", "iii", "jjj", "kkk", "lll",
		"mmm", "nnn", "ooo", "ppp", "qqq", "rrr",
		"sss", "ttt", "uuu", "vvv", "www", "xxx",
		"yyy", "zzz",
	}
	hm := NewHashTable(100)

	for i, key := range keys {
		hm.Insert(key, i)
	}

	for i := 0; i < b.N; i++ {
		hm.Get(keys[i%26])
	}
}

func BenchmarkHashTableDelete(b *testing.B) {
	keys := []string{
		"aaa", "bbb", "ccc", "ddd", "eee", "fff",
		"ggg", "hhh", "iii", "jjj", "kkk", "lll",
		"mmm", "nnn", "ooo", "ppp", "qqq", "rrr",
		"sss", "ttt", "uuu", "vvv", "www", "xxx",
		"yyy", "zzz",
	}
	hm := NewHashTable(100)

	for i, key := range keys {
		hm.Insert(key, i)
	}

	for i := 0; i < b.N; i++ {
		hm.Delete(keys[i%26])
	}
}

func TestHashTableSuite(t *testing.T) {
	suite.Run(t, new(HashTableSuite))
}
