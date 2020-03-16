package gluon

import "fmt"

//  #[import] ee4g.com/gluon/http
// #[import] . github.com/gluon/post
//
// blub

/*

  #[import] ee4g.com/gluon/http
 #[import] . github.com/gluon/post
 */


//
func blub() {

}

// @MariaDB
// #[Postgres()]
type UserRepository interface {
	FindAll(limit, offset int) ([]*User, error)
}

type GreetingResource struct {
	users UserRepository
}

/*


#[

	http.GET
	POST

	NewMethod(name+"_gen")
]
bla

    test	#[Repository()] // also allowed
*/

// #[Entity("MyName")]
/* #[Entity2("MyName")]*/

// @ee4g.com/gluon/GET
// @Produces("text/plain")
// @Path("/greeting/{name}?limit={limit}")
func (u *GreetingResource) Hello(name string, limit int) string {
	fmt.Println("asdf")
	return "hello"
}

type User struct {
	Name string
}

func define() {

}
