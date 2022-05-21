package controller

type Schema struct {
}

// NewController initialization controller
// go:generate: mockgen
func NewController() *Schema {
	return &Schema{}
}
