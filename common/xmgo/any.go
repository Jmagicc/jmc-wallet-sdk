package xmgo

func AnyPtr[T any](v T) *T {
	return &v
}

type EnumInSliceType interface {
	String() string
}
type AnyInSliceType interface {
	string | int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | float32 | float64 | bool
}

func EnumInSlice[T EnumInSliceType](v T, slice []T) bool {
	for _, item := range slice {
		if item.String() == v.String() {
			return true
		}
	}
	return false
}

func AnyInSlice[T AnyInSliceType](v T, slice []T) bool {
	for _, item := range slice {
		if item == v {
			return true
		}
	}
	return false
}

func AnySet[T AnyInSliceType](list []T) []T {
	// 去重
	set := make(map[T]int)
	for i, item := range list {
		_, ok := set[item]
		if !ok {
			set[item] = i
		}
	}
	result := make([]T, 0)
	for item, _ := range set {
		result = append(result, item)
	}
	//排序 根据 set[item] 的值
	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if set[result[i]] > set[result[j]] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return result
}
