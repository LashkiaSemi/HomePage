package helper

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestTime_Now(t *testing.T) {
	// t.Parallel()

	t.Run("success", func(t *testing.T) {
		mockTime := time.Date(2020, 2, 20, 1, 2, 3, 123456000, time.UTC)
		patch := monkey.Patch(time.Now, func() time.Time { return mockTime })
		defer patch.Unpatch()
		actual := Now()
		assert.Equal(t, mockTime, actual)
	})
}

func TestTime_FormattedDateTimeNow(t *testing.T) {
	// t.Parallel()

	t.Run("success", func(t *testing.T) {
		mockTime := time.Date(2020, 2, 20, 1, 2, 3, 123456000, time.UTC)
		patch := monkey.Patch(time.Now, func() time.Time { return mockTime })
		defer patch.Unpatch()
		expect := "2020/02/20 01:02:03"
		actual := FormattedDateTimeNow()
		assert.Equal(t, expect, actual)
	})
}
