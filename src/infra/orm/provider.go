package orm

type Provider interface {
	GetID() uint
	FindBy(string, interface{}) error
	ProviderName() string
}
