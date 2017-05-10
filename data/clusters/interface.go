package clusters

import "golang.org/x/net/context"

// Error type
type Error string

func (e Error) Error() string { return string(e) }

// Errors
const (
	InvalidName          = Error("name is invalid")
	InvalidProvider      = Error("provider is invalid")
	ClusterAlreadyExists = Error("cluster already exists")
	ClusterNotFound      = Error("cluster not found")
)

// Interface defines the cluster data access layer
type Interface interface {
	// CreateCluster creates a new cluster
	CreateCluster(ctx context.Context, name string, organizationName string) (cluster *Cluster, err error)

	// GetCluster fetches a cluster by id
	GetCluster(ctx context.Context, id string) (cluster *Cluster, err error)

	// GetClusterByName fetches a cluster by name
	GetClusterByName(ctx context.Context, name string, organizationName string) (cluster *Cluster, err error)

	// ListClusters lists clusters
	ListClusters(ctx context.Context, organizationName string) (clusters []*Cluster, err error)

	// DeleteCluster deletes a cluster by id
	DeleteCluster(ctx context.Context, id string) (err error)

	// Reset resets the cluster store
	Reset(ctx context.Context)
}
