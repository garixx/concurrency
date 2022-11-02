package main

import (
	"reflect"
	"testing"
)

func TestJSONEncode(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "wrong non-struct parameter",
			args:    args{v: "string"},
			want:    []byte{},
			wantErr: true,
		},
		{
			name:    "empty struct parameter",
			args:    args{v: City{}},
			want:    []byte(`{"Name": "", "Population": 0, "GDP": 0, "Mayor": ""}`),
			wantErr: false,
		},
		{
			name:    "User struct parameter",
			args:    args{v: User{"bob", 10}},
			want:    []byte(`{"Name": "bob", "Age": 10}`),
			wantErr: false,
		},
		{
			name:    "City struct parameter",
			args:    args{v: City{"sf", 5000000, 567896, "mr jones"}},
			want:    []byte(`{"Name": "sf", "Population": 5000000, "GDP": 567896, "Mayor": "mr jones"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JSONEncode(tt.args.v)
			//fmt.Println("expected:", string(tt.want), ";\n  actual:", string(got))
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONEncode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
