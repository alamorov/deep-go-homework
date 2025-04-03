package main

func Map[T any](data []T, action func(T) T) []T {
	if len(data) == 0 {
		return data
	}

	result := make([]T, len(data))
	for i, v := range data {
		result[i] = action(v)
	}
	return result
}

func Filter[T any](data []T, action func(T) bool) []T {
	if len(data) == 0 {
		return data
	}

	result := make([]T, 0, len(data))
	for _, v := range data {
		if action(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T any, R any](data []T, initial R, action func(R, T) R) R {
	if len(data) == 0 {
		return initial
	}

	result := initial
	for _, v := range data {
		result = action(result, v)
	}
	return result
}
