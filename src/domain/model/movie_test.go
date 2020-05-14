package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorCreateMovieNoName(t *testing.T) {
	m := Movie{
		Id:        1,
		User_id:   1,
		Title:     "",
		length:    12,
		Synopsis:  "",
		Image_url: "",
	}
	assert.NotNil(t,m.Validate())
}

func TestErrorCreateMovieWrongLength(t *testing.T) {
	m := Movie{
		Id:        1,
		User_id:   1,
		Title:     "movie Test",
		length:    -3,
		Synopsis:  "",
		Image_url: "",
	}
	assert.NotNil(t,m.Validate())
}

func TestCreateMovie(t *testing.T) {
	m := Movie{
		Id:        1,
		User_id:   1,
		Title:     "movie Test",
		length:    33,
		Synopsis:  "",
		Image_url: "",
	}
	assert.NotNil(t,m.Validate())
}