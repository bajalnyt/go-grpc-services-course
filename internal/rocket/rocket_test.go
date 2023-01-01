package rocket

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRocketService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("tests get rocket by id", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			GetRocketById(id).
			Return(Rocket{
				ID: id,
			}, nil)

		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.
			GetRocketById(
				context.Background(),
				id,
			)

		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", rkt.ID)
	})

	t.Run("tests insert rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			InsertRocket(Rocket{
				ID: id,
			}).
			Return(Rocket{
				ID: id,
			}, nil)

		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.
			InsertRocket(
				context.Background(),
				Rocket{
					ID: id,
				},
			)

		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", rkt.ID)
	})

	t.Run("tests delete rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			DeleteRocketById(id).
			Return(nil)

		rocketService := New(rocketStoreMock)
		err := rocketService.
			DeleteRocketById(
				context.Background(),
				id,
			)

		assert.NoError(t, err)
	})
}
