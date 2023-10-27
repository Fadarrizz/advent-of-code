use std::fmt;

#[derive(Eq, PartialEq)]
enum Part {
    One,
    Two,
}

fn main() -> color_eyre::Result<()> {
    let part = Part::One;

    let mut grid = Grid::parse(include_str!("../examples/day-14.txt"), &part);

    grid.simulate(&part);

    println!("{}", grid.num_settled());

    Ok(())
}

const SPAWN_POINT: Point = Point { x: 500, y: 0 };

#[derive(Debug, Copy, Clone, PartialEq, Eq, derive_more::Add, derive_more::AddAssign,derive_more::Sub)]
struct Point {
    x: i32,
    y: i32,
}

impl Point {
    fn parse(s: &str) -> Self {
        let mut coordinates = s.split(",");
        let (x, y) = (coordinates.next().unwrap(), coordinates.next().unwrap());
        Self {
            x: x.parse().unwrap(),
            y: y.parse().unwrap(),
        }
    }

    fn signum(self) -> Self {
        Self {
            x: self.x.signum(),
            y: self.y.signum(),
        }
    }
}

#[derive(Debug)]
struct RockLine {
    points: Vec<Point>,
}

impl RockLine {
    fn parse(s: &str) -> Self {
        Self {
            points: s.split(" -> ").map(Point::parse).collect(),
        }
    }

    fn path_points(&self) -> Vec<Point> {
        let mut points = self.points.iter().copied();
        let mut path_points = Vec::new();
        let Some(mut a) = points.next() else { return path_points };
        path_points.push(a);

        loop {
            let Some(b) = points.next() else { return path_points };
            let delta = (b - a).signum();

            loop {
                a += delta;
                path_points.push(a);
                if a == b {
                    break;
                }
            }
        }
    }
}

#[derive(Debug, Clone, Copy, Eq, PartialEq)]
enum Cell {
    Air,
    Rock,
    Sand,
}

struct Grid {
    origin: Point,
    width: usize,
    height: usize,
    cells: Vec<Cell>,
    current_cell: Point,
    settled: usize,
}

impl Grid {
    fn parse(input: &str, part: &Part) -> Self {
        let mut rock_lines: Vec<_> = input.lines().map(RockLine::parse).collect();

        let (mut min_x, mut max_x, mut min_y, mut max_y) = (i32::MAX, i32::MIN, i32::MAX, i32::MIN);

        for point in rock_lines
            .iter()
            .flat_map(|r| r.points.iter())
            .chain(std::iter::once(&SPAWN_POINT))
        {
            min_x = min_x.min(point.x);
            max_x = max_x.max(point.x);
            min_y = min_y.min(point.y);
            max_y = max_y.max(point.y);
        }

        if part == &Part::Two {
            let floor_y = max_y + 2;
            min_x = 300;
            max_x = 700;
            max_y = floor_y;

            rock_lines.push(RockLine {
                points: vec![
                    Point { x: min_x, y: floor_y },
                    Point { x: max_x, y: floor_y },
                ]
            });
        }

        let origin = Point { x: min_x, y: min_y };
        let width: usize = (max_x - min_x + 1).try_into().unwrap();
        let height: usize = (max_y - min_y + 1).try_into().unwrap();

        let mut grid = Grid {
            origin,
            width,
            height,
            cells: vec![Cell::Air; width * height],
            current_cell: SPAWN_POINT,
            settled: 0,
        };

        for point in rock_lines.iter().flat_map(|r| r.path_points()) {
            *grid.cell_mut(point).unwrap() = Cell::Rock;
        }

        grid
    }

    fn num_settled(&self) -> usize {
        self.settled
    }

    fn cell_index(&self, point: Point) -> Option<usize> {
        let Point { x, y } = point - self.origin;
        let x: usize = x.try_into().ok()?;
        let y: usize = y.try_into().ok()?;

        if x < self.width && y < self.height {
            Some(y * self.width + x)
        } else {
            None
        }
    }

    fn cell(&self, point: Point) -> Option<&Cell> {
        Some(&self.cells[self.cell_index(point)?])
    }

    fn cell_mut(&mut self, point: Point) -> Option<&mut Cell> {
        let index =self.cell_index(point)?; 
        Some(&mut self.cells[index])
    }

    fn simulate(&mut self, part: &Part) {
        loop {
            let straight_down = self.current_cell + Point { x: 0, y: 1 };
            let down_left = self.current_cell + Point { x: -1, y: 1 };
            let down_right = self.current_cell + Point { x: 1, y: 1 };
            let dirs = [straight_down, down_left, down_right];

            if part == &Part::Two 
                && matches!(self.cell(Point { x: 500, y: 0 }).unwrap(), Cell::Sand)
            {
                break;
            }

            // Move, if possible.
            if let Some(pos) = dirs
                .into_iter()
                .find(|pos| matches!(self.cell(*pos), Some(Cell::Air)))
            {
                self.current_cell = pos;
                dbg!(&self);
                continue;
            }

            // If not, check if we're offscreen
            if let Some(_) = dirs.into_iter().find(|pos| self.cell(*pos).is_none()) {
                break;
            }

            // Else, we've settled
            self.settled += 1;
            *self.cell_mut(self.current_cell).unwrap() = Cell::Sand;
            self.current_cell = SPAWN_POINT;
        }
    }
}

impl fmt::Debug for Grid {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        for y in 0..self.height {
            for x in 0..self.width {
                let point = Point {
                    x: x as _,
                    y: y as _,
                } + self.origin;
                let cell = self.cell(point).unwrap();
                let c = match cell {
                    Cell::Air => '.',
                    Cell::Rock => '#',
                    Cell::Sand => 'o',
                };
                write!(f, "{c}")?;
            }
            writeln!(f)?;
        }
        Ok(())
    }
}

#[test]
fn test_grid_cell_index() {
    let grid = Grid {
        origin: Point { x: 494, y: 0 },
        width: 10,
        height: 10,
        cells: vec![Cell::Air; 10 * 10],
        current_cell: SPAWN_POINT,
        settled: 0,
    };

    // first
    assert_eq!(Some(0), grid.cell_index(Point { x: 494, y: 0 }));
    // last
    assert_eq!(Some(99), grid.cell_index(Point { x: 503, y: 9 }));

    // outside of grid
    assert_eq!(None, grid.cell_index(Point { x: 493, y: 0 }));
    assert_eq!(None, grid.cell_index(Point { x: 504, y: 0 }));
    assert_eq!(None, grid.cell_index(Point { x: 494, y: -1 }));
    assert_eq!(None, grid.cell_index(Point { x: 504, y: 10 }));
}

#[test]
fn test_part1() {
    let part = Part::One;

    let mut grid = Grid::parse(include_str!("../examples/day-14.txt"), &part);

    grid.simulate(&part);

    assert_eq!(24, grid.num_settled());
}

#[test]
fn test_part2() {
    let part = Part::Two;

    let mut grid = Grid::parse(include_str!("../examples/day-14.txt"), &part);

    grid.simulate(&part);

    assert_eq!(93, grid.num_settled());
}
