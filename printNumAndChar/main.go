package main

import(
	"fmt"
	"sync"
	_"strings"
)

func  main()  {

	/*
	letter,number:=make(chan bool),make(chan bool)
	wait:=sync.WaitGroup{}
	go func(){
		i:=1
		for{
			select{
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter<-true
				break
			default:
				break
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup){
		str:="ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i:=0
		for{
			select{
			case<-letter:
				if i >= strings.Count(str,"")-1{
					wait.Done()
					return
				}
				fmt.Print(str[i:i+1])
				i++
				if i >= strings.Count(str,"") {
					i=0
				}
				fmt.Print(str[i:i+1])
				i++
				number<-true
				break
			default:
				break
			}
		}
	}(&wait)
	number<-true
	wait.Wait()

	*/

	flag:=make(chan int)
 
	var wg sync.WaitGroup
	wg.Add(2)
 
	go PrintNums(flag, &wg)
	go PrintChars(flag, &wg)
 
	wg.Wait()

}


func PrintNums(printChar chan int, wg *sync.WaitGroup)  {
	defer wg.Done()
	for i  := 0; i  < 1; i ++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d", 2*i+j+1)
		}
		printChar<-1
		fmt.Print("通道Num",<-printChar)
	}
}
 
func PrintChars(printChar chan int, wg *sync.WaitGroup)  {
	defer wg.Done()
	for i  := 0; i  < 1; i++ {
		fmt.Print("通道Char",<-printChar)
		for j := 0; j < 2; j++ {
			fmt.Printf("%c", 'A'+(2*i+j))
		}
		printChar<-1
	}
}