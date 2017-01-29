package worker

import (
	"context"

	oelastic "gopkg.in/olivere/elastic.v3"
)

//Option defines an elastic job store option
type Option func(*workerStore)

//SetContext set the context of this store
func SetContext(ctx context.Context) Option {
	return func(workerStore *workerStore) {
		workerStore.ctx = ctx
	}
}

//SetClient sets an elastic client for this store
func SetClient(client *oelastic.Client) Option {
	return func(workerStore *workerStore) {
		workerStore.client = client
	}
}

//SetIndex sets an elastic index for this store
func SetIndex(index string) Option {
	return func(workerStore *workerStore) {
		workerStore.index = index
	}
}

//SetDocType sets an elastic docType for this store
func SetDocType(docType string) Option {
	return func(workerStore *workerStore) {
		workerStore.docType = docType
	}
}

//SetRefreshIndex sets if the index should be immediatly refreshed
func SetRefreshIndex(refresh string) Option {
	return func(workerStore *workerStore) {
		workerStore.refreshIndex = refresh
	}
}
