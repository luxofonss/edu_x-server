package database

type DBConnection interface{}

type Database interface {
	Connect(config DBConfig) (DBConnection, error)
}

// DBConfig represents the configuration for a database.
type DBConfig struct {
	IdentificationName string // IdentificationName is used to obtain the specific database connection.
	DB                 string // Database name.
	User               string // Database user.
	Password           string `json:"_"` // Database password.
	Host               string // Database host.
	Port               string // Database port.
	Type               string // Type of the database ("mysql", "postgres", "mssql", etc.).
	SSLMode            string // SSL model for the database connection.
	TimeZone           string // Time zone for the database.
}

type DatabaseService struct {
	Database Database
}

func (db DatabaseService) Connect(config DBConfig) (DBConnection, error) {
	return db.Database.Connect(config)
}

func NewDatabase(dbType string) DatabaseService {
	var database Database
	switch dbType {
	case "mongo":
		database = &MongoDatabase{}
		break
	case "postgres":
		database = &PostgresDatabase{}
	}
	return DatabaseService{database}
}
