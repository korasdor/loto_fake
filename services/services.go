package services

import (
	"encoding/json"
	"github.com/korasdor/data"
)

func GetBalanceMessage() string {

	msg := data.BalanceMsgStruct{
		Сurrency:      data.CurrentCurrency,
		Balance:       data.TotalBalance,
		Bonus_balance: data.TotalBonus,
	}

	strb, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return string(strb)
}

func GetRulesMessage() string {
	msg := data.RulesMsgStruct{
		Single:   36,
		Split:    18,
		Street:   12,
		Square:   9,
		Color:    3.6,
		Even_odd: 1.8,
	}

	strb, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return string(strb)
}

func GetWinsMessage(gameId string) string {
	msg := data.WinMsgStruct{
		Total_win: 2500,
	}

	strb, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return string(strb)
}
