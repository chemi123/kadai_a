package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// 簡単のために、絶対に正しいデータが来るとしてエラーバリデーションはかけない
	dataSetNum, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	fmt.Println(dataSetNum)

	for i := 0; i < dataSetNum; i++ {
		scanner.Scan()
		specialLimit, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		fmt.Println(specialLimit)

		scanner.Scan()
		enemyList := strings.Split(scanner.Text(), " ")
		fmt.Println(enemyList)
	}
}
