package null

type Float64 struct{ *float64 }

func Float64From(val float64) Float64 {
	return Float64{&val}
}

func Float64FromPtr(val *float64) Float64 {
	return Float64{val}
}

func (f Float64) ValueOrZero() float64 {
	if f.float64 == nil {
		return 0
	}

	return *f.float64
}

func (f Float64) ValueOrDefault(def float64) float64 {
	if f.float64 == nil {
		return def
	}

	return *f.float64
}

func (f Float64) Ptr() *float64 {
	return f.float64
}
