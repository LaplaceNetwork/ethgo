package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/openzknetwork/ethgo/keystore"
	"github.com/urfave/cli"
)

func getEther(address string) error {
	resp, err := http.Get(fmt.Sprintf("http://faucet.ropsten.be:3001/donate/%s", address))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return err
}

func main() {
	app := cli.NewApp()
	app.Name = "faucet"
	app.Usage = "require ropsten eth from faucet website"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "wallet",
			Value: "./wallet.json",
			Usage: "output wallet path",
		},
		cli.StringFlag{
			Name:  "password",
			Value: "test",
			Usage: "wallet password",
		},
		cli.StringFlag{
			Name:  "address",
			Value: "",
			Usage: "special address",
		},
	}

	app.Action = func(c *cli.Context) error {
		walletpath := c.String("wallet")
		password := c.String("password")

		address := c.String("address")

		key, err := keystore.NewKey()

		if address == "" {

			if err != nil {

				return err
			}

			err = getEther(key.Address)

			if err != nil {
				return err
			}
		} else {
			err = getEther(address)

			if err != nil {
				return err
			}
		}

		if address == "" {
			data, err := keystore.WriteScryptKeyStore(key, password)

			if err != nil {
				return err
			}

			return ioutil.WriteFile(walletpath, data, 0777)
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("run app err %s\n", err)
		return
	}
}
