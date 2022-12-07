// detect first start-of-packet marker
// start of packet = 4 chars
// identify first position where most recent four chars are different

use std::fs;
use std::collections::HashMap;

fn are_different(chars: Vec<char>) -> bool {
    let len = chars.len();

    chars
        .iter()
        .map(|c| (c, true))
        .collect::<HashMap<_, _>>()
        .len() == len
}

fn start_of_packet_marker(s: &str, len: usize) -> usize {
    let chars: Vec<char> = s.chars().into_iter().collect();
    let mut min = 0;
    let mut max = len;

    while ! are_different(chars[min..max].to_vec()) {
        min += 1;
        max += 1;
    }

    max
}

fn part1(s: &str) -> usize {
    start_of_packet_marker(&s, 4)
}

fn part2(s: &str) -> usize {
    start_of_packet_marker(&s, 14)
}

fn main() {
    let input = fs::read_to_string("./input.txt").unwrap();

    println!("part 1: {}", part1(&input.trim()));
    println!("part 2: {}", part2(&input.trim()));
}

#[test]
fn test_are_different() {
    assert!(are_different(vec!['a', 'b']));
    assert!(are_different(vec!['a', 'b', 'c', 'd']));
    assert!(! are_different(vec!['a', 'a']));
    assert!(! are_different(vec!['a', 'b', 'c', 'a']));
}

#[test]
fn test_start_of_packet_marker() {
    let s = "bvwbjplbgvbhsrlpgdmjqwftvncz";
    assert_eq!(5, start_of_packet_marker(&s, 4));

    let s = "nppdvjthqldpwncqszvftbrmjlhg";
    assert_eq!(6, start_of_packet_marker(&s, 4));

    let s = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg";
    assert_eq!(10, start_of_packet_marker(&s, 4));

    let s = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw";
    assert_eq!(11, start_of_packet_marker(&s, 4));
}
