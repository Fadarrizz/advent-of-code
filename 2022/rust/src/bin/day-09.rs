use anyhow::Result;
use aoc_2022::helpers::grid::{Direction, Point, HashGrid};

type Step = (Direction, usize);

fn main() -> Result<()> {
    let input = std::fs::read_to_string("src/inputs/day-09.txt")?;

    let steps: Vec<Step> = parse(&input); 

    println!("{}", simulate(&steps, 2));
    println!("{}", simulate(&steps, 10));

    Ok(())
}

fn parse(s: &str) -> Vec<Step> {
    s
        .lines()
        .filter_map(|line| {
            let (dir, steps) = line.split_once(" ")?;
            let steps = steps.parse().ok()?;
            let step = match dir {
                "U" => (Direction::North, steps),
                "D" => (Direction::South, steps),
                "L" => (Direction::West, steps),
                "R" => (Direction::East, steps),
                _ => unreachable!(),
            };

            Some(step)
        })
        .collect()
}

fn simulate(steps: &Vec<Step>, knot_count: usize) -> usize {
    let mut grid = HashGrid::default();
    let mut knots: Vec<Point> = vec![Point {x: 0, y: 0}; knot_count];

    grid.insert(Point {x: 0, y: 0});

    steps.iter().for_each(|(dir, steps)| {
        for _ in 0..*steps {
            knots[0] = knots[0].get_neighbour(dir, 1);

            for i in 0..(knot_count - 1) {
                let distance = knots[i].chebyshev_distance(&knots[i + 1]) > 1;

                if distance {
                    knots[i + 1] = knots[i + 1].get_neighbour_towards(&knots[i]);
                }

                grid.insert(knots.last().unwrap().clone());
            }
        }
    });

    grid.points.len()
}
