package main

func main() {
	valid := isValidByStack("(])")
	println(valid)
}
func isValidByStack(s string) bool {
	var smallBraces []string
	for _, c := range s {
		if string(c) == "(" || string(c) == "[" || string(c) == "{" {
			smallBraces = append(smallBraces, string(c))
		} else {
			if len(smallBraces) == 0 {
				return false
			}
			top := smallBraces[len(smallBraces)-1]
			if (top == "(" && string(c) == ")") ||
				(top == "[" && string(c) == "]") ||
				(top == "{" && string(c) == "}") {
				smallBraces = (smallBraces)[:len(smallBraces)-1]
			} else {
				return false
			}
		}
	}

	return len(smallBraces) == 0
}
