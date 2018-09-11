package elasticsearch

import (
	"reflect"
	"testing"

	"github.com/olivere/elastic"
)

func TestAnalyzerText(t *testing.T) {
	InitialESConfig(esServer, true)
	client, err := GetSLClient()
	if err != nil {
		t.Error(err)
		return
	}
	type args struct {
		client  elastic.Client
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				client:  *client,
				content: "综合税费成本包括税收、社会保险费用、政府性基金和行政性收费、行政性灰色隐性寻租成本、办事难的制度性交易成本等。",
			},
			want:    []string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AnalyzerText(tt.args.client, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("AnalyzerText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnalyzerText() = %v, want %v", got, tt.want)
			}
		})
	}
}
