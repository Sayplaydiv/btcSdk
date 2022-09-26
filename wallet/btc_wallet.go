package wallet

import (
	"fmt"
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/blocktree/go-owcdrivers/btcLikeTxDriver"
	"github.com/dabankio/wallet-core/bip39"
	"github.com/dabankio/wallet-core/bip44"
	"new/data"
	"new/walletsdk/btc_client/core/btc"
	"strconv"
)

type BTC_Wallet struct {
	BaseService
}

func NewBTCWallet(main_symbol string) *BTC_Wallet {
	wallet := &BTC_Wallet{}
	wallet.main_symbol = main_symbol
	return wallet
}

func (this *BTC_Wallet) GenerateAddress(req data.ReqGenerateAddress) (*data.ResAccount, error) {
	chainId, err := strconv.ParseInt(req.ChainId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("chainId %v %v", req.ChainId, err)
	}

	if req.AddrType == "legacy" {
		seedByte, _ := bip39.NewEntropy(256)
		deriver, err := btc.NewBip44Deriver(bip44.PathFormat, false, seedByte, int(chainId))
		if err != nil {
			return nil, err
		}
		address, err := deriver.DeriveAddress()
		if err != nil {
			return nil, err
		}
		privateKey, err := deriver.DerivePrivateKey()
		if err != nil {
			return nil, err
		}
		return &data.ResAccount{Address: address, PrivateKey: privateKey, Label: req.Label}, nil
	} else if req.AddrType == "p2sh-segwit" {
		//SegWit（P2SH）,地址以“ 3 ”开头
		return nil, nil
	} else if req.AddrType == "bech32" {
		seedByte, _ := bip39.NewEntropy(256)
		deriver, err := btc.NewBip44Deriver(bip44.PathFormat, false, seedByte, int(chainId))
		if err != nil {
			return nil, err
		}
		address, err := deriver.DeriveAddress()
		if err != nil {
			return nil, err
		}

		privateKey, err := deriver.DerivePrivateKey()
		if err != nil {
			return nil, err
		}
		var BTCBech32Address = addressEncoder.AddressType{}
		var BTCAddress = addressEncoder.AddressType{}
		if chainId == btc.ChainMainNet {
			BTCBech32Address = addressEncoder.BTC_mainnetAddressBech32V0
			BTCAddress = addressEncoder.BTC_mainnetAddressP2PKH
		} else if chainId == btc.ChainTestNet3 {
			BTCBech32Address = addressEncoder.BTC_testnetAddressBech32V0
			BTCAddress = addressEncoder.BTC_testnetAddressP2PKH
		}

		//解码
		address_de, _ := addressEncoder.AddressDecode(address, BTCAddress)

		//编码为bech32地址
		var address_bech32 string
		if chainId == btc.ChainMainNet {
			address_bech32 = btcLikeTxDriver.Bech32Encode("bc", BTCBech32Address.Alphabet, address_de)
		} else if chainId == btc.ChainTestNet3 {
			address_bech32 = btcLikeTxDriver.Bech32Encode("tb", BTCBech32Address.Alphabet, address_de)
		}
		//SegWit（bech32）：地址以“ bc1 ”开头
		return &data.ResAccount{Address: address_bech32, PrivateKey: privateKey, Label: req.Label}, nil
	} else {
		return nil, fmt.Errorf("不支持:%v.", req.AddrType)
	}
}

func (this *BTC_Wallet) SignTxByUtxo(req data.ReqSignTxByUtxo) (*data.ResSig, error) {
	chainId, err := strconv.ParseInt(req.ChainId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("chainId %v %v", req.ChainId, err)
	}

	input := new(btc.BTCUnspent)
	for _, in := range req.InTxs {
		input.Add(in.TxId, int64(in.Vout), in.Amount, in.ScriptPubkey, "")
	}
	addr, err := btc.NewBTCAddressFromString(req.To, int(chainId))
	if err != nil {
		return nil, err
	}
	amtFloat, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil {
		return nil, err
	}
	amt, err := btc.NewBTCAmount(amtFloat)
	if err != nil {
		return nil, err
	}
	output := new(btc.BTCOutputAmount)
	output.Add(addr, amt)

	change, err := btc.NewBTCAddressFromString(req.Change, int(chainId))
	if err != nil {
		return nil, err
	}
	tt, err := btc.NewBTCTransaction(input, output, change, int64(req.FeeRate), int(chainId))
	if err != nil {
		return nil, err
	}
	hh, err := tt.EncodeToSignCmd()
	if err != nil {
		return nil, err
	}
	bb, _ := btc.New(bip44.PathFormat, false, nil, int(chainId))
	cc, err := bb.Sign(hh, req.PrivateKey)
	if err != nil {
		return nil, err
	}
	var sig = &data.ResSig{}
	sig.SigData = cc
	return sig, nil
}
