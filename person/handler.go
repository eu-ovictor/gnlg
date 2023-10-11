package person

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type FetchResponse struct {
	People []Person `json:"people"`
}

type editPersonRequest struct {
	Name string `json:"name"`
}

type EditPersonResponse struct {
	Edited int64 `json:"edited"`
}

type personHandler struct {
	usecase Usecase
}

func AddRoutes(router *router.Router, usecase Usecase) {
	handler := personHandler{usecase}

	router.GET("/person", handler.Fetch)
	router.POST("/person", handler.Add)
	router.PUT("/person/{id}", handler.Edit)
}

func (h personHandler) Add(ctx *fasthttp.RequestCtx) {
	person := Person{}

	if err := json.Unmarshal(ctx.PostBody(), &person); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)

		msg := fmt.Sprintf("error decoding request body: %s", err.Error())
		ctx.SetBodyString(msg)
		return
	}

	err := h.usecase.Add(person)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		msg := fmt.Sprintf("error adding person: %s", err.Error())
		ctx.SetBodyString(msg)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
}

func (h personHandler) Edit(ctx *fasthttp.RequestCtx) {
	idParam := ctx.UserValue("id").(string)

	personID, _ := strconv.Atoi(idParam)

	requestBody := editPersonRequest{}

	if err := json.Unmarshal(ctx.Request.Body(), &requestBody); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)

		msg := fmt.Sprintf("error decoding request body: %s", err.Error())
		ctx.SetBodyString(msg)
		return
	}

	person := Person{ID: personID, Name: requestBody.Name}

	edited, err := h.usecase.Edit(person)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		msg := fmt.Sprintf("error editing person: %s", err.Error())
		ctx.SetBodyString(msg)
		return
	}

	response, err := json.Marshal(EditPersonResponse{Edited: edited})
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		msg := fmt.Sprintf("error encoding response body: %s", err.Error())
		ctx.SetBodyString(msg)
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(response)
}

func (h personHandler) Fetch(ctx *fasthttp.RequestCtx) {
	queryArgs := ctx.QueryArgs()

	ID := queryArgs.GetUintOrZero("id")
	name := string(queryArgs.Peek("name"))

	people, err := h.usecase.Fetch(ID, name)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		msg := fmt.Sprintf("error fetching people: %s", err.Error())
		ctx.SetBodyString(msg)
		return
	}

	response, err := json.Marshal(FetchResponse{People: people})
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		msg := fmt.Sprintf("error encoding response body: %s", err.Error())
		ctx.SetBodyString(msg)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetBody(response)
}
