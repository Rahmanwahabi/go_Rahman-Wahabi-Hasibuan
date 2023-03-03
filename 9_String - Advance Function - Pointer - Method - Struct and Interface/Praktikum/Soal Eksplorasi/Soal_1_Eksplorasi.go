package main

import (
	"fmt"
	"strings"
)

type Chiper interface {
	Encode() string
	Decode() string
}

type student struct {
	name       string
	nameEncode string
	score      int
}

func (s *student) Encode() string {
	var nameEncode strings.Builder

	for _, c := range s.name {
		if c >= 'a' && c <= 'z' {
			c = 'a' + (c-'a'+3)%26
		}
		nameEncode.WriteRune(c)
	}

	s.nameEncode = nameEncode.String()

	return s.nameEncode
}

func (s *student) Decode() string {
	var nameDecode strings.Builder

	for _, c := range s.nameEncode {
		if c >= 'a' && c <= 'z' {
			c = 'a' + (c-'a'+23)%26
		}
		nameDecode.WriteRune(c)
	}

	return nameDecode.String()
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of student’s name " + a.nameEncode + " is : " + c.Decode())
	}
}
