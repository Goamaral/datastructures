package hash_list_test

import (
	"testing"

	"github.com/goamaral/data-structures/hash_list"
	"github.com/stretchr/testify/assert"
)

func TestHashList_New(t *testing.T) {
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
			one := hash_list.New(test.Data)
			two := hash_list.New(test.Data)
			assert.True(t, one.Match(two))

			invalid := hash_list.New(invalidData)
			assert.False(t, one.Match(invalid))
		})
	}
}

func TestHashList_Append(t *testing.T) {
	// TODO
}
