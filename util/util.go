package util

import (
	"bytes"
	crypto "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/big"
	"math/rand"
)

//int64转换成字节数组
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

//字节数组转换为int
func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

//生成随机数
func GenerateRealRandom() int64 {
	n, err := crypto.Int(crypto.Reader, big.NewInt(1000000000000000000))
	if err != nil {
		fmt.Println(err)
	}
	return n.Int64()
}

func RandomInRange(min int, max int) int {
	var tmp int
	if max < min {
		tmp = max
		max = min
		min = tmp
	}
	return rand.Intn(max-min+1) + min
}
