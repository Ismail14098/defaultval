package null

type Int struct{ *int }

func IntFrom(val int) Int {
	return Int{&val}
}

func IntFromPtr(val *int) Int {
	return Int{val}
}

func (i Int) ValueOrZero() int {
	if i.int == nil {
		return 0
	}

	return *i.int
}

func (i Int) ValueOrDefault(def int) int {
	if i.int == nil {
		return def
	}

	return *i.int
}

func (i Int) Ptr() *int {
	return i.int
}
