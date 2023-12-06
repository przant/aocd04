package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
    "strings"
)

type Card struct {
    WinNums []int
    MyNums  []int
}

func main() {
    pF, err := os.Open("example.txt")
    if err != nil {
        log.Fatalf("while opening the file %q: %s", pF.Name(), err)
    }
    defer pF.Close()

    sum := 0.0
    scnr := bufio.NewScanner(pF)

    for scnr.Scan() {
        line := scnr.Text()
        myNumsSet := make(map[int]bool)
        sections := strings.Split(line, "|")
        myNumsList := strings.Fields(sections[1])
        winNumsList := strings.Fields(strings.Split(sections[0], ":")[1])

        for _, strNum := range myNumsList {
            num, _ := strconv.Atoi(strings.TrimSpace(strNum))
            myNumsSet[num] = true
        }

        matchNumsCount := 0

        for _, strNum := range winNumsList {
            num, _ := strconv.Atoi(strings.TrimSpace(strNum))
            if myNumsSet[num] {
                matchNumsCount++
            }
        }

        if matchNumsCount > 0 {
            sum += math.Pow(2.0, float64(matchNumsCount-1))
        }
    }

    fmt.Println(int64(sum))

}
