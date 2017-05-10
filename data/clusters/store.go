package clusters

import (
	"path"
	"time"

	"github.com/appcelerator/amp/data/accounts"
	"github.com/appcelerator/amp/data/storage"
	"github.com/docker/docker/pkg/stringid"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

const clustersRootKey = "clusters"

// Store implements clusters data.Interface
type Store struct {
	store    storage.Interface
	accounts accounts.Interface
}

// NewStore returns an etcd implementation of clusters.Interface
func NewStore(store storage.Interface, accounts accounts.Interface) *Store {
	return &Store{
		store:    store,
		accounts: accounts,
	}
}

// Clusters

func (s *Store) toAccount(organizationName string) (account *accounts.Account, err error) {
	_, err = s.accounts.GetOrganization(context.Background(), organizationName)
	if err != nil {
		return nil, err
	}
	account = &accounts.Account{Type: accounts.AccountType_ORGANIZATION, Name: organizationName}
	return
}

// CreateCluster creates a new cluster
func (s *Store) CreateCluster(ctx context.Context, name string, organizationName string, provider CloudProvider) (cluster *Cluster, err error) {
	// Check if cluster already exists
	clusterAlreadyExists, err := s.GetClusterByName(ctx, name, organizationName)
	if err != nil {
		return nil, err
	}
	if clusterAlreadyExists != nil {
		return nil, ClusterAlreadyExists
	}

	owner, err := s.toAccount(organizationName)
	if err != nil {
		return nil, err
	}

	// Create the new cluster
	cluster = &Cluster{
		Id:       stringid.GenerateNonCryptoID(),
		Name:     name,
		Owner:    owner,
		Provider: provider,
		CreateDt: time.Now().Unix(),
	}
	if err := cluster.Validate(); err != nil {
		return nil, err
	}
	if err := s.store.Create(ctx, path.Join(clustersRootKey, cluster.Id), cluster, nil, 0); err != nil {
		return nil, err
	}
	return cluster, nil
}

// GetCluster fetches a cluster by id
func (s *Store) GetCluster(ctx context.Context, id string) (*Cluster, error) {
	cluster := &Cluster{}
	if err := s.store.Get(ctx, path.Join(clustersRootKey, id), cluster, true); err != nil {
		return nil, err
	}
	// If there's no "id" in the answer, it means the cluster has not been found, so return nil
	if cluster.GetId() == "" {
		return nil, nil
	}
	return cluster, nil
}

// GetClusterByName fetches a cluster by name
func (s *Store) GetClusterByName(ctx context.Context, name string, organizationName string) (cluster *Cluster, err error) {
	if name, err = CheckName(name); err != nil {
		return nil, err
	}
	clusters, err := s.ListClusters(ctx, organizationName, CloudProvider_ANY)
	if err != nil {
		return nil, err
	}
	for _, cluster := range clusters {
		if cluster.Name == name {
			return cluster, nil
		}
	}
	return nil, nil
}

type clusterFilter struct {
	owner    *accounts.Account
	provider CloudProvider
}

func (c clusterFilter) Filter(val proto.Message) bool {
	var keep bool = true
	msg := val.(*Cluster)
	if c.provider != CloudProvider_ANY {
		keep = keep && (msg.GetProvider() == c.provider)
	}
	keep = keep && (msg.GetOwner() == c.owner)
	return keep
}

// ListClusters lists clusters
func (s *Store) ListClusters(ctx context.Context, organizationName string, provider CloudProvider) ([]*Cluster, error) {
	owner, err := s.toAccount(organizationName)
	if err != nil {
		return nil, err
	}
	// TODO: check if the user is allowed to list these clusters (part of the organization)

	var filter storage.Filter = clusterFilter{provider: provider, owner: owner}
	protos := []proto.Message{}
	if err := s.store.List(ctx, clustersRootKey, filter, &Cluster{}, &protos); err != nil {
		return nil, err
	}
	clusters := []*Cluster{}
	for _, proto := range protos {
		clusters = append(clusters, proto.(*Cluster))
	}
	return clusters, nil
}

// DeleteCluster deletes a cluster by id
func (s *Store) DeleteCluster(ctx context.Context, id string) error {
	cluster, err := s.GetCluster(ctx, id)
	if err != nil {
		return err
	}
	if cluster == nil {
		return ClusterNotFound
	}

	// Check authorization
	if !s.accounts.IsAuthorized(ctx, accounts.GetRequesterAccount(ctx), accounts.DeleteAction, accounts.ClusterRN, cluster.Id) {
		return accounts.NotAuthorized
	}

	// Delete the cluster
	if err := s.store.Delete(ctx, path.Join(clustersRootKey, cluster.Id), false, nil); err != nil {
		return err
	}
	return nil
}

// Reset resets the account store
func (s *Store) Reset(ctx context.Context) {
	s.store.Delete(ctx, clustersRootKey, true, nil)
}
