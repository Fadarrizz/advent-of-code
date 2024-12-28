package coordinates

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

func (c *Coordinate) Direction(other *Coordinate) Direction {
	if (c.Y - other.Y < 0) && (c.X - other.X == 0)  {
		return NORTH
	}

	if (c.Y - other.Y > 0) && (c.X - other.X == 0)  {
		return SOUTH
	}

	if (c.X - other.X < 0) && (c.Y - other.Y == 0)  {
		return WEST
	}

	if (c.X - other.X > 0) && (c.Y - other.Y == 0)  {
		return EAST
	}

	return NONE
}

func (c *Coordinate) Edge(other *Coordinate) Edge {
	return Edge{
		X: float32(c.X + other.X) / 2,
		Y: float32(c.Y + other.Y) / 2,
	}
}

func (c *Coordinate) Corners() []Corner {
	fx := float32(c.X)
	fy := float32(c.Y)

	return []Corner{
		{X: fx - 0.5, Y: fy - 0.5 }, // NW
		{X: fx + 0.5, Y: fy - 0.5 }, // NE
		{X: fx + 0.5, Y: fy + 0.5 }, // SE
		{X: fx - 0.5, Y: fy + 0.5 }, // SW
	}
}

// either one has half value, e.g. (1.0, 0.5)
type Edge struct {
	X float32
	Y float32
}

// always has half values, e.g. (0.5, 1.5)
type Corner struct {
	X float32
	Y float32
}

func (c *Corner) SurroundingSquares() []Coordinate {
	return []Coordinate{
		{X: int(c.X - 0.5), Y: int(c.Y - 0.5) }, // NW
		{X: int(c.X + 0.5), Y: int(c.Y - 0.5) }, // NE
		{X: int(c.X + 0.5), Y: int(c.Y + 0.5) }, // SE
		{X: int(c.X - 0.5), Y: int(c.Y + 0.5) }, // SW
	}
}

type Line struct {
	A Coordinate
	B Coordinate
}

func (l Line) Length() Coordinate {
    dx := l.B.X - l.A.X
    dy := l.B.Y - l.A.Y

	return Coordinate{dx, dy}
}

func (l Line) Extend(length Coordinate) Line {
	a := l.A
	b := l.B

    c := Coordinate{
        X: a.X - length.X,
        Y: a.Y - length.Y,
    }

    d := Coordinate{
        X: b.X + length.X,
        Y: b.Y + length.Y,
    }

	return Line{c, d}
}

type Direction int

const (
	NONE = -1
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

func Directions() []Direction {
	return []Direction{
		NORTH,
		EAST,
		SOUTH,
		WEST,
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
