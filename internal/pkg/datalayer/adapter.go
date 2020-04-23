package datalayer

type Adapter interface {
	Get(output interface{}) error
	GetJSON() ([]byte, error)
}
