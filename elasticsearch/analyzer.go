package elasticsearch

import (
	"context"

	"github.com/olivere/elastic"
	"github.com/pkg/errors"
)

// AnalyzerText will seperate the content into a slice of
// meaning words with original order in content
// the default seperate mode is 'ik_smart'
func AnalyzerText(client elastic.Client, content string) ([]string, error) {
	res, err := client.IndexAnalyze().Analyzer(IKSmart).Text(content).Do(context.TODO())
	if err != nil {
		return nil, errors.Wrap(err, "analyzer text failed.")
	}
	token := make([]string, len(res.Tokens))
	for i, v := range res.Tokens {
		token[i] = v.Token
	}
	return token, nil
}
