package controller

import (
	"net/http"

	"github.com/IgnacyBialobrzewski/go-chat/utils"
	"github.com/IgnacyBialobrzewski/go-chat/views"
)

type Index struct{}

func (self Index) Routes() {
	http.HandleFunc("/", self.base)
}

func (self Index) base(w http.ResponseWriter, r *http.Request) {
	util.Render(w, views.Index())
}
