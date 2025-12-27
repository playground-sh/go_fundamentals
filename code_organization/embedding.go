package code_organization

import "fmt"

type user struct {
	id       int
	username string
	email    string
}

func (u *user) describe() {
	fmt.Printf("Username: %s, Email: %s\n", u.username, u.email)
}

type admin struct {
	user         // embedding `user`
	level        int
	canDeleteAll bool
}

func (a *admin) accessDatabase() {
	fmt.Println(a.username, "is accessing the DB with level:", a.level)
}

func EmbeddingDemo() {
	// Initialization looks a bit like composition
	superUser := admin{
		user: user{
			id:       1,
			username: "Tony Stark",
			email:    "tony@avengers.org",
		},
		level:        10,
		canDeleteAll: true,
	}

	// 1. Direct Field Access:
	// We don't need `superUser.user.username`
	fmt.Println("Admin:", superUser.username)

	// 2. Method Promotion:
	// `admin` didn't define `describe()`, but it "inherited" it from `user`
	superUser.describe()

	// 3. Specialized Logic:
	superUser.accessDatabase()
}
