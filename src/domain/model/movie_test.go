package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorCreateMovieNoName(t *testing.T) {
	var m Movie
	movie,err := m.Create(
		1,
		 "",
		12,
		"",
		"")
	assert.NotNil(t,err)
	assert.Empty(t,movie)
}

func TestErrorCreateMovieWrongLength(t *testing.T) {
	var m Movie
	movie,err := m.Create(
		1,
		"insidius",
		0,
		"",
		"")
	assert.NotNil(t,err)
	assert.Empty(t,movie)
}

func TestCreateMovie(t *testing.T) {
	var m Movie
	movie,err := m.Create(
		1,
		"insidius",
		60,
		"",
		"")
	assert.NotEmpty(t,movie)
	assert.Nil(t,err)
}