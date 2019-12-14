package mongo

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/ory/dockertest"
	"github.com/stretchr/testify/assert"
)

var mc *MongoClient

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mongo", "latest", []string{})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error

		portStr := resource.GetPort("27017/tcp")
		port, portErr := strconv.Atoi(portStr)
		if portErr != nil {
			return portErr
		}

		mc = New("localhost", port)
		err = mc.Ping()

		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestMongoClient(t *testing.T) {
	t.Run("ping", func(t *testing.T) {
		err := mc.Ping()
		assert.Equal(t, nil, err)
	})
	t.Run("dbs", func(t *testing.T) {
		names, err := mc.ListDatabaseNames()
		if err != nil {
			log.Fatalf("ListDatabaseNames failed with error: %s", err)
		}

		assert.EqualValues(t, []string{"admin", "local"}, names)
	})
	t.Run("collections", func(t *testing.T) {
		assert.Equal(t, 1, 1)
	})
	t.Run("finish", func(t *testing.T) {
		assert.Equal(t, 1, 1)
	})
}
