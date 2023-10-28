package sessions

import "designPatter/bridge/Lri"

type Session interface {
	Read()
	SetLri(lri Lri.Lri)
}

