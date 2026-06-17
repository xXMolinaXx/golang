package user

type Service interface {
	CreateUser(name, email, password string) error
	ReadUser(id string) (*User, error)
	ReadAllUsers() ([]User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}

type UserService struct {
	db Repository
}

func NewUserService(db Repository) *UserService {
	return &UserService{db: db} // crea la variable y retorna un puntero
}

func (s *UserService) CreateUser(name, email, password string) error {
	user := User{
		Fullname: name,
		Email:    email,
		Password: password,
	}
	err := s.db.CreateUser(&user)
	return err
}

func (s *UserService) ReadUser(id string) (*User, error) {
	return s.db.ReadUser(id)
}
func (s *UserService) ReadAllUsers() ([]User, error) {
	return s.db.ReadAllUsers()
}
func (s *UserService) UpdateUser(user *User, requestBody CreateUserRequest) error {
	user.Fullname = requestBody.Name
	user.Email = requestBody.Email
	user.Password = requestBody.Password
	return s.db.UpdateUser(user)
}
func (s *UserService) DeleteUser(id string) error {
	return s.db.DeleteUser(id)
}
