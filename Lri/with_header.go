package Lri

type WithHeader struct {
}

func (w *WithHeader) GetPacket() {
	print("WithHeader Lri Geting Packet ... ")
}
