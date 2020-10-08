package storage

//Storage is an interface to storage of Users
type Storage interface {
	Users() UserRepo
}
