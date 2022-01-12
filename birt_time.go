package datatypes

/*
 * @Descripttion:
 * @version:
 * @Author: devuser@gmail.com
 * @Wechat: PythonBJ
 * @Mobile: 13911223211
 * @Date: 2021-11-15 19:41:10
 * @LastEditors: devuser@gmail.com
 * @LastEditTime: 2022-01-12 14:06:03
 */

import (
	"encoding/xml"
	"fmt"
	"time"
)

// BirtTime 报表专用时间格式 2021-11-15 12:02:25.0
type BirtTime time.Time

// MarshalXML 支持XML编码
func (instance BirtTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const (
		zero    = "0001-01-01 00:00:00.0"
		dateFmt = "2006-01-02 15:04:05"
	)
	dateString := fmt.Sprintf("%s.0", time.Time(instance).Format(dateFmt))
	if dateString != zero {
		e.EncodeElement(dateString, start)
	} else {
		e.EncodeElement("", start)
	}

	return nil
}
