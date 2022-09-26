package data

type ResAccount struct {
	Address    string `json:"address"`
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
	Label      string `json:"label"`
	Debug      string `json:"debug"`
}

type ResSig struct {
	SigData string `json:"sigData"`
	Debug   string `json:"debug"`
}
