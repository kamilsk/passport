package postgres_test

import (
	"testing"

	"github.com/kamilsk/passport/pkg/dao/postgres"
	"github.com/kamilsk/passport/pkg/domain"
	"github.com/stretchr/testify/assert"
)

const (
	DSN  = "postgres://postgres:postgres@db:5432/postgres"
	UUID = domain.UUID("41ca5e09-3ce2-4094-b108-3ecc257c6fa4")
)

func TestDialect(t *testing.T) {
	assert.Equal(t, "postgres", postgres.Dialect())
}