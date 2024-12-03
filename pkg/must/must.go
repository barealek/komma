package must

func Must[e any](t e, err error) e {
	if err != nil {
		panic(err)
	}

	return t
}
