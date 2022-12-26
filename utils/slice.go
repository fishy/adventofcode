package utils

func Reverse[T any, S ~[]T](s S) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		s[i], s[j] = s[j], s[i]
	}
}
