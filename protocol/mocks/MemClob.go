// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	clobtypes "github.com/dydxprotocol/v4-chain/protocol/x/clob/types"
	mock "github.com/stretchr/testify/mock"

	perpetualstypes "github.com/dydxprotocol/v4-chain/protocol/x/perpetuals/types"

	subaccountstypes "github.com/dydxprotocol/v4-chain/protocol/x/subaccounts/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// MemClob is an autogenerated mock type for the MemClob type
type MemClob struct {
	mock.Mock
}

// CancelOrder provides a mock function with given fields: ctx, msgCancelOrder
func (_m *MemClob) CancelOrder(ctx types.Context, msgCancelOrder *clobtypes.MsgCancelOrder) (*clobtypes.OffchainUpdates, error) {
	ret := _m.Called(ctx, msgCancelOrder)

	var r0 *clobtypes.OffchainUpdates
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgCancelOrder) (*clobtypes.OffchainUpdates, error)); ok {
		return rf(ctx, msgCancelOrder)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgCancelOrder) *clobtypes.OffchainUpdates); ok {
		r0 = rf(ctx, msgCancelOrder)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clobtypes.OffchainUpdates)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, *clobtypes.MsgCancelOrder) error); ok {
		r1 = rf(ctx, msgCancelOrder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountSubaccountShortTermOrders provides a mock function with given fields: ctx, subaccountId
func (_m *MemClob) CountSubaccountShortTermOrders(ctx types.Context, subaccountId subaccountstypes.SubaccountId) uint32 {
	ret := _m.Called(ctx, subaccountId)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId) uint32); ok {
		r0 = rf(ctx, subaccountId)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// CreateOrderbook provides a mock function with given fields: ctx, clobPair
func (_m *MemClob) CreateOrderbook(ctx types.Context, clobPair clobtypes.ClobPair) {
	_m.Called(ctx, clobPair)
}

// DeleverageSubaccount provides a mock function with given fields: ctx, subaccountId, perpetualId, deltaQuantums, isFinalSettlement
func (_m *MemClob) DeleverageSubaccount(ctx types.Context, subaccountId subaccountstypes.SubaccountId, perpetualId uint32, deltaQuantums *big.Int, isFinalSettlement bool) (*big.Int, error) {
	ret := _m.Called(ctx, subaccountId, perpetualId, deltaQuantums, isFinalSettlement)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32, *big.Int, bool) (*big.Int, error)); ok {
		return rf(ctx, subaccountId, perpetualId, deltaQuantums, isFinalSettlement)
	}
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32, *big.Int, bool) *big.Int); ok {
		r0 = rf(ctx, subaccountId, perpetualId, deltaQuantums, isFinalSettlement)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, subaccountstypes.SubaccountId, uint32, *big.Int, bool) error); ok {
		r1 = rf(ctx, subaccountId, perpetualId, deltaQuantums, isFinalSettlement)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCancelOrder provides a mock function with given fields: ctx, orderId
func (_m *MemClob) GetCancelOrder(ctx types.Context, orderId clobtypes.OrderId) (uint32, bool) {
	ret := _m.Called(ctx, orderId)

	var r0 uint32
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.OrderId) (uint32, bool)); ok {
		return rf(ctx, orderId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.OrderId) uint32); ok {
		r0 = rf(ctx, orderId)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.OrderId) bool); ok {
		r1 = rf(ctx, orderId)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetMidPrice provides a mock function with given fields: ctx, clobPairId
func (_m *MemClob) GetMidPrice(ctx types.Context, clobPairId clobtypes.ClobPairId) (clobtypes.Subticks, clobtypes.Order, clobtypes.Order, bool) {
	ret := _m.Called(ctx, clobPairId)

	var r0 clobtypes.Subticks
	var r1 clobtypes.Order
	var r2 clobtypes.Order
	var r3 bool
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.ClobPairId) (clobtypes.Subticks, clobtypes.Order, clobtypes.Order, bool)); ok {
		return rf(ctx, clobPairId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.ClobPairId) clobtypes.Subticks); ok {
		r0 = rf(ctx, clobPairId)
	} else {
		r0 = ret.Get(0).(clobtypes.Subticks)
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.ClobPairId) clobtypes.Order); ok {
		r1 = rf(ctx, clobPairId)
	} else {
		r1 = ret.Get(1).(clobtypes.Order)
	}

	if rf, ok := ret.Get(2).(func(types.Context, clobtypes.ClobPairId) clobtypes.Order); ok {
		r2 = rf(ctx, clobPairId)
	} else {
		r2 = ret.Get(2).(clobtypes.Order)
	}

	if rf, ok := ret.Get(3).(func(types.Context, clobtypes.ClobPairId) bool); ok {
		r3 = rf(ctx, clobPairId)
	} else {
		r3 = ret.Get(3).(bool)
	}

	return r0, r1, r2, r3
}

