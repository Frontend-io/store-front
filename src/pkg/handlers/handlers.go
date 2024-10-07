package handlers

import (
	"net/http"

	api "github.com/Frontend-io/store-front/pkg/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	api.RenderTemplate(w, "home.page.html")
}
