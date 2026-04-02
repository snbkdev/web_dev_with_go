package services

import (
	"net/http"
	"testing"
	"web_app/models"

	"github.com/stretchr/testify/assert"
)

func TestValidateRunnerInvalidFirstName(t *testing.T) {
	runner := &models.Runner{
		LastName: "Smith",
		Age:      30,
		Country:  "United States",
	}

	responseErr := validateRunner(runner)
	assert.NotEmpty(t, responseErr)
	assert.Equal(t, "Invalid first name", responseErr.Message)
	assert.Equal(t, http.StatusBadRequest, responseErr.Status)
}

func TestValidateRunner(t *testing.T) {
	tests := []struct {
		name string
		runner *models.Runner
		want *models.ResponseError
	}{
		{
			name: "Invalid First Name",
			runner: &models.Runner{
				LastName: "Smith",
				Age: 30,
				Country: "United States",
			},
			want: &models.ResponseError{
				Message: "Invalid First Name",
				Status: http.StatusBadRequest,
			},
		},
		{
			name: "Invalid Last Name",
			runner: &models.Runner{
				FirstName: "John",
				Age: 30,
				Country: "United States",
			},
			want: &models.ResponseError{
				Message: "Invalid Last Name",
				Status: http.StatusBadRequest,
			},
		},
		{
			name: "Invalid Age",
			runner: &models.Runner{
				FirstName: "John",
				LastName: "Smith",
				Age: 250,
				Country: "United States",
			},
			want: &models.ResponseError{
				Message: "Invalid Age",
				Status: http.StatusBadRequest,
			},
		},
		{
			name: "Invalid Country",
			runner: &models.Runner{
				FirstName: "John",
				LastName: "Smith",
				Age: 30,
			},
			want: &models.ResponseError{
				Message: "Invalid country",
				Status: http.StatusBadRequest,
			},
		},
		{
			name: "Valid Runner",
			runner: &models.Runner{
				FirstName: "John",
				LastName: "Smith",
				Age: 30,
				Country: "United States",
			},
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			responseErr := validateRunner(test.runner)
			assert.Equal(t, test.want, responseErr)
		})
	}
}
