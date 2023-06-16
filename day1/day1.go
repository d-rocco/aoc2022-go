package main

import (
	"bufio"
	"fmt"
	"os"
    "strconv"
    "sort"
)

// used in part 1
// func maxCalories(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
       fmt.Println(err) 
    }
    
    defer file.Close()

    var intConvert, sum int
    topThree := []int{}
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        if scanner.Text() == "" {
            if len(topThree) < 3 {
                topThree = append(topThree, sum)
                sort.Ints(topThree) 
            }
            if sum > topThree[0] {
                topThree[0] = sum
                sort.Ints(topThree) 
            }
            sum = 0
        } else {
            intConvert, err = strconv.Atoi(scanner.Text())
            sum += intConvert
        }
    }

    if sum > topThree[0] {
        topThree[0] = sum
        sort.Ints(topThree) 
    }

    var total int
    for i := 0; i < len(topThree); i++ {
        total += topThree[i] 
    }
    fmt.Println(total) 
}
