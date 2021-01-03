package entity

type Pokemon struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
