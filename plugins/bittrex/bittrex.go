package bittrex

import (
	"fmt"
	"os"

	"github.com/go-chat-bot/bot"
	"github.com/toorop/go-bittrex"
)

func price(command *bot.Cmd) (string, error) {
	bx := bittrex.New(
		os.Getenv("BITTREX_API_KEY"),
		os.Getenv("BITTREX_SECRET_KEY"),
	)
	ticker, err := bx.GetTicker("BTC-XVG")
	return fmt.Sprintf("Bittrex: %0.8f XVG/BTC", ticker.Last), err
}

func init() {
	bot.RegisterCommand(
		"price",
		"Check market price for BTC-XVG",
		"",
		price,
	)
}
