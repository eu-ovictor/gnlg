package relationship

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type FetchByIDResponseItem struct {
	Name          string         `json:"name"`
	Relationships []Relationship `json:"relationships"`
}

type FetchByIDResponse struct {
	Members []FetchByIDResponseItem `json:"members"`
}

type relationshipHandler struct {
	usecase Usecase
}

func AddRoutes(router *router.Router, usecase Usecase) {
	h := relationshipHandler{usecase}

	router.POST("/relationship", h.Add)
    router.GET("/relationship/{id}", h.FetchByID)
}

func (h relationshipHandler) Add(ctx *fasthttp.RequestCtx) {
	members := Members{}

	if err := json.Unmarshal(ctx.PostBody(), &members); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)

		msg := fmt.Sprintf("error decoding request body: %s", err.Error())
		ctx.SetBodyString(msg)
		return
	}

	if err := h.usecase.Add(members); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		msg := fmt.Sprintf("error adding relationship: %s", err.Error())
		ctx.SetBodyString(msg)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
}

func (h relationshipHandler) FetchByID(ctx *fasthttp.RequestCtx) {
	idParam := ctx.UserValue("id").(string)

	personID, _ := strconv.Atoi(idParam)

	relationships, err := h.usecase.FetchByID(int64(personID))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		msg := fmt.Sprintf("error fetching person relationships: %s", err.Error())
		ctx.SetBodyString(msg)
		return
	}

	responseBody := FetchByIDResponse{
        Members: make([]FetchByIDResponseItem, len(relationships)),
    }

    memberIdx := 0
	for member, relationships := range relationships {
		responseItem := FetchByIDResponseItem{Name: member, Relationships: relationships}

		responseBody.Members[memberIdx] = responseItem
        memberIdx++
	}

	response, err := json.Marshal(responseBody)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		msg := fmt.Sprintf("error encoding response body: %s", err.Error())
		ctx.SetBodyString(msg)
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(response)
}
