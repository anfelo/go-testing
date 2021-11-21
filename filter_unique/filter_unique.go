package filterunique

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(dd []Developer) []string {
	var u []string
	check := make(map[string]int)
	for _, d := range dd {
		check[d.Name] = 1
	}

	for name := range check {
		u = append(u, name)
	}
	return u
}
