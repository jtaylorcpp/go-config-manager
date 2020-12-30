package types

type OptType = string

const (
	APPLY_CONFIG       OptType = "apply_config"
	APPLY_TRANSACTIONS OptType = "apply_transactions"
)

type StoreOpts interface {
	Type() OptType
	Transactions() []Transaction
	Config() Config
}
