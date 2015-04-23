package data

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Jp              []float32
	Coefs           [][]float32    = make([][]float32, 10)
	Balls           []int          = make([]int, 0)
	GameId          int            = 0
	Bets            []BetMsgStruct = make([]BetMsgStruct, 0)
	TotalBalance    float32        = 100000.00
	TotalBonus      float32        = 1000.00
	GameCounter     int            = 0
	IsFinished      bool           = true
	OnBet           bool           = false
	RandBalls       []int          = make([]int, 0)
	BetTime         time.Duration  = 30000
	BallTime        time.Duration  = 1000
	TotalWin        float32        = 0
	BetTimout       int            = int(BetTime) / 1000
	Token           string         = "b9d7683e179210d6bff6f3f4c4217a1b"
	GameName        string         = "aztec"
	CurrentCurrency string         = "rub"
)

type BalanceMsgStruct struct {
	Сurrency      string  `json:"currency"`
	Balance       float32 `json:"balance"`
	Bonus_balance float32 `json:"bonus_balance"`
}

type RulesMsgStruct struct {
	Single   float32 `json:"single"`
	Split    float32 `json:"split"`
	Street   float32 `json:"street"`
	Square   float32 `json:"square"`
	Color    float32 `json:"color"`
	Even_odd float32 `json:"even_odd"`
}

type BetsMsgStruct struct {
	Bets  []BetMsgStruct `json:"bets"`
	Token string         `json:"token"`
}

type BetMsgStruct struct {
	Bet_sum float32 `json:"bet_sum"`
	Content []int   `json:"content"`
	Bonus   bool    `json:"bonus"`
	Coeff   float32 `json:"coeff"`
}

type RandMsgStruct struct {
	GameId    int       `json:"game_id"`
	Balls     []int     `json:"balls"`
	OnBet     bool      `json:"on_bet"`
	BetTimout int       `json:"bet_timeout"`
	Finished  bool      `json:"finished"`
	Jp        []float32 `json:"jp"`
}

type WinMsgStruct struct {
	Total_win float32 `json:"total_win"`
}

type MatchStruct struct {
	Id         int
	MatchCount int
	MaxNums    int
	BetValue   float32
}

func CreateCoefs() [][]float32 {

	data := [55]float32{3.6, 1, 9, 0, 2, 45, 0, 1, 10, 80, 0, 1, 3, 20, 200, 0, 0, 2, 15, 50, 500, 0, 0, 2, 5, 25, 80, 1000, 0, 0, 0, 5, 15, 70, 250, 2000, 0, 0, 0, 2, 10, 35, 150, 1500, 5000, 0, 0, 0, 0, 5, 35, 100, 500, 2000, 10000}

	n := 0
	coefs := make([][]float32, 10)
	for i := 0; i < len(coefs); i++ {
		innerLen := i + 1
		coefs[i] = make([]float32, innerLen)
		for j := 0; j < innerLen; j++ {
			coefs[i][j] = data[n]
			n++
		}
	}

	return coefs

}

func CreateBalls(arr []int, length int) []int {
	for i := 0; i < length; i++ {
		arr = append(arr, i+1)
	}

	return arr
}

func Shuffle(arr []int) []int {
	var i = len(arr)
	r := rand.New(rand.NewSource(99))

	for i > 1 {
		var j int = int(r.Float32() * float32(i))
		tmp := arr[j]
		arr[j] = arr[i-1]
		arr[i-1] = tmp

		i--
	}

	return arr
}

func CalcWin() {
	matches := make([]MatchStruct, 0)
	for i := 0; i < len(Bets); i++ {
		content := Bets[i].Content

		matchStr := MatchStruct{
			Id:         i,
			MatchCount: 0,
			MaxNums:    len(Bets[i].Content),
			BetValue:   Bets[i].Bet_sum,
		}

		for j := 0; j < len(content); j++ {
			for k := 0; k < len(Balls); k++ {
				if content[j] == Balls[k] {
					matchStr.MatchCount += 1
				}
			}
		}

		if matchStr.MatchCount > 0 {
			matches = append(matches, matchStr)
		}
	}

	TotalWin = 0
	for i := 0; i < len(matches); i++ {
		matchStr := matches[i]

		row := Coefs[matchStr.MaxNums-1]
		curCoef := row[matchStr.MatchCount-1]

		fmt.Printf("row : %s, curCoef: %f\n, matchCount: %d\n", row, curCoef, matchStr.MatchCount)

		TotalWin += (curCoef * matchStr.BetValue)
	}
}
