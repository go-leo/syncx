package syncx

func BraveGo(f func(), r func(p any)) {
	go BraveDo(f, r)
}

func BraveDo(f func(), r func(p any)) {
	fe := func() error {
		f()
		return nil
	}
	_ = BraveDoE(fe, r)
}

func BraveDoE(f func() error, r func(p any)) error {
	defer func() {
		// if r is nil, which means panics are not recovered.
		if r == nil {
			return
		}
		if p := recover(); p != nil {
			r(p)
		}
	}()
	return f()
}
