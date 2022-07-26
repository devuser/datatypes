package datatypes

/*
 * @Descripttion:
 * @version:
 * @Author: devuser@gmail.com

 * @Date: 2021-08-10 11:15:00
 * @LastEditors: devuser@gmail.com
 * @LastEditTime: 2022-01-11 14:17:24
 */

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"errors"
	"reflect"

	log "github.com/sirupsen/logrus"
)

// IntBool 位布尔型实例
type IntBool struct {
	Int int64
}

// NewIntBool 创建位布尔型实例
func NewIntBool(v bool) IntBool {
	if v {
		return IntBool{
			Int: -1,
		}
	}
	return IntBool{
		Int: 0,
	}
}

// Scan implements the Scanner interface.
func (instance *IntBool) Scan(value interface{}) (err error) {
	// 如果 value 为空，则认为是 false
	if value == nil {
		instance.Int = 0
		return nil
	}
	// log.Infof("Scan value: %v\n", value)
	// log.Infof("Scan value: %q\n", value)
	// log.Infof("Scan value: %q\n", reflect.TypeOf(value))
	switch reflect.TypeOf(value).Kind() {
	case reflect.Array:
		{
			if v, ok := value.([]uint8); !ok {
				instance.Int = 0
				err = errors.New("failed to convert int to bool")
			} else {
				// log.Infof("Scan convert: %v\n", v)
				// intV, err := strconv.ParseInt(string(v), 10, 64)
				if err == nil && len(v) > 0 && v[0] != 0 {
					instance.Int = -1
				} else {
					instance.Int = 0
				}
			}
		}
	case reflect.Int64:
		{
			if v, ok := value.(int64); !ok {
				instance.Int = 0
				err = errors.New("failed to convert int to bool")
			} else {
				// log.Infof("Scan convert: %v\n", v)
				// intV, err := strconv.ParseInt(string(v), 10, 64)
				if v != 0 {
					instance.Int = -1
				}
			}

		}
	default:
		{
			instance.Int = 0
		}
	}
	return
}

// Value implements the driver Valuer interface.
func (instance IntBool) Value() (driver.Value, error) {

	return int64(instance.Int), nil
}

// MarshalJSON 打包JSON
func (instance IntBool) MarshalJSON() ([]byte, error) {
	j := "false"
	// log.Infof("MarshalJSON: %v\n", instance)
	if instance.Int != 0 {
		j = "true"
	}
	return []byte(j), nil
}

// MarshalXML 打包JSON
func (instance IntBool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	j := "false"
	// log.Infof("MarshalXML: %v\n", instance)
	if instance.Int != 0 {
		j = "true"
	}
	return e.EncodeElement(j, start)
}

// UnmarshalJSON 解析JSON
func (s *IntBool) UnmarshalJSON(data []byte) error {
	log.Infof("data: %s", string(data))
	var rawBool bool
	// rawBool := struct {
	// 	Hidden bool `json:"hidden"`
	// }{
	// 	//
	// }

	if err := json.Unmarshal(data, &rawBool); err != nil {
		return err
	}
	if rawBool {
		s.Int = -1
	} else {
		s.Int = 0
	}

	return nil
}
