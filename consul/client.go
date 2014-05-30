package consul

import (
	"github.com/mizzy/consul-catalog"
	"strings"
)

// Client provides a wrapper around the consulkv client
type Client struct {
	client *consulcatalog.Client
}

// NewConsulClient returns a new client to Consul for the given address
func NewConsulClient(addr string) (*Client, error) {
	conf := consulcatalog.DefaultConfig()
	conf.Address = addr
	client, err := consulcatalog.NewClient(conf)
	if err != nil {
		return nil, err
	}
	c := &Client{
		client: client,
	}
	return c, nil
}

// Queries Consul for services
func (c *Client) GetValues(services []string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	for _, service := range services {
		service = strings.Replace(service, "/", "", -1)
		var nodes consulcatalog.Nodes
		_, nodes, err := c.client.GetService(service)
		if err != nil {
			return result, err
		}
		result[service] = nodes
	}
	return result, nil
}
