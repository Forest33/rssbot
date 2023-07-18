package usecase

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/forest33/rssbot/business/entity"
)

func (uc *BotUseCase) createUser(ctx context.Context, msg *tgbotapi.Message, _ []string) error {
	_, err := uc.usersRepo.Create(ctx, &entity.User{
		ID:        msg.From.ID,
		FirstName: msg.From.FirstName,
		LastName:  msg.From.LastName,
		UserName:  msg.From.UserName,
		Language:  msg.From.LanguageCode,
	})
	if err != nil {
		return err
	}

	return uc.reply(msg, entity.MessageStart+"\n"+entity.MessageHelp)
}
