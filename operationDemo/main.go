package main

func main(){
	// x:=10
	// var p *int=&x

	// *p+=20
	// println(x,&x,p,*p)
	//30 0xc00003ff80 0xc00003ff80 30


	data:=[]string{"a","b","c"}

	for i,s:=range data {
		println(&i,&s)
		println(i,s)
	}
}