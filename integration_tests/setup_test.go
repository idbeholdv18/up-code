package integrationtests

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/idbeholdv18/up-code/app/user"
	"github.com/idbeholdv18/up-code/integration_tests/internal/helpers"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/network"
)

var testPool *pgxpool.Pool

func run(m *testing.M) int {
	ctx := context.Background()

	net, netCleanup, err := helpers.NewNetwork(ctx)
	if err != nil {
		log.Printf("network create error: %v", err)
		return 1
	}
	defer netCleanup()

	postgresAlias := "test_db"

	pgc, postgresCleanup, err := helpers.NewPostgresContainer(ctx,
		postgres.WithUsername("test"),
		postgres.WithPassword("test"),
		postgres.WithDatabase("test"),
		network.WithNetwork([]string{postgresAlias}, net),
	)
	if err != nil {
		log.Printf("postgres container creating error: %v", err)
		return 1
	}
	defer postgresCleanup()

	internalConnStr, err := pgc.InternalConnectionString(postgresAlias)
	if err != nil {
		log.Printf("postgres test containerinternal connection string getting error: %v", err)
		return 1
	}

	migrationsPath, err := filepath.Abs(filepath.Join("..", "migrations"))
	if err != nil {
		log.Printf("host migrations path resolving error: %v", err)
		return 1
	}

	_, migrationsCleanup, err := helpers.NewMigrationsContainer(ctx,
		internalConnStr,
		network.WithNetwork(nil, net),
		testcontainers.WithFiles(testcontainers.ContainerFile{
			HostFilePath:      migrationsPath,
			ContainerFilePath: "/migrations",
			FileMode:          0o400,
		}),
	)
	if err != nil {
		log.Printf("migrations container creating error: %v", err)
		return 1
	}
	defer migrationsCleanup()

	testPool, err = pgxpool.New(ctx, pgc.ConnectionString)
	if err != nil {
		log.Printf("pgxpool creating creating error: %v", err)
		return 1
	}
	defer testPool.Close()

	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(run(m))
}

func TestPgxpoolUserRepo_Create(t *testing.T) {
	ctx := context.Background()
	repo := user.NewPgxpoolUserRepo(testPool)

	// TODO add truncate helper
	created, err := repo.Create(ctx, "test@mail.ru", "test_password", "name surname")
	if err != nil {
		t.Fatalf("create user error: %v", err)
	}

	if created.ID == 0 ||
		created.Email != "test@mail.ru" ||
		created.FullName != "name surname" ||
		created.PasswordHash != "test_password" {
		t.Fatalf(
			"incorrect user create. Expected 0, test@mail.ru, name surname, test_password\ngot %d, %s, %s, %s",
			created.ID, created.Email, created.FullName, created.PasswordHash,
		)
	}
}
