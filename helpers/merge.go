package helpers

// Merge merges two maps by adding src to dst.
func Merge(src, dst map[string]interface{}) {
	for k, v := range src {
		dst[k] = v
	}
}
