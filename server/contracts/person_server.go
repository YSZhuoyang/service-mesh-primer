package msglib

import (
	"context"
	"fmt"
)

type PersonRegistryService struct {
	UnimplementedPersonRegistryServer
	people []*Person
}

func (prs *PersonRegistryService) Lookup(ctx context.Context, person *Person) (*Person, error) {
	for _, p := range prs.people {
		if p.Name == person.Name {
			return p, nil
		}
	}

	return &Person{}, nil
}

func (prs *PersonRegistryService) Create(ctx context.Context, person *Person) (*Person, error) {
	for _, p := range prs.people {
		if p.Name == person.Name {
			return p, fmt.Errorf("Person already exists")
		}
	}

	prs.people = append(prs.people, person)
	return person, nil
}
