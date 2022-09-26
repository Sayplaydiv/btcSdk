package data

type ReqGenerateAddress struct {
	ChainId  string //链ID或网络标识 比如:ETH链 1:mainnet 3:ropsten
	AddrType string
	Label    string
}

type ReqSignTxByFrom struct {
	ChainId      string //链ID或网络标识 比如:ETH链 1:mainnet 3:ropsten
	MainTransfer bool
	From         string
	PublicKey    string
	PrivateKey   string
	To           string
	Contract     string
	Amount       string
	GasPrice     int64
	GasLimit     int64
	Nonce        int64
	TipCap       int64 //sign1 必传
}

type ReqSignTxByUtxo struct {
	ChainId    string //链ID或网络标识 比如:ETH链 1:mainnet 3:ropsten
	From       []string
	PublicKey  []string
	PrivateKey []string
	InTxs      []InTx
	To         string
	Change     string //找零地址
	Amount     string
	FeeRate    float64
}

type InTx struct {
	TxId         string
	Vout         uint32
	Amount       float64
	ScriptPubkey string
	Address      string
}
