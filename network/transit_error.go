package network

import (
	"bytes"
	"encoding/gob"
	log "github.com/corgi-kx/logcustom"
)

type myerror struct {
	Error    string
	Addrfrom string
}

func (v myerror) serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(v)
	if err != nil {
		panic(err)
	}
	return result.Bytes()
}

func (v *myerror) deserialize(d []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(v)
	if err != nil {
		log.Panic(err)
	}
}
