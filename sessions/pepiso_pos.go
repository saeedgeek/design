package sessions

import (
	"designPatter/bridge/Lri"
	"fmt"
)

type SessionPepisoPos struct {
	lri Lri.Lri
}

func (s *SessionPepisoPos) Read() {
	fmt.Println("SessionPepisoPos")
	s.lri.GetPacket()
}

func (s *SessionPepisoPos) SetLri(l Lri.Lri) {
	s.lri = l
}
