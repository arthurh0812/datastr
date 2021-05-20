package hashtable

import "github.com/arthurh0812/datastruct/types"

// Function is the signature of a hash function that can also be implemented by clients.
type Function func(key types.Value) (idx int)

var DefaultFunction Function = func(key types.Value) (idx int) {
	return int(key.(types.Int)) * 10 % 2
}