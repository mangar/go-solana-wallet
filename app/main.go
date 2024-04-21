package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var privateKeyString = "5NeC7REGNGHpgarHVsU1cWMQnGMkRCt7ZMFyc66BDovQzQ1AFf4SuHPL5TCnny3W3ozFhLn3gZWa1YqnSM3pkzK8"
var publicKeyString = "HbGTjuZtr5Jyzwgi7wSsonmQyp4hrTtjF2D14QGXepox"
func main() {

	// 1
	// generateKeys()

	// 2
	// importKey()

	// 3
	getBalance()

}


func getBalance() {


    // Create a new account
    account := solana.NewWallet()
    fmt.Println("account private key:", account.PrivateKey)
    fmt.Println("account public key:", account.PublicKey())

    // Create a new RPC client on the DevNet
    rpcClient := rpc.New(rpc.DevNet_RPC)

    // Airdrop 1 SOL to the new account:
    out, err := rpcClient.RequestAirdrop(
        context.TODO(),
        account.PublicKey(),
		// publicKeyString,
        // 1 sol = 1000000000 lamports
        solana.LAMPORTS_PER_SOL*1,
        rpc.CommitmentFinalized,
    )

    if err != nil {
        panic(err)
    }

    fmt.Println("airdrop transaction signature:", out)
    time.Sleep(time.Second * 1)

    // Get the balance of the account
    balance, err := rpcClient.GetBalance(
        context.TODO(),
        account.PublicKey(),
        rpc.CommitmentFinalized,
    )

    if err != nil {
        panic(err)
    }

    var lamportsOnAccount = new(big.Float).SetUint64(uint64(balance.Value))
    var solBalance = new(big.Float).Quo(lamportsOnAccount, new(big.Float).SetUint64(solana.LAMPORTS_PER_SOL))

    fmt.Println("Wallet Balance:", solBalance, "SOL")


}


func importKey() {
    privateKey, err := solana.PrivateKeyFromBase58(privateKeyString)
    if err != nil {
        panic(err)
    }

    // Get the public key from the private key
    publicKey := privateKey.PublicKey()

    // Print the public and private keys
    println("Public Key:", publicKey.String())	
}


func generateKeys() {
    privateKey, err := solana.NewRandomPrivateKey()
    if err != nil {
        panic(err)
    }

    // Get the public key from the private key
    publicKey := privateKey.PublicKey()

    // Print the public and private keys (**keep private key secure!**)
    fmt.Println("Public Key:", publicKey)
    fmt.Println("Private Key:", privateKey)
}