// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	auth "gorepair-rest-api/internal/utils/auth"

	jwt "github.com/golang-jwt/jwt"

	mock "github.com/stretchr/testify/mock"
)

// JwtTokenInterface is an autogenerated mock type for the JwtTokenInterface type
type JwtTokenInterface struct {
	mock.Mock
}

// Sign provides a mock function with given fields: claims
func (_m *JwtTokenInterface) Sign(claims jwt.MapClaims) auth.TokenStruct {
	ret := _m.Called(claims)

	var r0 auth.TokenStruct
	if rf, ok := ret.Get(0).(func(jwt.MapClaims) auth.TokenStruct); ok {
		r0 = rf(claims)
	} else {
		r0 = ret.Get(0).(auth.TokenStruct)
	}

	return r0
}
