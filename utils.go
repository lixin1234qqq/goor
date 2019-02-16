package openrasp

func truncate(s string, n int) string {
	var j int
	for i := range s {
		if j == n {
			return s[:i]
		}
		j++
	}
	return s
}

func truncateString(s string, n int) string {
	return truncate(s, n)
}
