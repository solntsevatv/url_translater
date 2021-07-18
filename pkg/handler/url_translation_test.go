package handler

import (
	"bytes"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/solntsevatv/url_translater/internal/url_translater"
	"github.com/solntsevatv/url_translater/pkg/service"

	"net/http/httptest"
	"testing"

	service_mocks "github.com/solntsevatv/url_translater/pkg/service/mocks_pkg"
)

func TestHandler_longToShort(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *service_mocks.MockUrlTranslation, long_url url_translater.LongURL)

	tests := []struct {
		name                 string
		inputBody            string
		inputUrl             url_translater.LongURL
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"url": "some_long_url"}`,
			inputUrl: url_translater.LongURL{
				LinkUrl: "some_long_url",
			},
			mockBehavior: func(r *service_mocks.MockUrlTranslation, long_url url_translater.LongURL) {
				r.EXPECT().CreateShortURL(long_url).Return("AA", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"url":"AA"}`,
		},
		{
			name:      "Wrong Input",
			inputBody: `{"lds": ""}`,
			inputUrl: url_translater.LongURL{
				Id:      10,
				LinkUrl: "",
			},
			mockBehavior:         func(r *service_mocks.MockUrlTranslation, long_url url_translater.LongURL) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockUrlTranslation(c)
			test.mockBehavior(repo, test.inputUrl)

			services := &service.Service{UrlTranslation: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/short", handler.longToShort)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/short",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_shortToLong(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *service_mocks.MockUrlTranslation, short_url url_translater.ShortURL)

	tests := []struct {
		name                 string
		inputBody            string
		inputUrl             url_translater.ShortURL
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "wrong input",
			inputBody: `{"wrong": ""}`,
			inputUrl: url_translater.ShortURL{
				Id:      1,
				LinkUrl: "",
			},
			mockBehavior:         func(r *service_mocks.MockUrlTranslation, short_url url_translater.ShortURL) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "bd hasn't the url",
			inputBody: `{"url": "some_short_url"}`,
			inputUrl:  url_translater.ShortURL{Id: 1, LinkUrl: "some_short_url"},
			mockBehavior: func(r *service_mocks.MockUrlTranslation, short_url url_translater.ShortURL) {
				r.EXPECT().GetLongURL(short_url).Return("", errors.New("url was not found"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"url was not found"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockUrlTranslation(c)
			test.mockBehavior(repo, test.inputUrl)

			services := &service.Service{UrlTranslation: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/long", handler.ShortToLong)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/long",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
