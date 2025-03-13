package repository

type PgSqlRepository struct{}

func NewPgSqlRepository() *PgSqlRepository {
	return &PgSqlRepository{}
}
