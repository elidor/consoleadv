package world

type Direction uint8

const (
	North     Direction = iota + 1
	East      Direction = iota + 2
	South     Direction = iota + 3
	West      Direction = iota + 4
	Down      Direction = iota + 5
	Up        Direction = iota + 6
	Northeast Direction = iota + 7
	Southeast Direction = iota + 8
	Southwest Direction = iota + 9
	NorthWest Direction = iota + 10
)

type World struct {
	Zones []*Zone
}

type Zone struct {
	Rooms []*Room
}

type Room struct {
	Short_Description string
	Description       string
	Exits             []*Exit
}

type Exit struct {
	Room      *Room
	Direction Direction
}

func ConstructWorld(world *World) {
	exit := &Exit{
		Direction: North,
		Room:      &Room{},
	}

	world.Zones = append(world.Zones, &Zone{})
	world.Zones[0].Rooms = append(world.Zones[0].Rooms, &Room{
		Short_Description: "A room",
		Description:       "A really descriptive room",
		Exits:             []*Exit{exit},
	})
}
