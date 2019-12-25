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
	"reflect"
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

//JOHN
	// Open the file
	csvfile1, err1 := os.Open("/home/mpiuser/Desktop/multiplyByTwo.csv")
	if err1 != nil {
		log.Fatalln("Couldn't open the csv file", err1)
	}

	// Parse the file
	r1 := csv.NewReader(csvfile1)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records

	var studentMarks [10]string
	var count = 0
	for {
		// Read each record from csv
		record1, err1 := r1.Read()
		if err1 == io.EOF {
			break
		}
		if err1 != nil {
			log.Fatal(err1)
		}
		//fmt.Printf(record[0])
		
		studentMarks[count]=record1[0]
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
    	fmt.Println(studentMarksTotal)
	fmt.Println(reflect.TypeOf(studentMarksTotal))

//CLASS
	// Open the file
	csvfile2, err2 := os.Open("/home/mpiuser/Desktop/calculateAverage.csv")
	if err2 != nil {
		log.Fatalln("Couldn't open the csv file", err2)
	}

	// Parse the file
	r2 := csv.NewReader(csvfile2)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records

	var classAverage string
	for {
		// Read each record from csv
		record2, err2 := r2.Read()
		if err2 == io.EOF {
			break
		}
		if err2 != nil {
			log.Fatal(err2)
		}
		//fmt.Printf(record[0])
		
		classAverage=record2[0]
		//fmt.Println(count)
		//count = count + 1
	}

	fmt.Println(classAverage)
	fmt.Println(reflect.TypeOf(classAverage))

	//var s float64

	classAverageInt, _ := strconv.ParseFloat(classAverage, 8)
	fmt.Println(classAverageInt)
	fmt.Println(reflect.TypeOf(classAverageInt))

	outChannelModule4 := make(chan string, 1)

	if studentMarksTotal >= int(classAverageInt) {
		go pythonCall("workflow/"+commandsArray[3], outChannelModule4, "1")
		go messagePassing(outChannelModule3, outChannelModule4)
		fmt.Println(<- outChannelModule4)
    	} else {
		go pythonCall("workflow/"+commandsArray[4], outChannelModule4, "1")
		go messagePassing(outChannelModule3, outChannelModule4)
		fmt.Println(<- outChannelModule4)
   	}
	
/*
	if s, err := strconv.ParseFloat(classAverage, 32); err == nil {
    		fmt.Println(s)
		fmt.Println(reflect.TypeOf(s))
		
	}
	fmt.Println(s)


	var classAverageInt = []int{}

	for _, i := range classAverage {
        	j, err := strconv.Atoi(i)
        	if err != nil {
            		panic(err)
        	}
        	classAverageInt = append(classAverageInt, j)
    	}
    	//fmt.Println(studentMarksInt)

	var studentMarksTotal int
	
	for _, num := range studentMarksInt {
        	studentMarksTotal += num
    	}
    	//fmt.Println("sum:", studentMarksTotal)*/

	

	
}
