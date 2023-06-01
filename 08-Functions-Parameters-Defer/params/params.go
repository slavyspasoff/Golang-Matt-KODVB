package params

func Do(a *map[int]int) int {
	(*a)[0] = 0
	return (*a)[0]
}
