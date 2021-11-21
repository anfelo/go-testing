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

	expected := map[string]int{
		"Andres":    1,
		"David":     1,
		"Alexander": 1,
		"Eva":       1,
		"Alan":      1,
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

	expected := map[string]int{
		"Andres":    1,
		"Alexander": 1,
		"Eva":       1,
		"Alan":      1,
	}

	assert.NotEqual(t, expected, f.FilterUnique(input))
}