// GetOperationsRaw provides a mock function with given fields: ctx
func (_m *MemClob) GetOperationsRaw(ctx types.Context) []clobtypes.OperationRaw {
	ret := _m.Called(ctx)

	var r0 []clobtypes.OperationRaw
	if rf, ok := ret.Get(0).(func(types.Context) []clobtypes.OperationRaw); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]clobtypes.OperationRaw)
		}
	}

	return r0
}

// GetOperationsToReplay provides a mock function with given fields: ctx
func (_m *MemClob) GetOperationsToReplay(ctx types.Context) ([]clobtypes.InternalOperation, map[clobtypes.OrderHash][]byte) {
	ret := _m.Called(ctx)

	var r0 []clobtypes.InternalOperation
	var r1 map[clobtypes.OrderHash][]byte
	if rf, ok := ret.Get(0).(func(types.Context) ([]clobtypes.InternalOperation, map[clobtypes.OrderHash][]byte)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(types.Context) []clobtypes.InternalOperation); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]clobtypes.InternalOperation)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context) map[clobtypes.OrderHash][]byte); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[clobtypes.OrderHash][]byte)
		}
	}

	return r0, r1
}

// GetOrder provides a mock function with given fields: ctx, orderId
func (_m *MemClob) GetOrder(ctx types.Context, orderId clobtypes.OrderId) (clobtypes.Order, bool) {
	ret := _m.Called(ctx, orderId)

	var r0 clobtypes.Order
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.OrderId) (clobtypes.Order, bool)); ok {
		return rf(ctx, orderId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.OrderId) clobtypes.Order); ok {
		r0 = rf(ctx, orderId)
	} else {
		r0 = ret.Get(0).(clobtypes.Order)
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.OrderId) bool); ok {
		r1 = rf(ctx, orderId)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetOrderFilledAmount provides a mock function with given fields: ctx, orderId
func (_m *MemClob) GetOrderFilledAmount(ctx types.Context, orderId clobtypes.OrderId) subaccountstypes.BaseQuantums {
	ret := _m.Called(ctx, orderId)

	var r0 subaccountstypes.BaseQuantums
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.OrderId) subaccountstypes.BaseQuantums); ok {
		r0 = rf(ctx, orderId)
	} else {
		r0 = ret.Get(0).(subaccountstypes.BaseQuantums)
	}

	return r0
}

