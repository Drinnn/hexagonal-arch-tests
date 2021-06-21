package application_test

import (
	"testing"

	"github.com/Drinnn/hexagonal-arch-tests/application"
	mock_application "github.com/Drinnn/hexagonal-arch-tests/application/mocks"
	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}