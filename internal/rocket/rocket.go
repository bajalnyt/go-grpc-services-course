//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/bajalnyt/go-grpc-services-course/internal/rocket Store

package rocket

import "context"

type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

type Store interface {
	GetRocketById(id string) (Rocket, error)
	InsertRocket(rocket Rocket) (Rocket, error)
	DeleteRocketById(id string) error
}

// Service - will update rocket inventory
type Service struct {
	Store Store
}

func New(store Store) Service {
	return Service{
		Store: store,
	}
}

func (s Service) GetRocketById(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketById(id)
	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

func (s Service) InsertRocket(ctx context.Context, rocket Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rocket)
	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

func (s Service) DeleteRocketById(ctx context.Context, id string) error {
	err := s.Store.DeleteRocketById(id)
	if err != nil {
		return err
	}

	return nil
}
