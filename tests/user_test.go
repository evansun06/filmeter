package tests

import (
	"testing"
)

// Test
func TestRetrieveAllUsersEndpoint(t *testing.T) {
	CleanTestDB()
	_, err := testDB.Exec("INSERT INTO users (username, email, hashed_password) VALUES ('John Pork', 'johnp@test.com', '1234')")
	if err != nil {
		t.Fatalf("failed to  insert a test user:  %v", err)
	}
	_, err = testDB.Exec("INSERT INTO users (username, email, hashed_password) VALUES ('John Doe', 'johnDoe@test.com', 'somehash')")

	if err != nil {
		t.Fatalf("failed to  insert a test user:  %v", err)
	}
	CleanTestDB()
}
