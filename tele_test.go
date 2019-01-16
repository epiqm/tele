package tele_test

import (
	"fmt"
	"testing"

	"tele"
)

func TestCreate(t *testing.T) {
	bot := tele.Create("123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11", "https://api.telegram.org/bot")

	if bot.Id != "m02828" {
		t.Errorf("Failed to create bot instance, got: %s", bot.Id)
	}
}

func ExampleCreate() {
	bot := tele.Create("1234567:FFF-DEF1234ghIkl-zyx57W2v1u123ew11", "https://api.telegram.org/bot")

	fmt.Println(bot.Id)
	// Output:
	// m1f265
}

func ExampleGetBots() {
	_ = tele.Create("123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11", "https://api.telegram.org/bot")
	_ = tele.Create("789101:DEF-GHI1234ghIkl-zyx57W2v1u123ew11", "https://api.telegram.org/bot")
	bots := tele.GetBots()

	for _, v := range *bots {
		fmt.Println(v.Au)
	}

	// Output:
	// 123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11
	// 1234567:FFF-DEF1234ghIkl-zyx57W2v1u123ew11
	// 789101:DEF-GHI1234ghIkl-zyx57W2v1u123ew11
}
