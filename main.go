package main

func main() {
	NewBattle()
}

type Player struct {
	Name  string
	Level int
	Pet   Pet
}

type Pet struct {
	Name            string
	Level           int
	HP              int
	Characteristics Characteristics
}

type Characteristics struct {
	Strength   int
	Protection int
	Speed      int
}

func NewBattle() {
	player1 := Player{
		"Sander",
		3,
		Pet{
			"Onix",
			2,
			100,
			Characteristics{
				50,
				100,
				20,
			},
		},
	}

	player2 := Player{
		"Egor",
		2,
		Pet{
			"Pikachu",
			1,
			30,
			Characteristics{
				20,
				10,
				100,
			},
		},
	}

	player1.Attack(&player2)
}

func (p *Player) Attack(target *Player) {
	target.Pet.HP -= 5
}

func (p *Player) Dodge(playerAct, playerWait *Player) {
	if playerWait.Pet.Characteristics.Speed >= playerAct.Pet.Characteristics.Speed {
		//
	}
}
