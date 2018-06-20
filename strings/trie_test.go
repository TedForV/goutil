package strings

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie_IsExisted(t *testing.T) {
	nt := NewTrie()
	nt.InsertKey("无码专")
	nt.InsertKey("新建户")
	nt.InsertKey("玉蒲团")
	s := "阿斯顿发送到非常无码专喜爱的高发大地飞歌我让特如果"
	if existed, key := nt.IsExisted(s); existed {
		assert.Equal(t, "无码专", key, "dirty words")
	} else {
		t.Error("'无码专' is missing")
	}

}

func TestTrie_IsExisted_one_character(t *testing.T) {
	nt := NewTrie()
	nt.InsertKey("无码专")
	nt.InsertKey("新建户")
	nt.InsertKey("玉蒲团")
	s := "无"
	if existed, _ := nt.IsExisted(s); existed {
		assert.Error(t, errors.New("error"))
	}

}
