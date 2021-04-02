package controller

import (
	mock_interactor "homepage/mock/interactor"
	"homepage/pkg/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestActivity_GetAllGroupByYear(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockActivityInteractor)
		out      []*ActivitiesGroupByYearResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      []*ActivitiesGroupByYearResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().GetAll().Return([]*entity.Activity{
					{
						ID:   1,
						Date: "2020/03/27",
					},
					{
						ID:   2,
						Date: "2020/03/30",
					},
					{
						ID:   3,
						Date: "2021/02/10",
					},
				}, nil)
			},
			out: []*ActivitiesGroupByYearResponse{
				{
					Year: "2020",
					Activities: []*ActivityResponse{
						{
							ID:   1,
							Date: "2020/03/27",
						},
						{
							ID:   2,
							Date: "2020/03/30",
						},
					},
				},
				{
					Year: "2021",
					Activities: []*ActivityResponse{
						{
							ID:   3,
							Date: "2021/02/10",
						},
					},
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
			m := mock_interactor.NewMockActivityInteractor(ctrl)
			tt.injector(m)
			c := NewActivityController(m)
			out, err := c.GetAllGroupByYear()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestActivity_GetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockActivityInteractor)
		in       int
		out      *ActivityResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      &ActivityResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().GetByID(1).Return(&entity.Activity{
					ID: 1,
				}, nil)
			},
			in: 1,
			out: &ActivityResponse{
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
			m := mock_interactor.NewMockActivityInteractor(ctrl)
			tt.injector(m)
			c := NewActivityController(m)
			out, err := c.GetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestActivity_GetUpcoming(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockActivityInteractor)
		out      *ActivitiesResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().GetUpcoming().Return(nil, errors.New("some error"))
			},
			out:      &ActivitiesResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().GetUpcoming().Return([]*entity.Activity{
					{
						ID: 1,
					},
				}, nil)
			},
			out: &ActivitiesResponse{
				Activities: []*ActivityResponse{
					{
						ID: 1,
					},
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
			m := mock_interactor.NewMockActivityInteractor(ctrl)
			tt.injector(m)
			c := NewActivityController(m)
			out, err := c.GetUpcoming()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestActivity_GetForNotification(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockActivityInteractor)
		out      *ActivitiesResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().GetForNotification().Return(nil, errors.New("some error"))
			},
			out:      &ActivitiesResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().GetForNotification().Return([]*entity.Activity{
					{
						ID: 1,
					},
				}, nil)
			},
			out: &ActivitiesResponse{
				Activities: []*ActivityResponse{
					{
						ID: 1,
					},
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
			m := mock_interactor.NewMockActivityInteractor(ctrl)
			tt.injector(m)
			c := NewActivityController(m)
			out, err := c.GetForNotificaion()
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
		injector func(m *mock_interactor.MockActivityInteractor)
		in       in
		out      int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().Create("", "", "", "", 0, 0).Return(0, errors.New("some error"))
			},
			in:       in{},
			out:      0,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().Create("activity001", "showDate001", "date001", "anno001", 1, 1).Return(1, nil)
			},
			in: in{
				activity:    "activity001",
				showDate:    "showDate001",
				date:        "date001",
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
			m := mock_interactor.NewMockActivityInteractor(ctrl)
			tt.injector(m)
			c := NewActivityController(m)
			out, err := c.Create(tt.in.activity, tt.in.showDate, tt.in.date, tt.in.annotation, tt.in.isImportant, tt.in.isNotify)
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
		injector func(m *mock_interactor.MockActivityInteractor)
		in       in
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to create",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().UpdateByID(1, "", "", "", "", 0, 0).Return(errors.New("some error"))
			},
			in: in{
				id: 1,
			},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().UpdateByID(1, "activity001", "showDate001", "date001", "anno001", 1, 1).Return(nil)
			},
			in: in{
				id:          1,
				activity:    "activity001",
				showDate:    "showDate001",
				date:        "date001",
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
			m := mock_interactor.NewMockActivityInteractor(ctrl)
			tt.injector(m)
			c := NewActivityController(m)
			err := c.UpdateByID(tt.in.id, tt.in.activity, tt.in.showDate, tt.in.date, tt.in.annotation, tt.in.isImportant, tt.in.isNotify)
			tt.checkErr(t, err)
		})
	}
}

func TestActivity_DeleteByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockActivityInteractor)
		in       int
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to delete",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().DeleteByID(1).Return(errors.New("some error"))
			},
			in:       1,
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockActivityInteractor) {
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
			m := mock_interactor.NewMockActivityInteractor(ctrl)
			tt.injector(m)
			c := NewActivityController(m)
			err := c.DeleteByID(tt.in)
			tt.checkErr(t, err)
		})
	}
}

func TestActivity_AdminGetAll(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockActivityInteractor)
		out      []map[string]string
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().GetAll().Return(nil, errors.New("some error"))
			},
			out:      []map[string]string(nil),
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().GetAll().Return([]*entity.Activity{
					{
						ID:       1,
						Activity: "act001",
					},
					{
						ID:       2,
						Activity: "act002",
					},
				}, nil)
			},
			out: []map[string]string{
				{
					"id":    "1",
					"title": "act001",
				},
				{
					"id":    "2",
					"title": "act002",
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
			m := mock_interactor.NewMockActivityInteractor(ctrl)
			tt.injector(m)
			c := NewActivityController(m)
			out, err := c.AdminGetAll()
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}

func TestActivity_AdminGetByID(t *testing.T) {
	tests := []struct {
		name     string
		injector func(m *mock_interactor.MockActivityInteractor)
		in       int
		out      *FieldsResponse
		checkErr func(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool
	}{
		{
			name: "failed to get",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().GetByID(1).Return(nil, errors.New("some error"))
			},
			in:       1,
			out:      &FieldsResponse{},
			checkErr: assert.NotNil,
		},
		{
			name: "success",
			injector: func(m *mock_interactor.MockActivityInteractor) {
				m.EXPECT().GetByID(1).Return(&entity.Activity{
					ID:          1,
					Activity:    "act001",
					Annotation:  "anno001",
					ShowDate:    "show001",
					Date:        "date001",
					IsImportant: 0,
					IsNotify:    0,
				}, nil)
			},
			in: 1,
			out: &FieldsResponse{
				ID: 1,
				Fields: []*Field{
					{Key: "ID", Value: 1},
					{Key: "活動内容", Value: "act001"},
					{Key: "注釈", Value: "anno001"},
					{Key: "日付(表示用)", Value: "show001"},
					{Key: "日付(内部処理用)", Value: "date001"},
					{Key: "重要", Value: 0},
					{Key: "通知する", Value: 0},
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
			m := mock_interactor.NewMockActivityInteractor(ctrl)
			tt.injector(m)
			c := NewActivityController(m)
			out, err := c.AdminGetByID(tt.in)
			assert.Equal(t, tt.out, out)
			tt.checkErr(t, err)
		})
	}
}
