func isPalindrome(str string) bool {
	n := len(str)
	low := 0
	high := n-1

	s := []rune(str)

	for low < high{
		if !(unicode.IsDigit(s[low]) || unicode.IsLetter(s[low])){
			low++
		}else if !(unicode.IsDigit(s[high]) || unicode.IsLetter(s[high])){
			high--
		}else if unicode.ToLower(s[low]) != unicode.ToLower(s[high]){
			return false
		}else{
			low++
			high--
		}
	}

	return true
}
