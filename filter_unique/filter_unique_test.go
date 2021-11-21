package filterunique_test

import (
	"testing"

	f "github.com/anfelo/go-testing/filter_unique"
	"github.com/stretchr/testify/assert"
)

func TestFilterUnique(t *testing.T) {
	input := []f.Developer{
		{Name: "Andres"},
		{Name: "Andres"},
		{Name: "David"},
		{Name: "Alexander"},
		{Name: "Eva"},
		{Name: "Alan"},
	}

	expected := []string{
		"Andres",
		"David",
		"Alexander",
		"Eva",
		"Alan",
	}

	assert.Equal(t, expected, f.FilterUnique(input))
}

func TestNegativeFilterUnique(t *testing.T) {
	input := []f.Developer{
		{Name: "Andres"},
		{Name: "Andres"},
		{Name: "David"},
		{Name: "Alexander"},
		{Name: "Eva"},
		{Name: "Alan"},
	}

	expected := []string{
		"Andres",
		"Alexander",
		"Eva",
		"Alan",
	}

	assert.NotEqual(t, expected, f.FilterUnique(input))
}
