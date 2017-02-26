package linebot

import (
	"fmt"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

type Request struct {
	Result []*Result `json:"result"`
}

type Response struct {
	To        []string               `json:"to"`
	ToChannel int                    `json:"toChannel"`
	EventType string                 `json:"eventType"`
	Content   map[string]interface{} `json:"content"`
}

type Result struct {
	ID          string                 `json:"id"`
	EventType   string                 `json:"eventType"`
	From        string                 `json:"from"`
	FromChannel int                    `json:"fromChannel"`
	ToChannel   int                    `json:"toChannel"`
	To          []string               `json:"to"`
	CreateTime  int                    `json:"createTime"`
	Content     map[string]interface{} `json:"content"`
}

type Content struct {
	ContentType int    `json:"contentType"`
	ToType      int    `json:"toType"`
	Text        string `json:"text"`
}

func init() {
	http.HandleFunc("/", receiveLine)
}

func receiveLine(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	opt := linebot.WithHTTPClient(client)
	bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN"), opt)
	if err != nil {
		log.Errorf(c, "failed to init linebot Client: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	events, err := bot.ParseRequest(r)
	if err != nil {
		log.Errorf(c, "failed to parse webhook: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, ev := range events {
		switch ev.Type {
		case linebot.EventTypeMessage:
			switch message := ev.Message.(type) {
			case *linebot.TextMessage:
				status := GetTranslateStatus(message.Text)
				switch status {
				case StatusAlphabet, StatusKana, StatusNumber:
					morsecode := TranslateToMorse(message.Text)
					log.Infof(c, "msg: %s", message.Text)
					log.Infof(c, "morsecode: %s", morsecode)
					if _, err = bot.ReplyMessage(
						ev.ReplyToken,
						linebot.NewTextMessage(morsecode),
					).Do(); err != nil {
						log.Errorf(c, "failed to reply message: %s", err)
					}
				case StatusMorse:
					template := linebot.NewButtonsTemplate(
						"", "モールスシンゴウヲヘンカン！", "ドチラニヘンカンシマスカ？",
						linebot.NewPostbackTemplateAction("アルファベット", TranslateToWords(message.Text, LanguageAlphabet), "アルファベットに変換"),
						linebot.NewPostbackTemplateAction("カナ", TranslateToWords(message.Text, LanguageKana), "カナに変換"),
					)
					if _, err := bot.ReplyMessage(
						ev.ReplyToken,
						linebot.NewTemplateMessage("ヒタイオウデス", template),
					).Do(); err != nil {
						log.Errorf(c, "failed to reply Button template message: %s", err)
					}
				}
			}
		case linebot.EventTypePostback:
			if _, err = bot.ReplyMessage(
				ev.ReplyToken,
				linebot.NewTextMessage(ev.Postback.Data),
			).Do(); err != nil {
				log.Errorf(c, "failed to postback message: %s", err)
			}
		}
	}
	fmt.Fprint(w, "OK")
}
