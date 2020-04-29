package main

import (
    "fmt"
    "os"
	"bufio"
	"strconv"
	_"strings"
)

func main(){
    // scanner := bufio.NewScanner(os.Stdin)
    // scanner.Scan()
    // s := scanner.Text()
	// fmt.Println(lastWordLength(s))
	
	// inputReader := bufio.NewReader(os.Stdin)
	// str, err := inputReader.ReadString('\n')

    // str = strings.ToLower(str)
    // count := 0
    // if err != nil {
    //     print(0)
    // } else {
    //     target, _, err := inputReader.ReadRune()
    //     if err == nil {
    //         for _, c := range  str {
    //             if c == target {
    //                 count++
    //             }
    //         }
    //     }
    // }
	// fmt.Println(count)
	
	// count:=4
	// result:=count%2
	// fmt.Println(result)

	// scanner := bufio.NewScanner(os.Stdin)
    // scanner.Scan()
    // s := scanner.Text()
	// // fmt.Println("str",s)
	// intValue,err:=strconv.Atoi(s)
	// fmt.Println("int",intValue)
	// if err==nil{
	// 	fmt.Println(getResult(intValue))
	// }

	//num:=64577
	// str:="9876673"
	// filterRepeatNum(str)
	// value,err:=filterRepeatNum(str)
	// if err==nil{
	// 	fmt.Println(value)
	// }

	// str:="uqic^g`(s&jnl(m#vt!onwdj(ru+os&wx"
	// fmt.Println(printCharNum(str))

	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
	s := scanner.Text()
	value1,value2:=string(s[0]),string(s[2])
	s1 := strconv.Itoa(value1)
	s2:= strconv.Itoa(value2)
	fmt.Println(minComMultiple(s,s2))
	
}

func lastWordLength(str string) int{
    if len(str)==0||len(str)>5000{
        return -1
    }
    
    length:=len(str)
    lastlength:=0
    for i:=length-1;i>=0;i-- {
        if string(str[i])==" "{
            return lastlength
        }else{
            lastlength++
        }
    }
    return lastlength

}

func printTargetCount(str string,target string) int {
	return -1
}

func getResult(input int) []int {
	// result:=make([]int,0)
	var result []int
    temp:=input
    i:=2
    for i<=temp {
        if temp%i==0{
            result=append(result,i)
            temp=temp/i
        }else{
            i++
        }
	}
	return result  
}

func filterRepeatNum(str string) string {
	// intValue,_:=strconv.Atoi(str)
	var slice []string
	length:=len(str)
	for i:=length-1;i>=0;i--{
		if slice==nil{
			slice=append(slice,string(str[i]))
		}else{
			isExit:=false
			for _,value:=range slice{
				if value==string(str[i]){
					isExit=true
					break
				}
			}
			if isExit {
				continue
			}else{
				slice=append(slice,string(str[i]))
			}
		}
	}
	var res string
	for _,value:=range slice{
		res+=value
	}
	return res
	// return  strconv.Atoi(res)
}

func printCharNum(str string) int{
	if len(str)<=0{
		return 0
	}

	dir:=make(map[string]int)
	for i:=0;i<len(str);i++{
		if _,ok:=dir[string(str[i])];!ok{
			dir[string(str[i])]=i
		}else{
			continue
		}
	}
	for k,v:=range dir{
		fmt.Println(k,v)
	}
	return len(dir)
}

func minComMultiple(value1 int,value2 int) int{
	temp:=value1
  for i:=2;i<=value1 ;i++{
	  if temp%i==0&&value2%i==0{
		temp=temp/i
	  }
  }

  return temp*value2

}