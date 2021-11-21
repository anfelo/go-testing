package filterunique

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(dd []Developer) map[string]int {
	check := make(map[string]int)
	for _, d := range dd {
		check[d.Name] = 1
	}
	return check
}
