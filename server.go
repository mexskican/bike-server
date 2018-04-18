
package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"./bike"
	"./token"
)

const key = `paste the contents of your *testnet* key json here`

func bikeHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := ethclient.Dial("https://rinkeby.infura.io/SlVwORwAFmkf4EHiccur")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}
	contract, err := bike.NewBike(common.HexToAddress("0x903b552d9dba24bbb908a064d8568b76923d5501"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	// TODO: Add regular expression to validate the params
	method := r.URL.Path[len("/view-bike/"):]
	if strings.Compare("hourlycost", method) == 0 {
		hourlyCost, _ := contract.HourlyCost(&bind.CallOpts{})
		fmt.Fprintf(w, "<h1>Hourly Cost:</h1><div>%d</div>", hourlyCost)
	} else if strings.HasPrefix(method, "isavailable") {
		id, _ := strconv.ParseInt(r.URL.Path[len("/view/isavailable/"):], 10, 64)
		isAvailable, _ := contract.BikeIsAvailable(&bind.CallOpts{}, big.NewInt(id))
		fmt.Fprintf(w, "<h1>Is available:</h1><div>%t</div>", isAvailable)
	} else if strings.HasPrefix(method, "bikeid") {
		address := r.URL.Path[len("/view-bike/bikeid/"):]
		bikeId, _ := contract.GetRentalBikeId(&bind.CallOpts{}, common.HexToAddress(address))
		fmt.Fprintf(w, "<h1>Bike Id of %s:</h1><div>%d</div>", address, bikeId)
	} else if strings.HasPrefix(method, "returntime") {
		address := r.URL.Path[len("/view-bike/returntime/"):]
		returnTime, _ := contract.GetReturnTime(&bind.CallOpts{}, common.HexToAddress(address))
		fmt.Fprintf(w, "<h1>Return Time of %s:</h1><div>%d</div>", address, returnTime)
	} else {
		fmt.Fprintf(w, "<h1>Wrong Parameter<h1>")
	}
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := ethclient.Dial("https://rinkeby.infura.io/SlVwORwAFmkf4EHiccur")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}
	contract, err := token.NewBikeToken(common.HexToAddress("0x6bab74f46d46cbb495d61a3c730b657e46f605fa"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	method := r.URL.Path[len("/view-token/"):]
	if strings.Compare("name", method) == 0 {
		name, _ := contract.Name(&bind.CallOpts{})
		fmt.Fprintf(w, "<h1>Token Name:</h1><div>%s</div>", name)
	} else if strings.Compare("price", method) == 0 {
		price, _ := contract.PRICE(&bind.CallOpts{})
		fmt.Fprintf(w, "<h1>Token price:</h1><div>%d</div>", price)
	} else if strings.HasPrefix(method, "balance") {
		address := r.URL.Path[len("/view-bike/balance/"):]
		balance := big.NewInt(0)
		balance, _ = contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
		fmt.Fprintf(w, "<h1>Balance of %s:</h1><div>%d</div>",address, balance)
	} else {
		fmt.Fprintf(w, "<h1>Wrong Parameter<h1>")
	}
}

func addBikeHandler(w http.ResponseWriter, r *http.Request) {
	// Need to run an ETH node locally
	// Create an IPC based RPC connection to a remote node and instantiate a contract binding
	// Change the path with your own
	conn, err := ethclient.Dial("/home/mexskican/.ethereum/testnet/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	contract, err := bike.NewBike(common.HexToAddress("0x903b552d9dba24bbb908a064d8568b76923d5501"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}
	// Need to use the account of the owner to add a bike
	auth, err := bind.NewTransactor(strings.NewReader(key), "my password")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	tx, err := contract.AddBike(auth, big.NewInt(123))
	if err != nil {
		log.Fatalf("Failed to add a bike to the smart contract: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
	fmt.Fprintf(w, "Transfer pending: 0x%x\n", tx.Hash())
}

func main() {
	http.HandleFunc("/view-bike/", bikeHandler)
	http.HandleFunc("/view-token/", tokenHandler)
	http.HandleFunc("/add-bike/", addBikeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
