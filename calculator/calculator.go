package calculator

func calculate(s string) int {
	stack := []bool{true}
	sign := true
	n, r := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		switch s[i] {
		case '(':
			stack = append(stack, true)
		case '+', '-', ')':
			if sign {
				r += n
			} else {
				r -= n
			}
			n = 0
			sign = sign == stack[len(stack)-1]
			switch s[i] {
			case '+':
				stack[len(stack)-1] = true
			case '-':
				stack[len(stack)-1] = false
				sign = !sign
			default:
				stack = stack[:len(stack)-1]
			}
		default:
			n = n*10 + int(s[i]-'0')
		}
	}

	if sign {
		return r + n
	}
	return r - n
}
