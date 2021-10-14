package datatypes

/*
 * @Descripttion:
 * @version:
 * @Author: devuser@gmail.com
 * @Wechat: PythonBJ
 * @Mobile: 13911223211
 * @Date: 2021-10-13 22:53:56
 * @LastEditors: devuser@gmail.com
 * @LastEditTime: 2021-10-14 01:08:20
 */

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
)

type Bizdate time.Time

func (date *Bizdate) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = Bizdate(nullTime.Time)
	return
}

func (date Bizdate) Value() (driver.Value, error) {
	y, m, d := time.Time(date).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Time(date).Location()), nil
}

// GormDataType gorm common data type
func (date Bizdate) GormDataType() string {
	return "datetime"
}

func (date Bizdate) GobEncode() ([]byte, error) {
	return time.Time(date).GobEncode()
}

func (date *Bizdate) GobDecode(b []byte) error {
	return (*time.Time)(date).GobDecode(b)
}

func (date Bizdate) MarshalJSON() ([]byte, error) {
	// return time.Time(date).MarshalJSON()
	v := time.Time(date).Format("2006-01-02")
	log.Println(v)
	return json.RawMessage([]byte(v)).MarshalJSON()
}

func (date *Bizdate) UnmarshalJSON(b []byte) error {
	// return (*time.Time)(date).UnmarshalJSON(b)
	t, err := time.ParseInLocation("2006-01-02", string(b), time.Local)
	*date = Bizdate(t)
	return err
}
