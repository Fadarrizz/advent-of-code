use std::collections::HashMap;
use itertools::Itertools;

fn main() {
    let input = include_str!("../inputs/day-03.txt");      

    println!("{}", part1(&input));
    println!("{}", part2(&input));
}

fn part1(input: &str) -> usize {
    let mut x = 0;
    let mut y = 0;
    let mut seen: HashMap<Point, usize> = HashMap::new();
    seen.insert(Point { x: 0, y: 0}, 1);

    for c in input.trim().chars() {
        match c {
            '^' => y -= 1,
            '>' => x += 1,
            'v' => y += 1,
            '<' => x -= 1,
            _ => panic!("Unknown char"),
        };

        let point = Point { x, y };
        if seen.get(&point) == None {
            seen.insert(point.clone(), 0);
        }
        *seen.get_mut(&point).unwrap() += 1;
    }

    seen.iter().count()
}

fn part2(input: &str) -> usize {
    let mut santa = Point { x: 0, y: 0 };
    let mut robot = Point { x: 0, y: 0 };
    let mut seen: Map = Map::new();
    seen.insert(Point { x: 0, y: 0}, 2);

    for (c1, c2) in input.trim().chars().tuple_windows().step_by(2) {
        match c1 {
            '^' => santa.update(0, -1),
            '>' => santa.update(1, 0),
            'v' => santa.update(0, 1),
            '<' => santa.update(-1, 0),
            _ => panic!("Unknown char"),
        };

        match c2 {
            '^' => robot.update(0, -1),
            '>' => robot.update(1, 0),
            'v' => robot.update(0, 1),
            '<' => robot.update(-1, 0),
            _ => panic!("Unknown char"),
        };

        seen.update_or_insert(&santa);
        seen.update_or_insert(&robot);
    }

    seen.count()
}

#[derive(Clone, Debug, Eq, PartialEq, Hash)]
struct Point {
    x: i32,
    y: i32,
}

impl Point {
    fn update(&mut self, x: i32, y: i32) {
        self.x += x;
        self.y += y;
    }
}

#[derive(Debug)]
struct Map {
    items: HashMap<Point, usize>,
}

impl Map {
    fn new() -> Self {
        Self { items: HashMap::new() }
    }

    fn insert(&mut self, point: Point, count: usize) {
        self.items.insert(point, count);
    }

    fn update_or_insert(&mut self, point: &Point) {
        if self.items.get(point) == None {
            self.items.insert(point.clone(), 0);
        }
        *self.items.get_mut(point).unwrap() += 1;
    }

    fn count(&self) -> usize {
        self.items.iter().count()
    }
}

#[test]
fn test_part1() {
    let input = ">";
    assert_eq!(2, part1(&input));

    let input = "^>v<";
    assert_eq!(4, part1(&input));

    let input = "^v^v^v^v^v";
    assert_eq!(2, part1(&input));
}
