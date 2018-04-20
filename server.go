
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
	// This is the json key of the RINKEBY address 0x419a0372b09bc55e74e4ec1ee41dfcf952245774 (specially created to test this app)
	// You should NEVER push your mainnet key to any GitHub repository
	key = `{"address":"419a0372b09bc55e74e4ec1ee41dfcf952245774","crypto":{"cipher":"aes-128-ctr","ciphertext":"7ac7cb5e66f70368398a886eb4243e0beaf2e9515cb3f8d6bc59e69bf8f8d4c3","cipherparams":{"iv":"46a70bb5cd4caa3363a86e9fb54eea17"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"429e87e3e942822b698ced88a4991b0fa05f4de8539ca1a752b1f0fb60821798"},"mac":"67620de566b0949dde05b545c3f2ce54142f9958b37d7634cc7db33d7101319f"},"id":"bd8ca65e-4071-445e-9286-fb33f9c5c060","version":3}`
	infuraEndpoint = "https://rinkeby.infura.io/SlVwORwAFmkf4EHiccur"
	localEndpoint = "/home/mexskican/.ethereum/testnet/geth.ipc"
	bikeContractAddress = "0x68954caada95c39c268896907edde08e9ef86081"
	tokenContractAddress = "0xf039f64cc421f3a2016570578fe5d4983d73a5fe"
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
	// The function has to be called with the contract owner address
	tx, err := bikeSession.AddBike(big.NewInt(newBikeId))
	if err != nil {
		log.Fatalf("Failed to add a bike to the smart contract: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
	fmt.Fprintf(w, "Transfer pending: 0x%x\n", tx.Hash())
}

func setContracts() {
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
	// "hello" is the password for the RINKEBY address 0x419a0372b09bc55e74e4ec1ee41dfcf952245774
	// You should NEVER push your mainnet password to any GitHub repository
	auth, err := bind.NewTransactor(strings.NewReader(key), "hello")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	bikeSession = &bike.BikeSession{
		Contract: bikeContract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: 100000,
		},
	}
	tokenSession = &token.BikeTokenSession{
		Contract: tokenContract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{},
	}
}

func main() {
	setContracts()

	http.HandleFunc("/view-bike/", bikeHandler)
	http.HandleFunc("/view-token/", tokenHandler)
	http.HandleFunc("/add-bike/", addBikeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
