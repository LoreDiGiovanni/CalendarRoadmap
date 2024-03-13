package types

type User struct{
   ID string `json:"id" bson:"_id,omitempty"`
   Email string `json:"email" bson:"email"`
   PWD string `json:"pwd" bson:"pwd"`
   JWT string `json:"jwt" bson:"jwt"`
   Salt string  `json:"salt" bson:"salt"`
}
