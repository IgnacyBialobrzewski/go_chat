package util

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

func Render(w http.ResponseWriter, comp templ.Component) {
	comp.Render(context.Background(), w)
}
