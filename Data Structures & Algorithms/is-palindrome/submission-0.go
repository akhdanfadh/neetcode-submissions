func isPalindrome(s string) bool {
    // since only ascii
    var sb strings.Builder
    for _, c := range s {
        if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
            sb.WriteRune(c)
        }
    }
    s = strings.ToLower(sb.String())

    left, right := 0, len(s)-1
    for right > left {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}
