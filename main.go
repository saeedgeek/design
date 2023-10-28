package main

import (
	"designPatter/bridge/Lri"
	"designPatter/bridge/sessions"
)

func main() {
	session := &sessions.ShaparaksessionV1{

	}
	lriWitHeadert := &Lri.WithHeader{}
	session.SetLri(lriWitHeadert)

	session.Read()
}
