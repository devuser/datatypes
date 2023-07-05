package datatypes

/*
 * @Description:
 * @version:
 * @Author: devuser@gmail.com
 * @Date: 2023-07-05 11:04:51
 * @LastEditors: devuser@gmail.com
 * @LastEditTime: 2023-07-05 14:19:27
 */

// TSortableElement 支持排序因子的接口
// @Description:
type TSortableElement interface {
	GetSortingFactor() int
}

// ElementList 支持排序因子的实例切片
type ElementList []TSortableElement

// NewElementList
//
//	@Description:
//	@param n
//	@return ElementList
func NewElementList(n int) ElementList {
	return make([]TSortableElement, 0, n)
}

// Len
//
//	@Description:
//	@receiver el
//	@return int
func (el ElementList) Len() int {
	return len(el)
}

// Less
//
//	@Description:
//	@receiver el
//	@param i
//	@param j
//	@return bool
func (el ElementList) Less(i, j int) bool {
	return el[i].GetSortingFactor() < el[j].GetSortingFactor()
}

// Swap
//
//	@Description:
//	@receiver el
//	@param i
//	@param j
func (el ElementList) Swap(i, j int) {
	el[i], el[j] = el[j], el[i]
}
