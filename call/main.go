package main

import (
	"greeting"
	"greeting/kazakh"
	"greeting/russian"
)

func main() {
	greeting.SayHi()
	russian.SayRussianHi()
	kazakh.SayKazakhHi()
}
