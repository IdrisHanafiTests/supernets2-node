// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	pool "github.com/0xPolygon/cdk-validium-node/pool"

	time "time"

	types "github.com/ethereum/go-ethereum/core/types"
)

// PoolMock is an autogenerated mock type for the PoolInterface type
type PoolMock struct {
	mock.Mock
}

// AddTx provides a mock function with given fields: ctx, tx, ip
func (_m *PoolMock) AddTx(ctx context.Context, tx types.Transaction, ip string) error {
	ret := _m.Called(ctx, tx, ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.Transaction, string) error); ok {
		r0 = rf(ctx, tx, ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckPolicy provides a mock function with given fields: ctx, policy, address
func (_m *PoolMock) CheckPolicy(ctx context.Context, policy pool.PolicyName, address common.Address) (bool, error) {
	ret := _m.Called(ctx, policy, address)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pool.PolicyName, common.Address) (bool, error)); ok {
		return rf(ctx, policy, address)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pool.PolicyName, common.Address) bool); ok {
		r0 = rf(ctx, policy, address)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pool.PolicyName, common.Address) error); ok {
		r1 = rf(ctx, policy, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountPendingTransactions provides a mock function with given fields: ctx
func (_m *PoolMock) CountPendingTransactions(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGasPrices provides a mock function with given fields: ctx
func (_m *PoolMock) GetGasPrices(ctx context.Context) (pool.GasPrices, error) {
	ret := _m.Called(ctx)

	var r0 pool.GasPrices
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pool.GasPrices, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pool.GasPrices); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(pool.GasPrices)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNonce provides a mock function with given fields: ctx, address
func (_m *PoolMock) GetNonce(ctx context.Context, address common.Address) (uint64, error) {
	ret := _m.Called(ctx, address)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) (uint64, error)); ok {
		return rf(ctx, address)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) uint64); ok {
		r0 = rf(ctx, address)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPendingTxHashesSince provides a mock function with given fields: ctx, since
func (_m *PoolMock) GetPendingTxHashesSince(ctx context.Context, since time.Time) ([]common.Hash, error) {
	ret := _m.Called(ctx, since)

	var r0 []common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) ([]common.Hash, error)); ok {
		return rf(ctx, since)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) []common.Hash); ok {
		r0 = rf(ctx, since)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time) error); ok {
		r1 = rf(ctx, since)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPendingTxs provides a mock function with given fields: ctx, limit
func (_m *PoolMock) GetPendingTxs(ctx context.Context, limit uint64) ([]pool.Transaction, error) {
	ret := _m.Called(ctx, limit)

	var r0 []pool.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) ([]pool.Transaction, error)); ok {
		return rf(ctx, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []pool.Transaction); ok {
		r0 = rf(ctx, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pool.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxByHash provides a mock function with given fields: ctx, hash
func (_m *PoolMock) GetTxByHash(ctx context.Context, hash common.Hash) (*pool.Transaction, error) {
	ret := _m.Called(ctx, hash)

	var r0 *pool.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*pool.Transaction, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *pool.Transaction); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pool.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPoolMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewPoolMock creates a new instance of PoolMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPoolMock(t mockConstructorTestingTNewPoolMock) *PoolMock {
	mock := &PoolMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
