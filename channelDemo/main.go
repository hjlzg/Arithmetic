package main

func main(){

	data:=make(chan int)
	done:=make(chan bool)

	go consumer(data,done)
	go producer(data)

	
	println("结束:",<-done)
}

func consumer(data chan int,done chan bool){
	// for x:=range data{
	// 	println("reve",x)
	// }
	// done<-true

	//只会输出一个数
	// var i int
	// select {
	// 	case i=<-data:{
	// 		println("reve",i)
	// 	}

	// }
	// done<-true

	//匿名函数实现,不能打印
	go func(){
		for{
			i:=<-data
			println("reve",i)
		}
	}()
	done <-true
}

func producer(data chan int){
	for i:=0;i<4;i++{
		data<-i
	}

	close(data)
}