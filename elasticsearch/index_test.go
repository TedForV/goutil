package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/stretchr/testify/assert"
	"testing"
)

var esServer = []string{"http://127.0.0.1:9200"}

func TestNewIndexWithMapping(t *testing.T) {
	//esServer := []string{"http://127.0.0.1:9200"}
	InitialESConfig(esServer)
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
	ok, err := NewIndexWithMapping("news", typeName, mapping)
	assert.True(t, ok, "")
	fmt.Printf("%+v", err)
}

func TestIsIndexExisted(t *testing.T) {
	InitialESConfig(esServer)
	ok, err := IsIndexExisted(nil, "news")
	assert.True(t, ok)
	fmt.Printf("%+v", err)
}

func TestUpdateMapping(t *testing.T) {
	InitialESConfig(esServer)
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
	ok, err := UpdateMapping(nil, "news", "new", mapping)
	assert.True(t, ok)
	fmt.Printf("%+v", err)
}

func TestDeleteIndex(t *testing.T) {
	InitialESConfig(esServer)
	client, err := GetESClient()
	if err != nil {
		t.Errorf("%+v", err)
	}
	result, err := elastic.NewIndicesDeleteService(client).Index([]string{"news"}).Do(context.Background())
	if err != nil {
		t.Errorf("%+v", err)
	}
	assert.True(t, result.Acknowledged)
}
