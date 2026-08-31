package base62

func Encode(n int64) string {
	const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if n == 0 {
		return string(alphabet[0])
	}

	result := []byte{}
	for n > 0 {
		result = append([]byte{alphabet[n%62]}, result...)
		n /= 62
	}
	return string(result)
}
