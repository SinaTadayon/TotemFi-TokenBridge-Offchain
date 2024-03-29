package swap

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcom "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	//contractabi "github.com/TotemFi/totem-bridge-offchain/abi"
	"github.com/TotemFi/totem-bridge-offchain/model"
	"github.com/TotemFi/totem-bridge-offchain/util"
)

func buildSwapPairInstance(pairs []model.SwapPair) (map[ethcom.Address]*SwapPairIns, error) {
	swapPairInstances := make(map[ethcom.Address]*SwapPairIns, len(pairs))

	for _, pair := range pairs {

		lowBound := big.NewInt(0)
		_, ok := lowBound.SetString(pair.LowBound, 10)
		if !ok {
			panic(fmt.Sprintf("invalid lowBound amount: %s", pair.LowBound))
		}
		upperBound := big.NewInt(0)
		_, ok = upperBound.SetString(pair.UpperBound, 10)
		if !ok {
			panic(fmt.Sprintf("invalid upperBound amount: %s", pair.LowBound))
		}

		swapPairInstances[ethcom.HexToAddress(pair.BEP20Addr)] = &SwapPairIns{
			Symbol:     pair.Symbol,
			Name:       pair.Name,
			Decimals:   pair.Decimals,
			LowBound:   lowBound,
			UpperBound: upperBound,
			BEP20Addr:  ethcom.HexToAddress(pair.BEP20Addr),
			ERC20Addr:  ethcom.HexToAddress(pair.ERC20Addr),
		}

		util.Logger.Infof("Load swap pair, symbol %s, bep20 address %s, erc20 address %s", pair.Symbol, pair.BEP20Addr, pair.ERC20Addr)
	}

	return swapPairInstances, nil
}

//func GetKeyConfig(cfg *util.Config) (*util.KeyManagerConfig, error) {
//	return &util.KeyManagerConfig{
//		HMACKey:       cfg.KeyManagerConfig.HMACKey,
//		BSCPrivateKey: cfg.KeyManagerConfig.BSCPrivateKey,
//		MTSPrivateKey: cfg.KeyManagerConfig.MTSPrivateKey,
//	}, nil
//}

func abiEncodeFillMETIS2BSCSwap(amount, fee, exchange, nonce, deadline *big.Int, account ethcom.Address, metisTxHash ethcom.Hash, abi *abi.ABI) ([]byte, error) {

	message := struct {
		SwapType string
		Base     string   // MTS/WBNB
		Quote    string   // TOTM
		Amount   *big.Int // metis amount
		Fee      *big.Int // fee in totem
		Exchange *big.Int // totem amount
		Nonce    *big.Int
		Deadline *big.Int
		Account  ethcom.Address
	}{
		SwapType: "peg-in",
		Base:     "WBNB",
		Quote:    "TOTM",
		Amount:   amount,
		Fee:      fee,
		Exchange: exchange,
		Nonce:    nonce,
		Deadline: deadline,
		Account:  account,
	}

	data, err := abi.Pack("fillBNB2TOTMPegin", message, metisTxHash)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func abiEncodeFillMTS2TOTMPegIn(bscTxHash, dataHash ethcom.Hash,
	recipient ethcom.Address, fee, exchange *big.Int, swapType string, abi *abi.ABI) ([]byte, error) {
	data, err := abi.Pack("fillMTS2TOTMPegIn", recipient, dataHash, bscTxHash, swapType, fee, exchange)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func buildSignedTransaction(contract ethcom.Address, ethClient *ethclient.Client, txInput []byte,
	privateKey *ecdsa.PrivateKey, chainID, gasPrice *big.Int, gasLimit int64) (*types.Transaction, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	nonce, err := ethClient.PendingNonceAt(context.Background(), txOpts.From)
	if err != nil {
		return nil, err
	}

	if gasPrice == nil {
		gasPrice, err = ethClient.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, err
		}
	}
	value := big.NewInt(0)
	msg := ethereum.CallMsg{From: txOpts.From, To: &contract, GasPrice: gasPrice, Value: value, Data: txInput}
	estimateGasLimit, err := ethClient.EstimateGas(context.Background(), msg)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas needed: %v", err)
		//util.Logger.Errorf("failed to estimate gas needed: %s", err.Error(), err)
		//estimateGasLimit = uint64(gasLimit)
	}

	rawTx := types.NewTransaction(nonce, contract, value, estimateGasLimit, gasPrice, txInput)
	//signedTx, err := txOpts.Signer(types.HomesteadSigner{}, txOpts.From, rawTx)
	signedTx, err := txOpts.Signer(txOpts.From, rawTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func buildNativeCoinTransferTx(contract ethcom.Address, ethClient *ethclient.Client, value *big.Int, privateKey *ecdsa.PrivateKey, chainID *big.Int) (*types.Transaction, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	nonce, err := ethClient.PendingNonceAt(context.Background(), txOpts.From)
	if err != nil {
		return nil, err
	}
	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{From: txOpts.From, To: &contract, GasPrice: gasPrice, Value: value}
	gasLimit, err := ethClient.EstimateGas(context.Background(), msg)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas needed: %v", err)
	}

	rawTx := types.NewTransaction(nonce, contract, value, gasLimit, gasPrice, nil)
	signedTx, err := txOpts.Signer(txOpts.From, rawTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

//func queryDeployedBEP20ContractAddr(erc20Addr ethcom.Address, bscSwapAgentAddr ethcom.Address, txRecipient *types.Receipt, bscClient *ethclient.Client) (ethcom.Address, error) {
//	swapAgentInstance, err := contractabi.NewBSCSwapAgentImpl(bscSwapAgentAddr, bscClient)
//	if err != nil {
//		return ethcom.Address{}, err
//	}
//	//if len(txRecipient.Logs) != 2 {
//	//	return ethcom.Address{}, fmt.Errorf("Expected tx logs length in recipient is 2, actual it is %d", len(txRecipient.Logs))
//	//}
//	createSwapEvent, err := swapAgentInstance.ParseSwapPairCreated(*txRecipient.Logs[0])
//	if err != nil || createSwapEvent == nil {
//		for i := 1; i < len(txRecipient.Logs); i++ {
//			createSwapEvent, err = swapAgentInstance.ParseSwapPairCreated(*txRecipient.Logs[i])
//			if err == nil || createSwapEvent != nil {
//				break
//			}
//		}
//	}
//
//	//createSwapEvent, err := swapAgentInstance.ParseSwapPairCreated(*txRecipient.Logs[1])
//	if err != nil || createSwapEvent == nil {
//		return ethcom.Address{}, err
//	}
//
//	util.Logger.Debugf("Deployed bep20 contact %s for register erc20 %s", createSwapEvent.Bep20Addr.String(), erc20Addr.String())
//	return createSwapEvent.Bep20Addr, nil
//}

func BuildKeys(privateKeyStr string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	if strings.HasPrefix(privateKeyStr, "0x") {
		privateKeyStr = privateKeyStr[2:]
	}
	priKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, nil, err
	}
	publicKey, ok := priKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, fmt.Errorf("get public key error")
	}
	return priKey, publicKey, nil
}
