package healthctlr

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewHealthController(t *testing.T) {
	healthController := NewHealthController()
	assert.NotNil(t, healthController)
}

func Test_healthController_GetHealth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	type args struct {
		c echo.Context
	}

	type want struct {
		statusCode int
		response   map[string]string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    want
	}{
		{
			name:    "success",
			args:    args{c: c},
			wantErr: false,
			want: want{
				statusCode: http.StatusOK,
				response:   map[string]string{"status": "up"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ct := &healthController{}
			if err := ct.GetHealth(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("healthController.GetHealth() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.want.statusCode, rec.Code)

			var response map[string]string
			if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
				t.Errorf("healthController.GetHealth() cannot unmarshal response body")
			}
			assert.Equal(t, tt.want.response, response)
		})
	}
}
