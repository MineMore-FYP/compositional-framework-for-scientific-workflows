package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"log"
)

func pythonCall(progName string, inChannel chan <- string, workflowNumber string) {
	cmd := exec.Command("python3", progName, workflowNumber)
	out, err := cmd.CombinedOutput()
	log.Println(cmd.Run())

	if err != nil {
		fmt.Println(err)
    		os.Exit(3)
	}
	fmt.Println(string(out))
	msg := string(out)[:len(out)-1]
	inChannel <- msg
}

func simplePythonCall(progName string){
	cmd := exec.Command("python3", progName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os. Stderr
	log.Println(cmd.Run())
}

func messagePassing(inChannel <- chan string, outChannel chan <- string ){
	msg := <- inChannel
	outChannel <- msg
}

//reads a file and returns an array of comments beginning with ##
func readLines( progName string) [20]string{
		var commandsArray [20]string
    file, err := os.Open(progName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    i := 0
    for scanner.Scan() {
        command := scanner.Text()
				if len(command) >1{
					if command[0:2] == "##" {
						commandsArray[i] = command[2:len(command)]
						i++
					}
				}
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

		return commandsArray
}

func main(){
	fmt.Println(" +---------------------------------------------------------------+")
	fmt.Println(" |     ______        ______  		       		         |")
	fmt.Println(" |    \\    __|    _ |  ___ |                    ____ _     ___   |")
	fmt.Println(" |     \\  (   ___|_|| |_ | | ___ __  __  ___   / ___| |   |_ _|  |")
	fmt.Println(" |   __ \\  \\ /  _|_ |  _|| |/ _ \\  \\/  \\/  /  | |   | |    | |   |")
	fmt.Println(" |  \\  (_)  |  (_| || |  | | (_) |   /\\   /   | |___| |___ | |   |")
	fmt.Println(" |   \\ ___ / \\___|_||_|  |_|\\___/ \\_/  \\_/     \\____|_____|___|  |")
	fmt.Println(" |                                                               |")
	fmt.Println(" +---------------------------------------------------------------+")
    	simplePythonCall("workflow/parslConfig.py")

    	commandsArray := readLines("workflow/userScript.py")
    	fmt.Println(commandsArray)

    	//start module execution from here onwards
	inChannelModule1 := make(chan string, 1)
	outChannelModule1 := make(chan string, 1)
	go pythonCall("workflow/"+commandsArray[0], inChannelModule1,"1")
	go messagePassing(inChannelModule1, outChannelModule1)
	fmt.Println(<-outChannelModule1)

	outChannelModule2 := make(chan string, 1)
	go pythonCall("workflow/"+commandsArray[1], outChannelModule2, "1")
	go messagePassing(outChannelModule1, outChannelModule2)
	fmt.Println(<- outChannelModule2)

	outChannelModule3 := make(chan string, 1)
	go pythonCall("workflow/"+commandsArray[2], outChannelModule3, "1")
	go messagePassing(outChannelModule2, outChannelModule3)
	fmt.Println(<- outChannelModule3)

	outChannelModule4 := make(chan string, 1)
	go pythonCall("workflow/"+commandsArray[3], outChannelModule4, "1")
	go messagePassing(outChannelModule3, outChannelModule4)
	fmt.Println(<- outChannelModule4)
}
