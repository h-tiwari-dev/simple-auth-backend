package database

import "sample-auth-backend/app/queries"

type Queries struct {
	*queries.UserQuery
}

func OpenDbConnection() (*Queries, error) {
	db, err := PostgresSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		UserQuery: &queries.UserQuery{DB: db},
	}, nil
}
