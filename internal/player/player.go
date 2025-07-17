package player

type PlayerClass interface {
	GetClass() string
	GetStartingHealth() int
	GetStartingResourceType() string
	GetStartingResourceAmount() int
	performBasicAttack(p *Player) int
	PerformAbility(p *Player) int
}

type Player struct {
	name           string
	class          PlayerClass
	level          int
	experience     int
	health         int
	resourceType   string
	resourceAmount int
	// Item Name -> Quantity
	inventory map[string]int
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

func (w Warrior) GetStartingResourceName() string {
	return "Rage"
}

func (w Warrior) GetStartingResourceAmount() int {
	return 0
}

// Mage
type Mage struct{}

func (m Mage) GetClass() string {
	return "Mage"
}

func (m Mage) GetStartingHealth() int {
	return 100
}

func (m Mage) GetStartingResourceName() string {
	return "Mana"
}

func (m Mage) GetStartingResourceAmount() int {
	return 75
}

// Player
func NewPlayer(name string, class PlayerClass) *Player {
	return &Player{
		name:           name,
		class:          class,
		level:          1,
		experience:     0,
		health:         class.GetStartingHealth(),
		resourceType:   class.GetStartingResourceType(),
		resourceAmount: class.GetStartingResourceAmount(),
	}
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Level() int {
	return p.Level()
}

func (p *Player) AddExperience(amount int) {
	p.experience += amount
	// TODO: Calculate experience required to level and handle
	// leveling player
}

func (p *Player) AddItem(itemName string) {
	// Increase the amount of itemName the player has
	p.inventory[itemName]++
}
