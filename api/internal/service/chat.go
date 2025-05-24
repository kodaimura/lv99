package service

import (
	"lv99/internal/dto/input"
	"lv99/internal/model"
	"lv99/internal/query"
	"lv99/internal/repository"
)

type ChatService interface {
	Get(in input.GetChat) ([]model.Chat, error)
	CreateOne(in input.Chat) (model.Chat, error)
	UpdateOne(in input.Chat) (model.Chat, error)
	DeleteOne(in input.Chat) error
}

type chatService struct {
	chatRepository repository.ChatRepository
	chatQuery query.ChatQuery
}

func NewChatService(
	chatRepository repository.ChatRepository, 
	chatQuery query.ChatQuery,
) ChatService {
	return &chatService{
		chatRepository: chatRepository,
		chatQuery: chatQuery,
	}
}

func (srv *chatService) Get(in input.GetChat) ([]model.Chat, error) {
	return srv.chatQuery.Get(in.FromId, in.ToId, in.Before, in.Limit)
}

func (srv *chatService) CreateOne(in input.Chat) (model.Chat, error) {
	return srv.chatRepository.Insert(&model.Chat{
		FromId: in.FromId,
		ToId: in.ToId,
		Message: in.Message,
	})
}

func (srv *chatService) UpdateOne(in input.Chat) (model.Chat, error) {
	chat, err := srv.chatRepository.GetOne(&model.Chat{
		Id: in.Id,
	})
	if err != nil {
		return model.Chat{}, err
	}

	chat.Message = in.Message
	if in.IsRead {
		chat.IsRead = in.IsRead
	}	
	return srv.chatRepository.Update(&chat)
}

func (srv *chatService) DeleteOne(in input.Chat) error {
	return srv.chatRepository.Delete(&model.Chat{
		Id: in.Id,
	})
}