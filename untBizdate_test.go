package datatypes

/*
 * @Descripttion:
 * @version:
 * @Author: devuser@gmail.com
 * @Date: 2021-10-13 23:08:51
 * @LastEditors: devuser@gmail.com
 * @LastEditTime: 2021-10-14 01:03:34
 */

import (
	"reflect"
	"testing"
	"time"
)

func TestBizdate_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	dtNow := Bizdate(time.Now())
	tests := []struct {
		name    string
		date    *Bizdate
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test-Bizdate-A",
			date: &dtNow,
			args: args{
				b: []byte("2021-10-13"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date := &Bizdate{}
			if err := date.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Bizdate.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBizdate_MarshalJSON(t *testing.T) {
	dtNow := Bizdate(time.Now())
	tests := []struct {
		name    string
		date    Bizdate
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test-MarshalJSON-A",
			date:    dtNow,
			want:    []byte(time.Now().Format("2006-01-02")),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.date.MarshalJSON()
			t.Log(string(got))
			t.Log(string(tt.want))
			if (err != nil) != tt.wantErr {
				t.Errorf("Bizdate.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bizdate.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
