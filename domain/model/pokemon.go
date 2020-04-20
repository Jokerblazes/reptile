package model

type Pokemon struct {
	BasicProperty
	PowerProperty
}

type BasicProperty struct {
	Height string //m
	Weight string //kg
}

type PowerProperty struct {
	Name    string
	Hp      int
	Attack  int
	Defense int
	Speed   int
	SpAtk   int
	SpDef   int
}
