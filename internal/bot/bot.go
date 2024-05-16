package bot

import (
	dronozor2 "dronozor/protos/gen/go/obb.dronozor.v1"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(botchan chan dronozor2.PhotoRequest) {
	bot, err := tgbotapi.NewBotAPI("6278832157:AAHY2z1G_zLyFregZWr0sI8UEIUxkyoD0jk")
	if err != nil {
		panic("failed to start bot" + err.Error())
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	for {
		data := <-botchan
		//TODO: get chat_id from config
		img := tgbotapi.NewPhoto(481370222, tgbotapi.FileBytes{Bytes: data.Image})
		_, err = bot.Send(img)
		if err != nil {
			fmt.Println(err.Error())
		}
		msg := tgbotapi.NewMessage(481370222, fmt.Sprintf("coords: %s\ntime: %s\n", data.GetCords(), data.GetImageTS().AsTime().String()))
		_, err = bot.Send(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
