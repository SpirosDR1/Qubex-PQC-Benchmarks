package main

import (
 "context"
 "fmt"
 "log"
 "math/big"
 "os"
 "sync"

 "github.com/ethereum/go-ethereum/core/types"
 "github.com/ethereum/go-ethereum/crypto"
 "github.com/ethereum/go-ethereum/ethclient"
)

type TargetChain struct {
 Name    string
 RPCURL  string
 ChainID int64
}

var mu sync.Mutex

func main() {
 privKeyHex := os.Getenv("PRIVATE_KEY")
 privateKey, err := crypto.HexToECDSA(privKeyHex)
 if err != nil {
  log.Fatal("Invalid Private Key")
 }
 fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

 networks := []TargetChain{
  {"Base", "https://mainnet.base.org", 8453}, // Note: QUBEX Production Chaos Engine utilizes Alchemy Enterprise RPCs for mempool interception. Benchmark defaults to standard RPC.
  {"Ethereum", "https://ethereum.publicnode.com", 1},
  {"BSC", "https://bsc-dataseed.binance.org/", 56},
  {"Arbitrum", "https://arb1.arbitrum.io/rpc", 42161},
  {"Optimism", "https://mainnet.optimism.io", 10},
  {"Polygon", "https://polygon.drpc.org", 137},
 }

 var wg sync.WaitGroup
 for _, net := range networks {
  wg.Add(1)
  go func(target TargetChain) {
   defer wg.Done()
   client, _ := ethclient.Dial(target.RPCURL)
   nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)
   gasPrice, _ := client.SuggestGasPrice(context.Background())
   
   bumpedGas := new(big.Int).Mul(gasPrice, big.NewInt(120))
   bumpedGas.Div(bumpedGas, big.NewInt(100))

   tx := types.NewTx(&types.LegacyTx{
    Nonce:    nonce,
    GasPrice: bumpedGas,
    Gas:      100000,
    To:       &fromAddress,
    Value:    big.NewInt(0),
   })
   
   signedTx, _ := types.SignTx(tx, types.LatestSignerForChainID(big.NewInt(target.ChainID)), privateKey)
   err := client.SendTransaction(context.Background(), signedTx)
   
   mu.Lock()
   if err == nil {
    fmt.Printf("%-10s | %s\n", target.Name, signedTx.Hash().Hex())
   } else {
    fmt.Printf("%-10s | FAILED\n", target.Name)
   }
   mu.Unlock()
  }(net)
 }
 wg.Wait()
}
