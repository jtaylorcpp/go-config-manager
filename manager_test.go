package configmanager

import (
	"encoding/json"
	"testing"

	"github.com/jtaylorcpp/go-config-manager/stores"
	"github.com/jtaylorcpp/go-config-manager/types"
)

func TestManagerInit(t *testing.T) {
	memStore := &stores.InMemoryStore{}
	storeErr := memStore.Init()
	if storeErr != nil {
		t.Fatalf("error init'ing store: %v", storeErr.Error())
	}

	manager := &Manager{
		State: &State{
			Version: 0,
			Config:  json.RawMessage(``),
		},
		Store: memStore,
	}

	err := manager.Init()
	if err != nil {
		t.Fatalf("error init'int manager: %v", err.Error())
	}
}

func TestManagerInitWithConfig(t *testing.T) {
	memStore := &stores.InMemoryStore{}
	storeErr := memStore.Init(stores.InMemoryStoreOpts{
		Config: types.Config(`{"hello":"world"}`),
	})
	if storeErr != nil {
		t.Fatalf("error init'ing store: %v", storeErr.Error())
	}

	manager := &Manager{
		State: &State{
			Version: 0,
			Config:  json.RawMessage(``),
		},
		Store: memStore,
	}

	err := manager.Init()
	if err != nil {
		t.Fatalf("error init'int manager: %v", err.Error())
	}

	config := manager.GetConfig()
	t.Logf("returned config: %v", string(config))
}
