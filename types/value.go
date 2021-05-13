package types

// Value groups together a set of values of comparable data types
type Value interface {
	IsEqualTo(v Value) bool
	IsGreaterThan(v Value) bool
	IsLessThan(v Value) bool
}

func GetValue(val interface{}) Value {
	switch val.(type) {
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
		return GetInt(val)
	case uint:
	case uint8:
	case uint16:
	case uint32:
	case uint64:
		return GetUint(val)
	case string:
	case []byte:
		return GetString(val)
	}
	return nil
}

func GetString(s interface{}) String {
	switch s.(type) {
	case string:
	case []byte:
		return String(s.(string))
	}
	return String("")
}

func GetInt(n interface{}) Int {
	switch n.(type) {
	case int:
		return Int(n.(int))
	case int8:
		return Int(n.(int8))
	case int16:
		return Int(n.(int16))
	case int32:
		return Int(n.(int32))
	case int64:
		return Int(n.(int64))
	}
	return Int(0)
}

func GetUint(n interface{}) Uint {
	switch n.(type) {
	case uint:
		return Uint(n.(uint))
	case uint8:
		return Uint(n.(uint8))
	case uint16:
		return Uint(n.(uint16))
	case uint32:
		return Uint(n.(uint32))
	case uint64:
		return Uint(n.(uint64))
	}
	return Uint(0)
}

func GetFloat(f interface{}) Float {
	switch f.(type) {
	case float32:
		return Float(f.(float32))
	case float64:
		return Float(f.(float64))
	}
	return Float(0)
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

// FLOAT

type Float float64

func (f Float) IsEqualTo(v Value) bool {
	if vf, ok := v.(Float); ok {
		return f == vf
	}
	return false
}

func (f Float) IsGreaterThan(v Value) bool {
	if vf, ok := v.(Float); ok {
		return vf < f
	}
	return false
}

func (f Float) IsLessThan(v Value) bool {
	if vf, ok := v.(Float); ok {
		return f < vf
	}
	return false
}
