package api

// Ninja represents an individual ANW competitor.
type Ninja struct {
	ID        uint   `gorm:"column:ninja_id;primary_key"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Sex       string `gorm:"not null"`

	// `Age` and `Occupation` may be NULL since we don't always know them.
	Age        NullableInt
	Occupation NullableString

	// Social media may be NULL since some competitors don't have them.
	Instagram NullableString
	Twitter   NullableString
}

// CareerSummary provides a high-level view of a competitor's career, including
// the number of courses completed (qualifying, finals, and Mount Midoriyama
// stages), the number of seasons competed, their best finish, and their Ninja
// Rating.
type CareerSummary struct {
	ID          uint    `gorm:"column:summary_id;primary_key"`
	BestFinish  string  `gorm:"not null"`
	Speed       float64 `gorm:"not null"`
	Success     float64 `gorm:"not null"`
	Consistency float64 `gorm:"not null"`
	Seasons     int64   `gorm:"not null"`
	Qualifying  int64   `gorm:"not null"`
	Finals      int64   `gorm:"not null"`
	Stages      int64   `gorm:"not null"`
}
