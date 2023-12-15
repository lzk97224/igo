package islice

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

func ToMap[T any, K comparable, V any](data []T, k func(T) K, v func(T) V) map[K]V {
	result := make(map[K]V)
	for _, item := range data {
		result[k(item)] = v(item)
	}
	return result
}

func GroupToMap[T any, K comparable](data []T, k func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, item := range data {
		key := k(item)
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
