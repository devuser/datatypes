package datatypes

/*
 * @Descripttion:
 * @version:
 * @Author: devuser@gmail.com
 * @Wechat: PythonBJ
 * @Mobile: 13911223211
 * @Date: 2021-10-14 00:28:55
 * @LastEditors: devuser@gmail.com
 * @LastEditTime: 2021-10-14 00:33:31
 */

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

type BizdateTime time.Time

func (date *BizdateTime) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = BizdateTime(nullTime.Time)
	return
}

func (date BizdateTime) Value() (driver.Value, error) {
	y, m, d := time.Time(date).Date()
	hour := time.Time(date).Hour()
	min := time.Time(date).Minute()
	sec := time.Time(date).Second()
	return time.Date(y, m, d, hour, min, sec, 0, time.Time(date).Location()), nil
}

func (date BizdateTime) String() string {
	return time.Time(date).Format("2006-01-02 15:04:05")
}

// GormDataType gorm common data type
func (date BizdateTime) GormDataType() string {
	return "datetime"
}

func (date BizdateTime) GobEncode() ([]byte, error) {
	return time.Time(date).GobEncode()
}

func (date *BizdateTime) GobDecode(b []byte) error {
	return (*time.Time)(date).GobDecode(b)
}

func (date BizdateTime) MarshalJSON() ([]byte, error) {
	return time.Time(date).MarshalJSON()
}

func (date *BizdateTime) UnmarshalJSON(b []byte) error {
	return (*time.Time)(date).UnmarshalJSON(b)
}
