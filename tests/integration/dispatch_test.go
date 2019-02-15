package integration

import (
	"testing"

	"github.com/pokt-network/pocket-core/db"
	"github.com/pokt-network/pocket-core/node"
)

func DummyNode() node.Node {
	chains := []node.Blockchain{{Name: "ethereum", NetID: "1", Version: "1"}}
	n := node.Node{
		GID:         "test",
		IP:          "123",
		RelayPort:   "0",
		ClientPort:  "0",
		ClientID:    "0",
		CliVersion:  "0",
		Blockchains: chains,
	}
	return n
}

func TestDB(t *testing.T) {
	n := DummyNode()
	_, err := db.DB().Add(n)
	if err != nil {
		t.Fatalf(err.Error())
	}
	_, err = db.DB().Remove(n)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
