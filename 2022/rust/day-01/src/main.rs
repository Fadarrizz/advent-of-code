use std::fs;

fn sum_calories_per_elf(input: &String) -> Vec<i32> {
    let mut elves = Vec::new();

    let mut sum = 0;
    for line in input.lines() {
        if line.is_empty() {
            elves.push(sum);
            sum = 0;
            continue;
        }

        sum += line.parse::<i32>().unwrap();
    }
    elves.push(sum);

    elves.sort();
    elves.reverse();

    elves
}

fn main() {
    let input = fs::read_to_string("./input.txt").unwrap();

    let elves = sum_calories_per_elf(&input);

    println!("part 1 = {}", elves.iter().max().unwrap());
    println!("part 2 = {}", elves[..3].iter().sum::<i32>());
}
