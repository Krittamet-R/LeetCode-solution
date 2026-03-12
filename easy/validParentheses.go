package easy

func IsValid(s string) bool {
	a, b, c := 0, 0, 0
	pop := []string{}
	pop = append(pop, "zero")
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 40:
			a++
			pop = append(pop, "a")
		case 41:
			if pop[len(pop)-1] == "a"{
				a--
				pop = pop[:len(pop)-1]
			} else {
				return false
			}
		case 91:
			b++
			pop = append(pop, "b")
		case 93:
			if pop[len(pop)-1] == "b"{
				b--
				pop = pop[:len(pop)-1]
			} else {
				return false
			}
		case 123:
			c++
			pop = append(pop, "c")
		case 125:
			if pop[len(pop)-1] == "c"{
				c--
				pop = pop[:len(pop)-1]
			} else {
				return false
			}
		}
	}
	if a == 0 && b == 0 && c == 0 {
		return true
	} else {
		return false
	}
}
