# Merkle-Root-Tree
Create a file named transactions.txt containing the hex-encoded transactions, one per line. Then you can simply run the program with go run main.go.

This implementation uses recursion to compute the Merkle Tree Root, which is a common approach. It also handles the case where the number of transactions is odd by duplicating the last transaction, which is a common convention.
