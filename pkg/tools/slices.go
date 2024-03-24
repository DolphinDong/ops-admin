package tools

// RemoveDuplication[T any]
//
//	@Description: 删除重复的值，并且保持原有的顺序
//	@param arr
//	@return []T
func RemoveDuplication[T any](arr []T) []T {
	set := make(map[any]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}
	return arr[:j]
}
