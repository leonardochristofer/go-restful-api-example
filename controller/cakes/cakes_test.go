package cakes_test

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/leonardochristofer/go-restful-api-example/controller/cakes"
	"github.com/leonardochristofer/go-restful-api-example/mocks"
	"github.com/leonardochristofer/go-restful-api-example/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllCakes(t *testing.T) {

	mockCakeRepository := new(mocks.CakeInterface)

	server := &cakes.CakesServer{CakeDatabase: mockCakeRepository}

	// Create a new router and register the GetAllCakes handler
	r := mux.NewRouter()
	r.HandleFunc("/cakes", server.GetAllCakes).Methods("GET")

	// Create a mock HTTP server and recorder
	mockServer := httptest.NewServer(r)
	defer mockServer.Close()

	t.Run("success", func(t *testing.T) {

		a := 3.5

		mockCake := model.Cake{
			ID:          1,
			Title:       "Test Title",
			Description: "Test Description",
			Rating:      &a,
		}

		mockListCake := make([]model.Cake, 0)
		mockListCake = append(mockListCake, mockCake)

		mockCakeRepository.On("GetAllCakes", mock.Anything).Return(mockListCake, nil).Once()

		// Make a GET request to /cakes
		resp, err := http.Get(mockServer.URL + "/cakes")
		if err != nil {
			t.Fatalf("Failed to make GET request: %v", err)
		}
		defer resp.Body.Close()

		// Check the response status code
		assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

		// Decode the response body
		var actualCakes []model.Cake
		err = json.NewDecoder(resp.Body).Decode(&actualCakes)
		if err != nil {
			t.Fatalf("Failed to decode response body: %v", err)
		}

		// Compare the actual cakes with the expected cakes
		assert.Equal(t, mockListCake, actualCakes, "Expected and actual cakes do not match")
	})

	t.Run("error", func(t *testing.T) {

		mockCakeRepository.On("GetAllCakes", mock.Anything).Return(nil, errors.New("Unexpected")).Once()

		// Make a GET request to /cakes
		resp, err := http.Get(mockServer.URL + "/cakes")
		if err != nil {
			t.Fatalf("Failed to make GET request: %v", err)
		}
		defer resp.Body.Close()

		// Check the response status code
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "Expected status code 500")
	})

	t.Run("error", func(t *testing.T) {

		mockCakeRepository.On("GetAllCakes", mock.Anything).Return(nil, nil).Once()

		// Make a GET request to /cakes
		resp, err := http.Get(mockServer.URL + "/cakes")
		if err != nil {
			t.Fatalf("Failed to make GET request: %v", err)
		}
		defer resp.Body.Close()

		// Check the response status code
		assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Expected status code 404")
	})
}

func TestGetCake(t *testing.T) {
	mockCakeRepository := new(mocks.CakeInterface)

	server := &cakes.CakesServer{CakeDatabase: mockCakeRepository}

	// Create a new router and register the GetCake handler
	r := mux.NewRouter()
	r.HandleFunc("/cake/{id}", server.GetCake).Methods("GET")

	// Create a mock HTTP server and recorder
	mockServer := httptest.NewServer(r)
	defer mockServer.Close()

	t.Run("success", func(t *testing.T) {
		a := 3.5

		mockCake := model.Cake{
			ID:          1,
			Title:       "Test Title",
			Description: "Test Description",
			Rating:      &a,
		}

		mockCakeRepository.On("GetCake", mock.Anything, mock.Anything).Return(mockCake, nil).Once()

		// Make a GET request to /cake
		resp, err := http.Get(mockServer.URL + "/cake/1")
		if err != nil {
			t.Fatalf("Failed to make GET request: %v", err)
		}
		defer resp.Body.Close()

		// Check the response status code
		assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

		// Decode the response body
		var actualCakes model.Cake
		err = json.NewDecoder(resp.Body).Decode(&actualCakes)
		if err != nil {
			t.Fatalf("Failed to decode response body: %v", err)
		}

		// Compare the actual cakes with the expected cakes
		assert.Equal(t, mockCake, actualCakes, "Expected and actual cakes do not match")
	})

	t.Run("error", func(t *testing.T) {
		mockCakeRepository.On("GetCake", mock.Anything, mock.Anything).Return(model.Cake{}, errors.New("Unexpected")).Once()

		// Make a GET request to /cake
		resp, err := http.Get(mockServer.URL + "/cake/1")
		if err != nil {
			t.Fatalf("Failed to make GET request: %v", err)
		}
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		// Check the response status code
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "Expected status code 500")
	})

	t.Run("error", func(t *testing.T) {

		mockCakeRepository.On("GetCake", mock.Anything, mock.Anything).Return(model.Cake{}, nil).Once()

		// Make a GET request to /cake
		resp, err := http.Get(mockServer.URL + "/cake/abc")
		if err != nil {
			t.Fatalf("Failed to make GET request: %v", err)
		}
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		// Check the response status code
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Expected status code 400")
	})
}
