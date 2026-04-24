func isValid(s string) bool {
	arr := make([]rune, 0, len(s))
	for _, c := range s {
		if len(arr) > 0 {
			prev := arr[len(arr)-1]
			if prev == '(' && c == ')' ||
			   prev == '{' && c == '}' ||
			   prev == '[' && c == ']' {
				arr = arr[:len(arr)-1]
				continue
			 }
		}
		arr = append(arr, c)
	}
	return len(arr) == 0
}
