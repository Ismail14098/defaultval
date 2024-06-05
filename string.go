package null

type String struct{ *string }

func StringFrom(val string) String {
	return String{&val}
}

func StringFromPtr(val *string) String {
	return String{val}
}

func (s String) ValueOrZero() string {
	if s.string == nil {
		return ""
	}

	return *s.string
}

func (s String) ValueOrDefault(def string) string {
	if s.string == nil {
		return def
	}

	return *s.string
}
