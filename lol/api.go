// Package lol is a bind of API League Of Legends
package lol

// API interface of APIs
type API interface {
	APIVersion() string
	APIName() string
}
