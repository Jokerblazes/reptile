package model

type Pokemon struct {
	Id int
	Basic
	Power
}

type Basic struct {
	Height string //m

	Weight string //kg
}

type Power struct {
	Name    string
	Hp      int
	Attack  int
	Defense int
	Speed   int
	SpAtk   int
	SpDef   int
}
