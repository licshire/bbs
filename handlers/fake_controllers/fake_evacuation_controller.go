// Code generated by counterfeiter. DO NOT EDIT.
package fake_controllers

import (
	"sync"

	"code.cloudfoundry.org/bbs/handlers"
	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/lager"
)

type FakeEvacuationController struct {
	RemoveEvacuatingActualLRPStub        func(lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) error
	removeEvacuatingActualLRPMutex       sync.RWMutex
	removeEvacuatingActualLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 *models.ActualLRPKey
		arg3 *models.ActualLRPInstanceKey
	}
	removeEvacuatingActualLRPReturns struct {
		result1 error
	}
	removeEvacuatingActualLRPReturnsOnCall map[int]struct {
		result1 error
	}
	EvacuateClaimedActualLRPStub        func(lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) (error, bool)
	evacuateClaimedActualLRPMutex       sync.RWMutex
	evacuateClaimedActualLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 *models.ActualLRPKey
		arg3 *models.ActualLRPInstanceKey
	}
	evacuateClaimedActualLRPReturns struct {
		result1 error
		result2 bool
	}
	evacuateClaimedActualLRPReturnsOnCall map[int]struct {
		result1 error
		result2 bool
	}
	EvacuateCrashedActualLRPStub        func(lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, string) error
	evacuateCrashedActualLRPMutex       sync.RWMutex
	evacuateCrashedActualLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 *models.ActualLRPKey
		arg3 *models.ActualLRPInstanceKey
		arg4 string
	}
	evacuateCrashedActualLRPReturns struct {
		result1 error
	}
	evacuateCrashedActualLRPReturnsOnCall map[int]struct {
		result1 error
	}
	EvacuateRunningActualLRPStub        func(lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, *models.ActualLRPNetInfo) (error, bool)
	evacuateRunningActualLRPMutex       sync.RWMutex
	evacuateRunningActualLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 *models.ActualLRPKey
		arg3 *models.ActualLRPInstanceKey
		arg4 *models.ActualLRPNetInfo
	}
	evacuateRunningActualLRPReturns struct {
		result1 error
		result2 bool
	}
	evacuateRunningActualLRPReturnsOnCall map[int]struct {
		result1 error
		result2 bool
	}
	EvacuateStoppedActualLRPStub        func(lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) error
	evacuateStoppedActualLRPMutex       sync.RWMutex
	evacuateStoppedActualLRPArgsForCall []struct {
		arg1 lager.Logger
		arg2 *models.ActualLRPKey
		arg3 *models.ActualLRPInstanceKey
	}
	evacuateStoppedActualLRPReturns struct {
		result1 error
	}
	evacuateStoppedActualLRPReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEvacuationController) RemoveEvacuatingActualLRP(arg1 lager.Logger, arg2 *models.ActualLRPKey, arg3 *models.ActualLRPInstanceKey) error {
	fake.removeEvacuatingActualLRPMutex.Lock()
	ret, specificReturn := fake.removeEvacuatingActualLRPReturnsOnCall[len(fake.removeEvacuatingActualLRPArgsForCall)]
	fake.removeEvacuatingActualLRPArgsForCall = append(fake.removeEvacuatingActualLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 *models.ActualLRPKey
		arg3 *models.ActualLRPInstanceKey
	}{arg1, arg2, arg3})
	fake.recordInvocation("RemoveEvacuatingActualLRP", []interface{}{arg1, arg2, arg3})
	fake.removeEvacuatingActualLRPMutex.Unlock()
	if fake.RemoveEvacuatingActualLRPStub != nil {
		return fake.RemoveEvacuatingActualLRPStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeEvacuatingActualLRPReturns.result1
}

func (fake *FakeEvacuationController) RemoveEvacuatingActualLRPCallCount() int {
	fake.removeEvacuatingActualLRPMutex.RLock()
	defer fake.removeEvacuatingActualLRPMutex.RUnlock()
	return len(fake.removeEvacuatingActualLRPArgsForCall)
}

