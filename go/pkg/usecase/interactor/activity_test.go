package interactor

import (
	mock_service "homepage/mock/service"
	"homepage/pkg/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestActivity_GetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockActivity)
		out      []*entity.Activity
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get all",
			injector: func(m *mock_service.MockActivity) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockActivity) {
				m.EXPECT().GetAll().Return([]*entity.Activity{
					{
						ID: 1,
					},
				}, nil)
			},
			out: []*entity.Activity{
				{
					ID: 1,
				},
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_service.NewMockActivity(ctrl)
			tt.injector(m)
			usecase := NewActivityInteractor(m)
			out, err := usecase.GetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestActivity_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockActivity)
		in       int
		out      *entity.Activity
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get all",
			injector: func(m *mock_service.MockActivity) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockActivity) {
				m.EXPECT().GetByID(1).Return(&entity.Activity{
					ID: 1,
				}, nil)
			},
			in: 1,
			out: &entity.Activity{
				ID: 1,
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_service.NewMockActivity(ctrl)
			tt.injector(m)
			usecase := NewActivityInteractor(m)
			out, err := usecase.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestActivity_GetUpcoming(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockActivity)
		out      []*entity.Activity
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_service.MockActivity) {
				m.EXPECT().GetUpcoming().Return(nil, errors.New("some error"))
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockActivity) {
				m.EXPECT().GetUpcoming().Return([]*entity.Activity{
					{
						ID: 1,
					},
				}, nil)
			},
			out: []*entity.Activity{
				{
					ID: 1,
				},
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_service.NewMockActivity(ctrl)
			tt.injector(m)
			usecase := NewActivityInteractor(m)
			out, err := usecase.GetUpcoming()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestActivity_GetForNotification(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockActivity)
		out      []*entity.Activity
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_service.MockActivity) {
				m.EXPECT().GetForNotification().Return(nil, errors.New("some error"))
			},
			out:      nil,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockActivity) {
				m.EXPECT().GetForNotification().Return([]*entity.Activity{
					{
						ID: 1,
					},
				}, nil)
			},
			out: []*entity.Activity{
				{
					ID: 1,
				},
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_service.NewMockActivity(ctrl)
			tt.injector(m)
			usecase := NewActivityInteractor(m)
			out, err := usecase.GetForNotification()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestActivity_Create(t *testing.T) {

	type in struct {
		activity    string
		showDate    string
		date        string
		annotation  string
		isImportant int
		isNotify    int
	}
	tests := []struct {
		name     string
		injector func(m *mock_service.MockActivity)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_service.MockActivity) {
				m.EXPECT().Create("act001", "2020/02/02 00:00:00", "2020/02/02 00:00:00", "anno001", 1, 1).Return(0, errors.New("some error"))
			},
			in: in{
				activity:    "act001",
				showDate:    "2020/02/02 00:00:00",
				date:        "2020/02/02 00:00:00",
				annotation:  "anno001",
				isImportant: 1,
				isNotify:    1,
			},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockActivity) {
				m.EXPECT().Create("act001", "2020/02/02 00:00:00", "2020/02/02 00:00:00", "anno001", 1, 1).Return(1, nil)
			},
			in: in{
				activity:    "act001",
				showDate:    "2020/02/02 00:00:00",
				date:        "2020/02/02 00:00:00",
				annotation:  "anno001",
				isImportant: 1,
				isNotify:    1,
			},
			out:      1,
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_service.NewMockActivity(ctrl)
			tt.injector(m)
			usecase := NewActivityInteractor(m)
			out, err := usecase.Create(tt.in.activity, tt.in.showDate, tt.in.date, tt.in.annotation, tt.in.isImportant, tt.in.isNotify)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestActivity_UpdateByID(t *testing.T) {

	type in struct {
		id          int
		activity    string
		showDate    string
		date        string
		annotation  string
		isImportant int
		isNotify    int
	}
	tests := []struct {
		name     string
		injector func(m *mock_service.MockActivity)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to update",
			injector: func(m *mock_service.MockActivity) {
				m.EXPECT().UpdateByID(1, "act001", "2020/02/02 00:00:00", "2020/02/02 00:00:00", "anno001", 1, 1).Return(errors.New("some error"))
			},
			in: in{
				id:          1,
				activity:    "act001",
				showDate:    "2020/02/02 00:00:00",
				date:        "2020/02/02 00:00:00",
				annotation:  "anno001",
				isImportant: 1,
				isNotify:    1,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockActivity) {
				m.EXPECT().UpdateByID(1, "act001", "2020/02/02 00:00:00", "2020/02/02 00:00:00", "anno001", 1, 1).Return(nil)
			},
			in: in{
				id:          1,
				activity:    "act001",
				showDate:    "2020/02/02 00:00:00",
				date:        "2020/02/02 00:00:00",
				annotation:  "anno001",
				isImportant: 1,
				isNotify:    1,
			},
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_service.NewMockActivity(ctrl)
			tt.injector(m)
			usecase := NewActivityInteractor(m)
			err := usecase.UpdateByID(tt.in.id, tt.in.activity, tt.in.showDate, tt.in.date, tt.in.annotation, tt.in.isImportant, tt.in.isNotify)
			tt.checkErr(t, err)
		})
	}
}

func TestActivity_DeleteByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_service.MockActivity)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_service.MockActivity) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_service.MockActivity) {
				m.EXPECT().DeleteByID(1).Return(nil)
			},
			in:       1,
			checkErr: assert.Nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_service.NewMockActivity(ctrl)
			tt.injector(m)
			usecase := NewActivityInteractor(m)
			err := usecase.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}
