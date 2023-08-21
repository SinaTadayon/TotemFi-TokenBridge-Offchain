package app

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	contracts "github.com/TotemFi/totem-bridge-offchain/abi"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcmm "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/TotemFi/totem-bridge-offchain/swap"
	"github.com/TotemFi/totem-bridge-offchain/util"
)

const (
	DefaultListenAddr = "0.0.0.0:8080"
	MaxIconUrlLength  = 400
)

type ErrMsg struct {
	Msg string `json:"message"`
}

type App struct {
	DB *gorm.DB

	cfg *util.Config

	hmacSigner *util.HmacSigner
	swapEngine *swap.SwapEngine

	//accountMap map[string]*dto.PeginStateResponseDto
	nonceMap   map[string]*big.Int
	priceUtils *util.SwapPriceUtils

	MetisSwapAgentAddr    ethcmm.Address
	MetisBridgeAgent      *contracts.MetisBridgeAgentImpl
	MetisSwapAgentAbi     abi.ABI
	MetisChainID          int64
	MetisClient           *ethclient.Client
	MetisPrivateKey       *ecdsa.PrivateKey
	MetisPublicKey        *ecdsa.PublicKey
	MetisAdminAddress     ethcmm.Address
	MetisLowerBoundAmount *big.Int
	MetisUpperBoundAmount *big.Int
}

func NewApp(config *util.Config, db *gorm.DB, signer *util.HmacSigner, swapEngine *swap.SwapEngine,
	priceUtils *util.SwapPriceUtils, metisClient *ethclient.Client) *App {

	//mtsPrivateKey, mtsPublicKey, err := buildKeys(config.KeyManagerConfig.MTSPrivateKey)
	mtsPrivateKey, mtsPublicKey, err := buildKeys(config.KeyManagerConfig.MTSPrivateKey)
	if err != nil {
		panic(err.Error())
	}

	address := crypto.PubkeyToAddress(*mtsPublicKey)
	//util.Logger.Info("Signer Address: ", address)

	lowerBound := big.NewInt(config.ChainConfig.MTSLowerBoundAmount)
	lowerBound = lowerBound.Mul(big.NewInt(10).Exp(big.NewInt(10), big.NewInt(18), nil), lowerBound)

	upperBound := big.NewInt(config.ChainConfig.MTSUpperBoundAmount)
	upperBound = upperBound.Mul(big.NewInt(10).Exp(big.NewInt(10), big.NewInt(18), nil), upperBound)

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

	return &App{
		DB:                    db,
		cfg:                   config,
		hmacSigner:            signer,
		swapEngine:            swapEngine,
		nonceMap:              make(map[string]*big.Int),
		priceUtils:            priceUtils,
		MetisSwapAgentAddr:    ethcmm.HexToAddress(config.ChainConfig.MTSSwapAgentAddr),
		MetisBridgeAgent:      metisSwapAgentInst,
		MetisSwapAgentAbi:     metisAgentAbi,
		MetisClient:           metisClient,
		MetisPrivateKey:       mtsPrivateKey,
		MetisPublicKey:        mtsPublicKey,
		MetisAdminAddress:     address,
		MetisLowerBoundAmount: lowerBound,
		MetisUpperBoundAmount: upperBound,
		MetisChainID:          metisChainID.Int64(),
	}
}

func (app *App) checkAuth(r *http.Request) ([]byte, error) {
	apiKey := r.Header.Get("APIKey")
	hash := r.Header.Get("Authorization")

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if app.hmacSigner.ApiKey != apiKey {
		return nil, fmt.Errorf("api key mismatch")
	}

	if !app.hmacSigner.Verify(payload, hash) {
		return nil, fmt.Errorf("invalud auth")
	}
	return payload, nil
}

func buildKeys(privateKeyStr string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
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
