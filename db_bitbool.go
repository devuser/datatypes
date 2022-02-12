package datatypes

/*
 * @Descripttion:
 * @version:
 * @Author: devuser@gmail.com
 * @Wechat: PythonBJ
 * @Mobile: 13911223211
 * @Date: 2021-08-10 11:15:00
 * @LastEditors: devuser@gmail.com
 * @LastEditTime: 2022-02-12 18:17:03
 */

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// BitBool 位布尔型实例
type BitBool struct {
	Bool bool
}

// NewBitBool 创建位布尔型实例
func NewBitBool(v bool) *BitBool {
	return &BitBool{
		Bool: v,
	}
}

// Scan implements the Scanner interface.
func (instance *BitBool) Scan(value interface{}) (err error) {
	// 如果 value 为空，则认为是 false
	if value == nil {
		instance.Bool = false
		return nil
	}
	// log.Infof("%v\n", value)
	if v, ok := value.([]byte); !ok {
		instance.Bool = false
		err = errors.New("failed to convert bit to bool")
	} else {
		instance.Bool = (len(v) >= 0) && (int(v[0]) != 0)
	}
	return
}

// Value implements the driver Valuer interface.
func (instance BitBool) Value() (driver.Value, error) {

	return instance.Bool, nil
}

// MarshalJSON 打包JSON
func (instance BitBool) MarshalJSON() ([]byte, error) {
	j := "false"
	// log.Infof("MarshalJSON: %v\n", instance)
	if instance.Bool {
		j = "true"
	}
	return []byte(j), nil
}

// UnmarshalJSON 解析JSON
func (s *BitBool) UnmarshalJSON(data []byte) error {
	// log.Infof("data: %s", string(data))
	var rawBool bool
	// rawBool := struct {
	// 	Hidden bool `json:"hidden"`
	// }{
	// 	//
	// }

	if err := json.Unmarshal(data, &rawBool); err != nil {
		return err
	}
	s.Bool = rawBool

	return nil
}
