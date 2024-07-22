package base

type AddressUtil interface {
	// EncodePublicKeyToAddress 地址加密转换伴随0x
	EncodePublicKeyToAddress(publicKey string) (string, error)
	// DecodeAddressToPublicKey 地址解密转换伴随0x
	DecodeAddressToPublicKey(address string) (string, error)
	// IsValidAddress 是否是合法地址
	IsValidAddress(address string) bool
}

type Account interface {
	// PrivateKey 生成私钥
	PrivateKey() ([]byte, error)
	// PrivateKeyHex 生成私钥16进制
	PrivateKeyHex() (string, error)

	// PublicKey 生成公钥
	PublicKey() []byte
	// PublicKeyHex 公钥16进制
	PublicKeyHex() string

	// Address 转换地址
	Address() string

	// Sign 签名数据
	Sign(message []byte, password string) ([]byte, error)
	SignHex(messageHex string, password string) (*OptionalString, error)
}
