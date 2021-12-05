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
	"strings"
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
		fooErr                                   = errors.New("foo")
		successContext, successRec               = getContextOfActivePumpSuccess(pumpID, isActive, t)
		withoutIDParamContext, withoutIDParamRec = getContextOfActivePumpWithoutIDParam(pumpID)
		withoutReqBodyContext, withoutReqBodyRec = getContextOfActivePumpWithoutReqBody(pumpID)
		activePumpErrContext, activePumpErrRec   = getContextOfActivePumpError(pumpID, isActive, t)
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
					ID:       pumpID,
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
				c: withoutIDParamContext,
			},
			recorder: withoutIDParamRec,
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "error without req body",
			fields: fields{
				pumpService:         mockPumpserviceitf.NewMockPumpService(ctrl),
				pumpServiceBehavior: func(m *mockPumpserviceitf.MockPumpService) {},
			},
			args: args{
				c: withoutReqBodyContext,
			},
			recorder: withoutReqBodyRec,
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

func getContextOfActivePumpSuccess(pumpID int, isActive bool, t *testing.T) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	pumpActiveReq := pumpmodel.PumpActiveReq{IsActive: isActive}
	pumpActiveReqB, err := json.Marshal(pumpActiveReq)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(string(pumpActiveReqB)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/active")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", pumpID))

	return c, rec
}

func getContextOfActivePumpWithoutIDParam(pumpID int) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPatch, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c, rec
}

func getContextOfActivePumpWithoutReqBody(pumpID int) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(`{"is_active":"foo"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/active")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", pumpID))

	return c, rec
}

func getContextOfActivePumpError(pumpID int, isActive bool, t *testing.T) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	pumpActiveReq := pumpmodel.PumpActiveReq{IsActive: isActive}
	pumpActiveReqB, err := json.Marshal(pumpActiveReq)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(string(pumpActiveReqB)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/active")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", pumpID))

	return c, rec
}

func Test_pumpController_AskPump(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		pumpID = 1
	)

	var (
		fooErr                                   = errors.New("foo")
		expected                                 = &pumpmodel.Pump{ID: pumpID, IsAsk: false, IsWorking: true}
		successContext, successRec               = getContextOfAskPumpSuccess(pumpID)
		withoutIDParamContext, withoutIDParamRec = getContextOfAskPumpWithoutIDParam()
		askPumpErrContext, askPumpErrRec         = getContextOfAskPumpAndAskPumpServiceError(pumpID)
		getPumpErrContext, getPumpErrRec         = getContextOfAskPumpAndGetPumpServiceError(pumpID)
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
					m.EXPECT().AskPump(gomock.Any()).Return(nil)
					m.EXPECT().GetPump(gomock.Any()).Return(expected, nil)
				},
			},
			args: args{
				c: successContext,
			},
			recorder: successRec,
			want: want{
				statusCode: http.StatusOK,
				response:   expected,
			},
		},
		{
			name: "error without id param",
			fields: fields{
				pumpService:         mockPumpserviceitf.NewMockPumpService(ctrl),
				pumpServiceBehavior: func(m *mockPumpserviceitf.MockPumpService) {},
			},
			args: args{
				c: withoutIDParamContext,
			},
			recorder: withoutIDParamRec,
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "error asking pump",
			fields: fields{
				pumpService: mockPumpserviceitf.NewMockPumpService(ctrl),
				pumpServiceBehavior: func(m *mockPumpserviceitf.MockPumpService) {
					m.EXPECT().AskPump(gomock.Any()).Return(fooErr)
				},
			},
			args: args{
				c: askPumpErrContext,
			},
			recorder: askPumpErrRec,
			want: want{
				statusCode: http.StatusInternalServerError,
			},
		},
		{
			name: "error getting pump",
			fields: fields{
				pumpService: mockPumpserviceitf.NewMockPumpService(ctrl),
				pumpServiceBehavior: func(m *mockPumpserviceitf.MockPumpService) {
					m.EXPECT().AskPump(gomock.Any()).Return(nil)
					m.EXPECT().GetPump(gomock.Any()).Return(nil, fooErr)
				},
			},
			args: args{
				c: getPumpErrContext,
			},
			recorder: getPumpErrRec,
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
			ct.AskPump(tt.args.c)

			assert.Equal(t, tt.want.statusCode, tt.recorder.Code)
			if tt.recorder.Code == http.StatusOK {
				var response *pumpmodel.Pump
				if err := json.Unmarshal(tt.recorder.Body.Bytes(), &response); err != nil {
					t.Errorf("pumpController.AskPump() cannot unmarshal response body")
				}
			}
		})
	}
}

func getContextOfAskPumpSuccess(pumpID int) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPatch, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/ask")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", pumpID))

	return c, rec
}

func getContextOfAskPumpWithoutIDParam() (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPatch, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c, rec
}

func getContextOfAskPumpAndAskPumpServiceError(pumpID int) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPatch, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/ask")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", pumpID))

	return c, rec
}

func getContextOfAskPumpAndGetPumpServiceError(pumpID int) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPatch, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/ask")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", pumpID))

	return c, rec
}

func Test_pumpController_GetPump(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		pumpID = 1
	)

	var (
		fooErr                                   = errors.New("foo")
		expected                                 = &pumpmodel.Pump{ID: pumpID, IsAsk: false, IsWorking: true}
		successContext, successRec               = getContextOfGetPumpSuccess(pumpID)
		withoutIDParamContext, withoutIDParamRec = getContextOfGetPumpWithoutIDParam()
		getPumpErrContext, getPumpErrRec         = getContextOfGetPumpAndGetPumpServiceError(pumpID)
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
					m.EXPECT().GetPump(gomock.Any()).Return(expected, nil)
				},
			},
			args: args{
				c: successContext,
			},
			recorder: successRec,
			want: want{
				statusCode: http.StatusOK,
				response:   expected,
			},
		},
		{
			name: "error without id param",
			fields: fields{
				pumpService:         mockPumpserviceitf.NewMockPumpService(ctrl),
				pumpServiceBehavior: func(m *mockPumpserviceitf.MockPumpService) {},
			},
			args: args{
				c: withoutIDParamContext,
			},
			recorder: withoutIDParamRec,
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "error getting pump",
			fields: fields{
				pumpService: mockPumpserviceitf.NewMockPumpService(ctrl),
				pumpServiceBehavior: func(m *mockPumpserviceitf.MockPumpService) {
					m.EXPECT().GetPump(gomock.Any()).Return(nil, fooErr)
				},
			},
			args: args{
				c: getPumpErrContext,
			},
			recorder: getPumpErrRec,
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
			ct.GetPump(tt.args.c)

			assert.Equal(t, tt.want.statusCode, tt.recorder.Code)
			if tt.recorder.Code == http.StatusOK {
				var response *pumpmodel.Pump
				if err := json.Unmarshal(tt.recorder.Body.Bytes(), &response); err != nil {
					t.Errorf("pumpController.GetPump() cannot unmarshal response body")
				}
			}
		})
	}
}

func getContextOfGetPumpSuccess(pumpID int) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPatch, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", pumpID))

	return c, rec
}

func getContextOfGetPumpWithoutIDParam() (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPatch, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c, rec
}
func getContextOfGetPumpAndGetPumpServiceError(pumpID int) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPatch, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/ask")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", pumpID))

	return c, rec
}
