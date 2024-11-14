package blockchain

import (
	"bytes"
	"crypto/md5"
)

type Block struct {
	Hash     string
	Data     string
	PrevHash string
}

func (b *Block) ComputeHash() {
	concatenatedData := bytes.Join([][]byte{[]byte(b.Data), []byte(b.PrevHash)}, []byte{})
	computeHash := md5.Sum(concatenatedData)
	b.Hash = string(computeHash[:])
}

func CreateBlock(data, prevHash string) *Block {
	newBlock := &Block{"", data, prevHash}

	newBlock.ComputeHash()
	return newBlock
}

func Genesis() *Block {
	return CreateBlock("Genesis", "")
}
