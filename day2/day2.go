package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)


func main(){
    fmt.Println(partOneResult())
    fmt.Println(partTwoResult())
}

func partTwoResult() int{
    rounds := readFile()

    var totalScore []int

    for i := 0; i < len(rounds)-1; i++ {
        currentRound := rounds[i]
        if currentRound != "0"{
            totalScore = append(totalScore, evaluatePartTwoRound(currentRound))
        }
    }

    return sumSlice(totalScore)
}

func evaluatePartTwoRound(line string) int{

    splitString := strings.Split(line, " ")

    oppMove := strings.TrimSpace(splitString[0])

    myMove := strings.TrimSpace(splitString[1])

    var currentScore int

    switch oppMove {
    case "A":
        switch myMove {
        case "X":
            currentScore += 3
            break
        case "Y":
            currentScore += 4
            break
        case "Z":
            currentScore += 8
            break
        }
    case "B":
        switch myMove {
        case "X":
            currentScore += 1
            break
        case "Y":
            currentScore += 5
            break
        case "Z":
            currentScore += 9
            break
        }
    case "C":
        switch myMove {
        case "X":
            currentScore += 2
            break
        case "Y":
            currentScore += 6
            break
        case "Z":
            currentScore += 7
            break
        }
    }

    return currentScore
}

func partOneResult() int{

    rounds := readFile()

    var totalScore []int

    for i := 0; i < len(rounds)-1; i++ {
        currentRound := rounds[i]
        if currentRound != "0"{
            totalScore = append(totalScore, evaluateRounds(currentRound))
        }
    }

    return sumSlice(totalScore)
}

func readFile() []string{

    file, err := os.Open("input.txt")
    if err != nil{
        fmt.Println("There was an error opening the file")
    }
    defer file.Close()
 
    rawBytes, err := io.ReadAll(file)
    if err != nil{
        fmt.Println("Failed converting file to bytes")
    }

    return strings.Split(string(rawBytes), "\n")
}

func evaluateRounds(line string) int{
    splitString := strings.Split(line, " ")
    
    oppMove := strings.TrimSpace(splitString[0])

    myMove := strings.TrimSpace(splitString[1])

    var currentScore int

    switch myMove {
    case "X":
        currentScore += 1
        currentScore += evaluateCurrentRound(myMove, oppMove)
    case "Y":
        currentScore += 2
        currentScore += evaluateCurrentRound(myMove, oppMove)
    case "Z":
        currentScore += 3
        currentScore += evaluateCurrentRound(myMove, oppMove)
    }
    return currentScore
}

func evaluateCurrentRound(myMove string,oppMove string) int{
    switch myMove {
    case "X":
        switch oppMove {
        case "A":
            return 3
        case "B":
            return 0
        case "C":
            return 6
        }
    break
    case "Y":
        switch oppMove {
        case "A":
            return 6
        case "B":
            return 3
        case "C":
            return 0
        }
    break
    case "Z":
        switch oppMove {
        case "A":
            return 0
        case "B":
            return 6
        case "C":
            return 3
        }
    break
    }
    return 0
}


func sumSlice(slice []int) int{
    sum := 0
    for _, num := range slice{
        sum += num
    }
    return sum
}
