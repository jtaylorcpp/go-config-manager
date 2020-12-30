package configmanager

import "github.com/jtaylorcpp/go-config-manager/types"

type Store interface {
	Init(...types.StoreOpts) error
	Submit(types.Transaction) error
	ReadAll() ([]types.Transaction, error)
}
