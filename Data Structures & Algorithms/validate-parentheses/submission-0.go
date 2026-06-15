func isValid(s string) bool {
    var stack []byte

	for i := 0 ; i < len(s); i++{
		if len(stack) == 0{
			stack = append(stack, s[i])
		}else if s[i] == '(' || s[i] == '{' || s[i] == '['{
			stack = append(stack, s[i])
		}else if s[i] == ')'{
			top := stack[len(stack)-1]
			if top == '('{
				stack = stack[:len(stack)-1]
			}else{
				return false
			}
		}else if s[i] == '}'{
			top := stack[len(stack)-1]
			if top == '{'{
				stack = stack[:len(stack)-1]
			}else{
				return false
			}
			
		}else if s[i] == ']'{
			top := stack[len(stack)-1]
			if top == '['{
				stack = stack[:len(stack)-1]
			}else{
				return false
			}
		}
	}

	return len(stack) == 0
}
