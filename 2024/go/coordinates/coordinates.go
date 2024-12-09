package coordinates

type Direction int

const (
	NORTH Direction = iota
	NORTHEAST
	EAST
	SOUTHEAST
	SOUTH
	SOUTHWEST
	WEST
	NORTHWEST
)

func (d Direction) Name() string {
	switch d {
	case NORTH:
		return "N"
	case NORTHEAST:
		return "NE"
	case EAST:
		return "E"
	case SOUTHEAST:
		return "SE"
	case SOUTH:
		return "S"
	case SOUTHWEST:
		return "SW"
	case WEST:
		return "W"
	case NORTHWEST:
		return "NW"
	default:
		return "" // should not be possible
	}
}

func AllDirections() []Direction {
	return []Direction{
		NORTH,
		NORTHEAST,
		EAST,
		SOUTHEAST,
		SOUTH,
		SOUTHWEST,
		WEST,
		NORTHWEST,
	}
}

type Coordinate struct {
	X int
	Y int
}

func New(row, col int) Coordinate {
	return Coordinate{col, row}
}

func Position(c Coordinate, d Direction) Coordinate {
	a, b := adjustmentsFromDirection(d)

	return Coordinate{c.X + a, c.Y + b}
}

func adjustmentsFromDirection(d Direction) (int, int) {
	switch d {
	case NORTH:
		return 0, -1
	case NORTHEAST:
		return 1, -1
	case EAST:
		return 1, 0
	case SOUTHEAST:
		return 1, 1
	case SOUTH:
		return 0, 1
	case SOUTHWEST:
		return -1, 1
	case WEST:
		return -1, 0
	case NORTHWEST:
		return -1, -1
	default:
		return 0, 0 // should not be possible
	}
}
