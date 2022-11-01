package leetcode

func MaxCharInString(s string) ([]string, int) {
	ans := make(map[string]int)
	for _, s := range s {
		ans[string(s)] = ans[string(s)] + 1
	}
	max := 0
	fas := make([]string, 10)
	for s, v := range ans {
		if v > 0 && v > max {
			max = v
			fas = fas[:0]
			fas = append(fas, s)
			continue
		}
		if v == max {
			fas = append(fas, s)
		}
	}
	return fas, max
}
