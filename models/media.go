package models

//  Defualt Film Type
type Film struct {
	Name string
	Director string
	Reviews []Review
	ReleaseYr uint16
}

// Default Review
type Review struct {
	LetterRating string
	WritenReview string
}