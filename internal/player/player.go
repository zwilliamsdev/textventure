package player

type PlayerClass interface {
	GetClass() string
	GetStartingHealth() int
	// Return resource type and amount
	// Warrior: rage, mage: mana etc.
	GetStartingResource() (string, int)
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

// For this small project I will create the player classes
// in this single file in a real project I would break them
// into individual files.

// Warrior
type Warrior struct{}

func (w Warrior) GetClass() string {
	return "Warrior"
}

func (w Warrior) GetStartingHealth() int {
	return 150
}

func (w Warrior) GetStartingResource() (string, int) {
	return "Rage", 0
}

// Mage
type Mage struct{}

func (m Mage) GetClass() string {
	return "Mage"
}

func (m Mage) GetStartingHealth() int {
	return 100
}

func (m Mage) GetStartingResource() (string, int) {
	return "Mana", 75
}
