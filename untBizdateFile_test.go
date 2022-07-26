package datatypes

/*
 * @Descripttion:
 * @version:
 * @Author: devuser@gmail.com
 * @Date: 2022-07-26 15:58:56
 * @LastEditors: devuser@gmail.com
 * @LastEditTime: 2022-07-26 15:59:01
 */

import "testing"

func Test_GetSimpleBizdateFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name        string
		args        args
		wantBizdate string
		wantErr     bool
	}{
		{
			name: "test-GetSimpleBizdateFromFile-A",
			args: args{
				filename: "A.B.20220725.txt",
			},
			wantBizdate: "20220725",
			wantErr:     false,
		},
		{
			name: "test-GetSimpleBizdateFromFile-B",
			args: args{
				filename: "./A.B.20220725.txt",
			},
			wantBizdate: "20220725",
			wantErr:     false,
		},
		{
			name: "test-GetSimpleBizdateFromFile-C",
			args: args{
				filename: ".202207250.P",
			},
			wantBizdate: "",
			wantErr:     true,
		},
		{
			name: "test-GetSimpleBizdateFromFile-D",
			args: args{
				filename: "./A.B.202207250.txt",
			},
			wantBizdate: "",
			wantErr:     true,
		},

		{
			name: "test-GetSimpleBizdateFromFile-E",
			args: args{
				filename: ".202207250.P",
			},
			wantBizdate: "",
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBizdate, err := GetSimpleBizdateFromFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSimpleBizdateFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotBizdate != tt.wantBizdate {
				t.Errorf("GetSimpleBizdateFromFile() = %v, want %v", gotBizdate, tt.wantBizdate)
			}
		})
	}
}
