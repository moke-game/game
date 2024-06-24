package internal

import (
	cpb "github.com/moke-game/game/api/gen/common"
	"github.com/moke-game/game/services/bff/internal/biface"
)

type Registers struct {
	handlers map[cpb.C2S_EVENT]biface.IHandler
}

func (h *Registers) GetHandlers() map[cpb.C2S_EVENT]biface.IHandler {
	return h.handlers
}

func (h *Registers) GetHandler(event cpb.C2S_EVENT) biface.IHandler {
	return h.handlers[event]
}

func (h *Registers) RegisterHandler(event cpb.C2S_EVENT, handler biface.IHandler) {
	h.handlers[event] = handler
}

func (h *Registers) RegisterHandlers(start, end cpb.C2S_EVENT, handler biface.IHandler) {
	for i := start; i <= end; i++ {
		h.handlers[i] = handler
	}
}

func NewRegisters() *Registers {
	return &Registers{
		handlers: make(map[cpb.C2S_EVENT]biface.IHandler),
	}
}
