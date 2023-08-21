package util

import (
	"context"
	"fmt"
	"time"

	//"github.com/TotemFi/Totem-Bridge/pkg/contracts"
	contracts "github.com/TotemFi/totem-bridge-offchain/abi"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcmm "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"

	//"github.com/cockroachdb/apd"
	"math/big"
	"strings"
)

type PriceModel struct {
	BscFeeTx           *decimal.Decimal
	MetisFeeTx         *decimal.Decimal
	PriceImpactPercent *decimal.Decimal
	SlippagePercent    *decimal.Decimal
	TotmTradeFee       *decimal.Decimal
	MetisPrice         *decimal.Decimal
	MetisAmount        *decimal.Decimal
	MetisGasPrice      *decimal.Decimal
	BscGasPrice        *decimal.Decimal
	BNBPrice           *decimal.Decimal
	WBNBAmount         *decimal.Decimal
	WBNBReserved       *decimal.Decimal
	TOTMReserved       *decimal.Decimal
	TOTMPrice          *decimal.Decimal
	TOTMExchange       *decimal.Decimal
	TOTMFee            *decimal.Decimal
}

type SwapPriceUtils struct {
	Config             *Config
	SymbolPrice        SymbolPrice
	BscChain           string
	BscSwapAgentAddr   ethcmm.Address
	BscBridgeAgent     *contracts.BSCBridgeAgentImpl
	BscSwapAgentAbi    abi.ABI
	BscClient          *ethclient.Client
	MetisChainID       int64
	BscChainID         int64
	MetisChain         string
	MetisSwapAgentAddr ethcmm.Address
	MetisBridgeAgent   *contracts.MetisBridgeAgentImpl
	MetisSwapAgentAbi  abi.ABI
	MetisClient        *ethclient.Client
	//MetisPrivateKey    *ecdsa.PrivateKey
	//BscPrivateKey      *ecdsa.PrivateKey
	MetisAgentGasLimit *big.Int
	BSCAgentGasLimit   *big.Int
}

func NewSwapPriceUtils(config *Config, bscClient *ethclient.Client, metisClient *ethclient.Client) *SwapPriceUtils {

	bscAgentAbi, err := abi.JSON(strings.NewReader(contracts.BSCBridgeAgentImplMetaData.ABI))
	if err != nil {
		panic("marshal abi error")
	}

	bscBridgeAgent, err := contracts.NewBSCBridgeAgentImpl(ethcmm.HexToAddress(config.ChainConfig.BSCSwapAgentAddr), bscClient)
	if err != nil {
		panic(err.Error())
	}

	//bscPrivateKey, _, err := buildKeys("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//mtsPrivateKey, _, err := buildKeys("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	//if err != nil {
	//	panic(err.Error())
	//}

	bscChainID, err := bscClient.ChainID(context.Background())
	if err != nil {
		panic(err.Error())
	}

	metisChainID, err := metisClient.ChainID(context.Background())
	if err != nil {
		panic(err.Error())
	}

	metisAgentAbi, err := abi.JSON(strings.NewReader(contracts.MetisBridgeAgentImplMetaData.ABI))
	if err != nil {
		panic("marshal abi error")
	}
	metisSwapAgentInst, err := contracts.NewMetisBridgeAgentImpl(ethcmm.HexToAddress(config.ChainConfig.MTSSwapAgentAddr), metisClient)
	if err != nil {
		panic(err.Error())
	}

	metisAgentGasLimit := big.NewInt(config.ChainConfig.MTSAgentGasLimit)
	bscAgentGasLimit := big.NewInt(config.ChainConfig.BSCAgentGasLimit)

	return &SwapPriceUtils{
		Config:             config,
		BscChain:           "",
		BscSwapAgentAddr:   ethcmm.HexToAddress(config.ChainConfig.BSCSwapAgentAddr),
		BscBridgeAgent:     bscBridgeAgent,
		BscSwapAgentAbi:    bscAgentAbi,
		BscClient:          bscClient,
		MetisChainID:       metisChainID.Int64(),
		BscChainID:         bscChainID.Int64(),
		MetisChain:         "",
		MetisSwapAgentAddr: ethcmm.HexToAddress(config.ChainConfig.MTSSwapAgentAddr),
		MetisBridgeAgent:   metisSwapAgentInst,
		MetisSwapAgentAbi:  metisAgentAbi,
		MetisClient:        metisClient,
		MetisAgentGasLimit: metisAgentGasLimit,
		BSCAgentGasLimit:   bscAgentGasLimit,
		SymbolPrice: SymbolPrice{
			CoinMarketCapApiKey: config.KeyManagerConfig.CoinMarketCapApiKey,
			lastCallTimestamp:   time.Now(),
			firstCall:           true,
		},
	}
}

