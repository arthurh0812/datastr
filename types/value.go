package types

// Value groups together a set of values of comparable data types
type Value interface {
	IsEqualTo(v Value) bool
	IsGreaterThan(v Value) bool
	IsLessThan(v Value) bool
}

// STRING

type String string

func (s String) IsEqualTo(v Value) bool {
	if vs, ok := v.(String); ok {
		return s == vs
	}
	return false
}

func (s String) IsGreaterThan(v Value) bool {
	if vs, ok := v.(String); ok {
		return vs < s
	}
	return false
}

func (s String) IsLessThan(v Value) bool {
	if vs, ok := v.(String); ok {
		return s < vs
	}
	return false
}

// INT

type Int int64

func (n Int) IsEqualTo(v Value) bool {
	if vn, ok := v.(Int); ok {
		return n == vn
	}
	return false
}

func (n Int) IsGreaterThan(v Value) bool {
	if vn, ok := v.(Int); ok {
		return vn < n
	}
	return false
}

func (n Int) IsLessThan(v Value) bool {
	if vn, ok := v.(Int); ok {
		return n < vn
	}
	return false
}

// UINT

type Uint uint64

func (u Uint) IsEqualTo(v Value) bool {
	if vu, ok := v.(Uint); ok {
		return u == vu
	}
	return false
}

func (u Uint) IsGreaterThan(v Value) bool {
	if vu, ok := v.(Uint); ok {
		return vu < u
	}
	return false
}

func (u Uint) IsLessThan(v Value) bool {
	if vu, ok := v.(Uint); ok {
		return u < vu
	}
	return false
}