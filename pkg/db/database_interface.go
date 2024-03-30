package db

import (
	"database/sql"

	// calculatorpc "github.com/AskarKasimov/grpc-calculator/pkg/proto"
	_ "github.com/lib/pq"
)

type iDatabase interface {
	// GetOneAvailableExpression(workerId int64) (models.Expression, error)
	// AddExpression(e models.ExpressionAdding) (int64, error)
	// AllExpressions() ([]models.ExpressionGeneral, error)
	// IsWorkerAlive(workerId int64) (bool, error)
	// WakeUp(workerId int64) error
	GetWorkerIdByName(name string) (string, error)
	NewWorker(name string) (string, error)
	// AllAliveWorkers() ([]models.Worker, error)
	// FallAsleep(workerId int64) error
	// GetActiveExpressionsFromWorker(workerId int64) ([]models.Expression, error)
	// MakeExpressionAvailableAgain(expressionId int64) error
	// SolveExpression(workerId, expressionId int64, solution string) error
	// GetExpressionById(expressionId int64) (models.Expression, error)
	// AllWorkers() ([]models.Worker, error)
}

type database struct {
	db *sql.DB
}

var db iDatabase

func DB() iDatabase { return db }

func init() {
	connStr := "user=admin password=admin dbname=yg sslmode=disable host=postgres"
	newConn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	db = &database{db: newConn}
}

func (d *database) GetWorkerIdByName(name string) (string, error) {
	var id string

	err := d.db.QueryRow("SELECT id FROM workers WHERE name=$1", name).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (d *database) NewWorker(name string) (string, error) {
	var id string
	err := d.db.QueryRow("INSERT INTO workers (name, isAlive) VALUES ($1, true) RETURNING id", name).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
