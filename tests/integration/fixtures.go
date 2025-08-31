package integration

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/ikshavaku/catalogue/utils"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	fixtures   *testfixtures.Loader
	testDBConn *sql.DB
	once       sync.Once
)

func openDBConnection(connectionURL string) (*sql.DB, error) {
	// Open connection to the test database.
	// Do NOT import fixtures in a production database!
	// Existing data would be DELETED
	db, err := sql.Open("pgx", connectionURL)
	if err != nil {
		return nil, fmt.Errorf("[TEST] fatal %+v", err)
	}
	return db, nil
}

func LoadTestFixtures(cfg utils.GlobalConfig) {
	log.Println("[TEST] Loading integration test assets...")
	if cfg.Env != "local" && cfg.Env != "test" && cfg.Env != "ci" {
		panic("only run on test")
	}
	connectionURL := cfg.Postgres.PostgresConnectionURL()
	fmt.Printf("Connection URL: %s\n", connectionURL)
	if !strings.Contains(connectionURL, "localhost") || !strings.Contains(connectionURL, "postgres") {
		panic("only run on a test database to avoid deleting all data!")
	}

	// Ensure the DB connection is established only once
	once.Do(func() {
		var err error
		testDBConn, err = openDBConnection(connectionURL)
		if err != nil {
			log.Fatalf("[TEST] fatal: %+v", err)
		}
	})

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("[TEST] error getting working directory: %+v", err)
	}
	parts := strings.Split(wd, "/tests")
	svcRootFolder := parts[0]
	assetPath := fmt.Sprintf("%s/tests/assets", svcRootFolder)
	log.Printf("[TEST] Asset path: %s\n", assetPath)

	// Initialize test fixtures
	fixtures, err = testfixtures.New(
		testfixtures.Database(testDBConn),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(assetPath),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		log.Fatalf("[TEST] fatal: %+v", err)
		return
	}

	// Load test fixtures
	if err := fixtures.Load(); err != nil {
		log.Fatalf("[TEST] failed to load fixtures: %s", err)
	}
	log.Println("[TEST] Assets loading completed")
}
