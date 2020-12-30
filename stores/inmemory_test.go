package stores

import (
	"encoding/json"
	"reflect"
	"testing"

	cm "github.com/jtaylorcpp/go-config-manager/types"
)

func TestInit(t *testing.T) {
	mem := &InMemoryStore{}
	initErr := mem.Init()

	if initErr != nil {
		t.Fatal(initErr.Error())
	}
}

func TestSubmit(t *testing.T) {
	mem := &InMemoryStore{
		transactions: []cm.Transaction{},
	}

	if err := mem.Submit(cm.Transaction{0, cm.UPDATE, []byte("hello")}); err != nil {
		t.Fatal(err.Error())
	}

	if len(mem.transactions) != 1 {
		t.Fatal("not enough transactions")
	}

	if mem.transactions[0].Version != 0 {
		t.Fatal("incorrect version num")
	}

	if mem.transactions[0].Action != cm.UPDATE {
		t.Fatal("incorrect action")
	}

	if !reflect.DeepEqual(mem.transactions[0].PartialConfig, json.RawMessage("hello")) {
		t.Logf("actual: %#v, expected: %#v", mem.transactions[0].PartialConfig, []byte("hello"))
		t.Fatal("incorrect partial config")
	}
}

func TestReadAll(t *testing.T) {
	mem := &InMemoryStore{
		transactions: []cm.Transaction{
			{
				Version: 0,
				Action:  cm.UPDATE,
				PartialConfig: []byte(`
				{
					"hello": "world"
				}
				`),
			},
		},
	}

	list, err := mem.ReadAll()

	if err != nil {
		t.Fatal(err.Error())
	}
	if len(list) != 1 {
		t.Fatal("list incorrect length")
	}

	if list[0].Version != 0 || list[0].Action != cm.UPDATE {
		t.Fatal("transaction incorrect metadata")
	}
}
