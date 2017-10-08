package game_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/gapi/game"
	"github.com/gin-gonic/gin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(FakeGameDataContextMW())
	r.GET("/game/:Id", RetrieveSingleGame)
	r.GET("/games", RetrieveAllGames)

	Context("When RetrieveAllGames is called  ", func() {
		It("returns a 200", func() {

			req, _ := http.NewRequest("GET", "/games", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			fmt.Println(w.Body.String())
			statusOK := w.Code == http.StatusOK
			fmt.Println(w)

			expected := `{"data":[{"Id":1,"name":"(("},{"Id":2,"name":"))"},{"Id":3,"name":"(( ))"}],"status":200}`
			Expect(statusOK).To(Equal(true))
			Expect(w.Body.String()).To(Equal(expected))
		})
	})

	Context("When RetrieveSingleGame is called ", func() {

		It("Returns a 200 when the game exists", func() {
			req, _ := http.NewRequest("GET", "/game/1", nil)
			fmt.Println(req)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			statusOK := w.Code == http.StatusOK
			fmt.Println(w)

			expected := `{"data":{"Id":1,"name":"The Ungame"},"status":200}`
			Expect(statusOK).To(Equal(true))
			Expect(w.Body.String()).To(Equal(expected))

		})
		It("Returns a 404 when the game does not exist", func() {
			req, _ := http.NewRequest("GET", "/game/2", nil)
			fmt.Println(req)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			statusMatch := w.Code == http.StatusNotFound
			fmt.Println(w)

			Expect(statusMatch).To(Equal(true))
		})
		It("Returns a 500 when input is invalid", func() {
			req, _ := http.NewRequest("GET", "/game/a", nil)
			fmt.Println(req)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			statusMatch := w.Code == http.StatusInternalServerError
			fmt.Println(w)

			Expect(statusMatch).To(Equal(true))
		})
	})

})
