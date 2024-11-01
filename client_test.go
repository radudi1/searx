package searx

import (
	"testing"
)

func TestClient_Search(t *testing.T) {
	type fields struct {
		Url           string
		SearchOptions SearchOptions
	}
	type args struct {
		query   string
		options *SearchOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestSearchWithDefaultOptions",
			fields: fields{
				Url: "http://172.31.88.88:8080",
			},
			args: args{
				query:   "Go programming language",
				options: nil,
			},
			wantErr: false,
		},
		{
			name: "TestSearchWithAllOptionsSet",
			fields: fields{
				Url: "http://172.31.88.88:8080",
			},
			args: args{
				query: "Go programming language",
				options: &SearchOptions{
					Categories: []string{"general", "science"},
					Engines:    []string{"duckduckgo", "qwant"},
					Language:   "en",
					TimeRange:  "month",
					SafeSearch: 2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Url:           tt.fields.Url,
				SearchOptions: tt.fields.SearchOptions,
			}
			got, err := c.Search(tt.args.query, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got.Results) == 0 {
				t.Errorf("Client.Search() did not return any search results")
			}
		})
	}
}