func (priceUtils *SwapPriceUtils) SwapPriceCalc(amount *big.Int) (*PriceModel, error) {

	metisPrice := priceUtils.SymbolPrice.GetPrice("Metis")
	if metisPrice == decimal.Zero {
		value, err := priceUtils.SymbolPrice.GetSymbolPriceFromMarketCap("Metis")
		if err != nil {
			Logger.Error("Get Metis Price failed", err)
			return nil, err
		}
		Logger.Info("metis float price: ", value)
		metisPrice = decimal.NewFromFloatWithExponent(value, -18).Shift(18)
		priceUtils.SymbolPrice.SetPrice("Metis", metisPrice)
	}

	wbnbPrice := priceUtils.SymbolPrice.GetPrice("WBNB")
	if wbnbPrice == decimal.Zero {
		value, err := priceUtils.SymbolPrice.GetSymbolPriceFromMarketCap("WBNB")
		if err != nil {
			Logger.Error("Get BNB Price failed", err)
			return nil, err
		}
		Logger.Info("wbnb float price: ", value)
		wbnbPrice = decimal.NewFromFloatWithExponent(value, -18).Shift(18)
		priceUtils.SymbolPrice.SetPrice("WBNB", wbnbPrice)
	}

	bscGasPrice := priceUtils.SymbolPrice.GetPrice("BSCGASPrice")
	if bscGasPrice == decimal.Zero {
		value, err := priceUtils.BscClient.SuggestGasPrice(context.Background())
		if err != nil {
			Logger.Error("Get BSC gas price failed", err)
			return nil, err
		}
		bscGasPrice = decimal.NewFromBigInt(value, 0)
		priceUtils.SymbolPrice.SetPrice("BSCGASPrice", bscGasPrice)
	}

	metisGasPrice := priceUtils.SymbolPrice.GetPrice("METISGASPrice")
	if metisGasPrice == decimal.Zero {
		value, err := priceUtils.MetisClient.SuggestGasPrice(context.Background())
		if err != nil {
			Logger.Error("Get Metis gas price failed", err)
			return nil, err
		}
		metisGasPrice = decimal.NewFromBigInt(value, 0)
		priceUtils.SymbolPrice.SetPrice("METISGASPrice", metisGasPrice)
	}

	bscGasLimit := decimal.NewFromBigInt(priceUtils.BSCAgentGasLimit, 0)
	metisGasLimit := decimal.NewFromBigInt(priceUtils.MetisAgentGasLimit, 0)

	metisAmount := decimal.NewFromBigInt(amount, 0)
	wbnbAmount := metisAmount.Mul(metisPrice).DivRound(wbnbPrice, 0)
	Logger.Info("WBNB amount without wbnbfee: ", wbnbAmount)
	result, err := priceUtils.BscBridgeAgent.SwapRouterQuery(nil, wbnbAmount.BigInt())
	if err != nil {
		Logger.Error("SwapRouterQuery failed", err)
		return nil, err
	}

	totemExchange := decimal.NewFromBigInt(result.Exchange, 0)
	baseReserved := decimal.NewFromBigInt(result.BaseReserved, 0)
	quoteReserved := decimal.NewFromBigInt(result.QuoteReserved, 0)

	//totemExchange, _ := decimal.NewFromString("6058803318838433708609")
	//baseReserved, _ := decimal.NewFromString("170771741466742815428")
	//quoteReserved, _ := decimal.NewFromString("793750625327472603809880")

	wbnbfee := wbnbAmount.Mul(decimal.NewFromFloat(priceUtils.Config.ChainConfig.SwapRouterTradeFee+priceUtils.Config.ChainConfig.BridgeTaxFee)).DivRound(decimal.NewFromInt(100), 0)
	wbnbAmount = wbnbAmount.Sub(wbnbfee)
	constantProduct := baseReserved.Mul(quoteReserved)
	totemWbnbRate := baseReserved.DivRound(quoteReserved, 18)
	wbnbTotemRate := quoteReserved.DivRound(baseReserved, 18)
	newBaseReserved := baseReserved.Add(wbnbAmount)
	newQuoteReserved := constantProduct.DivRound(newBaseReserved, 0)
	totemReceivedExchange := quoteReserved.Sub(newQuoteReserved)
	wbnbTotemNewRate := newQuoteReserved.DivRound(newBaseReserved, 18)
	wbnbTotemExchangeRate := totemReceivedExchange.DivRound(wbnbAmount, 18) // price paid per wbnb
	diffWbnbTotemRate := wbnbTotemExchangeRate.Sub(wbnbTotemRate).Abs()
	totmPriceImpactPercent := diffWbnbTotemRate.Mul(decimal.NewFromInt(100)).DivRound(wbnbTotemRate, 18)
	totmPriceImpactValue := diffWbnbTotemRate.Mul(wbnbAmount).Truncate(0)
	totemTradeFee := wbnbfee.Mul(wbnbTotemRate).Truncate(0)

	bscTxFee := bscGasPrice.Mul(bscGasLimit)
	metisTxFee := metisGasPrice.Mul(metisGasLimit)

	totemBscTxFee := bscTxFee.Mul(wbnbTotemRate).Truncate(0)
	wbnbMetisTxFee := metisTxFee.Mul(metisPrice).Div(wbnbPrice)
	totemMetisTxFee := wbnbMetisTxFee.Mul(wbnbTotemRate).Truncate(0)

	totemSwapFee := totemBscTxFee.Add(totemMetisTxFee).Add(totemTradeFee).Add(totmPriceImpactValue).Truncate(0)
	totemSwapFinal := totemReceivedExchange.Sub(totemBscTxFee).Sub(totemMetisTxFee).Truncate(0)

	Logger.Info("WBNB trade fee: ", wbnbfee)
	Logger.Info("totem trade fee: ", totemTradeFee)
	Logger.Info("WBNB amount: ", wbnbAmount)
	Logger.Info("metis price: ", metisPrice)
	Logger.Info("wbnb price: ", wbnbPrice)
	Logger.Info("bsc gas price: ", bscGasPrice)
	Logger.Info("metis gas price: ", metisGasPrice)
	Logger.Info("bsc gas limit: ", bscGasLimit)
	Logger.Info("metis gas limit: ", metisGasLimit)
	Logger.Info("bsc tx fee: ", bscTxFee)
	Logger.Info("metis tx fee: ", metisTxFee)
	Logger.Info("totem bsc tx fee: ", totemBscTxFee)
	Logger.Info("totem metis tx fee: ", totemMetisTxFee)
	Logger.Info("totem exchange amount: ", totemExchange)
	Logger.Info("wbnb (base reserved) amount: ", baseReserved)
	Logger.Info("totm (qoute reserved) amount: ", quoteReserved)
	Logger.Info("constant product value: ", constantProduct)
	Logger.Info("wbnb_totem_rate: ", wbnbTotemRate)
	Logger.Info("totem_wbnb_rate: ", totemWbnbRate)
	Logger.Info("newBaseReserved: ", newBaseReserved)
	Logger.Info("newQuoteReserved: ", newQuoteReserved)
	Logger.Info("totemReceivedExchange: ", totemReceivedExchange)
	Logger.Info("wbnb_totem_new_rate: ", wbnbTotemNewRate)
	Logger.Info("wbnbTotemExchangeRate: ", wbnbTotemExchangeRate)
	Logger.Info("diffWbnbTotemRate: ", diffWbnbTotemRate)
	Logger.Info("totmPriceImpactPercent: ", totmPriceImpactPercent)
	Logger.Info("totmPriceImpactValue: ", totmPriceImpactValue)
	Logger.Info("Swap fee (totem): ", totemSwapFee)
	Logger.Info("Swap final (totem): ", totemSwapFinal)

	if totemSwapFinal.IsNegative() {
		Logger.Errorf("totem exchange is negative, totemSwapFinal: %s", totemSwapFinal.String())
		return nil, fmt.Errorf("totem exchange is negative")
	}

	priceModel := &PriceModel{
		BscFeeTx:           &bscTxFee,
		MetisFeeTx:         &metisTxFee,
		PriceImpactPercent: &totmPriceImpactPercent,
		//SlippagePercent:    &decimal.NewFromInt(3),
		TotmTradeFee:  &totemTradeFee,
		MetisPrice:    &metisPrice,
		MetisAmount:   &metisAmount,
		MetisGasPrice: &metisGasPrice,
		BscGasPrice:   &bscGasPrice,
		BNBPrice:      &wbnbPrice,
		WBNBAmount:    &wbnbAmount,
		WBNBReserved:  &baseReserved,
		TOTMReserved:  &quoteReserved,
		TOTMPrice:     &wbnbTotemRate,
		TOTMExchange:  &totemSwapFinal,
		TOTMFee:       &totemSwapFee,
	}

	return priceModel, nil
}

