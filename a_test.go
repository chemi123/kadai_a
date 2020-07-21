package main

import (
	"testing"
)

func Test_canDefeatAllEnemies(t *testing.T) {
	testCases := []struct {
		name                 string
		enemyList            []int
		specialLimit         int
		wantEnemyDefeatedNum int
		wantDefeatAllEnemies bool
	}{}
	for _, tc := range testCases {
		enemyDefeatedNum, defeatAllEnemies := canDefeatAllEnemies(tc.enemyList, tc.specialLimit)
		if enemyDefeatedNum != tc.wantEnemyDefeatedNum || defeatAllEnemies != tc.wantDefeatAllEnemies {
			t.Fatalf("")
		}
	}
}
