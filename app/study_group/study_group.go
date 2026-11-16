package studygroup

import "time"

type StudyGroup struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Year      int       `db:"year"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
