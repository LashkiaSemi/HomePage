package service

import (
	mock_repository "homepage/mock/repository"
	"homepage/pkg/domain/entity"
	"homepage/pkg/helper"
	"testing"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestActivity_GetAll(t *testing.T) {
	// t.Parallel()
	err := errors.New("some error")

	tests := []struct {
		name     string
		injector func(m *mock_repository.MockActivityRepository)
		out      []*entity.Activity
		err      error
	}{
		{
			name: "failed to get",
			injector: func(m *mock_repository.MockActivityRepository) {
				m.EXPECT().FindAll().Return(nil, err)
			},
			out: nil,
			err: err,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockActivityRepository) {
				m.EXPECT().FindAll().Return([]*entity.Activity{
					{
						ID: 1,
					},
					{
						ID: 2,
					},
				}, nil)
			},
			out: []*entity.Activity{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_repository.NewMockActivityRepository(ctrl)
			tt.injector(m)
			srv := NewActivity(m)
			out, err := srv.GetAll()
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestActivity_GetByID(t *testing.T) {
	// t.Parallel()
	err := errors.New("some error")

	tests := []struct {
		name     string
		injector func(m *mock_repository.MockActivityRepository)
		in       int
		out      *entity.Activity
		err      error
	}{
		{
			name: "failed to get",
			injector: func(m *mock_repository.MockActivityRepository) {
				m.EXPECT().FindByID(1).Return(nil, err)
			},
			in:  1,
			out: nil,
			err: err,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockActivityRepository) {
				m.EXPECT().FindByID(1).Return(&entity.Activity{
					ID: 1,
				}, nil)
			},
			in: 1,
			out: &entity.Activity{
				ID: 1,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_repository.NewMockActivityRepository(ctrl)
			tt.injector(m)
			srv := NewActivity(m)
			out, err := srv.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestActivity_GetUpcoming(t *testing.T) {
	// t.Parallel()
	err := errors.New("some error")

	tests := []struct {
		name     string
		injector func(m *mock_repository.MockActivityRepository)
		out      []*entity.Activity
		err      error
	}{
		{
			name: "failed to get",
			injector: func(m *mock_repository.MockActivityRepository) {
				m.EXPECT().FindUpcoming().Return(nil, err)
			},
			out: nil,
			err: err,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockActivityRepository) {
				m.EXPECT().FindUpcoming().Return([]*entity.Activity{
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
			err: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_repository.NewMockActivityRepository(ctrl)
			tt.injector(m)
			srv := NewActivity(m)
			out, err := srv.GetUpcoming()
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestActivity_GetForNotification(t *testing.T) {
	// t.Parallel()
	err := errors.New("some error")

	tests := []struct {
		name     string
		injector func(m *mock_repository.MockActivityRepository)
		out      []*entity.Activity
		err      error
	}{
		{
			name: "failed to get",
			injector: func(m *mock_repository.MockActivityRepository) {
				m.EXPECT().FindByNotify().Return(nil, err)
			},
			out: nil,
			err: err,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockActivityRepository) {
				m.EXPECT().FindByNotify().Return([]*entity.Activity{
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
			err: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_repository.NewMockActivityRepository(ctrl)
			tt.injector(m)
			srv := NewActivity(m)
			out, err := srv.GetForNotification()
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestActivity_Create(t *testing.T) {
	// t.Parallel()
	err := errors.New("some error")

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
		injector func(m *mock_repository.MockActivityRepository)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_repository.MockActivityRepository) {
				m.EXPECT().Create(&entity.Activity{}).Return(0, err)
			},
			in:       in{},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockActivityRepository) {
				m.EXPECT().Create(&entity.Activity{
					Activity:    "act001",
					ShowDate:    "2020/02/02",
					Date:        "2020/02/02",
					Annotation:  "annotation",
					IsImportant: 1,
					IsNotify:    1,
				}).Return(1, nil)
			},
			in: in{
				activity:    "act001",
				showDate:    "2020/02/02",
				date:        "2020/02/02",
				annotation:  "annotation",
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
			m := mock_repository.NewMockActivityRepository(ctrl)
			tt.injector(m)
			patch := monkey.Patch(entity.NewActivity, func(string, string, string, string, int, int) *entity.Activity {
				return &entity.Activity{
					Activity:    tt.in.activity,
					ShowDate:    tt.in.showDate,
					Date:        tt.in.date,
					Annotation:  tt.in.annotation,
					IsImportant: tt.in.isImportant,
					IsNotify:    tt.in.isNotify,
				}
			})
			defer patch.Restore()
			srv := NewActivity(m)
			out, err := srv.Create(tt.in.activity, tt.in.showDate, tt.in.date, tt.in.annotation, tt.in.isImportant, tt.in.isNotify)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestActivity_UpdateByID(t *testing.T) {
	// t.Parallel()
	err := errors.New("some error")

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
		injector func(m *mock_repository.MockActivityRepository)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get origin data",
			injector: func(m *mock_repository.MockActivityRepository) {
				m.EXPECT().FindByID(1).Return(nil, err)
			},
			in: in{
				id: 1,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "failed to update",
			injector: func(m *mock_repository.MockActivityRepository) {
				m.EXPECT().FindByID(1).Return(&entity.Activity{
					ID:          1,
					Activity:    "act001",
					ShowDate:    "2020/02/02 00:00:00",
					Date:        "2020/02/02 00:00:00",
					Annotation:  "annotation",
					IsImportant: 1,
					IsNotify:    1,
				}, nil)
				m.EXPECT().UpdateByID(&entity.Activity{
					ID:          1,
					Activity:    "act001",
					ShowDate:    "2020/02/02 00:00:00",
					Date:        "2020/02/02 00:00:00",
					Annotation:  "annotation",
					IsImportant: 1,
					IsNotify:    1,
					UpdatedAt:   "2020/02/02 00:00:00",
				}).Return(err)
			},
			in: in{
				id:          1,
				activity:    "act001",
				showDate:    "2020/02/02 00:00:00",
				date:        "2020/02/02 00:00:00",
				annotation:  "annotation",
				isImportant: 1,
				isNotify:    1,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockActivityRepository) {
				m.EXPECT().FindByID(1).Return(&entity.Activity{
					ID:          1,
					Activity:    "act001",
					ShowDate:    "2020/02/02 00:00:00",
					Date:        "2020/02/02 00:00:00",
					Annotation:  "annotation",
					IsImportant: 1,
					IsNotify:    1,
				}, nil)
				m.EXPECT().UpdateByID(&entity.Activity{
					ID:          1,
					Activity:    "act001",
					ShowDate:    "2020/02/02 00:00:00",
					Date:        "2020/02/02 00:00:00",
					Annotation:  "annotation",
					IsImportant: 1,
					IsNotify:    1,
					UpdatedAt:   "2020/02/02 00:00:00",
				}).Return(nil)
			},
			in: in{
				id:          1,
				activity:    "act001",
				showDate:    "2020/02/02 00:00:00",
				date:        "2020/02/02 00:00:00",
				annotation:  "annotation",
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
			m := mock_repository.NewMockActivityRepository(ctrl)
			tt.injector(m)
			patch := monkey.Patch(helper.FormattedDateTimeNow, func() string { return "2020/02/02 00:00:00" })
			defer patch.Restore()
			srv := NewActivity(m)
			err := srv.UpdateByID(tt.in.id, tt.in.activity, tt.in.showDate, tt.in.date, tt.in.annotation, tt.in.isImportant, tt.in.isNotify)
			tt.checkErr(t, err)
		})
	}
}

func TestActivity_DeleteByID(t *testing.T) {
	// t.Parallel()
	err := errors.New("some error")

	tests := []struct {
		name     string
		injector func(m *mock_repository.MockActivityRepository)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_repository.MockActivityRepository) {
				m.EXPECT().DeleteByID(1).Return(err)
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_repository.MockActivityRepository) {
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
			m := mock_repository.NewMockActivityRepository(ctrl)
			tt.injector(m)
			srv := NewActivity(m)
			e := srv.DeleteByID(tt.in)
			tt.checkErr(t, e)
		})
	}
}
