package customer

// IService is the interface to access our customer service
type IService interface {
	Add(Customer *Customer) map[string]interface{}
	Get() map[string]interface{}
}
