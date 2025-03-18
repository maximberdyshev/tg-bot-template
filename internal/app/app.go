package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"tg-bot-template/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type App struct {
	Bot *tgbotapi.BotAPI
}

func New(cfg *config.Config) (*App, error) {
	bot, err := tgbotapi.NewBotAPIWithAPIEndpoint(cfg.Bot.Token, cfg.Telegram.APIEndpoint)
	if err != nil {
		return nil, err
	}

	// debug mode only
	bot.Debug = true

	return &App{
		Bot: bot,
	}, nil
}

func (app *App) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	updateCfg := tgbotapi.UpdateConfig{
		Offset:         0,
		Limit:          0,
		Timeout:        60,
		AllowedUpdates: []string{},
	}

	fetCh := app.Fetcher(updateCfg)
	app.Processor(fetCh)

	log.Println("App launching..")

	<-ctx.Done()

	log.Println("Recieved: signal interrupt!")
	log.Println("App is stopped!")
}

func (app *App) Fetcher(cfg tgbotapi.UpdateConfig) chan tgbotapi.Update {
	ch := make(chan tgbotapi.Update, app.Bot.Buffer)

	log.Println("Fetcher is running..")

	go func() {
		defer close(ch)

		for {
			updates, err := app.Bot.GetUpdates(cfg)
			if err != nil {
				log.Println(err)
				log.Println("Failed to get updates, retrying in 3 seconds...")
				time.Sleep(time.Second * 3)

				continue
			}

			for _, update := range updates {
				if update.UpdateID >= cfg.Offset {
					cfg.Offset = update.UpdateID + 1
					ch <- update
				}
			}
		}
	}()

	return ch
}

func (app *App) Processor(ch <-chan tgbotapi.Update) {
	log.Println("Processor is running..")

	go func() {
		for u := range ch {
			switch {
			case u.Message != nil:
				app.message(u.Message)
			case u.CallbackQuery != nil:
				app.callback(u.CallbackQuery)
			default:
				log.Println("Unknown type update -_-")
			}
		}
	}()
}
