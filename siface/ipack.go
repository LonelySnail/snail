package siface

import "io"

type IPack interface {
	GetHeaderLen()  uint32
	GetDataLen()  uint16
	Pack(msg IMessage) ([]byte,error)
	UnPack(r io.Reader) (IMessage,error)
}
