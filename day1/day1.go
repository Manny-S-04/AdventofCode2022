package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
    part2Result(part1Result())
}

func part2Result(m map[int]int){
    
    values := make([]int,0, len(m))

    for _, v := range m{
        values = append(values, v)
    }

    sort.Slice(values, func(i, j int) bool {
        return values[i] > values[j]
    })
    
    result := values[0] + values[1] + values[2]

    fmt.Println(result)
}

func part1Result() map[int]int{
    file, err := os.Open("input.txt")
    if err != nil{
        fmt.Println("There was an error opening the file")
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    rawBytes, err := io.ReadAll(file)
    if err != nil{
        fmt.Println("Failed converting file to bytes")
    }

    lines := strings.Split(string(rawBytes), "\n")

    var summed map[int]int
    summed = make(map[int]int)

    var currentElf []int

    for i, line := range lines{
        if(len(line) == 0){
            summed[i] = sumSlice(currentElf) 
            currentElf = nil
            continue
        }else{
            lineAsInt, err := strconv.Atoi(line)
            if err != nil{
                fmt.Println("Input failed to convert string to int")
                return nil
            }
            currentElf = append(currentElf, lineAsInt)
        }
    }

    if err := scanner.Err(); err != nil{
        fmt.Println("There was an error reading the file")
    }

    result := result(summed)

    fmt.Printf("Final result: %v \n", result)

    return summed
}


func sumSlice(slice []int) int{
    sum := 0
    for _, num := range slice{
        sum += num
    }
    return sum
}

func result(m map[int]int) int{
    var largestValue int

    for _, value := range m{
        if value > largestValue{
            largestValue = value
        }
    }

    return largestValue
}
