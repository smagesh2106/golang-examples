/*
 * Swagger User store
 *
 * This is a sample server User server.  You can find out more about     Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).      For this sample, you can use the api key `special-key` to test the authorization     filters.
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

//mgo "examples/go-server-user_example/db"

//"gopkg.in/validator.v2"

type User struct {
	Id string `json:"_id,omitempty" bson:"_id,omitempty"`

	Username string `validate:"nonzero,nonnil" json:"username,omitempty" bson:"username, omitempty"`

	FirstName string `validate:"nonzero,nonnil" json:"firstName" bson:"firstname"`

	LastName string `validate:"nonzero,nonnil" json:"lastName" bson:"lastname"`

	Email string `validate:"nonzero" json:"email,omitempty" bson:"email, omitempty"`

	Password string `validate:"nonzero" json:"password,omitempty" bson:"password, omitempty"`

	Phone string `validate:"nonzero" json:"phone,omitempty" bson:"phone, omitempty"`

	Age int32 `validate:"min=1,max=120" json:"age,omitempty" bson:"age, omitempty"`

	UserStatus int32 `validate:"nonzero" json:"userStatus,omitempty" bson:"userStatus, omitempty"`

	Address Address `json:"address,omitempty" bson:"address, omitempty"`

	Hobbies []Hobby `json:"hobbies" bson:"hobbies"`
}

type Address struct {
	Id string `json:"_id,omitempty" bson:"_id,omitempty"`

	DoorNumber int32 `validate:"min=1,max=120" json:"doorNumber,omitempty" bson:"doorNumber, omitempty"`

	StreetName string `validate:"nonzero" json:"streetName,omitempty" bson:"streetName, omitempty"`

	City string `validate:"nonzero" json:"city,omitempty" bson:"city, omitempty"`
}

type Hobby struct {
	Name   string `json:"name", bson: "name"`
	Number int    `validate:"min=1,max=120" json:"number" bson:"number" `
}
