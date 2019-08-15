package String

func RepeatStrings(str string) []string {
	rlt := make([]string, 0)
	tmp := make(map[string]bool)
	for _, v := range str {
		s := string(v)
		if tmp[s] == true {
			rlt = append(rlt, s)
		}
		tmp[s] = true
	}

	return rlt
}
