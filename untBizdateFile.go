package datatypes

/*
 * @Descripttion:
 * @version:
 * @Author: devuser@gmail.com
 * @Date: 2022-07-26 15:55:44
 * @LastEditors: devuser@gmail.com
 * @LastEditTime: 2022-07-26 16:02:00
 */

import (
	"fmt"
	"path/filepath"
	"regexp"
)

// GetSimpleBizdateFromFile 识别文件名称中前后点号中间的8位阿拉伯数字作为业务日期
func GetSimpleBizdateFromFile(filename string) (bizdate string, err error) {
	//
	regexpX := regexp.MustCompile(`\.(\d{8})\.`)
	if arr := regexpX.FindStringSubmatch(filepath.Base(filename)); len(arr) == 2 {
		bizdate = arr[1]
	} else {
		err = fmt.Errorf("can not find the bizdate from file %s", filename)
	}
	return
}
