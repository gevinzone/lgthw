package reflection

import "time"

type User struct {
	Name     string
	age      int
	Birthday time.Time
}

func CreateUser(name string, age int) User {
	//now := time.Now()
	//birthday := time.Date(now.Year()-age, now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	return User{
		Name: name,
		age:  age,
	}
}

func CreateUserPointer(name string, age int) *User {
	u := CreateUser(name, age)
	return &u
}

func (u User) private() {

}

func (u User) GetAge() int {
	return u.age
}

func (u *User) ChangeName(name string) *User {
	u.Name = name
	return u
}

type Person struct {
	Id int
	User
}

type Actor struct {
	Id int
	*User
}
