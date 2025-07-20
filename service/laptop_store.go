package service

import (
	"errors"
	"fmt"
	"grpc/pb"
	"sync"

	"github.com/jinzhu/copier"
)

// go get -u github.com/jinzhu/copier
var ErrAllreadyExists = errors.New("recorf already exists")

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
}

type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAllreadyExists
	}

	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)

	if err != nil {
		return fmt.Errorf("can not copy laptop data %w", err)
	}
	store.data[other.Id] = other
	return nil
}
