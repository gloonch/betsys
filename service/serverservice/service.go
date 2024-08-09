package serverservice

import (
	"betsys/entity"
	"fmt"
)

type Repository interface {
	Register(s entity.Server) (entity.Server, error)
	FindServerByName(name string) (entity.Server, error)
	IsServerNameTaken(name string) (bool, error)
}

type Service struct {
	repository Repository
}

type RegisterRequest struct {
	Name string `json:"name"`
}

type ServerInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type RegisterResponse struct {
	ServerInfo ServerInfo `json:"server_info"`
}

func New(repository Repository) Service {
	return Service{
		repository: repository,
	}
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	// validation on future properties

	// validate name
	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name is too short")
	}

	// check uniqueness of Name
	isTaken, err := s.repository.IsServerNameTaken(req.Name)
	if err == nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error while checking if Name exists: %w", err)
	}
	if isTaken {
		return RegisterResponse{}, fmt.Errorf("serverservice name already taken: %w", err)
	}

	createdServer, err := s.repository.Register(entity.Server{
		Name: req.Name,
	})
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("failed to register serverservice: %w", err)
	}

	return RegisterResponse{ServerInfo: ServerInfo{
		ID:   createdServer.ID,
		Name: createdServer.Name,
	}}, nil
}
