use anyhow::{Result}; 
use aoc_2022::helpers::grid::{Direction, Grid};

static DIRECTIONS: [Direction; 4] = [
    Direction::North,
    Direction::South,
    Direction::West,
    Direction::East,
];

fn main() -> Result<()> {
    let input = std::fs::read_to_string("src/inputs/day-08.txt")?;
    let grid = parse(&input);

    // println!("part1 -> {}", part1(grid));
    println!("part1 -> {}", part2(grid));

    return Ok(());
}

fn parse(s: &str) -> Grid<usize> {
    Grid::from_str(s, &|c| c.to_string().parse::<usize>().unwrap())
}

fn part1(grid: Grid<usize>) -> usize {
     grid
        .points()
        .iter()
        .filter(|point| {
            if grid.is_edge(point) {
                return true;
            }

            let tree = grid.get(point);

            DIRECTIONS.iter().any(|dir| {
                grid.walk(point, dir).all(|p| grid.get(&p) < tree)
            })
        })
        .count()
}

fn part2(grid: Grid<usize>) -> usize {
    grid
        .points()
        .iter()
        .filter(|point| ! grid.is_edge(point))
        .map(|point| {
            let tree = grid.get(point);

            DIRECTIONS
                .iter()
                .map(|dir| {
                    let mut trees_in_sight = 0;

                    for p in grid.walk(point, dir) {
                        trees_in_sight += 1;

                        if grid.get(&p) >= tree {
                            break;
                        }
                    }

                    trees_in_sight
                })
                .product()
        })
        .max()
        .unwrap()
}
