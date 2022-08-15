package cmd

import (
	bot "gopkg.in/telebot.v3"
	"simigo/simisimi"
)

type Config struct {
	Bot     *bot.Bot
	OwnerID int
}

func New(conf Config) *Config {
	return &conf
}

func (a *Config) HelpCmd(c bot.Context) error {
	_, err := c.Bot().Reply(c.Message(), "Mulai kirim pesan sekarang")
	if err != nil {
		return err
	}

	return nil
}

func (a *Config) StartCmd(c bot.Context) error {
	_, err := c.Bot().Reply(c.Message(), "Hai Kak")
	if err != nil {
		return err
	}

	return nil
}

func (a *Config) OnTextHandler(c bot.Context) error {
	var text = c.Text()

	//fake typing
	err := c.Notify(bot.Typing)
	if err != nil {
		return err
	}

	res := simisimi.GetSimi(text)

	_, err = c.Bot().Reply(c.Message(), res)
	if err != nil {
		return err
	}
	return nil
}
