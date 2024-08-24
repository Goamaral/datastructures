package merkle_tree_test

import (
	"testing"

	"github.com/goamaral/data-structures/merkle_tree"
	"github.com/stretchr/testify/assert"
)

func TestMerkleTree_New(t *testing.T) {
	type Test struct {
		Name string
		Data []byte
	}

	invalidData := []byte("hello world")
	data := []byte("WVwV8gzzXdJ1XK5nGh1GjibUW6y54c5a1hle6Gc5PG5F0GyHq094VRN53ZWz4qeTpBh9Pcg6Je4HVmUHU2gy6oY8Zx658q04DZLeV6i07X8JU1uSRa3GlhKEzfW")

	tests := []Test{
		{"Empty", nil},
		{"Below 32 bytes", data[:31]},
		{"Above 32 bytes", data[:33]},
		{"Above 64 bytes", data[:65]},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			one := merkle_tree.New(test.Data)
			two := merkle_tree.New(test.Data)
			assert.True(t, one.Match(two))

			invalid := merkle_tree.New(invalidData)
			assert.False(t, one.Match(invalid))
		})
	}
}

func TestMerkleTree_Append(t *testing.T) {
	// TODO
}
