package datastr

import (
	"github.com/arthurh0812/datastr/types"
)

func NewString(s string) types.String {
	return types.String(s)
}

func NewStringSlice(s []string) []types.String {
	res := make([]types.String, 0, len(s))
	for _, e := range s {
		res = append(res, NewString(e))
	}
	return res
}

func NewInt(n interface{}) types.Int {
	switch n.(type) {
	case int:
		return types.Int(n.(int))
	case int8:
		return types.Int(n.(int8))
	case int16:
		return types.Int(n.(int16))
	case int32:
		return types.Int(n.(int32))
	case int64:
		return types.Int(n.(int64))
	}
	return types.Int(0)
}

func NewIntSlice(s []interface{}) []types.Int {
	res := make([]types.Int, 0, len(s))
	for _, e := range s {
		res = append(res,NewInt(e))
	}
	return res
}

func NewUint(n interface{}) types.Uint {
	switch n.(type) {
	case uint:
		return types.Uint(n.(uint))
	case uint8:
		return types.Uint(n.(uint8))
	case uint16:
		return types.Uint(n.(uint16))
	case uint32:
		return types.Uint(n.(uint32))
	case uint64:
		return types.Uint(n.(uint64))
	}
	return types.Uint(0)
}

func NewUintSlice(s []interface{}) []types.Uint {
	res := make([]types.Uint, 0, len(s))
	for _, e := range s {
		res = append(res, NewUint(e))
	}
	return res
}