package main

import (
	"testing"
)

func sum(nums []int) int {
	var s int
	for _, num := range nums {
		s += num
	}
	return s
}

func Test_canDefeatAllEnemies(t *testing.T) {
	testCases := []struct {
		name                 string
		enemyList            []int
		specialLimit         int
		wantEnemyDefeatedNum int
		wantDefeatAllEnemies bool
	}{
		{
			name:                 "Defeat all enemies",
			enemyList:            []int{1, 2, 3, 4, 5, 6},
			specialLimit:         3,
			wantEnemyDefeatedNum: sum([]int{1, 2, 3, 4, 5, 6}),
			wantDefeatAllEnemies: true,
		},
		{
			name:                 "Defeat all enemies without special",
			enemyList:            []int{1, 2, 3, 4, 5},
			specialLimit:         0,
			wantEnemyDefeatedNum: sum([]int{1, 2, 3, 4, 5}),
			wantDefeatAllEnemies: true,
		},
		{
			name:                 "lose...",
			enemyList:            []int{1, 2, 3, 4, 5, 6},
			specialLimit:         0,
			wantEnemyDefeatedNum: sum([]int{1, 2, 3, 4, 5}),
			wantDefeatAllEnemies: false,
		},
	}
	for _, tc := range testCases {
		enemyDefeatedNum, defeatAllEnemies := canDefeatAllEnemies(tc.enemyList, tc.specialLimit)
		if enemyDefeatedNum != tc.wantEnemyDefeatedNum || defeatAllEnemies != tc.wantDefeatAllEnemies {
			t.Fatalf("")
		}
	}
}
