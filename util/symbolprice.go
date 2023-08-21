package util

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"
)

type SymbolPrice struct {
	PriceMap            sync.Map
	CoinMarketCapApiKey string
	lastCallTimestamp   time.Time
	firstCall           bool
}

// GetSymbolPriceFromMarketCap - checks if there already is a price for the requested symbol in the priceMap
// if not, fetches the price from coinmarketcap api and stores it in the priceMap
func (symbolPrice *SymbolPrice) GetSymbolPriceFromMarketCap(token string) (float64, error) {
	symbol := strings.ToUpper(token)

	//now := time.Now().UTC().Unix()
	//before := time.Now().UTC().Add(-1 * time.Hour).Unix()
	//req, _ := http.NewRequest("GET", fmt.Sprintf("https://pro-api.coinmarketcap.com/v2/cryptocurrency/ohlcv/historical?time_period=hourly&symbole=%s&time_start=%d&end_time=%d&interval=1h", symbol, before, now), nil)
	if symbolPrice.firstCall {
		symbolPrice.firstCall = false
		symbolPrice.lastCallTimestamp = time.Now()

	} else {
		if time.Now().Sub(symbolPrice.lastCallTimestamp) < 60 {
			time.Sleep(60 * time.Second)
			symbolPrice.lastCallTimestamp = time.Now()
		}
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=%s", symbol), nil)
	req.Header.Add("X-CMC_PRO_API_KEY", symbolPrice.CoinMarketCapApiKey)
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(req)
	if err != nil {
		Logger.Errorf("The HTTP request failed with error %s", err)
		return 0, err
	}

	if response.StatusCode != http.StatusOK {
		data, _ := ioutil.ReadAll(response.Body)
		Logger.Errorf("The HTTP request failed with error %d, message: %s", response.StatusCode, data)
		return 0, fmt.Errorf("the HTTP request failed with error %s", response.Status)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	data, _ := ioutil.ReadAll(response.Body)
	mapRes := make(map[string]interface{})
	err = json.Unmarshal(data, &mapRes)
	if err != nil {
		Logger.Errorf("The HTTP request failed with error %s", err)
		return 0, err
	}

	dataA := mapRes["data"].(map[string]interface{})
	Symbol := dataA[token].(map[string]interface{})
	quote := Symbol["quote"].(map[string]interface{})
	USD := quote["USD"].(map[string]interface{})

	return USD["price"].(float64), nil
}

func (symbolPrice *SymbolPrice) GetPrice(symbol string) decimal.Decimal {
	curPrice, ok := symbolPrice.PriceMap.Load(symbol)
	if ok {
		curPriceMap := curPrice.(map[string]interface{})
		if time.Now().UTC().Unix()-curPriceMap["time"].(int64) < 30*60 {
			return curPriceMap["price"].(decimal.Decimal)
		}
	}
	return decimal.Zero
}

func (symbolPrice *SymbolPrice) SetPrice(symbol string, value decimal.Decimal) {
	symbolPrice.PriceMap.Store(symbol, map[string]interface{}{"time": time.Now().UTC().Unix(), "price": value})
}

func BigIntViaBigFloatString(flt float64) (b *big.Int) {

	// got one more precision error with this case

	var in = big.NewFloat(flt).Text('f', 20)

	const parts = 2

	var ss = strings.SplitN(in, ".", parts)

	// protect from numbers without period
	if len(ss) != parts {
		ss = append(ss, "0")
	}

	// protect from ".0" and "0." values
	if ss[0] == "" {
		ss[0] = "0"
	}

	if ss[1] == "" {
		ss[1] = "0"
	}

	const (
		base     = 10
		fraction = 18
	)

	// get fraction length
	var fract = len(ss[1])
	if fract > fraction {
		ss[1], fract = ss[1][:fraction], fraction
	}

	in = strings.Join([]string{ss[0], ss[1]}, "")
	// convert to big integer from the string
	b, _ = big.NewInt(0).SetString(in, base)
	if fract == fraction {
		return // ready
	}
	// fract < 20, * (20 - fract)
	var (
		ten = big.NewInt(base)
		exp = ten.Exp(ten, big.NewInt(fraction-int64(fract)), nil)
	)
	b = b.Mul(b, exp)
	return

}
