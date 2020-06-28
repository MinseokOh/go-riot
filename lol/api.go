package lol

type API interface {
	APIVersion() string
	APIName() string
}
