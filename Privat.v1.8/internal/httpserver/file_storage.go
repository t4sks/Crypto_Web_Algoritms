package httpserver

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type processedFileJob struct {
	ID           string
	OriginName   string
	SafeFileName string
	Algorithm    string
	Operation    string
	ResultText   string
	Code         string
	CreateTime   time.Time
}

type processedFileStore struct {
	mu   sync.RWMutex
	jobs map[string]processedFileJob
}

func newProcessedFileStore() *processedFileStore {
	return &processedFileStore{
		jobs: make(map[string]processedFileJob),
	}
}

func (s *processedFileStore) Save(job processedFileJob) processedFileJob {
	s.mu.Lock()
	defer s.mu.Unlock()

	if job.ID == "" {
		job.ID = uuid.NewString()
	}
	if job.CreateTime.IsZero() {
		job.CreateTime = time.Now()
	}

	if len(s.jobs) > 20 {
		clear(s.jobs)
	}

	s.jobs[job.ID] = job
	return job
}

func (s *processedFileStore) Get(id string) (processedFileJob, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	job, ok := s.jobs[id]
	return job, ok
}

var fileStore = newProcessedFileStore()