// GetOrderRemainingAmount provides a mock function with given fields: ctx, order
func (_m *MemClob) GetOrderRemainingAmount(ctx types.Context, order clobtypes.Order) (subaccountstypes.BaseQuantums, bool) {
	ret := _m.Called(ctx, order)

	var r0 subaccountstypes.BaseQuantums
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.Order) (subaccountstypes.BaseQuantums, bool)); ok {
		return rf(ctx, order)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.Order) subaccountstypes.BaseQuantums); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Get(0).(subaccountstypes.BaseQuantums)
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.Order) bool); ok {
		r1 = rf(ctx, order)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetPricePremium provides a mock function with given fields: ctx, clobPair, params
func (_m *MemClob) GetPricePremium(ctx types.Context, clobPair clobtypes.ClobPair, params perpetualstypes.GetPricePremiumParams) (int32, error) {
	ret := _m.Called(ctx, clobPair, params)

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.ClobPair, perpetualstypes.GetPricePremiumParams) (int32, error)); ok {
		return rf(ctx, clobPair, params)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.ClobPair, perpetualstypes.GetPricePremiumParams) int32); ok {
		r0 = rf(ctx, clobPair, params)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.ClobPair, perpetualstypes.GetPricePremiumParams) error); ok {
		r1 = rf(ctx, clobPair, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubaccountOrders provides a mock function with given fields: ctx, clobPairId, subaccountId, side
func (_m *MemClob) GetSubaccountOrders(ctx types.Context, clobPairId clobtypes.ClobPairId, subaccountId subaccountstypes.SubaccountId, side clobtypes.Order_Side) ([]clobtypes.Order, error) {
	ret := _m.Called(ctx, clobPairId, subaccountId, side)

	var r0 []clobtypes.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.ClobPairId, subaccountstypes.SubaccountId, clobtypes.Order_Side) ([]clobtypes.Order, error)); ok {
		return rf(ctx, clobPairId, subaccountId, side)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.ClobPairId, subaccountstypes.SubaccountId, clobtypes.Order_Side) []clobtypes.Order); ok {
		r0 = rf(ctx, clobPairId, subaccountId, side)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]clobtypes.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.ClobPairId, subaccountstypes.SubaccountId, clobtypes.Order_Side) error); ok {
		r1 = rf(ctx, clobPairId, subaccountId, side)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertZeroFillDeleveragingIntoOperationsQueue provides a mock function with given fields: ctx, subaccountId, perpetualId
func (_m *MemClob) InsertZeroFillDeleveragingIntoOperationsQueue(ctx types.Context, subaccountId subaccountstypes.SubaccountId, perpetualId uint32) {
	_m.Called(ctx, subaccountId, perpetualId)
}

// PlaceOrder provides a mock function with given fields: ctx, order
func (_m *MemClob) PlaceOrder(ctx types.Context, order clobtypes.Order) (subaccountstypes.BaseQuantums, clobtypes.OrderStatus, *clobtypes.OffchainUpdates, error) {
	ret := _m.Called(ctx, order)

	var r0 subaccountstypes.BaseQuantums
	var r1 clobtypes.OrderStatus
	var r2 *clobtypes.OffchainUpdates
	var r3 error
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.Order) (subaccountstypes.BaseQuantums, clobtypes.OrderStatus, *clobtypes.OffchainUpdates, error)); ok {
		return rf(ctx, order)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.Order) subaccountstypes.BaseQuantums); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Get(0).(subaccountstypes.BaseQuantums)
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.Order) clobtypes.OrderStatus); ok {
		r1 = rf(ctx, order)
	} else {
		r1 = ret.Get(1).(clobtypes.OrderStatus)
	}

	if rf, ok := ret.Get(2).(func(types.Context, clobtypes.Order) *clobtypes.OffchainUpdates); ok {
		r2 = rf(ctx, order)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*clobtypes.OffchainUpdates)
		}
	}

	if rf, ok := ret.Get(3).(func(types.Context, clobtypes.Order) error); ok {
		r3 = rf(ctx, order)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// PlacePerpetualLiquidation provides a mock function with given fields: ctx, liquidationOrder
func (_m *MemClob) PlacePerpetualLiquidation(ctx types.Context, liquidationOrder clobtypes.LiquidationOrder) (subaccountstypes.BaseQuantums, clobtypes.OrderStatus, *clobtypes.OffchainUpdates, error) {
	ret := _m.Called(ctx, liquidationOrder)

	var r0 subaccountstypes.BaseQuantums
	var r1 clobtypes.OrderStatus
	var r2 *clobtypes.OffchainUpdates
	var r3 error
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.LiquidationOrder) (subaccountstypes.BaseQuantums, clobtypes.OrderStatus, *clobtypes.OffchainUpdates, error)); ok {
		return rf(ctx, liquidationOrder)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.LiquidationOrder) subaccountstypes.BaseQuantums); ok {
		r0 = rf(ctx, liquidationOrder)
	} else {
		r0 = ret.Get(0).(subaccountstypes.BaseQuantums)
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.LiquidationOrder) clobtypes.OrderStatus); ok {
		r1 = rf(ctx, liquidationOrder)
	} else {
		r1 = ret.Get(1).(clobtypes.OrderStatus)
	}

	if rf, ok := ret.Get(2).(func(types.Context, clobtypes.LiquidationOrder) *clobtypes.OffchainUpdates); ok {
		r2 = rf(ctx, liquidationOrder)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*clobtypes.OffchainUpdates)
		}
	}

	if rf, ok := ret.Get(3).(func(types.Context, clobtypes.LiquidationOrder) error); ok {
		r3 = rf(ctx, liquidationOrder)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// PurgeInvalidMemclobState provides a mock function with given fields: ctx, fullyFilledOrderIds, expiredStatefulOrderIds, canceledStatefulOrderIds, removedStatefulOrderIds, existingOffchainUpdates
func (_m *MemClob) PurgeInvalidMemclobState(ctx types.Context, fullyFilledOrderIds []clobtypes.OrderId, expiredStatefulOrderIds []clobtypes.OrderId, canceledStatefulOrderIds []clobtypes.OrderId, removedStatefulOrderIds []clobtypes.OrderId, existingOffchainUpdates *clobtypes.OffchainUpdates) *clobtypes.OffchainUpdates {
	ret := _m.Called(ctx, fullyFilledOrderIds, expiredStatefulOrderIds, canceledStatefulOrderIds, removedStatefulOrderIds, existingOffchainUpdates)

	var r0 *clobtypes.OffchainUpdates
	if rf, ok := ret.Get(0).(func(types.Context, []clobtypes.OrderId, []clobtypes.OrderId, []clobtypes.OrderId, []clobtypes.OrderId, *clobtypes.OffchainUpdates) *clobtypes.OffchainUpdates); ok {
		r0 = rf(ctx, fullyFilledOrderIds, expiredStatefulOrderIds, canceledStatefulOrderIds, removedStatefulOrderIds, existingOffchainUpdates)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clobtypes.OffchainUpdates)
		}
	}

	return r0
}

