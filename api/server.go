package api

import (
	github.com/google/uuid
	github.com/gorilla/mux
)

type Item struct {
	ID   uuid.UUID
	Name string
}

type Server struct {
	*mux.Router
	shoppingItems []Item
}
