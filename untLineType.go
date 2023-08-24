package datatypes

/*
 * @Description:
 * @version:
 * @Author: devuser@gmail.com
 * @Date: 2023-08-24 16:27:07
 * @LastEditors: devuser@gmail.com
 * @LastEditTime: 2023-08-24 16:27:12
 */

// AppendLineTypeMap[T any] 行类型字典，扩展行类型字典，指定缺省键值
//
//	@Description:
//	@param rawMap
//	@param defaultValue
//	@param lineTypes
//	@return rlt
func AppendLineTypeMap[T any](rawMap map[string]T, defaultValue T, lineTypes ...string) (rlt map[string]T) {
	if len(lineTypes) == 0 {
		rlt = rawMap
		return
	}
	rlt = make(map[string]T, len(rawMap)+len(lineTypes))
	for k, v := range rawMap {
		rlt[k] = v
	}

	for _, lineType := range lineTypes {
		rlt[lineType] = defaultValue
	}
	return
}
