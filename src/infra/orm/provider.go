package orm

type Provider interface {
	GetID() uint
	Find(int) (err error)
	FindBy(string, interface{}) error
	ProviderName() string
	GetAccountID() string
	GetMypageURL() string
}