func (fake *FakeEvacuationController) RemoveEvacuatingActualLRPArgsForCall(i int) (lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) {
	fake.removeEvacuatingActualLRPMutex.RLock()
	defer fake.removeEvacuatingActualLRPMutex.RUnlock()
	return fake.removeEvacuatingActualLRPArgsForCall[i].arg1, fake.removeEvacuatingActualLRPArgsForCall[i].arg2, fake.removeEvacuatingActualLRPArgsForCall[i].arg3
}

func (fake *FakeEvacuationController) RemoveEvacuatingActualLRPReturns(result1 error) {
	fake.RemoveEvacuatingActualLRPStub = nil
	fake.removeEvacuatingActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEvacuationController) RemoveEvacuatingActualLRPReturnsOnCall(i int, result1 error) {
	fake.RemoveEvacuatingActualLRPStub = nil
	if fake.removeEvacuatingActualLRPReturnsOnCall == nil {
		fake.removeEvacuatingActualLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeEvacuatingActualLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEvacuationController) EvacuateClaimedActualLRP(arg1 lager.Logger, arg2 *models.ActualLRPKey, arg3 *models.ActualLRPInstanceKey) (error, bool) {
	fake.evacuateClaimedActualLRPMutex.Lock()
	ret, specificReturn := fake.evacuateClaimedActualLRPReturnsOnCall[len(fake.evacuateClaimedActualLRPArgsForCall)]
	fake.evacuateClaimedActualLRPArgsForCall = append(fake.evacuateClaimedActualLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 *models.ActualLRPKey
		arg3 *models.ActualLRPInstanceKey
	}{arg1, arg2, arg3})
	fake.recordInvocation("EvacuateClaimedActualLRP", []interface{}{arg1, arg2, arg3})
	fake.evacuateClaimedActualLRPMutex.Unlock()
	if fake.EvacuateClaimedActualLRPStub != nil {
		return fake.EvacuateClaimedActualLRPStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.evacuateClaimedActualLRPReturns.result1, fake.evacuateClaimedActualLRPReturns.result2
}

func (fake *FakeEvacuationController) EvacuateClaimedActualLRPCallCount() int {
	fake.evacuateClaimedActualLRPMutex.RLock()
	defer fake.evacuateClaimedActualLRPMutex.RUnlock()
	return len(fake.evacuateClaimedActualLRPArgsForCall)
}

func (fake *FakeEvacuationController) EvacuateClaimedActualLRPArgsForCall(i int) (lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) {
	fake.evacuateClaimedActualLRPMutex.RLock()
	defer fake.evacuateClaimedActualLRPMutex.RUnlock()
	return fake.evacuateClaimedActualLRPArgsForCall[i].arg1, fake.evacuateClaimedActualLRPArgsForCall[i].arg2, fake.evacuateClaimedActualLRPArgsForCall[i].arg3
}

func (fake *FakeEvacuationController) EvacuateClaimedActualLRPReturns(result1 error, result2 bool) {
	fake.EvacuateClaimedActualLRPStub = nil
	fake.evacuateClaimedActualLRPReturns = struct {
		result1 error
		result2 bool
	}{result1, result2}
}

func (fake *FakeEvacuationController) EvacuateClaimedActualLRPReturnsOnCall(i int, result1 error, result2 bool) {
	fake.EvacuateClaimedActualLRPStub = nil
	if fake.evacuateClaimedActualLRPReturnsOnCall == nil {
		fake.evacuateClaimedActualLRPReturnsOnCall = make(map[int]struct {
			result1 error
			result2 bool
		})
	}
	fake.evacuateClaimedActualLRPReturnsOnCall[i] = struct {
		result1 error
		result2 bool
	}{result1, result2}
}

func (fake *FakeEvacuationController) EvacuateCrashedActualLRP(arg1 lager.Logger, arg2 *models.ActualLRPKey, arg3 *models.ActualLRPInstanceKey, arg4 string) error {
	fake.evacuateCrashedActualLRPMutex.Lock()
	ret, specificReturn := fake.evacuateCrashedActualLRPReturnsOnCall[len(fake.evacuateCrashedActualLRPArgsForCall)]
	fake.evacuateCrashedActualLRPArgsForCall = append(fake.evacuateCrashedActualLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 *models.ActualLRPKey
		arg3 *models.ActualLRPInstanceKey
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("EvacuateCrashedActualLRP", []interface{}{arg1, arg2, arg3, arg4})
	fake.evacuateCrashedActualLRPMutex.Unlock()
	if fake.EvacuateCrashedActualLRPStub != nil {
		return fake.EvacuateCrashedActualLRPStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.evacuateCrashedActualLRPReturns.result1
}

func (fake *FakeEvacuationController) EvacuateCrashedActualLRPCallCount() int {
	fake.evacuateCrashedActualLRPMutex.RLock()
	defer fake.evacuateCrashedActualLRPMutex.RUnlock()
	return len(fake.evacuateCrashedActualLRPArgsForCall)
}

func (fake *FakeEvacuationController) EvacuateCrashedActualLRPArgsForCall(i int) (lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, string) {
	fake.evacuateCrashedActualLRPMutex.RLock()
	defer fake.evacuateCrashedActualLRPMutex.RUnlock()
	return fake.evacuateCrashedActualLRPArgsForCall[i].arg1, fake.evacuateCrashedActualLRPArgsForCall[i].arg2, fake.evacuateCrashedActualLRPArgsForCall[i].arg3, fake.evacuateCrashedActualLRPArgsForCall[i].arg4
}

func (fake *FakeEvacuationController) EvacuateCrashedActualLRPReturns(result1 error) {
	fake.EvacuateCrashedActualLRPStub = nil
	fake.evacuateCrashedActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEvacuationController) EvacuateCrashedActualLRPReturnsOnCall(i int, result1 error) {
	fake.EvacuateCrashedActualLRPStub = nil
	if fake.evacuateCrashedActualLRPReturnsOnCall == nil {
		fake.evacuateCrashedActualLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.evacuateCrashedActualLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEvacuationController) EvacuateRunningActualLRP(arg1 lager.Logger, arg2 *models.ActualLRPKey, arg3 *models.ActualLRPInstanceKey, arg4 *models.ActualLRPNetInfo) (error, bool) {
	fake.evacuateRunningActualLRPMutex.Lock()
	ret, specificReturn := fake.evacuateRunningActualLRPReturnsOnCall[len(fake.evacuateRunningActualLRPArgsForCall)]
	fake.evacuateRunningActualLRPArgsForCall = append(fake.evacuateRunningActualLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 *models.ActualLRPKey
		arg3 *models.ActualLRPInstanceKey
		arg4 *models.ActualLRPNetInfo
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("EvacuateRunningActualLRP", []interface{}{arg1, arg2, arg3, arg4})
	fake.evacuateRunningActualLRPMutex.Unlock()
	if fake.EvacuateRunningActualLRPStub != nil {
		return fake.EvacuateRunningActualLRPStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.evacuateRunningActualLRPReturns.result1, fake.evacuateRunningActualLRPReturns.result2
}

func (fake *FakeEvacuationController) EvacuateRunningActualLRPCallCount() int {
	fake.evacuateRunningActualLRPMutex.RLock()
	defer fake.evacuateRunningActualLRPMutex.RUnlock()
	return len(fake.evacuateRunningActualLRPArgsForCall)
}

func (fake *FakeEvacuationController) EvacuateRunningActualLRPArgsForCall(i int) (lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey, *models.ActualLRPNetInfo) {
	fake.evacuateRunningActualLRPMutex.RLock()
	defer fake.evacuateRunningActualLRPMutex.RUnlock()
	return fake.evacuateRunningActualLRPArgsForCall[i].arg1, fake.evacuateRunningActualLRPArgsForCall[i].arg2, fake.evacuateRunningActualLRPArgsForCall[i].arg3, fake.evacuateRunningActualLRPArgsForCall[i].arg4
}

func (fake *FakeEvacuationController) EvacuateRunningActualLRPReturns(result1 error, result2 bool) {
	fake.EvacuateRunningActualLRPStub = nil
	fake.evacuateRunningActualLRPReturns = struct {
		result1 error
		result2 bool
	}{result1, result2}
}

func (fake *FakeEvacuationController) EvacuateRunningActualLRPReturnsOnCall(i int, result1 error, result2 bool) {
	fake.EvacuateRunningActualLRPStub = nil
	if fake.evacuateRunningActualLRPReturnsOnCall == nil {
		fake.evacuateRunningActualLRPReturnsOnCall = make(map[int]struct {
			result1 error
			result2 bool
		})
	}
	fake.evacuateRunningActualLRPReturnsOnCall[i] = struct {
		result1 error
		result2 bool
	}{result1, result2}
}

func (fake *FakeEvacuationController) EvacuateStoppedActualLRP(arg1 lager.Logger, arg2 *models.ActualLRPKey, arg3 *models.ActualLRPInstanceKey) error {
	fake.evacuateStoppedActualLRPMutex.Lock()
	ret, specificReturn := fake.evacuateStoppedActualLRPReturnsOnCall[len(fake.evacuateStoppedActualLRPArgsForCall)]
	fake.evacuateStoppedActualLRPArgsForCall = append(fake.evacuateStoppedActualLRPArgsForCall, struct {
		arg1 lager.Logger
		arg2 *models.ActualLRPKey
		arg3 *models.ActualLRPInstanceKey
	}{arg1, arg2, arg3})
	fake.recordInvocation("EvacuateStoppedActualLRP", []interface{}{arg1, arg2, arg3})
	fake.evacuateStoppedActualLRPMutex.Unlock()
	if fake.EvacuateStoppedActualLRPStub != nil {
		return fake.EvacuateStoppedActualLRPStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.evacuateStoppedActualLRPReturns.result1
}

func (fake *FakeEvacuationController) EvacuateStoppedActualLRPCallCount() int {
	fake.evacuateStoppedActualLRPMutex.RLock()
	defer fake.evacuateStoppedActualLRPMutex.RUnlock()
	return len(fake.evacuateStoppedActualLRPArgsForCall)
}

func (fake *FakeEvacuationController) EvacuateStoppedActualLRPArgsForCall(i int) (lager.Logger, *models.ActualLRPKey, *models.ActualLRPInstanceKey) {
	fake.evacuateStoppedActualLRPMutex.RLock()
	defer fake.evacuateStoppedActualLRPMutex.RUnlock()
	return fake.evacuateStoppedActualLRPArgsForCall[i].arg1, fake.evacuateStoppedActualLRPArgsForCall[i].arg2, fake.evacuateStoppedActualLRPArgsForCall[i].arg3
}

func (fake *FakeEvacuationController) EvacuateStoppedActualLRPReturns(result1 error) {
	fake.EvacuateStoppedActualLRPStub = nil
	fake.evacuateStoppedActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEvacuationController) EvacuateStoppedActualLRPReturnsOnCall(i int, result1 error) {
	fake.EvacuateStoppedActualLRPStub = nil
	if fake.evacuateStoppedActualLRPReturnsOnCall == nil {
		fake.evacuateStoppedActualLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.evacuateStoppedActualLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEvacuationController) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.removeEvacuatingActualLRPMutex.RLock()
	defer fake.removeEvacuatingActualLRPMutex.RUnlock()
	fake.evacuateClaimedActualLRPMutex.RLock()
	defer fake.evacuateClaimedActualLRPMutex.RUnlock()
	fake.evacuateCrashedActualLRPMutex.RLock()
	defer fake.evacuateCrashedActualLRPMutex.RUnlock()
	fake.evacuateRunningActualLRPMutex.RLock()
	defer fake.evacuateRunningActualLRPMutex.RUnlock()
	fake.evacuateStoppedActualLRPMutex.RLock()
	defer fake.evacuateStoppedActualLRPMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEvacuationController) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ handlers.EvacuationController = new(FakeEvacuationController)
