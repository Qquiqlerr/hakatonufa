package bot

import (
	dronozor2 "dronozor/protos/gen/go/obb.dronozor.v1"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(botchan chan dronozor2.PhotoRequest, admins []int) {
	bot, err := tgbotapi.NewBotAPI("6278832157:AAHY2z1G_zLyFregZWr0sI8UEIUxkyoD0jk")
	if err != nil {
		panic("failed to start bot" + err.Error())
	}
	fmt.Println("Telegram bot started")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	for {
		data := <-botchan
		for _, val := range admins {
			img := tgbotapi.NewPhoto(int64(val), tgbotapi.FileBytes{Bytes: data.Image})
			_, err = bot.Send(img)
			if err != nil {
				fmt.Println(err.Error())
			}
			msg := tgbotapi.NewMessage(int64(val), fmt.Sprintf("DRON DETECTED!!!\n"+
				"coords: %s\n"+
				"time: %s\n"+
				"Sender: %s", data.GetCords(), data.GetImageTS().AsTime().Local(), data.GetPhone()))
			_, err = bot.Send(msg)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

	}
}
