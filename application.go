package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "io/ioutil"
)

func main() {
    // Read the transactions from the file
    txs, err := ioutil.ReadFile("transactions.txt")
    if err != nil {
        panic(err)
    }

    // Convert the hex-encoded transactions to byte arrays
    var txBytes [][]byte
    for _, tx := range bytes.Split(txs, []byte{'\n'}) {
        txBytes = append(txBytes, tx)
    }

    // Compute the Merkle Tree Root
    root := merkleRoot(txBytes)
    fmt.Printf("Merkle Root: %x\n", root)
}

func merkleRoot(txs [][]byte) []byte {
    if len(txs) == 0 {
        return nil
    }

    // Base case: if there's only one transaction, return its hash
    if len(txs) == 1 {
        return txs[0]
    }

    // If the number of transactions is odd, duplicate the last one
    if len(txs)%2 != 0 {
        txs = append(txs, txs[len(txs)-1])
    }

    // Compute the hashes of each pair of transactions and store them in a new slice
    var hashes [][]byte
    for i := 0; i < len(txs); i += 2 {
        txPair := append(txs[i], txs[i+1]...)
        hash := sha256.Sum256(txPair)
        hashes = append(hashes, hash[:])
    }

    // Recursively compute the Merkle Tree Root of the hash slice
    return merkleRoot(hashes)
}
