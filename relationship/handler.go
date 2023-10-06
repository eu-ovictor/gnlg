package relationship

import (
	"encoding/json"
	"fmt"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type relationshipHandler struct {
	usecase Usecase
}

func AddRoutes(router *router.Router, usecase Usecase) {
	h := relationshipHandler{usecase}

	router.POST("/relationship", h.Add)
}

func (h relationshipHandler) Add(ctx *fasthttp.RequestCtx) {
	rel := Relationship{}

	if err := json.Unmarshal(ctx.PostBody(), &rel); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)

		msg := fmt.Sprintf("error decoding request body: %s", err.Error())
		ctx.SetBodyString(msg)
        return
	}

	if err := h.usecase.Add(rel); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		msg := fmt.Sprintf("error adding relationship: %s", err.Error())
		ctx.SetBodyString(msg)
        return
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
}
