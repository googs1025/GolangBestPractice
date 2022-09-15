package Testing_practice

func networkCompute(a, b int) (int, error) {
	c := a + b

	return c, nil
}

func Compute(a, b int) (int, error) {
	sum, err := networkCompute(a, b)

	return sum, err
}


