
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

var (
	bikeSession *bike.BikeSession
	tokenSession *token.BikeTokenSession
)

const (
	key = ``
	infuraEndpoint = "https://rinkeby.infura.io/SlVwORwAFmkf4EHiccur"
	localEndpoint = "/home/mexskican/.ethereum/testnet/geth.ipc"
	bikeContractAddress = "0x903b552d9dba24bbb908a064d8568b76923d5501"
	tokenContractAddress = "0x6bab74f46d46cbb495d61a3c730b657e46f605fa"
)


func bikeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Add regular expression to validate the params
	method := r.URL.Path[len("/view-bike/"):]
	if strings.Compare("hourlycost", method) == 0 {
		hourlyCost, _ := bikeSession.HourlyCost()
		fmt.Fprintf(w, "<h1>Hourly Cost:</h1><div>%d</div>", hourlyCost)
	} else if strings.HasPrefix(method, "isavailable") {
		id, _ := strconv.ParseInt(r.URL.Path[len("/view/isavailable/"):], 10, 64)
		isAvailable, _ := bikeSession.BikeIsAvailable(big.NewInt(id))
		fmt.Fprintf(w, "<h1>Is available:</h1><div>%t</div>", isAvailable)
	} else if strings.HasPrefix(method, "bikeid") {
		address := r.URL.Path[len("/view-bike/bikeid/"):]
		bikeId, _ := bikeSession.GetRentalBikeId(common.HexToAddress(address))
		fmt.Fprintf(w, "<h1>Bike Id of %s:</h1><div>%d</div>", address, bikeId)
	} else if strings.HasPrefix(method, "returntime") {
		address := r.URL.Path[len("/view-bike/returntime/"):]
		returnTime, _ := bikeSession.GetReturnTime(common.HexToAddress(address))
		fmt.Fprintf(w, "<h1>Return Time of %s:</h1><div>%d</div>", address, returnTime)
	} else {
		fmt.Fprintf(w, "<h1>Wrong Parameter<h1>")
	}
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	method := r.URL.Path[len("/view-token/"):]
	if strings.Compare("name", method) == 0 {
		name, _ := tokenSession.Name()
		fmt.Fprintf(w, "<h1>Token Name:</h1><div>%s</div>", name)
	} else if strings.Compare("price", method) == 0 {
		price, _ := tokenSession.PRICE()
		fmt.Fprintf(w, "<h1>Token price:</h1><div>%d</div>", price)
	} else if strings.HasPrefix(method, "balance") {
		address := r.URL.Path[len("/view-token/balance/"):]
		balance := big.NewInt(0)
		balance, _ = tokenSession.BalanceOf(common.HexToAddress(address))
		fmt.Fprintf(w, "<h1>Balance of %s:</h1><div>%d</div>",address, balance)
	} else {
		fmt.Fprintf(w, "<h1>Wrong Parameter<h1>")
	}
}

func addBikeHandler(w http.ResponseWriter, r *http.Request) {
	newBikeId, _ := strconv.ParseInt(r.URL.Path[len("/add-bike/"):], 10, 64)
	if newBikeId == 0 {
		fmt.Fprintf(w, "<h1>Wrong Parameter<h1>")
	}
	tx, err := bikeSession.AddBike(big.NewInt(newBikeId))
	if err != nil {
		log.Fatalf("Failed to add a bike to the smart contract: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
	fmt.Fprintf(w, "Transfer pending: 0x%x\n", tx.Hash())
}

func main() {
	conn, err := ethclient.Dial(infuraEndpoint)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}
	bikeContract, err := bike.NewBike(common.HexToAddress(bikeContractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}
	tokenContract, err := token.NewBikeToken(common.HexToAddress(tokenContractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), "")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	bikeSession = &bike.BikeSession{
		Contract: bikeContract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{},
	}
	tokenSession = &token.BikeTokenSession{
		Contract: tokenContract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{},
	}

	http.HandleFunc("/view-bike/", bikeHandler)
	http.HandleFunc("/view-token/", tokenHandler)
	http.HandleFunc("/add-bike/", addBikeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
