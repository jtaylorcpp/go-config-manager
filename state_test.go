package configmanager

import (
	"encoding/json"
	"testing"

	"github.com/jtaylorcpp/go-config-manager/types"
)

func TestApplyUpdateToEmptyConfig(t *testing.T) {
	state := &State{
		Version: 0,
		Config:  json.RawMessage(``),
	}

	initTxn := types.Transaction{
		Version: 1,
		Action:  types.UPDATE,
		PartialConfig: json.RawMessage(`
		{
			"hello": "world"
		}
		`),
	}

	applyErr := state.ApplyTransactions(initTxn)

	if applyErr != nil {
		t.Fatalf("failed to apply transactions: %v", applyErr.Error())
	}

	t.Logf("state after transactions: %#v", state)
	t.Logf("config after transactions: %#v", string(state.Config))

	if string(state.Config) != "{\"hello\":\"world\"}" {
		t.Fatal("unexpected state")
	}
}

func TestApplyDeleteToEmptyConfig(t *testing.T) {
	state := &State{
		Version: 0,
		Config:  json.RawMessage(``),
	}

	initTxn := types.Transaction{
		Version: 1,
		Action:  types.DELETE,
		PartialConfig: json.RawMessage(`
		{
			"hello": "world"
		}
		`),
	}

	applyErr := state.ApplyTransactions(initTxn)

	if applyErr != nil {
		t.Fatalf("failed to apply transactions: %v", applyErr.Error())
	}

	t.Logf("state after transactions: %#v", state)
	t.Logf("config after transactions: %#v", string(state.Config))

	if string(state.Config) != "" {
		t.Fatal("unexpected state")
	}
}

func TestApplyUpdateToNestedConfig(t *testing.T) {
	state := &State{
		Version: 0,
		Config: json.RawMessage(`
		{
			"test": {
				"hello":"world"
			}
		}
		`),
	}

	initTxn := types.Transaction{
		Version: 1,
		Action:  types.UPDATE,
		PartialConfig: json.RawMessage(`
		{
			"test": {
				"number": 1
			}
		}
		`),
	}

	applyErr := state.ApplyTransactions(initTxn)

	if applyErr != nil {
		t.Fatalf("failed to apply transactions: %v", applyErr.Error())
	}

	t.Logf("state after transactions: %#v", state)
	t.Logf("config after transactions: %#v", string(state.Config))

	if string(state.Config) != "\n\t\t{\n\t\t\t\"test\": {\n\t\t\t\t\"hello\":\"world\"\n\t\t\t,\"number\":1}\n\t\t}\n\t\t" {
		t.Fatal("unexpected state")
	}
}

func TestApplyDeleteToNestedConfig(t *testing.T) {
	state := &State{
		Version: 0,
		Config: json.RawMessage(`
		{
			"test": {
				"hello":"world",
				"number":1
			}
		}
		`),
	}

	initTxn := types.Transaction{
		Version: 1,
		Action:  types.DELETE,
		PartialConfig: json.RawMessage(`
		{
			"test": {
				"number": 1
			}
		}
		`),
	}

	applyErr := state.ApplyTransactions(initTxn)

	if applyErr != nil {
		t.Fatalf("failed to apply transactions: %v", applyErr.Error())
	}

	t.Logf("state after transactions: %#v", state)
	t.Logf("config after transactions: %#v", string(state.Config))

	if string(state.Config) != "\n\t\t{\n\t\t\t\"test\": {\n\t\t\t\t\"hello\":\"world\"\n\t\t\t}\n\t\t}\n\t\t" {
		t.Fatal("unexpected state")
	}
}
