package server

import "errors"

type Client struct {
	UserID  string
	Balance int
}

type Server struct {
	clients map[string]*Client
}

func (s Server) Register(client *Client) (Client, error) {
	if _, exists := s.clients[client.UserID]; exists {
		return Client{}, errors.New("client already registered")
	}
	s.clients[client.UserID] = client
	return *client, nil
}

//func (s Server) List() ([]Client, error) {
//	if len(s.clients) != 0 {
//		return , nil
//	}
//}
