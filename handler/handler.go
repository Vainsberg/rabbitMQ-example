package handler

import (
	"fmt"
	"net/http"

	"github.com/Vainsberg/rabbitMQ-example/rabbitMQ"
)

type Handler struct {
	RabbitMQ *rabbitMQ.RepositoryRabbitMQ
}

func NewHandler(rabbitMQ *rabbitMQ.RepositoryRabbitMQ) *Handler {
	return &Handler{
		RabbitMQ: rabbitMQ,
	}
}

func (h *Handler) ExampleHanlder(w http.ResponseWriter, r *http.Request) {
	err := h.RabbitMQ.PublishMessage("text", "hello")
	if err != nil {
		http.Error(w, "Ошибка при публикации сообщения", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Сообщение отправлено")

	h.RabbitMQ.ConsumeMessages("text")

}
