//go:generate mockgen -source=$GOFILE -destination=../../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package repository

import "homepage/pkg/domain/entity"

type ActivityRepository interface {
	FindAll() ([]*entity.Activity, error)
	FindByID(id int) (*entity.Activity, error)

	// FindUpcoming お知らせ欄に乗る。
	FindUpcoming() ([]*entity.Activity, error)
	// FindByNotify is_notifyカラムがonのデータを取得
	FindByNotify() ([]*entity.Activity, error)

	Create(*entity.Activity) (int, error)
	UpdateByID(*entity.Activity) error

	DeleteByID(id int) error
}
