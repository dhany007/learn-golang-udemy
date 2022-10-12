package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

type SimpleService struct {
	*SimpleRepository
}

// provider selalu diawalai dengan New

// tidak ada parameter karena struct simplerepository diatas tidak ada parameter
func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
	}
}

// ada parameter karena struct simpleservice diatas ada parameter
// paramtere nya harus sama dengan yang di simpleservice
func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("Failed create service")
	} else {
		return &SimpleService{
			SimpleRepository: repository,
		}, nil
	}
}
