package main

import (
	"sync"
	"time"


)

import	"github.com/qyuhen/test"		//默认方式：test.A
import X "github.com/qyuhen/test"		//别名方式：X.A
import . "github.com/qyuhen/test"		//简便方式：A
import _ "github.com/qyuhen/test"		//初始化方式：无法引用，仅用来初始化目标包





func main(){
	// f:=hello
	// exce(f)

	// var  a *int=test()
	// println(a,*a)

	var p *int
	testPoint(&p)
	println(*p,&p)

	s:=`line\r\n
	line 2`
	println(s)
	//输出 line\r\n
	//		line2


	s:="abc"
	println(s[1]) //b
	println(&s[1]) //错误：cannot take the address of s[1]

	s:="金龙"
	for i:=0;i<lne(s);i++{				//byte
		fmt.Printf("%d:[%c]\n",i,s[i])
	}

	for i,c:=range c {
		fmt.Printf("%d:[%c]\n",i,c)  //0:[金] 3：[龙]
	}

	str:="雨.痕"
	println(len(s),utf8.RuneCountInString(s)) //7 3

	a:=[...]int{1,2}
	p:=&a

	p[1]+=10
	println(p[1])


	m:=map[string]int{
		"a":1,
		"b":2,
	}
	m["a"]=10  //修改
	m["c"]=30  //新增
	if v,ok:=m["d"];ok{  //使用 ok-idiom判断key是否存在，返回值
		println(v)
	}
	delete(m,"d") //删除键值对，不存在时，不会出错

	var lock sync.RWMutex
	m:=make(map[string]int)

	go func(){
		for{
			lock.Lock()
			m["a"]+=1
			lock.UnLock()
			time.Sleep(time.Microsecond)
		}
	}()


	exit:=make(chan struct{})

	go func(){
		println("hello world")
		exit<-struct{}{}
	}()
	<-exit
	println("end")

	u:=user{"Tom",1}
	v:=reflect.ValueOf(u)
	t:=v.Type()

	for i,n:=0,t.NumField();i<n;i++{
		fmt.Printf("%s:%v\n",t.Field(i).Tag,v.Field(i))
		//输出 昵称：Tom
		//		性别 1
	}

	v1:=struct{
		a byte
		b byte
		c int32   //对齐宽度 4
	}

	d:=data{}
	d.Lock()   //编译器会处理为 sync.(*Mutex).Lock()调用
	defer d.UnLock()

	var n N=25
	fmt.Printf("test.n:%p,%d\n",&n,n)

	f1:=N.test
	f1(n)		//func(n N)

	f2:=(*N).test //func(n *N)
	f2(&n)		//按方法集中的签名传递正确类型的参数

	d:=data{100}
	var t interface{}=&d
	t.(*data).x=200
	println(t.(*data).x)

	var x interface{}=func(x int) string{
		return fmt.Sprintf("d:%d",x)
	}

	switch v:=x.(type){
	case nil:
		println("nil")
	case *int:
		println(*v)
	case func(int) string:
		println(v(100))
	case fmt.Stringer:
		fmt.Println(v)
	default:
		println("unknown")
	}


	exit:=make(chan struct{})		//创建通道。因为仅是通知，数据并没有实际意义

	go func(){
		time.Sleep(time.Second)
		println("goroutine done.")

		close(exit)					//关闭通道，发出信号
	}()

	println("main ....")
	<-exit							//如果通道关闭，立即解除阻塞
	println("main exit")
	//输出结果： main....	/n goroutine done.	/n main exit


	var wg sync.WaitGroup

	for i:=0;i<10;i++ {
		wg.Add(1)						//累加计数

		go func(id int){
			defer wg.Done()				//递减计数

			time.Sleep(time.Second)
			println("groutime",id,"done.")
		}(i)
	}

	println("main...")
	wg.Wait()							//阻塞，直到计数归零
	println("main exit.")


	var wg Sync.WaitGroup

	go func(){
		wg.Add(1)		//来不及设置
		println("hi")
	}()

	wg.Wait()
	println("exit.")

	done:=make(chan struct{})	//结束事件
	c:=make(chan string)		//数据传输通道

	go func (){
		defer close(done)		//关闭通道

		s:=<-c					//接收数据
		println(s)
	}
	c<-"hi"						//发送消息
	<-done						//阻塞，直到有数据或管道关闭


	a,b:=make(chan int),make(chan int,3)

	b<-1
	b<-2
	println("a:",len(a),cap(a))
	println("b:",len(b),cap(b))
	//输出： a:0 0  /n  b:2 3

	done:=make(chan struct{})
	c:=make(chan int)

	go func(){
		defer close(done)  //确保发出结束通知

		for{
			x,ok:=<-c	//据此判断通道是否被关闭
			if !ok {
				return
			}
			println(x)
		}

		for x:=range c {	//循环获取消息，直到通道被关闭
			println(x)
		}
	}()

	c<-1
	c<-2
	close(c)
	<-done


}

type data struct{
	x int
}

type N int
func (n N) test(){
	fmt.Printf("test.n:%p,%d\n",&n,n)
}

type iface struct{
	tab *itab			//类型信息
	data unsafe.Pointer	//实际对象指针
}

type itab struct{
	inter	*interfacetype	//接口类型
	_type	*_type			//实际对象类型
	fun		[1]uintptr		//实际对象方法地址
}


type data struct{
	sync.Mutex
	buf [1024]byte
}

type user struct{
	name string	`昵称`
	sex byte	`性别`
}

type N int

func (N) test(){
	println("hello world")
}

type att struct{
	perm int
}

type file struct{
	name string
	att				//仅有类型名
}

type node struct{
	_ int
	id int
	next *node
}

type file struct{
	name string
	arr struct{
		owner int 
		perm int
	}
}

type slice struct{
	array unsafe.Pointer
	len int
	cap int
}

type stringStruct struct{
	str unsafe.Pointer
	len int
}

func testPoint(p **int){
	x:=100
	*p=&x
}

func testMethod(a ...int){
	Println(a)
}

func test() *int{
	a:=0x100
	return &a
}

func hello(){
	println("hello world")
}

func exce(f func()){
	f()
}