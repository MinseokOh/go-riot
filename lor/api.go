package tft

// API interface of APIs
type API interface {
	APIVersion() string
	APIName() string
}
