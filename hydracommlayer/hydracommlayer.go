package hydracommlayer

import(
	// "MasteringGoTutorial/HYDRA/hydracommlayer/hydraproto"
)
const (
	Protobuf uint8 = iota
)

type HydraConnection interface{
	EncodeAndSend(obj interface{}, destination string) error
	ListenAndDecode(Listenaddress string) (chan interface{}, error)
}

//factory design pattern
func NewConnection(connType uint8) HydraConnection{
	switch connType{
	case Protobuf:
		// return hydraproto.NewProtoHandler()
	}
	return nil
}