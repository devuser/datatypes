package datatypes

/*
 * @Descripttion:
 * @version:
 * @Author: devuser@gmail.com
 * @Wechat: PythonBJ
 * @Mobile: 13911223211
 * @Date: 2021-11-15 19:41:10
 * @LastEditors: devuser@gmail.com
 * @LastEditTime: 2022-01-12 14:08:50
 */

import (
	"encoding/xml"
	"testing"
)

// 2021-11-15 12:02:25.0
func TestBirtTime_MarshalXML(t *testing.T) {
	type args struct {
		e     *xml.Encoder
		start xml.StartElement
	}
	tests := []struct {
		name     string
		instance BirtTime
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instance := BirtTime{}
			if err := instance.MarshalXML(tt.args.e, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("BirtTime.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
