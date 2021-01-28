package rpc

import "goo/example/rpc/codec"

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber int
	CodeType    codec.Type
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodeType:    codec.GobType,
}

type Serve struct{}
