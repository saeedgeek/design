package sessions

import (
	"designPatter/bridge/Lri"
	"fmt"
)

type ShaparaksessionV1 struct {
	lri Lri.Lri
}

func (s *ShaparaksessionV1) Read() {
	fmt.Println("ShaparaksessionV1")
	s.lri.GetPacket()
}

func (s *ShaparaksessionV1) SetLri(l Lri.Lri) {
	s.lri = l
}
