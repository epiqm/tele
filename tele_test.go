package tele_test

import (
	"fmt"
	"testing"

	"tele"
)

func TestCreate(t *testing.T) {
	bot, err := tele.Create("123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11", "https://api.telegram.org/bot", 0)
	if err != nil {
		fmt.Println(err.Error())
	}

	if bot.Id != "02828fb18e9bb275590a4108a0bc61ec" {
		t.Errorf("Failed to create bot instance, got: %s", bot.Id)
	}
}

func ExampleCreate() {
	bot, err := tele.Create("1234567:FFF-DEF1234ghIkl-zyx57W2v1u123ew11", "https://api.telegram.org/bot", 0)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(bot.Id)
	// Output:
	// 1f265e35ec091d9f1258ca30ca80bd00
}

func ExampleGetBots() {
	_, _ = tele.Create("123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11", "https://api.telegram.org/bot", 0)
	_, _ = tele.Create("789101:DEF-GHI1234ghIkl-zyx57W2v1u123ew11", "https://api.telegram.org/bot", 0)
	bots := tele.GetBots()

	for _, v := range *bots {
		fmt.Println(v.Au)
	}

	// Output:
	// 123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11
	// 1234567:FFF-DEF1234ghIkl-zyx57W2v1u123ew11
	// 789101:DEF-GHI1234ghIkl-zyx57W2v1u123ew11
}
