package orm

type Provider interface {
	Insert() (err error)

	Find(string) (err error)

	ProviderName() string

	GetAccountID() string

	GetMypageURL() string

	// Exists はプライマリーキーを渡し、レコードが存在するか確認します
	Exists(string) bool

	// NewAvatarURL はアバターurlを取得します
	NewAvatarURL() string

	// SetMyPageURL はmypageurlを生成します
	SetMyPageURL()

	SetUserID(uint)
}
