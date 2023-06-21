package main

import (
	"bufio"
	"fmt"
	"os"
)

var combos map[string]int

func main() {
    combos = make(map[string]int) 
    //part 1
    // combos["B X"] = 1
    // combos["C Y"] = 2
    // combos["A Z"] = 3
    // combos["A X"] = 4
    // combos["B Y"] = 5
    // combos["C Z"] = 6
    // combos["C X"] = 7
    // combos["A Y"] = 8
    // combos["B Z"] = 9

    combos["B X"] = 1
    combos["C X"] = 2
    combos["A X"] = 3
    combos["A Y"] = 4
    combos["B Y"] = 5
    combos["C Y"] = 6
    combos["C Z"] = 7
    combos["A Z"] = 8
    combos["B Z"] = 9

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err.Error())
    }

    defer file.Close()
    
    var sum int;
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        sum += combos[scanner.Text()] 
    }
    fmt.Println(sum)
}
