package repository

import (
	"errors"
	"fulfillment-service/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSaveDeliveryExecutive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockDeliveryExecutiveRepository(ctrl)

	t.Run("Success", func(t *testing.T) {
		repo.EXPECT().
			SaveDeliveryExecutive("New York").
			Return(1, nil)

		location := "New York"
		id, err := repo.SaveDeliveryExecutive(location)
		assert.NoError(t, err)
		assert.NotZero(t, id, "The returned ID should not be zero")
	})

	t.Run("DBConnectionFailure", func(t *testing.T) {
		repo.EXPECT().
			SaveDeliveryExecutive(gomock.Any()).
			Return(0, errors.New("failed to get database connection"))

		location := "New York"
		id, err := repo.SaveDeliveryExecutive(location)
		assert.Error(t, err, "Expected an error due to DB connection failure")
		assert.Equal(t, 0, id, "Expected returned ID to be zero on failure")
		assert.Contains(t, err.Error(), "failed to get database connection")
	})

	t.Run("QueryFailure", func(t *testing.T) {
		repo.EXPECT().
			SaveDeliveryExecutive(gomock.Any()).
			Return(0, errors.New("failed to add Delivery Executive"))

		location := "Los Angeles"
		id, err := repo.SaveDeliveryExecutive(location)
		assert.Error(t, err, "Expected an error due to query execution failure")
		assert.Equal(t, 0, id, "Expected returned ID to be zero on failure")
		assert.Contains(t, err.Error(), "failed to add Delivery Executive")
	})
}

func TestGetDeliveryExecutive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockDeliveryExecutiveRepository(ctrl)

	t.Run("Success", func(t *testing.T) {
		repo.EXPECT().
			GetDeliveryExecutive(int32(1)).
			Return(true, int32(123), nil)

		isAvailable, assignedOrderID, err := repo.GetDeliveryExecutive(1)
		assert.NoError(t, err)
		assert.True(t, isAvailable, "The delivery executive should be available")
		assert.Equal(t, int32(123), assignedOrderID, "The assigned order ID should be 123")
	})

	t.Run("DBConnectionFailure", func(t *testing.T) {
		repo.EXPECT().
			GetDeliveryExecutive(gomock.Any()).
			Return(false, int32(0), errors.New("failed to get database connection"))

		isAvailable, assignedOrderID, err := repo.GetDeliveryExecutive(1)
		assert.Error(t, err, "Expected an error due to DB connection failure")
		assert.False(t, isAvailable, "Expected isAvailable to be false on failure")
		assert.Equal(t, int32(0), assignedOrderID, "Expected assignedOrderID to be zero on failure")
		assert.Contains(t, err.Error(), "failed to get database connection")
	})

	t.Run("QueryFailure", func(t *testing.T) {
		repo.EXPECT().
			GetDeliveryExecutive(gomock.Any()).
			Return(false, int32(0), errors.New("failed to retrieve Delivery Executive"))

		isAvailable, assignedOrderID, err := repo.GetDeliveryExecutive(1)
		assert.Error(t, err, "Expected an error due to query execution failure")
		assert.False(t, isAvailable, "Expected isAvailable to be false on failure")
		assert.Equal(t, int32(0), assignedOrderID, "Expected assignedOrderID to be zero on failure")
		assert.Contains(t, err.Error(), "failed to retrieve Delivery Executive")
	})
}

func TestUpdateDeliveryExecutiveStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockDeliveryExecutiveRepository(ctrl)

	t.Run("Success", func(t *testing.T) {
		repo.EXPECT().
			UpdateDeliveryExecutiveStatus(true, int32(123), int32(1)).
			Return(nil)

		err := repo.UpdateDeliveryExecutiveStatus(true, 123, 1)
		assert.NoError(t, err, "Expected no error on successful update")
	})

	t.Run("DBConnectionFailure", func(t *testing.T) {
		repo.EXPECT().
			UpdateDeliveryExecutiveStatus(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(errors.New("failed to get database connection"))

		err := repo.UpdateDeliveryExecutiveStatus(true, 123, 1)
		assert.Error(t, err, "Expected an error due to DB connection failure")
		assert.Contains(t, err.Error(), "failed to get database connection")
	})

	t.Run("QueryFailure", func(t *testing.T) {
		repo.EXPECT().
			UpdateDeliveryExecutiveStatus(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(errors.New("failed to update Delivery Executive status"))

		err := repo.UpdateDeliveryExecutiveStatus(true, 123, 1)
		assert.Error(t, err, "Expected an error due to query execution failure")
		assert.Contains(t, err.Error(), "failed to update Delivery Executive status")
	})
}
