package strings

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie_IsExisted(t *testing.T) {
	nt := NewTrie()
	nt.InsertKey("无码专")
	nt.InsertKey("新建户")
	nt.InsertKey("玉蒲团")
	s := "阿斯顿发送到非常喜爱的无码专高发大地飞歌我让特如果"
	if existed, _, key := nt.IsExisted(s); existed {
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
	if existed, _, _ := nt.IsExisted(s); existed {
		assert.Error(t, errors.New("error"))
	}

}

func TestTrie_IsExisted_last_character(t *testing.T) {
	nt := NewTrie()
	nt.InsertKey("无码专")
	nt.InsertKey("新建户")
	nt.InsertKey("玉蒲团")
	s := "大多数码"
	if existed, _, _ := nt.IsExisted(s); existed {
		assert.Error(t, errors.New("error"))
	}

}

func TestTrie_IsExisted2(t *testing.T) {
	nt := NewTrie()
	nt.InsertKey("人权理事会")
	nt.InsertKey("个税起征点")
	nt.InsertKey("日球迷赛后捡垃圾")
	nt.InsertKey("加多宝红罐回归")

	nt.InsertKey("乌龙杯哥伦比亚")

	nt.InsertKey("武汉大熊猫遭虐待")

	nt.InsertKey("第一所核高校诞生")

	nt.InsertKey("日本球迷跳河")

	nt.InsertKey("10")

	nt.InsertKey("7")
	nt.InsertKey("9")
	nt.InsertKey("8")
	nt.InsertKey("世界杯竞猜")
	s1 := "专栏测测测测试的新闻3"
	s2 := "sina"
	s3 := "hehe"
	s4 := "fdsafdsa,jjkk,kkkkj"
	nt.IsExisted(s1)
	nt.IsExisted(s2)
	nt.IsExisted(s3)
	nt.IsExisted(s4)

}

func TestTrie_IsExisted3(t *testing.T) {
	nt := NewTrie()
	nt.InsertKey("222")
	nt.InsertKey("shabi")
	nt.InsertKey("W对的")
	nt.InsertKey("我去去去去群")
	nt.InsertKey("我说的")
	nt.InsertKey("草泥马")
	nt.InsertKey("蛤")
	nt.InsertKey("问问")
	nt.InsertKey("中国")

	result, _, word := nt.IsExisted("中国")
	t.Log(result)
	t.Log(word)
}

func TestTrie_FindAllExisted(t *testing.T) {

	type args struct {
		content string
	}

	tests := []struct {
		name string
		trie *Trie
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "1",
			trie: newTrie(),
			args: args{
				getNews(),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.trie.FindAllExisted(tt.args.content)
			//t.Error(got)
		})
	}
}

func Benchmark_FindAllExisted(b *testing.B) {
	t := newTrie()
	for i := 0; i < b.N; i++ {
		go func() {
			if result := t.FindAllExisted(getNews()); len(result) != 5 {
				b.Error(result)
			}
		}()

	}
}

func getNews() string {
	return `大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v
	大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v
	大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v
	大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰我去去去去群范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v
	大大是的发送中国到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发shabi违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v
	大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v
	大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v
	大大是的中国发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发我说的生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v
	大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v
	大大是的发送到发违法深V自行车V字形从草泥马是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v大大是的发送到发违法深V自行车V字形从是安慰范围发生的自行车v`
}

func newTrie() *Trie {
	nt := NewTrie()
	nt.InsertKey("222")
	nt.InsertKey("shabi")
	nt.InsertKey("W对的")
	nt.InsertKey("我去去去去群")
	nt.InsertKey("我说的")
	nt.InsertKey("草泥马")
	nt.InsertKey("蛤")
	nt.InsertKey("问问")
	nt.InsertKey("中国")
	return &nt
}
