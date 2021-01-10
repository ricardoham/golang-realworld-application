package presenter

type Pokemon struct {
	ID   int    `json:"id" bson:"-"`
	Name string `json:"name" bson:"-"`
}
