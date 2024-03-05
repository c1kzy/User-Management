package handlers

import (
	"net/http"
	"net/http/httptest"
	"restapi/internal"
	"restapi/internal/domain"
	"restapi/internal/mocks"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_getID(t *testing.T) {

	var getCtx = func(id int) *gin.Context {
		ctx := &gin.Context{}
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: strconv.Itoa(id),
			},
		}
		return ctx
	}

	tests := []struct {
		name    string
		c       *gin.Context
		want    int
		wantErr bool
	}{
		{
			name:    "ok",
			c:       getCtx(1),
			want:    1,
			wantErr: false,
		},
		{
			name:    "not ok",
			c:       &gin.Context{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getID(tt.c)
			if tt.wantErr == true {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestHandler_userIdentity(t *testing.T) {
	ctrl := gomock.NewController(t)
	userService := mocks.NewUserService(ctrl)
	service := &internal.Service{UserService: userService}

	handler := Handler{service: service}

	r := gin.New()
	r.GET("/identity", handler.UserIdentity, func(c *gin.Context) {
		id, _ := c.Get("id")
		c.String(200, "%d", id)
	})

	tests := []struct {
		name                 string
		token                string
		headerName           string
		headerValue          string
		mock                 func(identity *mocks.UserService, token string)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "ok",
			token:       "token",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			mock: func(identity *mocks.UserService, token string) {
				identity.EXPECT().ParseToken(token).Return(&domain.TokenClaims{
					RegisteredClaims: jwt.RegisteredClaims{},
					UserID:           1,
					Role:             1,
				}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: "1",
		},
		{
			name:                 "invalid header name",
			token:                "token",
			headerName:           "",
			headerValue:          "Bearer ",
			mock:                 func(identity *mocks.UserService, token string) {},
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: `{"message":"empty auth header"}`,
		},
		{
			name:                 "invalid header value",
			token:                "token",
			headerName:           "Authorization",
			headerValue:          "Brr ",
			mock:                 func(identity *mocks.UserService, token string) {},
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name:                 "empty token",
			token:                "",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			mock:                 func(identity *mocks.UserService, token string) {},
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: `{"message":"token is empty"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(userService, tt.token)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/identity", nil)
			req.Header.Set(tt.headerName, tt.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, tt.expectedStatusCode)
			assert.Equal(t, w.Body.String(), tt.expectedResponseBody)
		})
	}
}
