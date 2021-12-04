package pumpctrl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	pumpmodel "plant-system-api/api/models/pump"
	mockPumpserviceitf "plant-system-api/api/services/pump/mock"
	mockConfig "plant-system-api/config/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewPumpController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		firebaseClient         *mockConfig.MockClient
		firebaseClientBehavior func(m *mockConfig.MockClient)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				firebaseClient: mockConfig.NewMockClient(ctrl),
				firebaseClientBehavior: func(m *mockConfig.MockClient) {
					m.EXPECT().NewRef(gomock.Any())
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.args.firebaseClientBehavior(tt.args.firebaseClient)

		t.Run(tt.name, func(t *testing.T) {
			got := NewPumpController(tt.args.firebaseClient)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func Test_pumpController_ActivePump(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		pumpID   = 1
		isActive = true
	)

	var (
		fooErr                                     = errors.New("foo")
		successContext, successRec                 = getContextOfActivePumpSuccess(pumpID, isActive)
		withoutIDContext, withoutIDRec             = getContextOfActivePumpWithoutIDParam()
		withoutIsActiveContext, withoutIsActiveRec = getContextOfActivePumpWithoutIsActiveQuery(pumpID)
		activePumpErrContext, activePumpErrRec     = getContextOfActivePumpSuccess(pumpID, isActive)
	)

	type fields struct {
		pumpService         *mockPumpserviceitf.MockPumpService
		pumpServiceBehavior func(m *mockPumpserviceitf.MockPumpService)
	}
	type args struct {
		c echo.Context
	}
	type want struct {
		statusCode int
		response   *pumpmodel.Pump
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		recorder *httptest.ResponseRecorder
		want     want
	}{
		{
			name: "success",
			fields: fields{
				pumpService: mockPumpserviceitf.NewMockPumpService(ctrl),
				pumpServiceBehavior: func(m *mockPumpserviceitf.MockPumpService) {
					m.EXPECT().ActivePump(gomock.Any()).Return(nil)
				},
			},
			args: args{
				c: successContext,
			},
			recorder: successRec,
			want: want{
				statusCode: http.StatusOK,
				response: &pumpmodel.Pump{
					ID:   pumpID,
					IsActive: isActive,
				},
			},
		},
		{
			name: "error without id param",
			fields: fields{
				pumpService:         mockPumpserviceitf.NewMockPumpService(ctrl),
				pumpServiceBehavior: func(m *mockPumpserviceitf.MockPumpService) {},
			},
			args: args{
				c: withoutIDContext,
			},
			recorder: withoutIDRec,
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "error without isActive query",
			fields: fields{
				pumpService:         mockPumpserviceitf.NewMockPumpService(ctrl),
				pumpServiceBehavior: func(m *mockPumpserviceitf.MockPumpService) {},
			},
			args: args{
				c: withoutIsActiveContext,
			},
			recorder: withoutIsActiveRec,
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "error ActivePump service",
			fields: fields{
				pumpService: mockPumpserviceitf.NewMockPumpService(ctrl),
				pumpServiceBehavior: func(m *mockPumpserviceitf.MockPumpService) {
					m.EXPECT().ActivePump(gomock.Any()).Return(fooErr)
				},
			},
			args: args{
				c: activePumpErrContext,
			},
			recorder: activePumpErrRec,
			want: want{
				statusCode: http.StatusInternalServerError,
			},
		},
	}
	for _, tt := range tests {
		tt.fields.pumpServiceBehavior(tt.fields.pumpService)

		t.Run(tt.name, func(t *testing.T) {
			ct := &pumpController{
				pumpService: tt.fields.pumpService,
			}
			ct.ActivePump(tt.args.c)

			assert.Equal(t, tt.want.statusCode, tt.recorder.Code)
			if tt.recorder.Code == http.StatusOK {
				var response *pumpmodel.Pump
				if err := json.Unmarshal(tt.recorder.Body.Bytes(), &response); err != nil {
					t.Errorf("pumpController.ActivePump() cannot unmarshal response body")
				}
				assert.Equal(t, tt.want.response, response)
			}
		})
	}
}

func getContextOfActivePumpSuccess(pumpID int, isActive bool) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	query := req.URL.Query()
	query.Add("isActive", fmt.Sprintf("%v", isActive))

	req.URL.RawQuery = query.Encode()

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/pumps/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", pumpID))

	return c, rec
}

func getContextOfActivePumpWithoutIDParam() (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c, rec
}

func getContextOfActivePumpWithoutIsActiveQuery(pumpID int) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/pumps/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", pumpID))

	return c, rec
}
