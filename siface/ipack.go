package siface

type IPack interface {
	GetHeaderLen()  uint32
	Pack(msg IMessage) ([]byte,error)
	UnPack([]byte) (IMessage,error)
}
