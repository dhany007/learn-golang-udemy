package test

import (
	"dhany007/belajar-golang-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("database")
	assert.NotNil(t, connection)

	// dua2 nya dari 2 provider itu diekseksui langsung dengan satu panggilan
	cleanup()
}
