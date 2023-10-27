fn main() {
    let input = include_str!("../inputs/day-01.txt");

    let mut floor = 0;
    for (i, c) in input.chars().enumerate() {
        match c {
            '(' => floor += 1,
            ')' => floor -= 1,
            '\n' => (),// ends here,
            _ => panic!("undefined char"),
        }

        if floor == -1 {
            println!("{}", i + 1);
            break;
        }
    }

    println!("{floor}");
}
