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
			n := tr.SearchNode([]byte(key), true)
			assert.True(t, n.Valid)
			assert.EqualValues(t, n.Prefix, leafPrefix)
		}
	}

	tr.InsertNode([]byte("slow"))
	leafPrefixMap["slow"] = "slow"
	testLeafNodePrefix()

	tr.InsertNode([]byte("test"))
	leafPrefixMap["test"] = "test"
	testLeafNodePrefix()

	tr.InsertNode([]byte("tester"))
	leafPrefixMap["tester"] = "er"
	testLeafNodePrefix()

	tr.InsertNode([]byte("team"))
	leafPrefixMap["test"] = "st"
	leafPrefixMap["team"] = "am"
	testLeafNodePrefix()

	tr.InsertNode([]byte("testing"))
	leafPrefixMap["testing"] = "ing"
	testLeafNodePrefix()

	tr.InsertNode([]byte("toast"))
	leafPrefixMap["toast"] = "oast"
	leafPrefixMap["test"] = "st"
	leafPrefixMap["team"] = "am"
	testLeafNodePrefix()
}
