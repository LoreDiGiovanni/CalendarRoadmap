package types

type User struct{
   ID string `json:"id" bson:"_id,omitempty"`
   Username string `json:"username" bson:"username"`
   Email string `json:"email" bson:"email"`
   PWD string `json:"pwd" bson:"pwd"`
}
