use anyhow::Result;
use aoc_2022::helpers::grid::{Grid, Point, Direction};
use std::collections::BinaryHeap;
use std::cmp::Ordering;

fn main() -> Result<()> {
    let input = include_str!("../inputs/day-12.txt");    

    let parsed = parse(&input);

    // println!("{:?}", part1(parsed));
    println!("{:?}", part2(parsed));

    Ok(())
}

fn part1((grid, start, end): (Grid<char>, Point, Point)) -> Option<usize> {
    shortest_path(&grid, &vec!(start), &end)
}

fn part2((grid, _, end): (Grid<char>, Point, Point)) -> Option<usize> {
    let start_points: Vec<Point> = grid
        .points()
        .into_iter()
        .filter(|p| *grid.get(p) == 'a')
        .collect();

    shortest_path(&grid, &start_points, &end)
}

fn parse(s: &str) -> (Grid<char>, Point, Point) {
    let mut start: Option<Point> = None;
    let mut end: Option<Point> = None;

   let grid: Grid<char> = Grid::from_str(s, &mut |c, x, y| match c {
        'S' => {
            start = Some(Point {
                x: x as isize,
                y: y as isize,
            });
            'a'
        }
        'E' => {
            end = Some(Point {
                x: x as isize,
                y: y as isize,
            });
            'z'
        }
        c => c
    });

    (grid, start.unwrap(), end.unwrap())
}

fn shortest_path(grid: &Grid<char>, start_points: &Vec<Point>, end: &Point) -> Option<usize> {
    static DIRECTIONS: [Direction; 4] = [
        Direction::North,
        Direction::South,
        Direction::East,
        Direction::West,
    ];

    let mut dist: Vec<_> = (0..(grid.width * grid.height))
            .map(|_| usize::MAX)
            .collect();
    let mut queue = BinaryHeap::new();
    
    for start_point in start_points {
        let start_id = grid.id_for_point(start_point);
        dist[start_id] = 0;
        queue.push(Candidate {cost: 0, position: start_id});
    }

    let end_id = grid.id_for_point(end);

    while let Some(Candidate {cost, position}) = queue.pop() {
        if position == end_id {
            return Some(cost);
        }

        // Skip since there is a better position already.
        if cost > dist[position] {
            continue;
        }

        let current = grid.point_for_id(position);

        let neighbours: Vec<Point> = DIRECTIONS
            .iter()
            .map(|dir| current.get_neighbour(dir, 1))
            .filter(|neighbour| grid.lies_within(neighbour))
            .filter(|neighbour| (*grid.get(&neighbour) as isize - *grid.get(&current) as isize) < 2)
            .collect();

        for neighbour in neighbours {
            let next = Candidate {cost: cost + 1, position: grid.id_for_point(&neighbour)};

            if next.cost < dist[next.position] {
                queue.push(next);
                dist[next.position] = next.cost;
            }
        }
    }

    None
}

#[derive(Eq, PartialEq, Clone, Copy)]
struct Candidate {
    cost: usize,
    position: usize,
}

impl Ord for Candidate {
    fn cmp(&self, other: &Self) -> Ordering {
       other
        .cost
        .cmp(&self.cost)
        .then_with(|| self.position.cmp(&other.position))
    }
}

impl PartialOrd for Candidate {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}
