package wallet

import (
	"fmt"
	"new/data"
	"new/walletsdk/btc_client/core/btc"
	"testing"
)

func Test_BTC_GenerateAddress(t *testing.T) {

	wallet := NewBTCWallet("BTC")

	req := data.ReqGenerateAddress{
		ChainId:  fmt.Sprintf("%v", btc.ChainTestNet3),
		AddrType: "legacy",
	}
	t.Log(wallet.GenerateAddress(req))

	req_bech32 := data.ReqGenerateAddress{
		ChainId:  fmt.Sprintf("%v", btc.ChainTestNet3),
		AddrType: "bech32",
	}
	t.Log(wallet.GenerateAddress(req_bech32))
}

func Test_BTC_SignTxByFrom(t *testing.T) {
	//测试浏览器:https://live.blockcypher.com/btc-testnet/address/n42j2CYeL4sEX9Hk6PagJvfzVV3bxnqLgD/
	//测试节点:https://blockstream.info/testnet/tx/push
	//获取TxId Vout Amount： https://blockstream.info/testnet/api/address/mvtvSF4RgeVDVzVwUPY7RD4pFYLXBMmgih/utxo
	//获取对应地址交易的scriptpubkey ：https://blockstream.info/testnet/api/tx/707c4de4e4ac653e2d4e28a3ae43d2fc1df86741539f713f8be0878bc67bac2e

	//测试币水龙头：https://testnet-faucet.com/btc-testnet/

	//n42j2CYeL4sEX9Hk6PagJvfzVV3bxnqLgD
	//cSEjqm6BFuBdfpnZWmSVMaFYcPmdiXHsUaJ1wBsMdArFZJPLnzEM
	//
	//mvtvSF4RgeVDVzVwUPY7RD4pFYLXBMmgih
	//cRE3RDwVi6Y42LF1dPKepD88y9frjUji9RKG3viPC1vyNnb5Y8Kh
	//
	//mpG7txCFf9x3REeezozKCjoGyy2UZhGE9b <nil>
	//cNHAVTRoWyTmFv1UsmedkffEUrHWsK8BdSzks9cfvLpajx9rfPJz <nil>
	//03c0612e7ec95913d497db3aa749ec1558a7c2ae4d145002449dfdeabc356852ff <nil>
	//
	//mhmSYCobZp7pTuHopRMCLuk921VgSWoUN6 <nil>
	//cReHGePVWUraRiRtxHfXBec66UUe6m9Q6ZoT5vGbexhUWc2Z8nWW <nil>
	//02f1f0e83632191014aa38541e10a3c1438a3a033fdb225982e620f3a88784b29c <nil>

	wallet := NewBTCWallet("BTC")

	req := data.ReqSignTxByUtxo{
		ChainId:    fmt.Sprintf("%v", btc.ChainTestNet3),
		From:       []string{"tb1qwgg3mzd3f6gkmygd2cvvhnq8cxxmqv7fylpxwj"},
		PublicKey:  []string{},
		PrivateKey: []string{"cNq4msygBZvVqVErf4Vk4FwyzT6hHcJPJCRAdsMJJ34xCNNavDy8"},
		InTxs: []data.InTx{
			data.InTx{TxId: "e79bd28b7ce64f799dfb9ce0ce56a635e6ceeb3c4e21c8da1b0930df5241a4ae", Vout: 1, Amount: 0.01825056, ScriptPubkey: "001472111d89b14e916d910d5618cbcc07c18db033c9", Address: "tb1qwgg3mzd3f6gkmygd2cvvhnq8cxxmqv7fylpxwj"},
			//	data.InTx{TxId: "ff97d5c9d3ae47b577c2673bd82a868c7681724b11df9c6153101fd7257513e3", Vout: 0, Amount: 0.000102, ScriptPubkey: "76a9145feae8605734db3b1062d6ee0aec5a25b345f19a88ac", Address: "mpG7txCFf9x3REeezozKCjoGyy2UZhGE9b"},
		},
		To:      "mvtvSF4RgeVDVzVwUPY7RD4pFYLXBMmgih",
		Change:  "tb1qwgg3mzd3f6gkmygd2cvvhnq8cxxmqv7fylpxwj",
		Amount:  "0.01",
		FeeRate: 2, //2= 2*1000
	}
	//已发送交易： https://live.blockcypher.com/btc-testnet/address/n42j2CYeL4sEX9Hk6PagJvfzVV3bxnqLgD/
	t.Log(wallet.SignTxByUtxo(req))
}
