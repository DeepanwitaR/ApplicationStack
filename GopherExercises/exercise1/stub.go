package main

import (
	"exercise" //must include it in src, and import relative to it post go install exercise, in C: dir
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

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
	// fmt.Print(data)
	// fmt.Printf("data type %T\n", data[0])
	num := len(data)
	// ans := make([]string, 100)
	i := 0
	flag := 0
	total := 0
	correct := 0
	userString := ""
	compare := ""
	for ; i < num-1; i = i + 1 {
		if string(data[i]) == "\n" {
			// fmt.Println("newline")
			continue
		} else {
			// fmt.Println("not newline")
			if flag == 0 {
				if string(data[i]) == "," {
					fmt.Println()
					flag = 1
					continue
				} else {
					fmt.Print(string(data[i]))
				}
			} else {
				fmt.Println("answer ?")

				// var wait sync.WaitGroup
				// wait.Add(1)
				// c := make(chan string)
				// go timerFunc(c)
				// timerFunc(c)
				userString = ""
				fmt.Scanln(&userString)
				// if userString != "" {
				// 	wait.Done()
				// 	// wait.Wait() ; not needed here.
				// } else {
				// msg := <-c
				// if msg == "time is up" {
				// 	fmt.Println(msg + "! Next question.")
				// 	total = total + 1
				// 	flag = 0
				// 	compare = ""
				// 	userString = ""
				// 	wait.Done()
				// 	continue
				// }
				compare = ""
				for string(data[i]) != "\n" && i < num {
					compare = compare + string(data[i])
					i = i + 1
				}
				compare = compare[:2]

				// check out formatting only 2 digits are correct
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
func timerFunc(c chan string) {
	time.Sleep(time.Second * 5)
	c <- "time is up"
}
