package main

import (
	"github.com/huin/goserial"
	"log"
)

func do() {
	parity:=goserial.ParityNone
	if Parity==1{
		parity=goserial.ParityEven
	}else if Parity==1 {
		parity = goserial.ParityOdd
	}
	c := &goserial.Config{Name: ComName, Baud: Baud,Parity:parity}
	s, err := goserial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("%q", buf[:n])
}
