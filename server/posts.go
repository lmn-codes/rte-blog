package server

import (
	"log"
	"net/http"
	"rte-blog/services"
	"rte-blog/templates"
	"rte-blog/types"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func (server *server) handleGetIndex(context echo.Context) error {
	return templates.Render(context, http.StatusOK, templates.DefaultLayout(templates.NewPostButton()))
}

func (server *server) handleCreatePost(context echo.Context) error {
	title := context.FormValue("title")
	_, err := server.postModel.Create(title)

	return err
}

func (server *server) handleGetPost(context echo.Context) error {
	id, err := services.ExtractIdFromContext(context)
	if err != nil {
		return err
	}

	title, err := server.postModel.GetById(id)
	post := types.Post{
		Id:          id,
		Title:       title,
		AuthorName:  "",
		PublishedAt: time.Now(),
		Contents:    []types.Content{},
	}

	if err != nil {
		log.Fatal(err)
	}

	return templates.Render(context, http.StatusOK, templates.PostLayout(post))
}

func (server *server) handlePutPostTitle(context echo.Context) error {
	id, err := services.ExtractIdFromContext(context)
	if err != nil {
		return err
	}

	title := context.FormValue("title")
	post := types.Post{
		Title: title,
		Id:    id,
	}

	_, err = server.postModel.PutTitle(post)

	return err
}

func (server *server) handleParagraphCreate(context echo.Context) error {
	postId, err := services.ExtractIdFromContext(context)
	if err != nil {
		return err
	}

	orderInPost, err := strconv.Atoi(context.Param("orderInPost"))
	if err != nil {
		return err
	}

	paragraph, err := server.postModel.CreatePostContent(postId, orderInPost)
	if err != nil {
		return err
	}

	contents := [1]types.Content{}
	contents[0] = *paragraph

	return templates.Render(context, http.StatusOK, templates.Main(types.Post{Title: "new", AuthorName: "Linh", Contents: contents[:]}))
}
