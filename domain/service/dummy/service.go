package dummy

import (
	"github.com/chaosystema/go-microservice/domain/model"
)

type DummyService struct{}

func (d *DummyService) FindById(id int) (*model.Dummy, error) {
	return &model.Dummy{Info: "info"}, nil
}
