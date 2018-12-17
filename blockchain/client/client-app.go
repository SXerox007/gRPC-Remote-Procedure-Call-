package main

import (
	"context"
	"log"

	blockchain "gRPC-Remote-Procedure-Call-/blockchain/protos"

	"google.golang.org/grpc"
)

func main() {
	Init()
}

func Init() {
	ClientSetup()
}

func ClientSetup() {
	client, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error in connection : ", err)
	}
	defer client.Close()

	msg := blockchain.NewBlockChainClient(client)

	//Add block
	AddBlockInBlockChain(msg)

	//Get All blocks
	GetBlocks(msg)

}

func AddBlockInBlockChain(msg blockchain.BlockChainClient) {
	req := &blockchain.AddBlockChainRequest{}
	res, err := msg.AddBlockChain(context.Background(), req)
	if err == nil {
		log.Println("Hash Value: ", res.GetHashValue())
	} else {
		log.Println("Error: ", err)
	}
}

func GetBlocks(msg blockchain.BlockChainClient) {
	req := &blockchain.GetBlockChainRequest{
		CommonRequest: &blockchain.CommonRequest{
			IsFetchAllBlocks: true,
		},
	}
	res, err := msg.GetBlockChainData(context.Background(), req)
	if err == nil {
		log.Println("Data: ", res.GetCommonResponse())
	} else {
		log.Println("Error: ", err)
	}

}
