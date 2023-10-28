package sessions

import (
	"designPatter/bridge/Lri"
	"fmt"
)

type ShaparaksessionV2 struct {
	lri Lri.Lri
}

func (s *ShaparaksessionV2) Read() {
	fmt.Println("ShaparaksessionV2")
	s.lri.GetPacket()
}

func (s *ShaparaksessionV2) SetLri(l Lri.Lri) {
	s.lri = l
}
