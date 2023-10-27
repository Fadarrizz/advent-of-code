fn main() {
    let input = include_str!("../inputs/day-05.txt");

    println!("{}", part1(&input));
}

fn part1(input: &str) -> usize {
    input
        .lines()
        .filter(|s| is_nice_string(s))
        .count()
}

fn part2(input: &str) -> usize {
    input
        .lines()
}

fn is_nice_string(s: &str) -> bool {
    has_at_least_three_vowels(s)
    && has_repeat(s)
    && does_not_contain(s)
}

fn has_at_least_three_vowels(s: &str) -> bool {
    let vowels = ['a', 'e', 'i', 'o', 'u'];

    s
        .chars()
        .filter(|c| vowels.contains(&c))
        .count() >= 3
}

fn has_repeat(s: &str) -> bool {
    for (a, b) in s.chars().zip(s.chars().skip(1)) {
        if a == b {
            return true;
        }
    }

    false
}

fn does_not_contain(s: &str) -> bool {
    let strings = vec!["ab", "cd", "pq", "xy"];

    strings.iter().all(|&string| !s.contains(string))
}

fn is_nicer(s: &str) -> bool {
    true
}

fn has_pair_twice(s: &str) -> bool {
    let slice = &s[..];
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = concat!(
            "ugknbfddgicrmopn\n",
            "aaa\n",
            "jchzalrnumimnmhp\n",
            "haegwjzuvuyypxyu\n",
            "dvszwmarrgswjxmb\n"
        );

        assert_eq!(2, part1(&input));
    }

    #[test]
    fn test_has_at_least_three_vowels() {
        assert_eq!(true, has_at_least_three_vowels("ugknbfddgicrmopn"));
        assert_eq!(true, has_at_least_three_vowels("aaa"));
        assert_eq!(true, has_at_least_three_vowels("jchzalrnumimnmhp"));
        assert_eq!(true, has_at_least_three_vowels("haegwjzuvuyypxyu"));
        assert_eq!(false, has_at_least_three_vowels("dvszwmarrgswjxmb"));
    }

    #[test]
    fn test_has_repeat() {
        assert_eq!(true, has_repeat("ugknbfddgicrmopn"));
        assert_eq!(true, has_repeat("aaa"));
        assert_eq!(false, has_repeat("jchzalrnumimnmhp"));
        assert_eq!(true, has_repeat("haegwjzuvuyypxyu"));
        assert_eq!(true, has_repeat("dvszwmarrgswjxmb"));
    }

    #[test]
    fn test_does_not_contain() {
        assert_eq!(true, does_not_contain("ugknbfddgicrmopn"));
        assert_eq!(true, does_not_contain("aaa"));
        assert_eq!(true, does_not_contain("jchzalrnumimnmhp"));
        assert_eq!(false, does_not_contain("haegwjzuvuyypxyu"));
        assert_eq!(true, does_not_contain("dvszwmarrgswjxmb"));
    }
}
