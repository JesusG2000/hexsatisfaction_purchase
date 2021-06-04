// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mock

import (
	context "context"
	time "time"

	model "github.com/JesusG2000/hexsatisfaction_purchase/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// Purchase is an autogenerated mock type for the Purchase type
type Purchase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, purchase
func (_m *Purchase) Create(ctx context.Context, purchase model.PurchaseDTO) (string, error) {
	ret := _m.Called(ctx, purchase)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, model.PurchaseDTO) string); ok {
		r0 = rf(ctx, purchase)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.PurchaseDTO) error); ok {
		r1 = rf(ctx, purchase)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Purchase) Delete(ctx context.Context, id string) (string, error) {
	ret := _m.Called(ctx, id)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByFileID provides a mock function with given fields: ctx, id
func (_m *Purchase) DeleteByFileID(ctx context.Context, id string) (string, error) {
	ret := _m.Called(ctx, id)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAfterDate provides a mock function with given fields: ctx, start
func (_m *Purchase) FindAfterDate(ctx context.Context, start time.Time) ([]model.PurchaseDTO, error) {
	ret := _m.Called(ctx, start)

	var r0 []model.PurchaseDTO
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) []model.PurchaseDTO); ok {
		r0 = rf(ctx, start)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PurchaseDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Time) error); ok {
		r1 = rf(ctx, start)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: ctx
func (_m *Purchase) FindAll(ctx context.Context) ([]model.PurchaseDTO, error) {
	ret := _m.Called(ctx)

	var r0 []model.PurchaseDTO
	if rf, ok := ret.Get(0).(func(context.Context) []model.PurchaseDTO); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PurchaseDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllByUserID provides a mock function with given fields: ctx, id
func (_m *Purchase) FindAllByUserID(ctx context.Context, id string) ([]model.PurchaseDTO, error) {
	ret := _m.Called(ctx, id)

	var r0 []model.PurchaseDTO
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.PurchaseDTO); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PurchaseDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBeforeDate provides a mock function with given fields: ctx, end
func (_m *Purchase) FindBeforeDate(ctx context.Context, end time.Time) ([]model.PurchaseDTO, error) {
	ret := _m.Called(ctx, end)

	var r0 []model.PurchaseDTO
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) []model.PurchaseDTO); ok {
		r0 = rf(ctx, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PurchaseDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Time) error); ok {
		r1 = rf(ctx, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByFileID provides a mock function with given fields: ctx, id
func (_m *Purchase) FindByFileID(ctx context.Context, id string) ([]model.PurchaseDTO, error) {
	ret := _m.Called(ctx, id)

	var r0 []model.PurchaseDTO
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.PurchaseDTO); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PurchaseDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, id
func (_m *Purchase) FindByID(ctx context.Context, id string) (*model.PurchaseDTO, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.PurchaseDTO
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.PurchaseDTO); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PurchaseDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByPeriod provides a mock function with given fields: ctx, start, end
func (_m *Purchase) FindByPeriod(ctx context.Context, start time.Time, end time.Time) ([]model.PurchaseDTO, error) {
	ret := _m.Called(ctx, start, end)

	var r0 []model.PurchaseDTO
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time) []model.PurchaseDTO); ok {
		r0 = rf(ctx, start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PurchaseDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Time, time.Time) error); ok {
		r1 = rf(ctx, start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserIDAfterDate provides a mock function with given fields: ctx, id, start
func (_m *Purchase) FindByUserIDAfterDate(ctx context.Context, id string, start time.Time) ([]model.PurchaseDTO, error) {
	ret := _m.Called(ctx, id, start)

	var r0 []model.PurchaseDTO
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time) []model.PurchaseDTO); ok {
		r0 = rf(ctx, id, start)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PurchaseDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, time.Time) error); ok {
		r1 = rf(ctx, id, start)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserIDAndFileID provides a mock function with given fields: ctx, userID, fileID
func (_m *Purchase) FindByUserIDAndFileID(ctx context.Context, userID string, fileID string) ([]model.PurchaseDTO, error) {
	ret := _m.Called(ctx, userID, fileID)

	var r0 []model.PurchaseDTO
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []model.PurchaseDTO); ok {
		r0 = rf(ctx, userID, fileID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PurchaseDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userID, fileID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserIDAndPeriod provides a mock function with given fields: ctx, id, start, end
func (_m *Purchase) FindByUserIDAndPeriod(ctx context.Context, id string, start time.Time, end time.Time) ([]model.PurchaseDTO, error) {
	ret := _m.Called(ctx, id, start, end)

	var r0 []model.PurchaseDTO
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time, time.Time) []model.PurchaseDTO); ok {
		r0 = rf(ctx, id, start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PurchaseDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, time.Time, time.Time) error); ok {
		r1 = rf(ctx, id, start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserIDBeforeDate provides a mock function with given fields: ctx, id, end
func (_m *Purchase) FindByUserIDBeforeDate(ctx context.Context, id string, end time.Time) ([]model.PurchaseDTO, error) {
	ret := _m.Called(ctx, id, end)

	var r0 []model.PurchaseDTO
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time) []model.PurchaseDTO); ok {
		r0 = rf(ctx, id, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PurchaseDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, time.Time) error); ok {
		r1 = rf(ctx, id, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLast provides a mock function with given fields: ctx
func (_m *Purchase) FindLast(ctx context.Context) (*model.PurchaseDTO, error) {
	ret := _m.Called(ctx)

	var r0 *model.PurchaseDTO
	if rf, ok := ret.Get(0).(func(context.Context) *model.PurchaseDTO); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PurchaseDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLastByUserID provides a mock function with given fields: ctx, id
func (_m *Purchase) FindLastByUserID(ctx context.Context, id string) (*model.PurchaseDTO, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.PurchaseDTO
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.PurchaseDTO); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PurchaseDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
