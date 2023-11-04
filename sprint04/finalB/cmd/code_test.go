package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HashMap(t *testing.T) {
	t.Run("Can get error", func(t *testing.T){
		m := NewHashMap()
		_, err := m.Get(23)
		assert.Error(t, err)
	})

	t.Run("Can save value", func(t *testing.T){
		m := NewHashMap()
		m.Put(23, 100)
		v, err := m.Get(23)
		assert.NoError(t, err)
		assert.Equal(t, 100,v)
	})

	t.Run("Can resolve colision", func(t *testing.T){
		m := NewHashMap()
		var v int
		var err error
		m.Put(23, 100)
		m.Put(23+10000, 1)
		v, err = m.Get(23)
		assert.NoError(t, err)
		assert.Equal(t, 100,v)
		v, err = m.Get(23+10000)
		assert.NoError(t, err)
		assert.Equal(t, 1,v)
	})

	t.Run("Can delete key single element", func(t *testing.T){
		m := NewHashMap()
		var err error
		m.Put(23, 100)
		m.Delete(23)
		_, err = m.Get(23)
		assert.Error(t, err)
	})
	t.Run("Can delete key in the middle in chain", func(t *testing.T){
		m := NewHashMap()
		var err error
		m.Put(23, 100)
		m.Put(23+10000, 1)
		m.Put(23+20000, 3)
		m.Delete(23 + 10000)
		_, err = m.Get(23+10000)
		assert.Error(t, err)
		_, err = m.Get(23)
		assert.Nil(t, err)
		_, err = m.Get(23 + 20000)
		assert.Nil(t, err)
	})
	t.Run("Can delete key last in chain", func(t *testing.T){
		m := NewHashMap()
		var v int
		var err error
		m.Put(23, 100)
		m.Put(23+10000, 1)
		v, err = m.Get(23)
		assert.NoError(t, err)
		assert.Equal(t, 100,v)
		v, err = m.Get(23+10000)
		assert.NoError(t, err)
		assert.Equal(t, 1,v)

		m.Delete(23+10000)
		v, err = m.Get(23+10000)
		assert.Error(t, err)

		m.Delete(23)
		v, err = m.Get(23)
		assert.Error(t, err)
	})

	t.Run("Remove All buttomup", func(t *testing.T){
		m := NewHashMap()
		var v int
		var err error
		m.Put(23, 100)
		m.Put(23+10000, 1)
		m.Put(23+20000, 2)

		v, err = m.Delete(23)
		assert.Equal(t, 100, v)
		assert.NoError(t, err)

		v, err = m.Delete(23+10000)
		assert.Equal(t, 1, v)
		assert.NoError(t, err)

		v, err = m.Delete(23+20000)
		assert.Equal(t, 2, v)
		assert.NoError(t, err)
	})
	t.Run("Remove All ", func(t *testing.T){
		m := NewHashMap()
		var v int
		var err error
		m.Put(23, 100)
		m.Put(23+10000, 1)
		m.Put(23+20000, 2)

		v, err = m.Delete(23+20000)
		assert.Equal(t, 2, v)
		assert.NoError(t, err)
		v, err = m.Delete(23+20000)
		assert.Error(t, err)

		v, err = m.Delete(23+10000)
		assert.Equal(t, 1, v)
		assert.NoError(t, err)
		v, err = m.Delete(23+10000)
		assert.Error(t, err)

		v, err = m.Delete(23)
		assert.Equal(t, 100, v)
		assert.NoError(t, err)
		v, err = m.Delete(23)
		assert.Error(t, err)
		v, err = m.Delete(23)
		assert.Error(t, err)
		
		v, err = m.Delete(-23)
		assert.Error(t, err)
	})

}

