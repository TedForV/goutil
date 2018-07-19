package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"strings"
)

var esConfig ESConfig

func InitialESConfig(serverUrls []string) {
	esConfig = ESConfig{
		ServerUrls: serverUrls,
	}
}

func NewIndexWithMapping(indexName string, typeName string, mapping string) (bool, error) {
	client, err := GetESClient()
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
		client, err = GetESClient()
		if err != nil {
			return false, err
		}
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
		client, err = GetESClient()
		if err != nil {
			return false, err
		}
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

func InsertData(client elastic.Client, indexName string, typeName string, item interface{}, id string) (bool, int, error) {
	//result, err := client.Index().Index("news").Type(typeName).Id(id).BodyJson(item).Do(context.TODO())

}

func GetESClient() (*elastic.Client, error) {
	return elastic.NewClient(elastic.SetURL(esConfig.ServerUrls...))
}

type ESConfig struct {
	ServerUrls []string
}
