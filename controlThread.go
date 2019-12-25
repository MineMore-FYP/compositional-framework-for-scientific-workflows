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

	//CALCULATE STUDENT TOTAL MARKS
	csvfile1, err1 := os.Open("/home/mpiuser/Desktop/multiplyByTwo.csv")
	if err1 != nil {
		log.Fatalln("Couldn't open the csv file", err1)
	}

	r1 := csv.NewReader(csvfile1)
	var studentMarks [10]string
	var count = 0
	//read student marks csv file (record wise) and put in studentMarks array
 	for {
		// Read each record from csv
		record1, err1 := r1.Read()
		if err1 == io.EOF {
			break
		}
		if err1 != nil {
			log.Fatal(err1)
		}
		studentMarks[count]=record1[0]
		count = count + 1
	}

	var studentMarksInt = []int{}

	//convert student marks array to int
	for _, i := range studentMarks {
        	j, err := strconv.Atoi(i)
        	if err != nil {
            		panic(err)
        	}
        	studentMarksInt = append(studentMarksInt, j)
    	}

	var studentMarksTotal int
	
	//calculate student's total marks
	for _, num := range studentMarksInt {
        	studentMarksTotal += num
    	}

	//RETREIVE CLASS AVERAGE 
	csvfile2, err2 := os.Open("/home/mpiuser/Desktop/calculateAverage.csv")
	if err2 != nil {
		log.Fatalln("Couldn't open the csv file", err2)
	}

	r2 := csv.NewReader(csvfile2)
	var classAverage string
	//read class average string value from csv file
	for {
		record2, err2 := r2.Read()
		if err2 == io.EOF {
			break
		}
		if err2 != nil {
			log.Fatal(err2)
		}
		classAverage=record2[0]
	}

	//convert class average to float64
	classAverageInt, _ := strconv.ParseFloat(classAverage, 8)

	//Dynamic decision making
	outChannelModule4 := make(chan string, 1)

	if studentMarksTotal >= int(classAverageInt) {
		//send to 'amazing' module
		go pythonCall("workflow/"+commandsArray[3], outChannelModule4, "1")
		go messagePassing(outChannelModule3, outChannelModule4)
		fmt.Println(<- outChannelModule4)
    	} else {
		//send to 'terrible' module
		go pythonCall("workflow/"+commandsArray[4], outChannelModule4, "1")
		go messagePassing(outChannelModule3, outChannelModule4)
		fmt.Println(<- outChannelModule4)
   	}	
}
