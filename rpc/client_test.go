package rpc

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var client *Client

func init() {
	// cnf, _ = config.NewFromFile("../../conf/test.json")
	client = NewClient("https://ropsten.infura.io/OTFK50Z1PCljMOeEAlA9")
}

func TestBalance(t *testing.T) {

	accoutState, err := client.GetBalance("0xfbb1b73c4f0bda4f67dca266ce6ef42f520fbb98")

	assert.NoError(t, err)

	printResult(accoutState)
}

func TestBlockNumber(t *testing.T) {

	blocknumber, err := client.BlockNumber()

	assert.NoError(t, err)

	printResult(blocknumber)
}

func TestBlocksPerSecond(t *testing.T) {

	blocknumber, err := client.BlockPerSecond()

	assert.NoError(t, err)

	printResult(blocknumber)
}

func TestGetBlockByNumber(t *testing.T) {

	blocknumber, err := client.GetBlockByNumber(3514168)

	assert.NoError(t, err)

	printResult(blocknumber)
}

func TestGetTransactionByHash(t *testing.T) {

	tx, err := client.GetTransactionByHash("0x0bcbe6ba69251b7269df067849784628595c527962d5f94d57199697d5085cd2")

	assert.NoError(t, err)

	printResult(tx)
}

func TestGetTransactionReceipt(t *testing.T) {

	tx, err := client.GetTransactionReceipt("0x0bcbe6ba69251b7269df067849784628595c527962d5f94d57199697d5085cd2")

	assert.NoError(t, err)

	printResult(tx)
}

func TestNonce(t *testing.T) {

	tx, err := client.NonceWithTag("0x8e488bf5e68f54e8553e827dad3a615967628d11", "latest")

	assert.NoError(t, err)

	printResult(tx)
}

func printResult(result interface{}) {

	data, _ := json.MarshalIndent(result, "", "\t")

	fmt.Println(string(data))
}
