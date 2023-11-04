package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanAddDocument(t *testing.T) {
	t.Run("1. Increment last document id", func(t *testing.T) {
		s := NewStorage()
		s.Add("Word1 Word2 Word3")
		assert.Equal(t, 1, s.LastDocumentID)
	})

	t.Run("2. All words linked to first document", func(t *testing.T) {
		s := NewStorage()
		s.Add("Word1 Word2 Word3")

		assert.Equal(t, map[int]int{1:1}, s.Index["Word1"])
		assert.Equal(t, map[int]int{1:1}, s.Index["Word2"])
		assert.Equal(t, map[int]int{1:1}, s.Index["Word3"])
	})

	t.Run("3. Document can have same words", func(t *testing.T) {
		s := NewStorage()
		s.Add("Word1 Word1 Word1")

		assert.Equal(t, map[int]int{1:3}, s.Index["Word1"])
	})

	t.Run("4. Can add more than 1 document", func(t *testing.T) {
		s := NewStorage()
		s.Add("Word1 Word2")
		s.Add("Word1 Word3")

		assert.Equal(t, map[int]int{1:1, 2:1}, s.Index["Word1"])
		assert.Equal(t, map[int]int{1:1}, s.Index["Word2"])
		assert.Equal(t, map[int]int{2:1}, s.Index["Word3"])
	})
}

func Test_CanQuery(t *testing.T) {
	t.Run("query will use uniq words from txt", func(t *testing.T) {
		df := uniqWords
		defer func(){uniqWords = df}()
		uniqWordsCalled := false
		uniqWords = func(txt string) []string {
			uniqWordsCalled = true	
			return []string{}
		}

		s := NewStorage()
		sQueryUWordsCalled := false
		s.QueryUWords = func([]string)[]int{
			sQueryUWordsCalled = true
			return []int{}
		}

		s.Query("Query")

		assert.True(t, uniqWordsCalled)
		assert.True(t, sQueryUWordsCalled)
	})

	t.Run("0. Can get empty result", func(t *testing.T) {
		s := NewStorage()
		res := s.Query("Unknown")

		assert.Equal(t, []int{}, res)
	})

	t.Run("1. Can query single word", func(t *testing.T) {
		s := NewStorage()
		s.Add("Word1 Word2")
		s.Add("Word1 Word3")
		res := s.Query("Word1")

		assert.Equal(t, []int{1, 2}, res)
	})

	t.Run("2. Can query two word", func(t *testing.T) {
		s := NewStorage()
		s.Add("Word1 Word2")
		s.Add("Word1 Word3")
		res := s.Query("Word1 Word2")

		assert.Equal(t, []int{1, 2}, res)
	})

	t.Run("3. Docs are sorted by range", func(t *testing.T) {
		s := NewStorage()
		s.Add("Word1 Word2")
		s.Add("Word1 Word3 Word1 Word1")
		res := s.Query("Word1 Word2")

		assert.Equal(t, []int{2, 1}, res)
	})
}

func Test_UniqWords(t *testing.T) {
	txt := "Word1 Word2 Word3 Word1" 
	res := uniqWords(txt)
	assert.Equal(t, 3, len(res))
}
