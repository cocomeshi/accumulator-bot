package bot

import (
	"context"
	"log"
	"os"

	repo "github.com/cocomeshi/accumulator-bot/infrastructure"
	"github.com/cocomeshi/accumulator-bot/internal"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func BotEntry(ctx context.Context, m PubSubMessage) error {

	message := string(m.Data)
	log.Println(message)
	// Pub/Subから送信されたメッセージが異なる場合、処理を終了する
	if message != "bot start" {
		return nil
	}
	apiKey := os.Getenv("GOOGLEAPI_KEY")
	internal.Exec(apiKey)
	internal.AdditionalUpdate(apiKey)
	db := repo.GetInstance()
	defer db.Close()

	return nil
}
