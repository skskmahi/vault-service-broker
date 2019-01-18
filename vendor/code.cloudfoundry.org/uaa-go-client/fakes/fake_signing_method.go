// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/dgrijalva/jwt-go"
)

type FakeSigningMethod struct {
	VerifyStub        func(signingString, signature string, key interface{}) error
	verifyMutex       sync.RWMutex
	verifyArgsForCall []struct {
		signingString string
		signature     string
		key           interface{}
	}
	verifyReturns struct {
		result1 error
	}
	SignStub        func(signingString string, key interface{}) (string, error)
	signMutex       sync.RWMutex
	signArgsForCall []struct {
		signingString string
		key           interface{}
	}
	signReturns struct {
		result1 string
		result2 error
	}
	AlgStub        func() string
	algMutex       sync.RWMutex
	algArgsForCall []struct{}
	algReturns     struct {
		result1 string
	}
}

func (fake *FakeSigningMethod) Verify(signingString string, signature string, key interface{}) error {
	fake.verifyMutex.Lock()
	fake.verifyArgsForCall = append(fake.verifyArgsForCall, struct {
		signingString string
		signature     string
		key           interface{}
	}{signingString, signature, key})
	fake.verifyMutex.Unlock()
	if fake.VerifyStub != nil {
		return fake.VerifyStub(signingString, signature, key)
	} else {
		return fake.verifyReturns.result1
	}
}

func (fake *FakeSigningMethod) VerifyCallCount() int {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return len(fake.verifyArgsForCall)
}

func (fake *FakeSigningMethod) VerifyArgsForCall(i int) (string, string, interface{}) {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return fake.verifyArgsForCall[i].signingString, fake.verifyArgsForCall[i].signature, fake.verifyArgsForCall[i].key
}

func (fake *FakeSigningMethod) VerifyReturns(result1 error) {
	fake.VerifyStub = nil
	fake.verifyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSigningMethod) Sign(signingString string, key interface{}) (string, error) {
	fake.signMutex.Lock()
	fake.signArgsForCall = append(fake.signArgsForCall, struct {
		signingString string
		key           interface{}
	}{signingString, key})
	fake.signMutex.Unlock()
	if fake.SignStub != nil {
		return fake.SignStub(signingString, key)
	} else {
		return fake.signReturns.result1, fake.signReturns.result2
	}
}

func (fake *FakeSigningMethod) SignCallCount() int {
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	return len(fake.signArgsForCall)
}

func (fake *FakeSigningMethod) SignArgsForCall(i int) (string, interface{}) {
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	return fake.signArgsForCall[i].signingString, fake.signArgsForCall[i].key
}

func (fake *FakeSigningMethod) SignReturns(result1 string, result2 error) {
	fake.SignStub = nil
	fake.signReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSigningMethod) Alg() string {
	fake.algMutex.Lock()
	fake.algArgsForCall = append(fake.algArgsForCall, struct{}{})
	fake.algMutex.Unlock()
	if fake.AlgStub != nil {
		return fake.AlgStub()
	} else {
		return fake.algReturns.result1
	}
}

func (fake *FakeSigningMethod) AlgCallCount() int {
	fake.algMutex.RLock()
	defer fake.algMutex.RUnlock()
	return len(fake.algArgsForCall)
}

func (fake *FakeSigningMethod) AlgReturns(result1 string) {
	fake.AlgStub = nil
	fake.algReturns = struct {
		result1 string
	}{result1}
}

var _ jwt.SigningMethod = new(FakeSigningMethod)
