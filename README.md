# Bike Server
The Bike contract is deployed on Rinkeby at [0x68954caada95c39c268896907edde08e9ef86081](https://rinkeby.etherscan.io/address/0x68954caada95c39c268896907edde08e9ef86081)   
The Bike Token contract is deployed on Rinkeby at [0xf039f64cc421f3a2016570578fe5d4983d73a5fe](https://rinkeby.etherscan.io/address/0xf039f64cc421f3a2016570578fe5d4983d73a5fe)

## Read only functions
The read only functions are using Infura as an Ethereum provider, therefore the user does not need to run a node locally to call those functions.

## Send a tansaction
A bike can be added to the smart contract using this server. To run this function a Ethereum node needs to be runed locally and the eth account private key should be added.
 
## Usage
1. Build
```
$ go build server.go
```
2. Run the server
```
$ ./server
```
3. Send some requests
```
$ curl http://localhost:8080/view-bike/hourlycost
$ curl http://localhost:8080/view-bike/isavailable/55
$ curl http://localhost:8080/view-bike/bikeid/0xe8076f19c37f318f59dffb71870f7f9b1ac817ba
$ curl http://localhost:8080/view-bike/returntime/0xe8076f19c37f318f59dffb71870f7f9b1ac817ba

$ curl http://localhost:8080/view-token/name
$ curl http://localhost:8080/view-token/price
$ curl http://localhost:8080/view-token/balance/0xe8076f19c37f318f59dffb71870f7f9b1ac817ba
```
Transactions can be send using this server. The key and password field can be changed to use a different account, but the addBike method will only suceed with the contract owner account.
```
$ curl http://localhost:8080/add-bike/123
```
The server will prepare and sign the transaction offchain and send the raw transaction to the Rinkeby network using Infura.
