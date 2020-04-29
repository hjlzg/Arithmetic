// 本题为考试多行输入输出规范示例，无需提交，不计分。
package main

import (
    "fmt"
)
func main() {
    M:=0
    N:=0

    fmt.Scan(&M)
    fmt.Scan(&N)
    var array [][]int
    for i := 0; i < M; i++ {
        for j := 0; j < N; j++ {
            x:=0
            fmt.Scan(&x)
            array[i][j]=x
        }
    }

    for i:=0; i<M; i++ {
        for j:=0; j<N;j++ {
            fmt.Print(array[i][j])
            fmt.Print(" ")
        }
        fmt.Println()

    }
}