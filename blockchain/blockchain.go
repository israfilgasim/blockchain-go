package blockchain

type Blockchain struct {
	Blocks []*Block
}

func InitBlockchain() *Blockchain {
	veryFirstBlock := Genesis()

	blockchainObj := &Blockchain{[]*Block{veryFirstBlock}}
	return blockchainObj
}

func (chain *Blockchain) AddNewBlock(data string) {
	l := len(chain.Blocks) - 1

	prevBlock := chain.Blocks[l]

	newBlock := CreateBlock(data, prevBlock.PrevHash)
	chain.Blocks = append(chain.Blocks, newBlock)

}
