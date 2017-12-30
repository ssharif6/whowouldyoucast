package handlers

import (
	"net/http"
	"github.com/ssharif6/whowouldyoucast/servers/gateway/models"
	"encoding/json"
)

type HandlerCtx struct {
	key string
	reviewStore *models.ReviewStore
}

func NewHandlerCtx (key string, reviewStore *models.ReviewStore) *HandlerCtx {
	return &HandlerCtx{
		reviewStore: reviewStore,
		key: key,
	}
}

func (ctx *HandlerCtx) ReviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "action not supported", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	review := &models.Review{}
	if err := json.NewDecoder(r.Body).Decode(review); err != nil {
		http.Error(w, "error decoding body", http.StatusBadRequest)
		return
	}

	// Post to mongo
	if _, err := ctx.reviewStore.PostReview(review); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}