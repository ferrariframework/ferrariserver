package worker

import (
	"context"

	"encoding/json"

	"github.com/ferrariframework/ferrariserver/models"
	"github.com/ferrariframework/ferrariserver/store"
	"github.com/ferrariframework/ferrariserver/store/errortypes"
	"github.com/mattheath/kala/snowflake"
	"github.com/pkg/errors"
	oelastic "gopkg.in/olivere/elastic.v3"
)

//WorkerStore implementation of a job store
type workerStore struct {
	client       *oelastic.Client
	index        string
	docType      string
	idGenerator  *snowflake.Snowflake
	refreshIndex string
}

//New returns a new instance of a worker store
func New(options ...Option) store.Worker {

	j := &workerStore{
		refreshIndex: "true",
	}

	for _, option := range options {
		option(j)
	}

	return j
}

func (w *workerStore) Save(worker *models.Worker) (*models.Worker, error) {
	id, err := w.idGenerator.Mint()

	if err != nil {
		return nil, errors.Wrap(err, "elastic.Store Failed to generate id for job ")
	}

	_, err = w.client.Index().
		Index(w.index).
		Type(w.docType).
		Refresh(w.refreshIndex).
		Id(id).
		BodyJson(worker).
		Do(context.Background())

	if err != nil {
		return nil, errors.Wrapf(err, "elastic.Store Failed to index job %v", worker)
	}

	worker.ID = id
	return worker, nil
}

func (w *workerStore) Get(id string) (*models.Worker, error) {
	res, err := w.client.Get().
		Index(w.index).
		Type(w.docType).
		Refresh(w.refreshIndex).
		Id(id).
		Do(context.Background())

	if err != nil {
		if oelastic.IsNotFound(err) {
			return nil, errors.Wrapf(errortypes.NewNotFound(), "elastic.Store.Get(%id)", id)
		}
		return nil, errors.Wrapf(err, "elastic.Store Failed to get jobID=%s", id)
	}

	var job models.Worker

	err = json.Unmarshal(*res.Source, &job)

	if err != nil {
		return nil, errors.Wrapf(err, "elastic.Store.Get(%id)", id)
	}
	job.ID = res.Id
	return &job, nil
}

func (w *workerStore) Update(id string, worker *models.Worker) error {

	_, err := w.client.Update().
		Index(w.index).
		Type(w.docType).
		Refresh(w.refreshIndex).
		Id(id).
		Doc(worker).
		Do(context.Background())

	if err != nil {
		if oelastic.IsNotFound(err) {
			return errors.Wrapf(errortypes.NewNotFound(), "elastic.Store.Update(%id, %v)", worker)
		}
		return errors.Wrapf(err, "elastic.Store.Update(%id, %v) ", id, worker)
	}

	return nil
}