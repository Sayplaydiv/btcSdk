package wallet

import (
	"fmt"
	"new/data"
)

type BaseService struct {
	main_symbol string
}

func (this *BaseService) GenerateAddress(req data.ReqGenerateAddress) (*data.ResAccount, error) {
	return nil, fmt.Errorf("子类未实现")
}

func (this *BaseService) SignTxByUtxo(req data.ReqSignTxByUtxo) (*data.ResSig, error) {
	return nil, fmt.Errorf("子类未实现")
}
