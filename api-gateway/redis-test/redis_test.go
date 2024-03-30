package redistest

import (
	redismock "EXAM3/api-gateway/redis-test/redis-mock"
	"testing"

	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
)

func TestRedisOperations(t *testing.T) {
	conn := redigomock.NewConn()

	client := redismock.NewRedisClient(conn)

	conn.Command("SET", "key", "value").Expect("OK")

	conn.Command("GET", "key").Expect("value")

	err := client.SetValue("key", "value")
	assert.NoError(t, err)

	value, err := client.GetValue("key")
	assert.NoError(t, err)
	assert.Equal(t, "value", value)

	conn.ExpectationsWereMet()
}
