package pumpservice

import (
	"errors"
	pumpmodel "plant-system-api/api/models/pump"
	mockPumprepoitf "plant-system-api/api/repositories/pump/mock"
	mockConfig "plant-system-api/config/mock"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewPumpService(t *testing.T) {
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
			got := NewPumpService(tt.args.firebaseClient)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func Test_pumpService_ActivePump(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		pumpID = 1
	)

	type fields struct {
		pumpRepository         *mockPumprepoitf.MockPumpRepository
		pumpRepositoryBehavior func(m *mockPumprepoitf.MockPumpRepository)
	}
	type args struct {
		pump *pumpmodel.Pump
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				pumpRepository: mockPumprepoitf.NewMockPumpRepository(ctrl),
				pumpRepositoryBehavior: func(m *mockPumprepoitf.MockPumpRepository) {
					m.EXPECT().ActivePump(gomock.Any()).Return(nil)
				},
			},
			args: args{
				pump: &pumpmodel.Pump{
					ID:       pumpID,
					IsActive: true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.fields.pumpRepositoryBehavior(tt.fields.pumpRepository)

		t.Run(tt.name, func(t *testing.T) {
			s := &pumpService{
				pumpRepository: tt.fields.pumpRepository,
			}
			if err := s.ActivePump(tt.args.pump); (err != nil) != tt.wantErr {
				t.Errorf("pumpService.ActivePump() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pumpService_AskPump(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		pumpID = 1
	)

	var (
		fooErr = errors.New("foo")
	)

	type fields struct {
		pumpRepository         *mockPumprepoitf.MockPumpRepository
		pumpRepositoryBehavior func(m *mockPumprepoitf.MockPumpRepository)
	}
	type args struct {
		pump *pumpmodel.Pump
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				pumpRepository: mockPumprepoitf.NewMockPumpRepository(ctrl),
				pumpRepositoryBehavior: func(m *mockPumprepoitf.MockPumpRepository) {
					m.EXPECT().IsPumpWorking(gomock.Any()).Return(nil)
					m.EXPECT().AskPump(gomock.Any()).Return(nil)
				},
			},
			args: args{
				pump: &pumpmodel.Pump{
					ID:       pumpID,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "error getting is pump working",
			fields: fields{
				pumpRepository: mockPumprepoitf.NewMockPumpRepository(ctrl),
				pumpRepositoryBehavior: func(m *mockPumprepoitf.MockPumpRepository) {
					m.EXPECT().IsPumpWorking(gomock.Any()).Return(fooErr)
				},
			},
			args: args{
				pump: &pumpmodel.Pump{
					ID:       pumpID,
					IsActive: true,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt.fields.pumpRepositoryBehavior(tt.fields.pumpRepository)

		t.Run(tt.name, func(t *testing.T) {
			s := &pumpService{
				pumpRepository: tt.fields.pumpRepository,
			}
			if err := s.AskPump(tt.args.pump); (err != nil) != tt.wantErr {
				t.Errorf("pumpService.AskPump() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pumpService_GetPump(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		pumpID = 1
	)

	var (
		expected = &pumpmodel.Pump{ID: pumpID}
	)

	type fields struct {
		pumpRepository         *mockPumprepoitf.MockPumpRepository
		pumpRepositoryBehavior func(m *mockPumprepoitf.MockPumpRepository)
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pumpmodel.Pump
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				pumpRepository: mockPumprepoitf.NewMockPumpRepository(ctrl),
				pumpRepositoryBehavior: func(m *mockPumprepoitf.MockPumpRepository) {
					m.EXPECT().GetPump(gomock.Any()).Return(expected, nil)
				},
			},
			args: args{
				id: pumpID,
			},
			want:    expected,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.fields.pumpRepositoryBehavior(tt.fields.pumpRepository)

		t.Run(tt.name, func(t *testing.T) {
			s := &pumpService{
				pumpRepository: tt.fields.pumpRepository,
			}
			got, err := s.GetPump(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("pumpService.GetPump() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pumpService.GetPump() = %v, want %v", got, tt.want)
			}
		})
	}
}
