package stores

import (
	"github.com/jtaylorcpp/go-config-manager/types"
	cm "github.com/jtaylorcpp/go-config-manager/types"
)

type InMemoryStoreOpts struct {
	Config       cm.Config
	Transactions []cm.Transaction
}

func (o InMemoryStoreOpts) Type() types.OptType {
	return types.APPLY_CONFIG
}

func (o InMemoryStoreOpts) Transactions() []types.Transaction {
	return o.Transactions
}

type InMemoryStore struct {
	transactions []cm.Transaction
}

func (s *InMemoryStore) Init(opts ...types.StoreOpts) error {
	s.transactions = []cm.Transaction{}

	for _, opt := range opts {
		switch opt.Type() {
		case cm.APPLY_CONFIG:
			s.transactions = append(s.transactions, cm.Transaction{
				Version:       0,
				Action:        cm.UPDATE,
				PartialConfig: opt.Config(),
			})
		case cm.APPLY_TRANSACTIONS:
			s.transactions = append(s.transactions, opt.Transactions()...)
		}
	}
	return nil
}

func (s *InMemoryStore) Submit(r cm.Transaction) error {
	s.transactions = append(s.transactions, r)
	return nil
}

func (s *InMemoryStore) ReadAll() ([]cm.Transaction, error) {
	return s.transactions, nil
}
