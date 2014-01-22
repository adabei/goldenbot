package helpers

func Merge(src, dst map[string]interface{}) {
	for k, v := range src {
		dst[k] = v
	}
}
