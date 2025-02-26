package integration_test

import (
	"app/internal/config"
	"app/internal/service"
	"app/internal/storage"
	httpClient "app/internal/transport/http"
	"context"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"testing"
)

type TestSuite struct {
	suite.Suite
	server  *httptest.Server
	handler *httpClient.Handler
	router  *mux.Router
	db      *sqlx.DB
	storage *storage.Storage
	service *service.Service
}

func (s *TestSuite) SetupSuite() {
	cfg, err := config.New[config.Config]()
	s.Require().NoError(err, "Failed to load config")

	cfg.Storage.DSN = "postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable"

	log := logrus.New()

	dbPool, err := sqlx.Open("postgres", cfg.Storage.DSN)
	s.Require().NoError(err, "Failed to connect to database")

	s.db = dbPool

	s.storage = storage.New(log, dbPool)
	s.service = service.New(context.Background(), log, s.storage)
	s.handler = httpClient.New(context.Background(), log, s.service)

	s.router = mux.NewRouter()
	httpClient.RegisterRoutes(s.router, s.handler)
	s.server = httptest.NewServer(s.router)

}

func (s *TestSuite) TearDownSuite() {
	s.db.Close()
	s.server.Close()
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
