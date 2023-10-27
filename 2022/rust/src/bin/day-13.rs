use serde::Deserialize;
use std::cmp::Ordering;

fn main() -> color_eyre::Result<()> {
    dbg!(part1());

    dbg!(part2());

    Ok(())
}

fn part1() -> usize {
    let mut sum = 0;
    for (i, groups) in include_str!("../inputs/day-13.txt").split("\n\n").enumerate() {
        let i = i + 1;

        let mut nodes = groups.lines().map(|line| serde_json::from_str::<Packet>(line).unwrap());
        let left = nodes.next().unwrap();
        let right = nodes.next().unwrap();

        if left < right {
            sum += i;
        }
    }

    sum
}

fn part2() -> usize {
    let dividers = vec![
        Packet::List(vec![Packet::Integer(2)]),
        Packet::List(vec![Packet::Integer(6)]),
    ];

    let mut packets = include_str!("../inputs/day-13.txt")
        .lines()
        .filter(|line| !line.is_empty())
        .map(|line| serde_json::from_str::<Packet>(line).unwrap())
        .chain(dividers.iter().cloned())
        .collect::<Vec<_>>();

    packets.sort();

    dividers
        .iter()
        .map(|d| packets.binary_search(d).unwrap() + 1)
        .product::<usize>()
}

#[derive(Deserialize, Clone, PartialEq, Eq, Debug)]
#[serde(untagged)]
enum Packet {
    List(Vec<Packet>),
    Integer(u64),
}

impl Packet {
    fn with_slice<T>(&self, f: impl FnOnce(&[Packet]) -> T) -> T {
        match self {
            Self::List(n) => f(&n[..]),
            Self::Integer(n) => f(&[Self::Integer(*n)]),
        }
    }
}

impl std::cmp::PartialOrd for Packet {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        match (self, other) {
            (Packet::Integer(a), Packet::Integer(b)) => a.partial_cmp(b),
            (left, right) => Some(left.with_slice(|left| {
                right.with_slice(|right| {
                    left
                        .iter()
                        .zip(right.iter())
                        .map(|(a, b)| a.cmp(b))
                        .find(|&ord| ord != Ordering::Equal)
                        .unwrap_or_else(|| left.len().cmp(&right.len()))
                })
            })),
        }
    }
}

impl std::cmp::Ord for Packet {
    fn cmp(&self, other: &Self) -> Ordering {
        self.partial_cmp(other).unwrap()
    }
}
