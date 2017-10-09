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
	ID          uint `gorm:"column:summary_id;primary_key"`
	BestFinish  string
	Speed       float64
	Success     float64
	Consistency float64
	Seasons     int64
	Qualifying  int64
	Finals      int64
	Stages      int64
	NinjaID     uint
	Ninja       Ninja
}

// TableName returns the name used in our database.
//
// By default, Gorm converts `TableName` to `table_name`.
//
// TODO: Update our database to use `table_name` and remove.
func (CareerSummary) TableName() string {
	return "careersummary"
}
