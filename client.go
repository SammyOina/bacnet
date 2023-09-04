package bacnet

type client struct {
	// Fields for client configuration
}

type Client interface {
	ReadProperty(objectID ObjectIdentifier, propertyID PropertyIdentifier) (interface{}, error)
	WriteProperty(objectID ObjectIdentifier, propertyID PropertyIdentifier, value interface{}) error
}

func NewClient(address string, port int) (Client, error) {
	return &client{}, nil
}

func (c *client) ReadProperty(objectID ObjectIdentifier, propertyID PropertyIdentifier) (interface{}, error) {
	return nil, nil
}

func (c *client) WriteProperty(objectID ObjectIdentifier, propertyID PropertyIdentifier, value interface{}) error {
	return nil
}