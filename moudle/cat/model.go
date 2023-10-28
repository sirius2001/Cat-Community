package cat

type Cat struct {
	Name string
}

func (c *Cat) TableName() string {
	return "Cats"
}
