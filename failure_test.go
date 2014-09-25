package failure_test

import (
	"errors"
	"fmt"

	"github.com/dagoof/failure"
)

func ExampleFail() {
	const max = 1<<8 - 1

	SumByte := func(ns ...uint) (sum uint8, err error) {
		defer failure.Recover(&err)

		for _, n := range ns {
			if n+uint(sum) > max {
				failure.Fail(errors.New(
					fmt.Sprintf("%d overflows byte", n+uint(sum)),
				))
			}

			sum += uint8(n)
		}

		return
	}

	var bytes []uint
	for i := uint(6); i < 10; i++ {
		bytes = append(bytes, 1<<i)
		fmt.Println(SumByte(bytes...))
	}

	// Output:
	// 64 <nil>
	// 192 <nil>
	// 192 448 overflows byte
	// 192 448 overflows byte
}

func ExampleFailFunc() {
	type User struct {
		Name string
		Age  int
	}

	createUsers := func(users ...User) (database map[uint]User, err error) {
		const maxAge = 100
		var pkey uint
		database = map[uint]User{}

		defer failure.Recover(&err)
		fail := failure.FailFunc(func() {
			fmt.Println("rolling back changes")
			for ID := range database {
				delete(database, ID)
			}
		})

		create := func(user User) error {
			if user.Age > maxAge {
				return errors.New(fmt.Sprintf("%d is too old", user.Age))
			}

			pkey++
			database[pkey] = user
			return nil
		}

		for _, user := range users {
			fail(create(user))
		}

		return database, err
	}

	users := []User{
		{"Fred", 15},
		{"Jill", 81},
	}
	fmt.Println(createUsers(users...))
	fmt.Println()

	users = append(users, User{"Anna", 121})
	fmt.Println(createUsers(users...))

	// Output:
	// map[1:{Fred 15} 2:{Jill 81}] <nil>
	//
	// rolling back changes
	// map[] 121 is too old
}
