package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/stretchr/testify/assert"
	"testing"
)

var esServer = []string{"http://10.10.11.200:9200"}

func TestNewIndexWithMapping(t *testing.T) {
	//esServer := []string{"http://127.0.0.1:9200"}
	InitialESConfig(esServer, true)
	typeName := "new"
	mapping := `{
		"new":{
			"properties":{
				"id":{
					"type":"keyword"
				},
				"p_a":{
					"type":"text",
					"analyzer":"ik_max_word",
					"search_analyzer":"ik_max_word",
					"copy_to":"full_text"
				},
				"p_b":{
					"type":"text",
					"analyzer":"ik_max_word",
					"search_analyzer":"ik_max_word",
					"copy_to":"full_text"
				},
				"p_c":{
					"type":"text",
					"analyzer":"ik_max_word",
					"search_analyzer":"ik_max_word",
					"copy_to":"full_text"
				},
				"full_text":{
					"type":"text"
					"analyzer":"ik_max_word",
					"search_analyzer":"ik_max_word"
				}
			}
		}
	
	}`
	ok, err := NewIndexWithMapping("news", typeName, mapping)
	assert.True(t, ok, "")
	fmt.Printf("%+v", err)
}

func TestIsIndexExisted(t *testing.T) {
	InitialESConfig(esServer, true)
	ok, err := IsIndexExisted(nil, "news")
	assert.True(t, ok)
	fmt.Printf("%+v", err)
}

func TestUpdateMapping(t *testing.T) {
	InitialESConfig(esServer, true)
	mapping := `{
		"new":{
			"properties":{
				"id":{
					"type":"keyword"
				},
				"p_a":{
					"type":"text",
					"analyzer":"ik_max_word",
					"search_analyzer":"ik_max_word"
				},
				"p_b":{
					"type":"text",
					"analyzer":"ik_max_word",
					"search_analyzer":"ik_max_word"
				},
				"p_c":{
					"type":"text",
					"analyzer":"ik_max_word",
					"search_analyzer":"ik_max_word"
				}
			}
		}
	}`
	client, err := GetSLClient()
	if err != nil {
		t.Errorf("%+v", err)
	}
	ok, err := UpdateMapping(client, "news", "new", mapping)
	assert.True(t, ok)
	fmt.Printf("%+v", err)
}

func TestDeleteIndex(t *testing.T) {
	InitialESConfig(esServer, true)
	client, err := GetSLClient()
	if err != nil {
		t.Errorf("%+v", err)
	}
	result, err := elastic.NewIndicesDeleteService(client).Index([]string{"news"}).Do(context.Background())
	if err != nil {
		t.Errorf("%+v", err)
	}
	assert.True(t, result.Acknowledged)
}

func TestInsertData(t *testing.T) {
	datas := []News{
		News{1, "阿斯顿发送到", "阿斯蒂芬", "请问是大法官发"},
		News{2, "退开奖号狂欢节", "安慰我而过 ", "集合uehfbnv"},
		News{3, "去玩儿群翁而且问题", "一夜润体乳请问", "把你们，胡椒粉统一"},
		News{4, "其二冬虫夏草标的", "安慰法行政总厨不在", "宣传部的人如果"},
		News{5, "阿斯顿发送到", " 玩儿我", "驱蚊器翁群翁群无"},
		News{6, "阿斯顿发送到", "疼我二哥个", "下次是第三方个人工"},
		News{7, "全国性交流会议在武汉举办", "疼我二哥个", "下次是第三方个人工"},
	}
	InitialESConfig(esServer, true)
	client, err := GetSLClient()
	if err != nil {
		t.Errorf("%+v", err)
	}
	for _, v := range datas {
		ok, id, err := InsertData(client, "news", "new", &v, string(v.Id))
		if err != nil {
			t.Errorf("%+v", err)
		}
		if !ok {
			t.Error("insert failed.")
		}
		t.Log(id)
	}
}

func TestSelectById(t *testing.T) {
	InitialESConfig(esServer, true)
	client, err := GetSLClient()
	if err != nil {
		t.Errorf("%+v", err)
	}
	getResult, err := client.Get().Index("news").Type("article").Id("100374").Do(context.TODO())
	if err != nil {
		t.Errorf("%+v", err)
	}
	t.Logf("%+v", getResult)
}

func TestCatIndex(t *testing.T) {
	InitialESConfig(esServer, true)
	client, err := GetSLClient()
	if err != nil {
		t.Errorf("%+v", err)
	}
	getResult, err := elastic.NewCatIndicesService(client).Do(context.TODO())
	datas := []elastic.CatIndicesResponseRow(getResult)
	for _, v := range datas {
		t.Log(v.Index)
	}
}

func TestCommonSearch(t *testing.T) {
	q := elastic.NewCommonTermsQuery("p_a", "阿送到")
	InitialESConfig(esServer, true)
	client, err := GetSLClient()
	if err != nil {
		t.Errorf("%+v", err)
	}
	countResult, err := elastic.NewCatCountService(client).Index("news").Do(context.TODO())
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", countResult)

	searchResult, err := client.Search().Index("news").Query(q).Size(10).Do(context.TODO())
	if err != nil {
		t.Error(err)
	}

	for _, hit := range searchResult.Hits.Hits {
		var item News
		err := json.Unmarshal(*hit.Source, &item)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%+v", item)
	}
}

type News struct {
	Id int    `json:"id"`
	PA string `json:"p_a"`
	PB string `json:"p_b"`
	PC string `json:"p_c"`
}