//func (priceUtils *SwapPriceUtils) abiEncodeSwapRouterQuery(amount *big.Int) ([]byte, error) {
//	data, err := priceUtils.BscSwapAgentAbi.Pack("swapRouterQuery", amount)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}

//func (priceUtils *SwapPriceUtils) calcBscGasLimit(amount, fee, exchange, nonce, deadline *big.Int,
//	account ethcmm.Address, metisTxHash ethcmm.Hash) (*big.Int, error) {
//
//	message := struct {
//		SwapType string
//		Base     string   // MTS
//		Quote    string   // TOTM
//		Amount   *big.Int // metis amount
//		Fee      *big.Int // fee in totem
//		Exchange *big.Int // totem amount
//		Nonce    *big.Int
//		Deadline *big.Int
//		Account  ethcmm.Address
//	}{
//		SwapType: "peg-in",
//		Base:     "WBNB",
//		Quote:    "TOTM",
//		Amount:   amount,
//		Fee:      fee,
//		Exchange: exchange,
//		Nonce:    nonce,
//		Deadline: deadline,
//		Account:  account,
//	}
//
//	data, err := priceUtils.BscSwapAgentAbi.Pack("fillBNB2TOTMPegin", message, metisTxHash)
//	if err != nil {
//		return nil, err
//	}
//
//	txOpts, err := bind.NewKeyedTransactorWithChainID(priceUtils.BscPrivateKey, big.NewInt(priceUtils.BscChainID))
//	if err != nil {
//		return nil, err
//	}
//
//	gasPrice, err := priceUtils.BscClient.SuggestGasPrice(context.Background())
//	if err != nil {
//		return nil, err
//	}
//
//	value := big.NewInt(0)
//	msg := ethereum.CallMsg{From: txOpts.From, To: &priceUtils.BscSwapAgentAddr, GasPrice: gasPrice, Value: value, Data: data}
//	gasLimit, err := priceUtils.BscClient.EstimateGas(context.Background(), msg)
//	if err != nil {
//		return nil, fmt.Errorf("failed to estimate gas needed: %v", err)
//	}
//
//	return big.NewInt(int64(gasLimit)), nil
//}
//
//func (priceUtils *SwapPriceUtils) calcMetisGasLimit(dataHash ethcmm.Hash, recipient ethcmm.Address, fee,
//	exchange *big.Int, swapType string) (*big.Int, error) {
//	data, err := priceUtils.MetisSwapAgentAbi.Pack("fillMTS2TOTMPegIn", recipient, dataHash, dataHash, swapType, fee, exchange)
//	if err != nil {
//		return nil, err
//	}
//
//	txOpts, err := bind.NewKeyedTransactorWithChainID(priceUtils.BscPrivateKey, big.NewInt(priceUtils.BscChainID))
//	if err != nil {
//		return nil, err
//	}
//
//	gasPrice, err := priceUtils.MetisClient.SuggestGasPrice(context.Background())
//	if err != nil {
//		return nil, err
//	}
//
//	value := big.NewInt(0)
//	msg := ethereum.CallMsg{From: txOpts.From, To: &priceUtils.BscSwapAgentAddr, GasPrice: gasPrice, Value: value, Data: data}
//	gasLimit, err := priceUtils.MetisClient.EstimateGas(context.Background(), msg)
//	if err != nil {
//		return nil, fmt.Errorf("failed to estimate gas needed: %v", err)
//	}
//
//	return big.NewInt(int64(gasLimit)), nil
//}

//func buildKeys(privateKeyStr string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
//	if strings.HasPrefix(privateKeyStr, "0x") {
//		privateKeyStr = privateKeyStr[2:]
//	}
//	priKey, err := crypto.HexToECDSA(privateKeyStr)
//	if err != nil {
//		return nil, nil, err
//	}
//	publicKey, ok := priKey.Public().(*ecdsa.PublicKey)
//	if !ok {
//		return nil, nil, fmt.Errorf("get public key error")
//	}
//	return priKey, publicKey, nil
//}
