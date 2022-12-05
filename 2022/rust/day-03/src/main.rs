use std::fs;

fn split(s: &str) -> Vec<&str> {
    let half = s.len() / 2;
    vec!(&s[0..half], &s[half..])
}

fn find_common(s: Vec<&str>) -> char {
    for c in 'a'..='z' {
        if s.iter().all(|&s| s.contains(c)) {
            return c;
        }
    }

    for c in 'A'..='Z' {
        if s.iter().all(|&s| s.contains(c)) {
            return c;
        }
    }

    panic!("Cannot find common char for {:?}", s);
}

fn priority(c: char) -> i32 {
    if c.is_ascii_lowercase() {
        (c as i32) + 1 - ('a' as i32)
    } else if c.is_ascii_uppercase() {
        (c as i32) + 27 - ('A' as i32)
    } else {
        panic!("{} is not a alphabetic letter", c)
    }
}

fn part1(s: &str) -> i32 {

    let mut priority_sum = 0;

    for line in s.lines() {
        let split = split(line);
        let common = find_common(split);
        priority_sum += priority(common);
    }

    priority_sum
}

fn part2(s: &str) -> i32 {
    s.lines()
        .into_iter()
        .map(|s| s.into())
        .collect::<Vec<String>>()
        .chunks(3)
        .into_iter()
        .map(|l| {
            let l: Vec<&str> = l.into_iter().map(|line| line.as_str()).collect();
            let c = find_common(l);
            priority(c)
        })
        .sum::<i32>()
}

fn main() {
    let input = fs::read_to_string("./input.txt").unwrap();

    println!("part 1 => {}", part1(&input));
    println!("part 2 => {}", part2(&input));
}

#[test]
fn test_find_common() {
    assert_eq!('p', find_common(split("vJrwpWtwJgWrhcsFMMfFFhFp")));
    assert_eq!('L', find_common(split("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")));
    assert_eq!('P', find_common(split("PmmdzqPrVvPwwTWBwg")));
    assert_eq!('v', find_common(split("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn")));
    assert_eq!('t', find_common(split("ttgJtRGJQctTZtZT")));
    assert_eq!('s', find_common(split("CrZsJsPPZsGzwwsLwLmpwMDw")));
}

#[test]
fn test_priority() {
    assert_eq!(1, priority('a'));
    assert_eq!(2, priority('b'));
    assert_eq!(26, priority('z'));
    assert_eq!(27, priority('A'));
    assert_eq!(52, priority('Z'));
}
