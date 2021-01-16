// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	presenter "github.com/ricardoham/pokedex-api/api/presenter"
)

// Pokemon is an autogenerated mock type for the Pokemon type
type Pokemon struct {
	mock.Mock
}

// GetAllResultPokemonFromPokeApi provides a mock function with given fields:
func (_m *Pokemon) GetAllResultPokemonFromPokeApi() (*presenter.Result, error) {
	ret := _m.Called()

	var r0 *presenter.Result
	if rf, ok := ret.Get(0).(func() *presenter.Result); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*presenter.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPokemonFromPokeApi provides a mock function with given fields: _a0
func (_m *Pokemon) GetPokemonFromPokeApi(_a0 string) (*presenter.Pokemon, error) {
	ret := _m.Called(_a0)

	var r0 *presenter.Pokemon
	if rf, ok := ret.Get(0).(func(string) *presenter.Pokemon); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*presenter.Pokemon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}