package imap

func IsEmpty[K comparable, V any](m map[K]V) bool {
	return len(m) <= 0
}
func IsNotEmpty[K comparable, V any](m map[K]V) bool {
	return !IsEmpty(m)
}
