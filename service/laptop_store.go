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
	Find(id string) (*pb.Laptop, error)
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

func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	// 읽기 잠금 획득
	store.mutex.RLock()
	// 읽기 잠금 해제
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("can not laptop data : %w", err)
	}

	return other, nil
}
