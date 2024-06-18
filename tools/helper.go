package tools

func GetFirstValue[T, U any](val T, _ U) T {
	return val
}
