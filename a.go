package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func canDefeatAllEnemies(enemyList []int, specialLimit int) (int, bool) {
	var defeatNum int
	for _, enemyNum := range enemyList {
		if enemyNum > 5 {
			if specialLimit == 0 {
				return defeatNum, false
			}
			specialLimit--
		}
		defeatNum += enemyNum
	}
	return defeatNum, true
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
		enemyStrList := strings.Split(scanner.Text(), " ")
		enemyList := make([]int, 0, len(enemyStrList))
		for _, enemyNumStr := range enemyStrList {
			enemyNum, _ := strconv.Atoi(enemyNumStr)
			enemyList = append(enemyList, enemyNum)
		}
		enemyDefeatedNum, defeatAllEnemies := canDefeatAllEnemies(enemyList, specialLimit)

		// "true, 35"などのように出力される。例では"True, 35"だが分かれば良いと思うのでここはそのままにする
		fmt.Println(defeatAllEnemies, enemyDefeatedNum)
	}
}
