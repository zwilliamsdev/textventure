package player

type PlayerClass interface {
	ClassName() string
	GetStartingHealth() int
	GetStartingResource() int
	PerformAbility(p *Player) string
}

type Player struct {
	name       string
	class      PlayerClass
	level      int
	experience int
	health     int
	resource   int
	inventory  map[string]int
}
