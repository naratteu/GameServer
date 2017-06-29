package main

import (
	"net"
)

type room struct {
	MapId int
}

func NewRoom() *room {
	return new(room)
}

func (r *room) AddVisitor() {

}
