package chat

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"

	chatModule "lv99/internal/module/chat"
)

type Usecase interface {
	Get(in GetDto) ([]chatModule.Chat, error)
	CreateOne(in CreateOneDto) (chatModule.Chat, error)
	UpdateOne(in UpdateOneDto) (chatModule.Chat, error)
	DeleteOne(in DeleteOneDto) error
	Read(in ReadDto) error
	Paginate(in PaginateDto) ([]chatModule.Chat, error)
}

type usecase struct {
	db          *gorm.DB
	dbx         *sqlx.DB
	chatService chatModule.Service
}

func NewUsecase(
	db *gorm.DB,
	dbx *sqlx.DB,
	chatService chatModule.Service,
) Usecase {
	return &usecase{
		db:          db,
		dbx:         dbx,
		chatService: chatService,
	}
}

func (uc *usecase) Get(in GetDto) ([]chatModule.Chat, error) {
	return uc.chatService.Get(chatModule.Chat{
		FromId: in.FromId,
		ToId:   in.ToId,
	}, uc.db)
}

func (uc *usecase) CreateOne(in CreateOneDto) (chatModule.Chat, error) {
	return uc.chatService.CreateOne(chatModule.Chat{
		FromId:  in.FromId,
		ToId:    in.ToId,
		Message: in.Message,
	}, uc.db)
}

func (uc *usecase) UpdateOne(in UpdateOneDto) (chatModule.Chat, error) {
	chat, err := uc.chatService.GetOne(chatModule.Chat{Id: in.Id}, uc.db)
	if err != nil {
		return chatModule.Chat{}, err
	}

	chat.Message = in.Message
	if in.IsRead {
		chat.IsRead = in.IsRead
	}
	return uc.chatService.UpdateOne(chat, uc.db)
}

func (uc *usecase) DeleteOne(in DeleteOneDto) error {
	return uc.chatService.DeleteOne(chatModule.Chat{Id: in.Id}, uc.db)
}

func (uc *usecase) Read(in ReadDto) error {
	return uc.chatService.Read(chatModule.Chat{FromId: in.FromId, ToId: in.ToId}, uc.db)
}

func (uc *usecase) Paginate(in PaginateDto) ([]chatModule.Chat, error) {
	return uc.chatService.Paginate(in.FromId, in.ToId, in.Before, in.Limit, uc.dbx)
}
