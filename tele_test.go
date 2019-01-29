package tele_test

import (
	"fmt"
	"testing"

	"tele"
)

func TestCreate(t *testing.T) {
	bot, err := tele.Create("123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11", "https://api.telegram.org/bot")
	if err != nil {
		t.Errorf(err.Error())
	}

	if bot.Id != "m02828" {
		t.Errorf("Failed to create bot instance, got: %s", bot.Id)
	}
}

func ExampleCreate() {
	bot, err := tele.Create("1234567:FFF-DEF1234ghIkl-zyx57W2v1u123ew11", "https://api.telegram.org/bot")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(bot.Id)
	// Output:
	// m1f265
}

func ExampleGetBots() {
	_, _ = tele.Create("123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11", "https://api.telegram.org/bot")
	_, _ = tele.Create("789101:DEF-GHI1234ghIkl-zyx57W2v1u123ew11", "https://api.telegram.org/bot")
	bots := tele.GetBots()

	for _, v := range *bots {
		fmt.Println(v.Au)
	}

	// Output:
	// 123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11
	// 1234567:FFF-DEF1234ghIkl-zyx57W2v1u123ew11
	// 789101:DEF-GHI1234ghIkl-zyx57W2v1u123ew11
}
