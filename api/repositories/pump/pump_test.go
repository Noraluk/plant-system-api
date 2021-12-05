package pumprepo

import (
	"context"
	"errors"
	pumpmodel "plant-system-api/api/models/pump"
	mockConfig "plant-system-api/config/mock"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewPumpRepository(t *testing.T) {
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
		},
	}
	for _, tt := range tests {
		tt.args.firebaseClientBehavior(tt.args.firebaseClient)

		t.Run(tt.name, func(t *testing.T) {
			got := NewPumpRepository(tt.args.firebaseClient)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func Test_pumpRepository_ActivePump(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		pumpID = 1
	)

	var (
		mockRef = mockConfig.NewMockRef(ctrl)
		fooErr  = errors.New("foo")
		ctx     = context.Background()
	)

	type fields struct {
		ctx         context.Context
		ref         *mockConfig.MockRef
		refBehavior func(m *mockConfig.MockRef)
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
				ctx: ctx,
				ref: mockRef,
				refBehavior: func(m *mockConfig.MockRef) {
					m.EXPECT().Child(gomock.Any()).Return(mockRef)
					m.EXPECT().Set(gomock.Any(), gomock.Any()).Return(nil)
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
			name: "error setting pump value to firebase db",
			fields: fields{
				ctx: ctx,
				ref: mockRef,
				refBehavior: func(m *mockConfig.MockRef) {
					m.EXPECT().Child(gomock.Any()).Return(mockRef)
					m.EXPECT().Set(gomock.Any(), gomock.Any()).Return(fooErr)
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
		tt.fields.refBehavior(tt.fields.ref)

		t.Run(tt.name, func(t *testing.T) {
			repo := &pumpRepository{
				ctx: tt.fields.ctx,
				ref: tt.fields.ref,
			}
			if err := repo.ActivePump(tt.args.pump); (err != nil) != tt.wantErr {
				t.Errorf("pumpRepository.ActivePump() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pumpRepository_AskPump(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		pumpID = 1
	)

	var (
		mockRef = mockConfig.NewMockRef(ctrl)
		fooErr  = errors.New("foo")
		ctx     = context.Background()
	)

	type fields struct {
		ctx         context.Context
		ref         *mockConfig.MockRef
		refBehavior func(m *mockConfig.MockRef)
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
				ctx: ctx,
				ref: mockRef,
				refBehavior: func(m *mockConfig.MockRef) {
					m.EXPECT().Child(gomock.Any()).Return(mockRef)
					m.EXPECT().Set(gomock.Any(), gomock.Any()).Return(nil)
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
			name: "error setting pump is_ask value to firebase db",
			fields: fields{
				ctx: ctx,
				ref: mockRef,
				refBehavior: func(m *mockConfig.MockRef) {
					m.EXPECT().Child(gomock.Any()).Return(mockRef)
					m.EXPECT().Set(gomock.Any(), gomock.Any()).Return(fooErr)
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
		tt.fields.refBehavior(tt.fields.ref)

		t.Run(tt.name, func(t *testing.T) {
			repo := &pumpRepository{
				ctx: tt.fields.ctx,
				ref: tt.fields.ref,
			}
			if err := repo.AskPump(tt.args.pump); (err != nil) != tt.wantErr {
				t.Errorf("pumpRepository.AskPump() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pumpRepository_IsPumpWorking(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		pumpID = 1
	)

	var (
		mockRef = mockConfig.NewMockRef(ctrl)
		fooErr  = errors.New("foo")
		ctx     = context.Background()
	)

	type fields struct {
		ctx         context.Context
		ref         *mockConfig.MockRef
		refBehavior func(m *mockConfig.MockRef)
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
				ctx: ctx,
				ref: mockRef,
				refBehavior: func(m *mockConfig.MockRef) {
					m.EXPECT().Child(gomock.Any()).Return(mockRef)
					m.EXPECT().Set(gomock.Any(), gomock.Any()).Return(nil)
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
			name: "error setting pump is_ask value to firebase db",
			fields: fields{
				ctx: ctx,
				ref: mockRef,
				refBehavior: func(m *mockConfig.MockRef) {
					m.EXPECT().Child(gomock.Any()).Return(mockRef)
					m.EXPECT().Set(gomock.Any(), gomock.Any()).Return(fooErr)
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
		tt.fields.refBehavior(tt.fields.ref)

		t.Run(tt.name, func(t *testing.T) {
			repo := &pumpRepository{
				ctx: tt.fields.ctx,
				ref: tt.fields.ref,
			}
			if err := repo.IsPumpWorking(tt.args.pump); (err != nil) != tt.wantErr {
				t.Errorf("pumpRepository.IsPumpWorking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pumpRepository_GetPump(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		pumpID = 1
	)

	var (
		mockRef  = mockConfig.NewMockRef(ctrl)
		fooErr   = errors.New("foo")
		ctx      = context.Background()
		expected = &pumpmodel.Pump{ID: pumpID}
	)

	type fields struct {
		ctx         context.Context
		ref         *mockConfig.MockRef
		refBehavior func(m *mockConfig.MockRef)
	}
	type args struct {
		pump *pumpmodel.Pump
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
				ctx: ctx,
				ref: mockRef,
				refBehavior: func(m *mockConfig.MockRef) {
					m.EXPECT().Child(gomock.Any()).Return(mockRef)
					m.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil)
				},
			},
			args: args{
				pump: &pumpmodel.Pump{
					ID: pumpID,
				},
			},
			want:    expected,
			wantErr: false,
		},
		{
			name: "error getting pump firebase db",
			fields: fields{
				ctx: ctx,
				ref: mockRef,
				refBehavior: func(m *mockConfig.MockRef) {
					m.EXPECT().Child(gomock.Any()).Return(mockRef)
					m.EXPECT().Get(gomock.Any(), gomock.Any()).Return(fooErr)
				},
			},
			args: args{
				pump: &pumpmodel.Pump{
					ID:       pumpID,
					IsActive: true,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt.fields.refBehavior(tt.fields.ref)

		t.Run(tt.name, func(t *testing.T) {
			repo := &pumpRepository{
				ctx: tt.fields.ctx,
				ref: tt.fields.ref,
			}
			got, err := repo.GetPump(tt.args.pump)
			if (err != nil) != tt.wantErr {
				t.Errorf("pumpRepository.GetPump() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pumpRepository.GetPump() = %v, want %v", got, tt.want)
			}
		})
	}
}
