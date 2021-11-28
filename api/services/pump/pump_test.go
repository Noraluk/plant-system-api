package pumpservice

import (
	pumpmodel "plant-system-api/api/models/pump"
	mockPumprepoitf "plant-system-api/api/repositories/pump/mock"
	mockConfig "plant-system-api/config/mock"
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
					PumpID:   pumpID,
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
