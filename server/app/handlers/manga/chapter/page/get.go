package page

import (
	"github.com/gin-gonic/gin"
	_type "github.com/imriyaya/mango/app/type"
	"net/http"
	"regexp"
)

func Get(server *_type.Server) {
	server.Router.GET("/manga/:manga_id/:chapter/:page", func(c *gin.Context) {
		mangaIdMatched, mangaIdError := regexp.MatchString(`^[a-fA-F0-9]{24}$`, c.Param("manga_id"))

		if !mangaIdMatched || mangaIdError != nil {
			_ = c.AbortWithError(http.StatusBadRequest, mangaIdError)
			return
		}

		chapterMatched, chapterError := regexp.MatchString(`[+-]?\d+`, c.Param("chapter"))

		if !chapterMatched || chapterError != nil {
			_ = c.AbortWithError(http.StatusBadRequest, mangaIdError)
			return
		}

		pageMatched, pageError := regexp.MatchString(`[+-]?\d+`, c.Param("page"))

		if !pageMatched || pageError != nil {
			_ = c.AbortWithError(http.StatusBadRequest, mangaIdError)
			return
		}

		c.File("./mangas/" + c.Param("manga_id") + "/" + c.Param("chapter") + "/" + c.Param("page") + ".jpg")
	})
}
