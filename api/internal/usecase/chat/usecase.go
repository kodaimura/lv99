package chat

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"

	chatModule "lv99/internal/module/chat"
)

type Usecase interface {
	Get(in GetDto, db *gorm.DB) ([]chatModule.Chat, error)
	CreateOne(in CreateOneDto, db *gorm.DB) (chatModule.Chat, error)
	UpdateOne(in UpdateOneDto, db *gorm.DB) (chatModule.Chat, error)
	DeleteOne(in DeleteOneDto, db *gorm.DB) error
	Read(in ReadDto, db *gorm.DB) error
	Paginate(in PaginateDto, db *sqlx.DB) ([]chatModule.Chat, error)
}

type usecase struct {
	chatService chatModule.Service
}

func NewUsecase(chatService chatModule.Service) Usecase {
	return &usecase{
		chatService: chatService,
	}
}

func (srv *usecase) Get(in GetDto, db *gorm.DB) ([]chatModule.Chat, error) {
	return srv.chatService.Get(chatModule.Chat{
		FromId: in.FromId,
		ToId:   in.ToId,
	}, db)
}

func (srv *usecase) CreateOne(in CreateOneDto, db *gorm.DB) (chatModule.Chat, error) {
	return srv.chatService.CreateOne(chatModule.Chat{
		FromId:  in.FromId,
		ToId:    in.ToId,
		Message: in.Message,
	}, db)
}

func (srv *usecase) UpdateOne(in UpdateOneDto, db *gorm.DB) (chatModule.Chat, error) {
	chat, err := srv.chatService.GetOne(chatModule.Chat{Id: in.Id}, db)
	if err != nil {
		return chatModule.Chat{}, err
	}

	chat.Message = in.Message
	if in.IsRead {
		chat.IsRead = in.IsRead
	}
	return srv.chatService.UpdateOne(chat, db)
}

func (srv *usecase) DeleteOne(in DeleteOneDto, db *gorm.DB) error {
	return srv.chatService.DeleteOne(chatModule.Chat{Id: in.Id}, db)
}

func (srv *usecase) Read(in ReadDto, db *gorm.DB) error {
	return srv.chatService.Read(chatModule.Chat{FromId: in.FromId, ToId: in.ToId}, db)
}

func (srv *usecase) Paginate(in PaginateDto, db *sqlx.DB) ([]chatModule.Chat, error) {
	return srv.chatService.Paginate(in.FromId, in.ToId, in.Before, in.Limit, db)
}
