package server

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"rte-blog/types"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type StubPostModel struct {
	Store *sql.DB
}

func (model *StubPostModel) Create(title string) (int, error) {
	return 1, nil
}

func (model *StubPostModel) GetById(id int) (string, error) {
	return "Sample post", nil
}

func (model *StubPostModel) PutTitle(post types.Post) (types.Post, error) {
	return post, nil
}

func TestHandleGetPost(t *testing.T) {
	t.Run("returns a post with title, meta-data and content", func(t *testing.T) {
		postModel := &StubPostModel{Store: &sql.DB{}}

		server := server{
			config:    &http.Server{},
			postModel: postModel,
		}

		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/posts/1", nil)
		response := httptest.NewRecorder()

		context := e.NewContext(request, response)
		context.SetParamNames("id")
		context.SetParamValues("1")

		// Assertions
		if assert.NoError(t, server.handleGetPost(context)) {
			assert.Equal(t, http.StatusOK, response.Code)

			assert.Equal(t, "<p>Sample post</p>", response.Body.String())
		}
	})
}
