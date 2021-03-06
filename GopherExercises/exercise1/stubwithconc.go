package main

import (
	"exercise" //must include it in src, and import relative to it post go install exercise, in C: dir
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var userString string = ""

func askUser(c chan string) {
	//u have to send in a channel and it adds it to its func, and u can send and recieve data.
	fmt.Println("answer?")
	fmt.Scanln(&userString)
	c <- "recieved"
}
func ansTimer(c chan string) {
	fmt.Println("Your timer has begun.")
	time.Sleep(time.Second * 5)
	c <- "time up"
}
func main() {
	fmt.Println("inside main")
	exercise.Part1() // mock test.
	fmt.Println("executing the main program:")
	file, err := os.Open("sum.csv")
	if err != nil {
		panic("error in opening file")
	}
	b := make([]byte, 1000)
	count, err2 := file.Read(b)
	if err != nil {
		panic("error while taking in file count")
	}
	fmt.Printf("error: %v and bytes read: %v\n", err2, count)
	data, err3 := ioutil.ReadFile("sum.csv")
	if err3 != nil {
		panic("error while reading file.")
	}
	num := len(data)
	i := 0
	flag := 0
	total := 0
	correct := 0
	compare := ""
	for ; i < num-1; i = i + 1 {
		if string(data[i]) == "\n" {
			continue
		} else {
			if flag == 0 {
				if string(data[i]) == "," {
					fmt.Println()
					flag = 1
					continue
				} else {
					fmt.Print(string(data[i]))
				}
			} else {
				askUserChan := make(chan string)
				ansTimerChan := make(chan string)
				go askUser(askUserChan)
				go ansTimer(ansTimerChan)
				//insert select statement here.
				select {
				case msg := <-askUserChan:
					{
						fmt.Println(msg)
					}
				case msg := <-ansTimerChan:
					{
						fmt.Println(msg)
						for string(data[i]) != "\n" && i < num {
							// compare = compare + string(data[i])
							i = i + 1
						}
						total = total + 1
						flag = 0
						compare = ""
						userString = ""
						continue
					}
				}
				compare = ""
				for string(data[i]) != "\n" && i < num {
					compare = compare + string(data[i])
					i = i + 1
				}
				compare = compare[:2]
				fmt.Print(len(compare))
				fmt.Printf("compare is:%v \n", compare)
				if strings.Compare(userString, compare) == 0 {
					fmt.Println("correct answer!")
					correct = correct + 1
				} else {
					fmt.Println("wrong answer!")
				}
				total = total + 1
				flag = 0
				compare = ""
				userString = ""
			}
		}
	}
	fmt.Printf("you scored %v / %v\n", correct, total)
	defer fmt.Println("closing file")
	defer file.Close() //embedded struct extention thing.
}
