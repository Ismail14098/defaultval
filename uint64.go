package null

type UInt64 struct{ *uint64 }

func UInt64From(val uint64) UInt64 {
	return UInt64{&val}
}

func UInt64FromPtr(val *uint64) UInt64 {
	return UInt64{val}
}

func (i UInt64) ValueOrZero() uint64 {
	if i.uint64 == nil {
		return 0
	}

	return *i.uint64
}

func (i UInt64) ValueOrDefault(def uint64) uint64 {
	if i.uint64 == nil {
		return def
	}

	return *i.uint64
}

func (i UInt64) Ptr() *uint64 {
	return i.uint64
}
