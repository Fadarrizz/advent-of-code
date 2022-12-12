use std::fs;

#[derive(Debug)]
struct Pair {
    low: i32,
    high: i32,
}

impl Pair {
   fn from_string(s: &str) -> Pair {
        let vec: Vec<i32> = s
            .split("-")
            .into_iter()
            .map(|c| c.parse().unwrap())
            .collect();

        return Pair {low: vec[0], high: vec[1]};
    }
}

fn parse_pairs(s: &str) -> Vec<Pair> {
    s.split(",")
        .into_iter()
        .map(|p| Pair::from_string(p))
        .collect()
}

fn fully_contains(p1: &Pair, p2: &Pair) -> bool {
    if p1.low <= p2.low && p1.high >= p2.high {
        true
    } else if p2.low <= p1.low && p2.high >= p1.high {
        true
    } else {
        false
    }
}

fn overlaps(p1: &Pair, p2: &Pair) -> bool {
    if p1.low >= p2.low && p1.low <= p2.high {
        true
    } else if p1.high <= p2.high && p1.high >= p2.low {
        true
    } else if p2.low >= p1.low && p2.low <= p1.high {
        true
    } else if p2.high <= p1.high && p2.high >= p1.low {
        true
    } else {
        false
    }
}

fn part1(s: &str) -> usize {
    s.lines()
        .into_iter()
        .map(|l| parse_pairs(l))
        .filter(|v| fully_contains(&v[0], &v[1]))
        .count()
}

fn part2(s: &str) -> usize {
    s.lines()
        .into_iter()
        .map(|l| parse_pairs(l))
        .filter(|v| overlaps(&v[0], &v[1]))
        .count()
}

fn main() {
    let input = fs::read_to_string("./input.txt").unwrap();

    println!("part 1 => {}", part1(&input));
    println!("part 2 => {}", part2(&input));
}

#[test]
fn test_pair_from_string() {
    let expected = Pair { low: 1, high: 2, };
    let actual = Pair::from_string("1-2");
    assert_eq!(expected.low, actual.low);
    assert_eq!(expected.high, actual.high);

    let expected = Pair { low: 4, high: 4, };
    let actual = Pair::from_string("4-4");
    assert_eq!(expected.low, actual.low);
    assert_eq!(expected.high, actual.high);
}

#[test]
fn test_parse_pairs() {
    let expected = vec!(
        Pair { low: 1, high: 2 },
        Pair { low: 3, high: 4 },
    );
    let actual = parse_pairs("1-2,3-4");
    assert_eq!(expected[0].low, actual[0].low);
    assert_eq!(expected[0].high, actual[0].high);
    assert_eq!(expected[1].low, actual[1].low);
    assert_eq!(expected[1].high, actual[1].high);
}

#[test]
fn test_fully_contains() {
    let p1 = Pair {low: 1, high: 4};
    let p2 = Pair {low: 2, high: 3};
    assert!(fully_contains(&p1, &p2));

    let p1 = Pair {low: 1, high: 4};
    let p2 = Pair {low: 4, high: 4};
    assert!(fully_contains(&p1, &p2));

    let p1 = Pair {low: 6, high: 7};
    let p2 = Pair {low: 1, high: 9};
    assert!(fully_contains(&p1, &p2));
}

#[test]
fn test_overlaps() {
    let p1 = Pair {low: 1, high: 4};
    let p2 = Pair {low: 2, high: 3};
    assert!(overlaps(&p1, &p2));

    let p1 = Pair {low: 1, high: 4};
    let p2 = Pair {low: 2, high: 6};
    assert!(overlaps(&p1, &p2));

    let p1 = Pair {low: 6, high: 7};
    let p2 = Pair {low: 1, high: 7};
    assert!(overlaps(&p1, &p2));

    let p1 = Pair {low: 2, high: 7};
    let p2 = Pair {low: 1, high: 4};
    assert!(overlaps(&p1, &p2));

    let p1 = Pair {low: 3, high: 5};
    let p2 = Pair {low: 6, high: 7};
    assert!(overlaps(&p1, &p2) == false);
}
