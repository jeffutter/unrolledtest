package checkout

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CartContents map[string]int

func TestIndex(t *testing.T) {

	Convey("Given a cart with no items in it", t, func() {

		Convey("The checkout page should show an error", func() {
			r, _ := http.NewRequest("GET", "", nil)
			w := httptest.NewRecorder()

			Checkout(w, r)

			So(w.Body.String(), ShouldContainSubstring, "Your Cart Is Empty")
		})

	})
}
