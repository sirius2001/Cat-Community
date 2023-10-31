package area

type Area struct {
	Name string
}

func (c *Area) TableName() string {
	return "Cats"
}
