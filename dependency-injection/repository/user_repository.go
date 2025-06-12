package repository

type IUserRepository interface {
	FindAll() []string
}

type mongoRepository struct {
}

func (r mongoRepository) FindAll() []string {
	return []string{"João", "Maria"}
}

func NewMongoRepository() IUserRepository {
	return mongoRepository{}
}
