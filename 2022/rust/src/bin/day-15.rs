use std::{fmt, path::absolute};
use nom::{
    bytes::streaming::tag,
    combinator::{all_consuming, map},
    sequence::{preceded, separated_pair},
    Finish, IResult,
};

fn main() {
    let input = include_str!("../examples/day-15.txt");

    let map = Map::parse(input);
    map.dump();
}

struct Map {
    records: Vec<Record>,
}

impl Map {
    fn parse(input: &str) -> Self {
        let records = input.lines().map(Record::must_parse).collect();
        Self { records }
    }

    fn dump(&self) {
        for record in &self.records {
            println!("{record:?}");
        }
    }
}

#[derive(Debug)]
struct Record {
    sensor: Point,
    beacon: Point,
}

impl Record {
    fn must_parse(i: &str) -> Self {
        all_consuming(Self::parse)(i)
            .finish()
            .expect("failed to parse input")
            .1
    }

    fn parse(i: &str) -> IResult<&str, Self> {
        // example line:
        // Sensor at x=2, y=18: closest beacon is at x=-2, y=15
        map(
            separated_pair(
                preceded(tag("Sensor at "), Point::parse),
                tag(": closest beacon is at "),
                Point::parse,
            ),
            |(sensor, beacon)| Record { sensor, beacon }
        )(i)
    }
}

#[derive(Clone, Copy, PartialEq, Eq, Hash)]
struct Point {
    x: i64,
    y: i64,
}

impl fmt::Debug for Point {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}, {}", self.x, self.y)
    }
}

impl Point {
    fn parse(i: &str) -> IResult<&str, Point> {
        map(
            separated_pair(
                preceded(tag("x="), nom::character::complete::i64),
                tag(", "),
                preceded(tag("y="), nom::character::complete::i64),
            ),
            |(x, y)| Point { x, y },
        )(i)
    }

    // https://en.wikipedia.org/wiki/Taxicab_geometry
    fn manhattan_distance(self, other: Self) -> i64 {
        (self.x.abs_diff(other.x) + self.y.abs_diff(other.y)) as i64
    }
}
