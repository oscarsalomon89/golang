package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teamcubation/neocamp-meli/testing/domain"
	"github.com/teamcubation/neocamp-meli/testing/mocks"
	"github.com/teamcubation/neocamp-meli/testing/usecase"
)

//SaveItem(name string, stock int) error
//GetItemByID(itemID uint) error

type itemRepositoryMock struct {
	err error
}

func (repo itemRepositoryMock) SaveItem(name string, stock int) error {
	return repo.err
}

func (repo itemRepositoryMock) GetItemByID(itemID uint) error {
	return repo.err
}

// func Test_itemUsecase_CreateItem_Manual(t *testing.T) {
// 	type args struct {
// 		name  string
// 		stock int
// 	}

// 	tests := []struct {
// 		name      string
// 		repo      repository.ItemRepository
// 		args      args
// 		wantedErr error
// 	}{
// 		{
// 			name: "Should works correctly",
// 			repo: itemRepositoryMock{
// 				err: nil,
// 			},
// 			args: args{
// 				name:  "tablet",
// 				stock: 12,
// 			},
// 			wantedErr: nil,
// 		},
// 		{
// 			name: "Should return an error when the name is empty",
// 			repo: itemRepositoryMock{
// 				err: nil,
// 			},
// 			args: args{
// 				name:  "",
// 				stock: 12,
// 			},
// 			wantedErr: fmt.Errorf("item name could not be empty"),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			svc := usecase.NewItemUsecase(tt.repo)

// 			err := svc.CreateItem(tt.args.name, tt.args.stock)
// 			assert.Equal(t, err, tt.wantedErr, "itemUsecase.CreateItem() error = %v, wantErr %v", err, tt.wantedErr)
// 		})
// 	}
// }

func Test_itemService_CreateItem(t *testing.T) {
	assert := assert.New(t)

	type input struct {
		name  string
		stock int
	}

	type dependencies struct {
		repo *mocks.ItemRepository
	}

	tests := []struct {
		name         string
		args         input
		wantedErr    error
		dependencies func(in input, d *dependencies)
	}{
		{
			name:      "Should work correctly",
			wantedErr: nil,
			args: input{
				name:  "tablet",
				stock: 10,
			},
			dependencies: func(in input, d *dependencies) {
				d.repo.On("SaveItem", domain.Item{
					Name:  in.name,
					Stock: in.stock,
				}).
					Return(nil).Once()
			},
		},
		{
			name:      "Should return error when item name is empty",
			wantedErr: fmt.Errorf("item name could not be empty"),
			args: input{
				name:  "",
				stock: 10,
			},
			dependencies: func(in input, d *dependencies) {},
		},
		{
			name:      "Should return error when item stock is zero",
			wantedErr: fmt.Errorf("stock could not be zero"),
			args: input{
				name:  "tablet",
				stock: 0,
			},
			dependencies: func(in input, d *dependencies) {},
		},
		{
			name:      "Should return error when repository returns an error",
			wantedErr: fmt.Errorf("error in repository: %w", errors.New("the repository error")),
			args: input{
				name:  "tablet",
				stock: 10,
			},
			dependencies: func(in input, d *dependencies) {
				d.repo.On("SaveItem", in.name, in.stock).
					Return(errors.New("the repository error")).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dependencies{
				repo: mocks.NewItemRepository(t),
			}

			tt.dependencies(tt.args, d)

			svc := usecase.NewItemUsecase(d.repo)

			err := svc.CreateItem(tt.args.name, tt.args.stock)
			if tt.wantedErr != nil {
				assert.NotNil(err)
				assert.Equal(tt.wantedErr.Error(), err.Error(), "Error message is not the expected")
				return
			}

			assert.Nil(err)
		})
	}
}
