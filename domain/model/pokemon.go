package model

type Pokemon struct {
	BasicProperty
	PowerProperty
}

type BasicProperty struct {
	height float32 //m
	weight float32 //kg
}

type PowerProperty struct {
	name    string
	hp      int
	attack  int
	defense int
	speed   int
	spAtk   int
	spDef   int
}
