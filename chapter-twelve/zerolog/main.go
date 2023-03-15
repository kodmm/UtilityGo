package main

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const DBConnErrCD = "E10001"

func businessLogic() {
	db, err := openDB()
	if err != nil {
		// エラーコード付きで出力
		log.Error().Str("code", DBConnErrCD).Err(err).Msg("db connection failed")
	}
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// それぞれのレベルで出力
	log.Error().Msg("errorで出力")
	log.Info().Msg("infoで出力")
	log.Debug().Msgf("debugで出力: %v", "ただし出力されない")
	log.Info().
		Str("app", "awesome-app").
		Int("user_id", 1234).Send()

	logger := log.With().Int("user_id", 1024).Str("path", "/api/user").Str("method", "post").
		Logger()

	ctx := context.Background()
	// context.Contextにロガーを登録
	ctx = logger.WithContext(ctx)

	// context.Contextに設定したロガーを取り出し
	newLogger := zerolog.Ctx(ctx)
	newLogger.Print("debug message")
}
