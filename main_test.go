package main

import "testing"

func Test_extractName(t *testing.T) {
	type args struct {
		URL string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "url", args: args{URL: "https://github.com/tillknuesting/rmclone"}, want: "rmclone", wantErr: false},
		{name: "urlSuffix", args: args{URL: "https://github.com/tillknuesting/rmclone.git"}, want: "rmclone", wantErr: false},
		{name: "urlNotEnoughSlash", args: args{URL: "https://github.com/tillknuestingrmclone.git"}, want: "", wantErr: true},
		{name: "urlNotEnoughSlashNoURL", args: args{URL: "tillknuestingrmclone.git"}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractName(tt.args.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
