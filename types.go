package datastruct

import (
	"github.com/arthurh0812/datastruct/types"
)

func NewString(s string) types.String {
	return types.GetString(s)
}

func NewStringSlice(s []string) []types.String {
	res := make([]types.String, 0, len(s))
	for _, e := range s {
		res = append(res, types.GetString(e))
	}
	return res
}

func NewInt(n interface{}) types.Int {
	return types.GetInt(n)
}

func NewIntSlice(s []interface{}) []types.Int {
	res := make([]types.Int, 0, len(s))
	for _, e := range s {
		res = append(res, types.GetInt(e))
	}
	return res
}

func NewUint(n interface{}) types.Uint {
	return types.GetUint(n)
}

func NewUintSlice(s []interface{}) []types.Uint {
	res := make([]types.Uint, 0, len(s))
	for _, e := range s {
		res = append(res, types.GetUint(e))
	}
	return res
}

func NewFloat(f interface{}) types.Float {
	return types.GetFloat(f)
}

func NewFloatSlice(s []interface{}) []types.Float {
	res := make([]types.Float, 0, len(s))
	for _, e := range s {
		res = append(res, types.GetFloat(e))
	}
	return res
}