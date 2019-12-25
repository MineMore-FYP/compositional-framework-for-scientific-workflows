package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"log"
	"io"
	"strconv"
	"encoding/csv"
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

func check(e error) {
    if e != nil {
        panic(e)
    }
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
    	simplePythonCall("logo.py")
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

	// Open the file
	csvfile, err := os.Open("/home/mpiuser/Desktop/multiplyByTwo.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records

	var studentMarks [10]string
	var count = 0
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf(record[0])
		
		studentMarks[count]=record[0]
		//fmt.Println(count)
		count = count + 1
	}

	//fmt.Println(studentMarks)

	var studentMarksInt = []int{}

	for _, i := range studentMarks {
        	j, err := strconv.Atoi(i)
        	if err != nil {
            		panic(err)
        	}
        	studentMarksInt = append(studentMarksInt, j)
    	}
    	//fmt.Println(studentMarksInt)

	var studentMarksTotal int
	
	for _, num := range studentMarksInt {
        	studentMarksTotal += num
    	}
    	//fmt.Println("sum:", studentMarksTotal)

	outChannelModule4 := make(chan string, 1)
	go pythonCall("workflow/"+commandsArray[3], outChannelModule4, "1")
	go messagePassing(outChannelModule3, outChannelModule4)
	fmt.Println(<- outChannelModule4)

	outChannelModule5 := make(chan string, 1)
	go pythonCall("workflow/"+commandsArray[4], outChannelModule5, "1")
	go messagePassing(outChannelModule4, outChannelModule5)
	fmt.Println(<- outChannelModule5)
}
