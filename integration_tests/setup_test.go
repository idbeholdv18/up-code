package integrationtests

import (
	"context"
	"log"
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

func TestMain(m *testing.M) {
	ctx := context.Background()

	net, err := network.New(ctx)
	if err != nil {
		log.Fatalf("error during network creating: %v", err)
		return
	}
	defer func() {
		if err := net.Remove(ctx); err != nil {
			log.Fatalf("error during network remove: %v", err)
		}
	}()

	postgresAlias := "test_db"

	pgc, err := helpers.CreatePostgresContainer(ctx,
		postgres.WithUsername("test"),
		postgres.WithPassword("test"),
		postgres.WithDatabase("test"),
		network.WithNetwork([]string{postgresAlias}, net),
	)
	defer func() {
		if err := pgc.Terminate(ctx); err != nil {
			log.Fatalf("error during postgres container terminating: %v", err)
		}
	}()
	if err != nil {
		log.Fatalf("error during postgres container creating: %v", err)
		return
	}

	testPool, err = pgxpool.New(ctx, pgc.ConnectionString)
	if err != nil {
		log.Fatalf("error during connection to db: %v", err)
		return
	}

	migrationsPath, err := filepath.Abs(filepath.Join("..", "migrations"))
	if err != nil {
		log.Fatalf("error during resolving migrations dir path: %v", err)
		return
	}

	internalConnStr, err := pgc.InternalConnectionString(postgresAlias)
	if err != nil {
		log.Fatalf("error during getting internal conn str: %v", err)
		return
	}
	mc, err := helpers.CreateMigrationsContainer(ctx,
		internalConnStr,
		network.WithNetwork(nil, net),
		testcontainers.WithFiles(testcontainers.ContainerFile{
			HostFilePath:      migrationsPath,
			ContainerFilePath: "/migrations",
			FileMode:          0o400,
		}),
	)
	defer func() {
		if err := mc.Terminate(ctx); err != nil {
			log.Fatalf("error during migration container terminating: %v", err)
		}
	}()
	if err != nil {
		log.Fatalf("error during migration container creating: %v", err)
		return
	}

	code := m.Run()
	if code != 0 {
		log.Fatalf("tests failed with exit code %d", code)
		return
	}
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
