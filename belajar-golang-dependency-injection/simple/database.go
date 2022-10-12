package simple

type Database struct {
	Name string
}

type DatabasePostgreeSQL Database
type DatabaseMongoDB Database

// dipakasa ke databatasepostgreessql
func NewDatabasePostgreeSQL() *DatabasePostgreeSQL {
	return (*DatabasePostgreeSQL)(&Database{
		Name: "PostgreeSQL",
	})
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{
		Name: "MongoDB",
	})
}

type DatabaseRepository struct {
	DatabasePostgreeSQL *DatabasePostgreeSQL
	DatabaseMongoDB     *DatabaseMongoDB
}

func NewDatabaseRepository(postgreeSQL *DatabasePostgreeSQL, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgreeSQL: postgreeSQL,
		DatabaseMongoDB:     mongoDB,
	}
}
