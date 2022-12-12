use std::fs;
use regex::Regex;

type Stack = Vec<char>;

#[derive(Debug, Clone)]
struct Instruction(usize, usize, usize);

fn build_stacks(s: &str) -> Vec<Stack> {
    let mut stacks = Vec::new();
    for line in s.lines().rev() {
        for (i, c) in line.chars().enumerate() {
            if c.is_alphabetic() {
                let idx = (i - 1) / 4;

                if idx >= stacks.len() {
                    stacks.push(Stack::new());
                }
                stacks[idx].push(c);
            }
        }
    }

    stacks
}

fn parse_instructions(s: &str) -> Vec<Instruction> {
    let mut instructions: Vec<Instruction> = Vec::new();
    for line in s.lines() {
        let r = Regex::new(r"^move ([0-9]+) from ([0-9]+) to ([0-9]+)$")
            .unwrap()
            .captures(line)
            .unwrap();

        instructions.push(
            Instruction(
                r[1].parse().unwrap(), 
                r[2].parse().unwrap(), 
                r[3].parse().unwrap(),
            )
        );
    }

    instructions
}

fn top_crates(stacks: &Vec<Stack>) -> String {
    stacks
        .iter()
        .map(|stack| stack.last().unwrap())
        .collect()
}

fn part1(mut stacks: Vec<Stack>, instructions: Vec<Instruction>) -> String {
    for instruction in instructions {
        let (amount, source, target) = match instruction {
            Instruction(amount, source, target) => (amount, source, target)
        };
        for _ in 0..amount {
            let c = stacks[source - 1].pop().unwrap();
            stacks[target - 1].push(c);
        }
    }

    top_crates(&stacks)
}

fn part2(mut stacks: Vec<Stack>, instructions: Vec<Instruction>) -> String {
    for instruction in instructions {
        let (amount, source, target) = match instruction {
            Instruction(amount, source, target) => (amount, source, target)
        };
        let len = stacks[source - 1].len();
        let mut c = stacks[source - 1]
            .drain((len - amount)..)
            .collect();
        stacks[target - 1].append(&mut c);
    }

    top_crates(&stacks)
}

fn main() {
    let start = fs::read_to_string("./start.txt").unwrap();
    let stacks: Vec<Stack> = build_stacks(&start);

    let input = fs::read_to_string("./input.txt").unwrap();
    let instructions = parse_instructions(&input);

    println!("part 1 => {}", part1(stacks.clone(), instructions.clone()));
    println!("part 2 => {}", part2(stacks.clone(), instructions.clone()));
}

#[test]
fn test_top_crates() {
    let stacks = vec![vec!('C'), vec!('M'), vec!['P', 'D', 'N', 'Z']];
    assert_eq!("CMZ", top_crates(&stacks));
}
