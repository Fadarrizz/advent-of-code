use std::cmp;
use std::collections::HashSet;

#[derive(Clone, Debug)]
pub enum Direction {
    North,
    NorthEast,
    East,
    SouthEast,
    South,
    SouthWest,
    West,
    NorthWest,
}

#[derive(Clone, Debug, Eq, PartialEq, Hash)]
pub struct Point {
    pub x: isize,
    pub y: isize,
}

impl Point {
    pub fn get_neighbour(&self, direction: &Direction, steps: isize) -> Self {
        match direction {
            Direction::North => Self {
                x: self.x,
                y: self.y - steps,
            },
            Direction::NorthEast => Self {
                x: self.x + steps,
                y: self.y - steps,
            },
            Direction::East => Self {
                x: self.x + steps,
                y: self.y,
            },
            Direction::SouthEast => Self {
                x: self.x + steps,
                y: self.y + steps,
            },
            Direction::South => Self {
                x: self.x,
                y: self.y + steps,
            },
            Direction::SouthWest => Self {
                x: self.x - steps,
                y: self.y + steps,
            },
            Direction::West => Self {
                x: self.x - steps,
                y: self.y,
            },
            Direction::NorthWest => Self {
                x: self.x - steps,
                y: self.y - steps,
            },
        }
    }

    // Get distance between two points in any direction
    pub fn chebyshev_distance(&self, other: &Point) -> isize {
        cmp::max((self.x - other.x).abs(), (self.y - other.y).abs())
    }

    pub fn get_neighbour_towards(&self, other: &Point) -> Point {
        Point {
            x: self.x + (other.x - self.x).clamp(-1, 1),
            y: self.y + (other.y - self.y).clamp(-1, 1),
        }
    }
}

#[derive(Clone, Debug)]
pub struct HashGrid<Point> {
    pub points: HashSet<Point>,
}

impl HashGrid<Point> {
    pub fn get(&self, point: &Point) -> Option<&Point> {
        self.points.get(point)
    }

    pub fn insert(&mut self, point: Point) {
        self.points.insert(point);
    }
}

impl Default for HashGrid<Point> {
    fn default() -> Self {
        Self {
            points: HashSet::new(),
        }
    }
}

#[derive(Debug)]
pub struct Grid<T> {
    pub width: usize,
    pub height: usize,
    pub data: Vec<Vec<T>>,
}

impl<T> Grid<T> {
    pub fn from_str(input: &str, parse: &mut dyn FnMut(char, usize, usize) -> T) -> Self {
        let data: Vec<Vec<T>> = input
            .trim()
            .lines()
            .enumerate()
            .map(|(y, line)| {
                line.chars()
                    .enumerate()
                    .map(|(x, c)| parse(c, x, y))
                    .collect()
            })
            .collect();

        Grid {
            width: data[0].len(),
            height: data.len(),
            data,
        }
    }

    pub fn points(&self) -> Vec<Point> {
        let mut points = vec![];

        for y in 0..self.height {
            for x in 0..self.width {
                points.push(Point {
                    x: x as isize,
                    y: y as isize,
                })
            }
        }

        points
    }

    pub fn get(&self, point: &Point) -> &T {
        &self.data[point.y as usize][point.x as usize]
    }

    pub fn is_edge(&self, point: &Point) -> bool {
        point.x == 0
            || point.y == 0
            || point.x as usize == self.width - 1
            || point.y as usize == self.height - 1
    }

    pub fn is_off(&self, point: &Point) -> bool {
        point.x < 0
            || point.y < 0
            || point.x as usize >= self.width
            || point.y as usize >= self.height
    }

    pub fn lies_within(&self, point: &Point) -> bool {
        !self.is_off(point)
    }

    pub fn id_for_point(&self, point: &Point) -> usize {
        point.x as usize + self.width * point.y as usize
    }

    pub fn point_for_id(&self, id: usize) -> Point {
        Point {
            x: (id % self.width) as isize,
            y: (id / self.width) as isize,
        }
    }

    pub fn walk<'a>(&'a self, current: &'a Point, direction: &'a Direction) -> WalkingIterator<T> {
        WalkingIterator {
            current: current.clone(),
            grid: self,
            direction,
        }
    }
}

pub struct WalkingIterator<'a, T> {
    current: Point,
    direction: &'a Direction,
    grid: &'a Grid<T>,
}

impl<T> Iterator for WalkingIterator<'_, T> {
    type Item = Point;

    fn next(&mut self) -> Option<Self::Item> {
        let next_point = self.current.get_neighbour(self.direction, 1);

        if self.grid.is_off(&next_point) {
            return None;
        }

        self.current = next_point;

        Some(self.current.clone())
    }
}
