// Information about the algorithm is available on Wikipedia
//
// https://en.wikipedia.org/wiki/Blockchain
//
package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"time"
)

var Blockchain []Block

type Block struct {
	Index     int
	Timestamp string
	Data      []byte
	Hash      []byte
	PrevHash  []byte
}

func calculateHash(block Block) []byte {
	record := new(bytes.Buffer)
	binary.Write(record, binary.LittleEndian, block.Index)
	binary.Write(record, binary.LittleEndian, block.Timestamp)
	binary.Write(record, binary.LittleEndian, block.Data)
	binary.Write(record, binary.LittleEndian, block.PrevHash)

	h := sha256.New()
	h.Write(record.Bytes())
	hashed := h.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(hashed)))
	hex.Encode(dst, hashed)
	return dst
}

func generateBlock(current Block, data []byte) (Block, error) {
	var next Block

	t := time.Now()

	next.Index = current.Index + 1
	next.Timestamp = t.String()
	next.Data = data
	next.PrevHash = current.Hash
	next.Hash = calculateHash(next)

	return next, nil
}

func isBlockValid(next, current Block) bool {
	if current.Index+1 != next.Index {
		return false
	}

	if bytes.Compare(current.Hash, next.PrevHash) != 0 {
		return false
	}

	if bytes.Compare(calculateHash(next), next.Hash) != 0 {
		return false
	}

	return true
}

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

func WriteBlock(data []byte) error {
	block, err := generateBlock(Blockchain[len(Blockchain)-1], data)
	if err != nil {
		return err
	}

	if isBlockValid(block, Blockchain[len(Blockchain)-1]) {
		newBlockchain := append(Blockchain, block)
		replaceChain(newBlockchain)
	}

	return nil
}
