package islice

import "sort"

func IsEmpty[T any](t []T) bool {
	return len(t) <= 0
}

func IsNotEmpty[T any](t []T) bool {
	return !IsEmpty(t)
}

// Filter 过滤掉切片中满足条件的项目
func Filter[T any](data []T, fun func(item T) bool) []T {
	removeIndex := make([]int, 0, 0)
	for i, item := range data {
		if fun(item) {
			removeIndex = append(removeIndex, i)
		}
	}

	for i := len(removeIndex) - 1; i >= 0; i-- {
		data = append(data[:removeIndex[i]], data[removeIndex[i]+1:]...)
	}
	return data
}

func EquallyDivide[T any](size int, arr []T) [][]T {
	if size <= 0 {
		return [][]T{arr}
	}
	result := make([][]T, 0, len(arr)/size+1)
	groupNum := len(arr) / size

	begin := 0
	for i := 0; i <= groupNum; i++ {
		end := begin + size

		if begin == len(arr) {
			break
		}
		if begin+size > len(arr) {
			result = append(result, arr[begin:])
			break
		}

		result = append(result, arr[begin:end])
		begin = end
	}
	return result
}

func ToMap[T any, K comparable, V any](data []T, t2k func(T) K, t2v func(T) V) map[K]V {
	result := make(map[K]V)
	for _, item := range data {
		result[t2k(item)] = t2v(item)
	}
	return result
}

func ToMapByKey[T any, K comparable](data []T, t2k func(T) K) map[K]T {
	return ToMap(data, t2k, func(t T) T {
		return t
	})
}

func GroupToMap[T any, K comparable](data []T, t2k func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, item := range data {
		key := t2k(item)
		if _, ok := result[key]; !ok {
			result[key] = make([]T, 0)
		}
		result[key] = append(result[key], item)
	}
	return result
}

func ToNewSlice[T any, NT any](data []T, fun func(T) NT) []NT {
	result := make([]NT, 0)
	for _, item := range data {
		result = append(result, fun(item))
	}
	return result
}

func SortByKeySlice[T any, K comparable](data []T, keySlice []K, t2k func(T) K) []T {
	//result := make([]T, 0)
	//mapByKey := ToMapByKey(data, t2k)

	keySliceIndexMap := map[K]int{}
	for index, k := range keySlice {
		if _, ok := keySliceIndexMap[k]; !ok {
			keySliceIndexMap[k] = index
		}
	}

	sort.Slice(data, func(i, j int) bool {
		return keySliceIndexMap[t2k(data[i])] < keySliceIndexMap[t2k(data[j])]
	})

	return data
	//for _, k := range keySlice {
	//	if t1, ok := mapByKey[k]; ok {
	//		result = append(result, t1)
	//	}
	//}
	//return result
}
