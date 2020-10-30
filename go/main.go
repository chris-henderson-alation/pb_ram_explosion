package main // import "github.com/chris-henderson-alation/pb_ram_explosion"

//go:generate protoc --proto_path=. --go_out=. example.proto

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	pid := os.Getpid()
	fmt.Printf("htop --pid %v\n", pid)
	b := bufio.NewReader(os.Stdin)
	b.ReadLine()
	//message := MySmallMessage{}
	//fmt.Printf("%v", message)
	b.ReadLine()
}
