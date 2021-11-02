package data_structures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MapSuite struct {
	suite.Suite
}

func (s *MapSuite) TestNewMap() {
	s.Equal(&Map{}, NewMap())
}

func (s *MapSuite) TestInsertGetDelelete() {
	mp := NewMap()
	mp.Set("key1", "value1")
	s.Equal("value1", mp.Get("key1"))

	mp.Set("key2", 23)
	mp.Set(3, "val 3")
	s.Equal(23, mp.Get("key2"))
	s.Equal("val 3", mp.Get(3))

	mp.Set("key1", "value 1 updated")
	s.Equal("value 1 updated", mp.Get("key1"))

	s.Equal(3, len(mp.items))
	mp.Delete("key1")
	s.Equal(2, len(mp.items))
	s.Nil(mp.Get("key1"))
	mp.Delete(3)
	s.Equal(1, len(mp.items))
	s.Nil(mp.Get(3))
	mp.Delete("doesnt_exist")
	s.Equal(1, len(mp.items))
}

func TestMapSuite(t *testing.T) {
	suite.Run(t, new(MapSuite))
}
