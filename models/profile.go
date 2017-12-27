package models

//reform:profiles
type Profile struct {
	ID        int32   `reform:"id,pk"`
	FirstName string  `reform:"first_name"`
	LastName  string  `reform:"last_name"`
	Avatar    *string `reform:"avatar"`
}
