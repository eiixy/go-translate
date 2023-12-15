package translate

import (
	"reflect"
	"testing"
)

func TestClient_Translates(t *testing.T) {
	type args struct {
		texts []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"", args{[]string{"hello", "world"}}, []string{"你好", "世界"}, false},
	}
	r := NewClient()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.Translates(tt.args.texts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Translates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Translates() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Translate(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"", args{"hello world"}, "你好世界", false},
	}
	r := NewClient()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.Translate(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Translate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
