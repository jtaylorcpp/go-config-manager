package configmanager

import (
	"github.com/jtaylorcpp/go-config-manager/types"
	"github.com/tidwall/sjson"
)

type State struct {
	Version uint64
	Config  types.Config
}

func (s *State) ApplyTransactions(transactions ...types.Transaction) error {
	for _, transaction := range transactions {
		keyvalues, transErr := transaction.GetKeyValueSet()
		if transErr != nil {
			return transErr
		}

		newConfig := s.Config
		var actionErr error
		for _, kv := range keyvalues {
			switch transaction.Action {
			case types.UPDATE:
				newConfig, actionErr = sjson.SetBytes(newConfig, kv.Key, kv.Value)

			case types.DELETE:
				newConfig, actionErr = sjson.DeleteBytes(newConfig, kv.Key)
			}
			if actionErr != nil {
				return actionErr
			}
		}
		s.Config = newConfig
		s.Version = transaction.Version
	}

	return nil
}
