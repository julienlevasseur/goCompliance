package consul

import (
	"context"

	consul "github.com/hashicorp/consul/api"
)

func consulCatalog() (*consul.Catalog, error) {
	consulClient, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		return nil, err
	}
	return consulClient.Catalog(), nil
}

// GetServices implements the Consul Catalog Services call.
func GetServices(ctx context.Context) (map[string][]string, error) {
	catalog, err := consulCatalog()
	if err != nil {
		return make(map[string][]string), err
	}

	services, _, err := catalog.Services(nil)
	if err != nil {
		return make(map[string][]string), err
	}

	return services, nil
}

// GetService implements the Consul Catalog Service call.
func GetService(ctx context.Context, id string) ([]*consul.CatalogService, error) {
	catalog, err := consulCatalog()
	if err != nil {
		return []*consul.CatalogService{}, err
	}

	catalogService, _, err := catalog.Service(id, "", nil)
	if err != nil {
		return []*consul.CatalogService{}, err
	}

	return catalogService, nil
}