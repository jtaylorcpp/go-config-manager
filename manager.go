package configmanager

import (
	"github.com/jtaylorcpp/go-config-manager/types"
)

type Manager struct {
	Store Store
	State *State
}

func (m *Manager) Init() error {
	transactions, err := m.Store.ReadAll()
	if err != nil {
		return err
	}

	txnErr := m.State.ApplyTransactions(transactions...)
	return txnErr
}

func (m *Manager) Update(partialConfig types.Config) error {
	// create new transaction
	newTransaction := types.Transaction{
		Version:       m.State.Version + 1,
		Action:        types.UPDATE,
		PartialConfig: partialConfig,
	}
	// submit transaction to store
	if storeErr := m.Store.Submit(newTransaction); storeErr != nil {
		return storeErr
	}
	// apply transaction to state
	if stateErr := m.State.ApplyTransactions(newTransaction); stateErr != nil {
		return stateErr
	}
	return nil
}

func (m *Manager) Delete(partialConfig types.Config) error {
	// create new transaction
	newTransaction := types.Transaction{
		Version:       m.State.Version + 1,
		Action:        types.DELETE,
		PartialConfig: partialConfig,
	}
	// submit transaction to store
	if storeErr := m.Store.Submit(newTransaction); storeErr != nil {
		return storeErr
	}
	// apply transaction to state
	if stateErr := m.State.ApplyTransactions(newTransaction); stateErr != nil {
		return stateErr
	}
	return nil
}

func (m *Manager) GetConfig() types.Config {
	return m.State.Config
}
