package types

import (
	"encoding/json"
	"testing"
)

func TestGetKeyValueSet(t *testing.T) {
	txn := Transaction{
		Version: 0,
		Action:  UPDATE,
		PartialConfig: json.RawMessage(`
		{
			"test": {
				"hello":"world",
				"number": 1
			}
		}
		`),
	}

	keyvalues, err := txn.GetKeyValueSet()
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(keyvalues) != 2 {
		t.Logf("expected 2 values, returned %v: %#v", len(keyvalues), keyvalues)
		t.Fatal("not enough returned in set")
	}

	t.Logf("recieved transaction KV pairs: %#v", keyvalues)

	switch keyvalues[0].key {
	case "test.hello":
		if keyvalues[0].value.(string) != "world" {
			t.Fatalf("expected value \"world\" got %s", keyvalues[0].value)
		}

		if keyvalues[1].key != "test.number" {
			t.Fatalf("expected key test.number got %v", keyvalues[0].key)
		}
	case "test.number":
		if keyvalues[0].value.(int) != 1 {
			t.Fatalf("expected value int(1) got %v", keyvalues[0].value)
		}

		if keyvalues[1].key != "test.hello" {
			t.Fatalf("expected key test.hello got %v", keyvalues[0].key)
		}
	default:
		t.Fatalf("unexpected key value pair: %#v", keyvalues[0])
	}
}
