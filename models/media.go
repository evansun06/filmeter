package models

// Defualt Film Type
type Film struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Director  string `json:"dirctor"`
	ReleaseYr uint16 `json:"release_yr"`
}

// Default Review
type Review struct {
	ID           int    `json:"id"`
	MovieID      int    `json:"movie_id"`
	UserID       int    `json:"user_id"`
	LetterGrade  string `json:"letter_grade"`
	WritenReview string `json:"written_review"`
}

