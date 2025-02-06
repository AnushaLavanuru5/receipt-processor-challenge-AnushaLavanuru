package storage

import "receipt-processor/internal/models"

type MemoryStore struct {
	data map[string]models.Receipt
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]models.Receipt),
	}
}

func (s *MemoryStore) Save(id string, receipt models.Receipt) {
	s.data[id] = receipt
}

func (s *MemoryStore) Get(id string) (models.Receipt, bool) {
	receipt, exists := s.data[id]
	return receipt, exists
}
