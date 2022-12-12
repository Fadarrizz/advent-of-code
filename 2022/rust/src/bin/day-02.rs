use std::fs;

const ROCK: i32 = 1;
const PAPER: i32 = 2;
const SCISSORS: i32 = 3;

const LOST: i32 = 0;
const DRAW: i32 = 3;
const WIN: i32 = 6;

fn to_shape(s: &str) -> i32 {
    match s {
        "A" | "X" => ROCK,
        "B" | "Y" => PAPER,
        "C" | "Z" => SCISSORS,
        _ => panic!("Cannot match {}", s),
    }
}

fn to_desired_outcome(s : &str) -> i32 {
    match s {
        "X" => LOST,
        "Y" => DRAW,
        "Z" => WIN,
        _ => panic!("Cannot match {}", s),
    }
}

fn round_outcome(me: i32, opponent: i32) -> i32 {
    match (me, opponent) {
        (ROCK, SCISSORS) => WIN,
        (PAPER, ROCK) => WIN,
        (SCISSORS, PAPER) => WIN,
        (a, b) if a == b => DRAW,
        (_, _) => LOST,
    }
}

fn play_outcome(desired: i32, theirs: i32) -> i32 {
    match desired {
        WIN => match theirs {
            ROCK => PAPER,
            PAPER => SCISSORS,
            SCISSORS => ROCK,
            _ => panic!(),
        },
        LOST => match theirs {
            ROCK => SCISSORS,
            PAPER => ROCK,
            SCISSORS => PAPER,
            _ => panic!(),
        },
        DRAW => theirs,
        _ => panic!(),
    }
}

fn part1(s: &str) -> i32 {
    let mut score = 0;
    for line in s.lines() {
        let moves: Vec<&str> = line.split(" ").collect();
        let theirs = to_shape(moves[0]);
        let ours = to_shape(moves[1]);

        score += ours + round_outcome(ours, theirs);
    }

    score
}

fn part2(s: &str) -> i32 {
    let mut score = 0;
    for line in s.lines() {
        let moves: Vec<&str> = line.split(" ").collect();
        let theirs = to_shape(moves[0]);
        let outcome = to_desired_outcome(moves[1]);
        let ours = play_outcome(outcome, theirs);

        score += ours + outcome;
    }

    score
}

fn main() {
    let input = fs::read_to_string("./input.txt").unwrap();

    println!("part 1 => {}", part1(&input));
    println!("part 2 => {}", part2(&input));
}
