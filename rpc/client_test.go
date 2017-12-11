package rpc

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dynamicgo/config"
	"github.com/stretchr/testify/assert"
)

var cnf *config.Config
var client *Client

func init() {
	cnf, _ = config.NewFromFile("./test.json")
	client = NewClient(cnf.GetString("server", "http://47.100.110.75:8545"))
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

	blocknumber, err := client.GetBlockByNumber(11)

	assert.NoError(t, err)

	printResult(blocknumber)
	printResult(fmt.Sprintf("0x%2x", 0))
}

func TestGetTransactionByHash(t *testing.T) {

	tx, err := client.GetTransactionByHash("0xa28142d727fe9c0f86d26023cc26f3a80c89d3860c36fb98d66ce5a0bd7a706d")

	assert.NoError(t, err)

	printResult(tx)
}

func printResult(result interface{}) {

	data, _ := json.MarshalIndent(result, "", "\t")

	fmt.Println(string(data))
}
