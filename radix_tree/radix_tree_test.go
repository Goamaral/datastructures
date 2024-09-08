package radix_tree_test

import (
	"testing"

	"github.com/goamaral/data-structures/radix_tree"
	"github.com/stretchr/testify/assert"
)

func TestRadixTree(t *testing.T) {
	tr := radix_tree.New()

	leafPrefixMap := map[string]string{}
	testLeafNodePrefix := func() {
		for key, leafPrefix := range leafPrefixMap {
			n := tr.ExactSearch([]byte(key))
			assert.True(t, n.Valid)
			assert.EqualValues(t, n.Prefix, leafPrefix)
		}
	}

	tr.Insert([]byte("slow"))
	leafPrefixMap["slow"] = "slow"
	testLeafNodePrefix()

	tr.Insert([]byte("test"))
	leafPrefixMap["test"] = "test"
	testLeafNodePrefix()

	tr.Insert([]byte("tester"))
	leafPrefixMap["tester"] = "er"
	testLeafNodePrefix()

	tr.Insert([]byte("team"))
	leafPrefixMap["test"] = "st"
	leafPrefixMap["team"] = "am"
	testLeafNodePrefix()

	tr.Insert([]byte("testing"))
	leafPrefixMap["testing"] = "ing"
	testLeafNodePrefix()

	tr.Insert([]byte("toast"))
	leafPrefixMap["toast"] = "oast"
	leafPrefixMap["test"] = "st"
	leafPrefixMap["team"] = "am"
	testLeafNodePrefix()
}