// RemoveAndClearOperationsQueue provides a mock function with given fields: ctx, localValidatorOperationsQueue
func (_m *MemClob) RemoveAndClearOperationsQueue(ctx types.Context, localValidatorOperationsQueue []clobtypes.InternalOperation) {
	_m.Called(ctx, localValidatorOperationsQueue)
}

// RemoveOrderIfFilled provides a mock function with given fields: ctx, orderId
func (_m *MemClob) RemoveOrderIfFilled(ctx types.Context, orderId clobtypes.OrderId) {
	_m.Called(ctx, orderId)
}

// ReplayOperations provides a mock function with given fields: ctx, localOperations, shortTermOrderTxBytes, existingOffchainUpdates
func (_m *MemClob) ReplayOperations(ctx types.Context, localOperations []clobtypes.InternalOperation, shortTermOrderTxBytes map[clobtypes.OrderHash][]byte, existingOffchainUpdates *clobtypes.OffchainUpdates) *clobtypes.OffchainUpdates {
	ret := _m.Called(ctx, localOperations, shortTermOrderTxBytes, existingOffchainUpdates)

	var r0 *clobtypes.OffchainUpdates
	if rf, ok := ret.Get(0).(func(types.Context, []clobtypes.InternalOperation, map[clobtypes.OrderHash][]byte, *clobtypes.OffchainUpdates) *clobtypes.OffchainUpdates); ok {
		r0 = rf(ctx, localOperations, shortTermOrderTxBytes, existingOffchainUpdates)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clobtypes.OffchainUpdates)
		}
	}

	return r0
}

// SetClobKeeper provides a mock function with given fields: keeper
func (_m *MemClob) SetClobKeeper(keeper clobtypes.MemClobKeeper) {
	_m.Called(keeper)
}

// SetMemclobGauges provides a mock function with given fields: ctx
func (_m *MemClob) SetMemclobGauges(ctx types.Context) {
	_m.Called(ctx)
}

type mockConstructorTestingTNewMemClob interface {
	mock.TestingT
	Cleanup(func())
}

// NewMemClob creates a new instance of MemClob. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMemClob(t mockConstructorTestingTNewMemClob) *MemClob {
	mock := &MemClob{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
