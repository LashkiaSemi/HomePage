package interactor

import (
	"homepage/pkg/domain/entity"
)

// TODO: ここのリポジトリ定義は使わない

// UserRepository ユーザのDBアクセスを実現
type UserRepository interface {
	FindAll() ([]*entity.User, error)
	FindByID(userID int) (*entity.User, error)
	FindByStudentID(studentID string) (*entity.User, error)
	FindAuthInfoByStudentID(studentID string) (*entity.User, error)

	UpdateByID(*entity.User) error
	UpdatePasswordByStudentID(studentID, password string) error

	AdminCreate(*entity.User) (int, error)
	AdminUpdateByID(*entity.User) error

	DeleteByID(userID int) error
}

// ActivityRepository 活動内容のDBアクセスを実現
// type ActivityRepository interface {
// 	FindAll() ([]*entity.Activity, error)
// 	FindByID(id int) (*entity.Activity, error)

// 	// FindUpcoming お知らせ欄に乗る。
// 	FindUpcoming() ([]*entity.Activity, error)
// 	// FindByNotify is_notifyカラムがonのデータを取得
// 	FindByNotify() ([]*entity.Activity, error)

// 	Create(*entity.Activity) (int, error)
// 	UpdateByID(*entity.Activity) error

// 	DeleteByID(id int) error
// }

// SocietyRepository 学会発表のDBアクセスを実現
type SocietyRepository interface {
	FindAll() ([]*entity.Society, error)
	FindByID(id int) (*entity.Society, error)

	Create(*entity.Society) (int, error)
	UpdateByID(*entity.Society) error

	DeleteByID(id int) error
}

// JobRepository 就職先のDBアクセスを実現
type JobRepository interface {
	FindAll() ([]*entity.Job, error)
	FindByID(id int) (*entity.Job, error)

	Create(*entity.Job) (int, error)
	UpdateByID(*entity.Job) error

	DeleteByID(id int) error
}

// LectureRepository レクチャーのDBアクセスを実現
type LectureRepository interface {
	FindAll() ([]*entity.Lecture, error)
	FindByID(id int) (*entity.Lecture, error)
	FindAuthorByStudentID(studentID string) (*entity.User, error)

	Create(*entity.Lecture) (int, error)
	UpdateByID(*entity.Lecture) error
	DeleteByID(id int) error
}

// ResearchRepository 卒業研究関連のDBアクセスを実現
type ResearchRepository interface {
	FindAll() ([]*entity.Research, error)
	FindByID(id int) (*entity.Research, error)

	Create(*entity.Research) (int, error)
	UpdateByID(*entity.Research) error

	DeleteByID(id int) error
}

// EquipmentRepository 備品関連のDBアクセスを実現
type EquipmentRepository interface {
	FindAll() ([]*entity.Equipment, error)
	FindByID(id int) (*entity.Equipment, error)

	Create(*entity.Equipment) (int, error)
	UpdateByID(*entity.Equipment) error

	DeleteByID(id int) error
}

// TagRepository タグ関連のDBアクセスを実現
type TagRepository interface {
	FindAll() ([]*entity.Tag, error)
	FindByID(id int) (*entity.Tag, error)
	Create(data *entity.Tag) (int, error)
	UpdateByID(data *entity.Tag) error
	DeleteByID(id int) error
}
