package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func canDefeatAllEnemies(enemyList []int, specialLimit int) (int, bool) {
	return 0, false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// 簡単のために、絶対に正しいデータが来るとしてエラーバリデーションはかけない
	dataSetNum, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	for i := 0; i < dataSetNum; i++ {
		scanner.Scan()
		specialLimit, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		scanner.Scan()
		enemyList := strings.Split(scanner.Text(), " ")
	}
}
