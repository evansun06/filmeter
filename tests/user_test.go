package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"restful-movie-api/models"
)

// GET endpoint test @ /users
func TestGETAllUsers(t *testing.T) {
	CleanTestDB()
	// Insert Test Users
	_, err := testDB.Exec("INSERT INTO users (username, email, hashed_password) VALUES ('John Pork', 'johnp@test.com', '1234')")
	if err != nil {
		t.Fatalf("failed to  insert a test user:  %v", err)
	}
	_, err = testDB.Exec("INSERT INTO users (username, email, hashed_password) VALUES ('John Doe', 'johnDoe@test.com', 'somehash')")

	if err != nil {
		t.Fatalf("failed to  insert a test user:  %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Create HTTP recorder
	recorder := httptest.NewRecorder()

	// Serve
	testRouter.ServeHTTP(recorder, req)

	// Ensure OK code
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status OK, got %v", recorder.Code)
	}

	var users []*models.User

	unmarshalErr := json.Unmarshal(recorder.Body.Bytes(), &users)

	if unmarshalErr != nil {
		t.Fatalf("failed to unmarshal JSON into []*models.User: %v", unmarshalErr)
	}

	// Validate
	if len(users) != 2 {
		t.Fatalf("Expected 2 users, got %d", len(users))
	}

	expectedUsers := map[string]string{
		"John Pork": "johnp@test.com",
		"John Doe":  "johnDoe@test.com",
	}

	for _, user := range users {
		expectedEmail, exists := expectedUsers[user.Username]

		if !exists {
			t.Fatalf("unexpected user in response: %s", user.Username)
		}

		if user.Email != expectedEmail {
			t.Fatalf("expected email %s for user %s, instead got %s", expectedEmail, user.Username, user.Email)
		}
	}

	CleanTestDB()
}

// PUT endpoint test @ /user
// Validate single user insert functionality
func TestPutSingleUser(t *testing.T) {
	CleanTestDB()

	newUser := map[string]string{
		"username":        "John Doe",
		"email":           "johnDoe@gmail.com",
		"hashed_password": "12345",
	}

	requestBody, err := json.Marshal(newUser)

	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	// Create the PUT request
	req, err := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create PUT request: %v", err)
	}

	req.Header.Set("Content-Tpye", "application/json")

	// Create an HTTP recorder to capture the response
	recorder := httptest.NewRecorder()
	// Serve
	testRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("Expected status OK, instead recieved %d", recorder.Code)
	}

	var verifyUser models.User
	err = testDB.QueryRow("SELECT id, username, email, hashed_password FROM users WHERE username = $1", "John Doe").
		Scan(&verifyUser.ID, &verifyUser.Username, &verifyUser.Email, &verifyUser.HashedPassword)
		
	
	if err != nil {
		t.Fatalf("Failed to query back the user for verification: %v", err)
	}

	if verifyUser.Username != "John Doe" || verifyUser.Email != "johnDoe@gmail.com" || verifyUser.HashedPassword != "" {
		t.Fatalf("PUT and GET user data mismatch, %s, %s, %s", verifyUser.Username, verifyUser.Email, verifyUser.HashedPassword)
	}

	CleanTestDB()
}

// GET endpoint @ /user/:id
// Validate Fetch User By ID
func TestGETUserByID(t *testing.T) {
	CleanTestDB()
	// Insert Test Users
	_, err := testDB.Exec("INSERT INTO users (username, email, hashed_password) VALUES ('John Doe', 'johnp@test.com', '1234')")
	if err != nil {
		t.Fatalf("failed to  insert a test user:  %v", err)
	}
	
	var userID int64 = 1

	url := fmt.Sprintf("/user/%d", userID)
	// Create HTTP request
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Create HTTP recorder
	recorder := httptest.NewRecorder()

	// Serve
	testRouter.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
	
	var fetchedUser models.User

	err = json.NewDecoder(recorder.Body).Decode(&fetchedUser)
	if err != nil {
		t.Fatalf("failed to decode json: %v", err)
		return
	}

	// Verify the fetched user data
    if fetchedUser.ID != userID || fetchedUser.Username != "John Doe" || fetchedUser.Email != "johnp@test.com" {
        t.Fatalf("Fetched user data mismatch: %+v", fetchedUser)
    }
	CleanTestDB()

}
