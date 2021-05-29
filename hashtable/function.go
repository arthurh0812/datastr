package hashtable

import (
	"github.com/arthurh0812/datastruct/types"
)

// Function is the signature of a hash function that can also be implemented by clients.
type Function func(key types.Value) (idx int)

var DefaultFunction Function = func(key types.Value) (idx int) {
	return int(key.(types.Int)) * 10 % 2
}

var StringFunction Function = func(key types.Value) (idx int) {
	s := string(key.(types.String))
	return (len(s) * 560) % 5
}

var IntFunction Function = func(key types.Value) (idx int) {
	i := int64(key.(types.Int))
	return int((i * 560) % 5)
}

var UintFunction Function = func(key types.Value) (idx int) {
	u := uint64(key.(types.Uint))
	return int((u * 30) % 3)
}

var FloatFunction Function = func(key types.Value) (idx int) {
	f := float64(key.(types.Float))
	return (int(f * 560)) % 2
}

var FuncMap = map[string]Function{
	"STRING": StringFunction,
	"INT": IntFunction,
	"INT8": IntFunction,
	"INT16": IntFunction,
	"INT32": IntFunction,
	"INT64": IntFunction,
	"UINT": UintFunction,
	"UINT8": UintFunction,
	"UIN16": UintFunction,
	"UINT32": UintFunction,
	"UINT64": UintFunction,
	"FLOAT32": FloatFunction,
	"FLOAT64": FloatFunction,
}

func GetFunction(kind string) Function {
	return FuncMap[kind]
}