package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
    sort.Slice(arr, func(i,j int) bool {
        return arr[i] < arr[j]
    })
    var sumMax, sumMin int64
    
    for i :=0 ; i < 4 ; i++ {
        sumMin += int64(arr[i])
    }
    for i :=1 ; i < 5 ; i++ {
        sumMax += int64(arr[i])
    }
    
    fmt.Println(sumMin, sumMax)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
