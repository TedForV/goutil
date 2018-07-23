package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"strings"
)

var esConfig ESConfig
var isInitaled bool

func InitialESConfig(serverUrls []string, forceReload bool) {
	if !isInitaled || forceReload {
		esConfig = ESConfig{
			ServerUrls: serverUrls,
		}
	}
}

func NewIndexWithMapping(indexName string, typeName string, mapping string) (bool, error) {
	client, err := GetSLClient()
	if err != nil {
		return false, errors.Wrap(err, "create client failed")
	}

	ok, err := IsIndexExisted(client, indexName)

	if ok || err != nil {
		return false, errors.Wrap(err, fmt.Sprintf("index (%s) is already existed. ", indexName))
	}

	result, err := client.CreateIndex(indexName).Do(context.Background())
	if err != nil {
		return false, errors.Wrap(err, "create index failed.")
	}
	if !result.Acknowledged {
		return false, errors.New("create index failed.")
	}

	return UpdateMapping(client, indexName, typeName, mapping)

}

func IsIndexExisted(client *elastic.Client, indexName string) (bool, error) {
	var err error
	if client == nil {
		return false, errors.New("client is nil.")
	}
	ok, err := client.IndexExists(indexName).Do(context.Background())
	if err != nil {
		return false, errors.Wrap(err, "check index failed")
	}
	if ok {
		return true, nil
	}
	return false, nil
}

func UpdateMapping(client *elastic.Client, indexName string, typeName string, mapping string) (bool, error) {
	//if mapping is empty, then return
	if len(strings.Trim(mapping, " ")) == 0 {
		return true, nil
	}
	var err error
	if client == nil {
		return false, errors.New("client is nil.")
	}

	result, err := client.PutMapping().Index(indexName).Type(typeName).BodyString(mapping).Do(context.Background())
	if err != nil {
		return false, errors.Wrap(err, "update mapping failed.")
	}
	if !result.Acknowledged {
		return false, errors.New("update mapping failed.")
	}
	return true, nil
}

func InsertData(client elastic.Client, indexName string, typeName string, item interface{}, id string) (bool, string, error) {
	result, err := client.Index().Index(indexName).Type(typeName).Id(id).BodyJson(item).Do(context.TODO())
	if err != nil {
		return false, "-1", errors.Wrap(err, "Insert failed.")
	}
	if result == nil {
		return false, "-1", errors.New("Insert failed.")
	}
	return true, result.Id, nil
}

// GetLLClient returns a long-lived client
func GetLLClient() (*elastic.Client, error) {
	return elastic.NewClient(elastic.SetURL(esConfig.ServerUrls...))
}

// GetSLClient return a short-lived client
func GetSLClient() (*elastic.Client, error) {
	return elastic.NewSimpleClient(elastic.SetURL(esConfig.ServerUrls...))
}

type ESConfig struct {
	ServerUrls []string
}
