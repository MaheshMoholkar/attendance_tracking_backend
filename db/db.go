package db

const (
	DB_NAME = "attendance-db"
	DB_URI  = "mongodb://localhost:27017"
)

type Store struct {
	UserStore
	//StudentStore
	//AttendanceStore
}
