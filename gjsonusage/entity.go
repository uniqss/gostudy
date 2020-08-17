package main

type Entity struct {
	EntityName string
	Components []*Component
}

func (e *Entity) LoadDefinition(filename string) error {
	return nil
}
