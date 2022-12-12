use std::fs;
use std::collections::HashMap;

fn are_different(chars: Vec<char>) -> bool {
    chars
        .iter()
        .map(|c| (c, true))
        .collect::<HashMap<_, _>>()
        .len() == chars.len()
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

fn main() {
    let input = fs::read_to_string("./input.txt").unwrap();

    println!("part 1: {}", start_of_packet_marker(&input.trim(), 4));
    println!("part 2: {}", start_of_packet_marker(&input.trim(), 14));
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
