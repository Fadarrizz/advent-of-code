use anyhow::Result;

type Instructions = Vec<Operation>;
type Image<'a> = Vec<Vec<&'a str>>;

#[derive(Debug, Clone, PartialEq)]
enum Operation {
    Noop,
    Add(i32),
}

fn main() -> Result<()> {
    let input = std::fs::read_to_string("src/inputs/day-10.txt")?;

    let parsed = parse(&input);

    let (signal_strengths, image) = solve(&parsed);

    println!("{:?}", signal_strengths);

    let image = image
        .iter()
        .map(|line| line.join(""))
        .collect::<Vec<String>>()
        .join("\n");

    println!("{}", image);

    Ok(())
}

fn parse(s: &str) -> Instructions {
    s
        .lines()
        .map(|line| {
            if line.starts_with("noop") {
                Operation::Noop
            } else {
                let (_, x) = line.split_once(' ').unwrap();

                Operation::Add(x.parse::<i32>().unwrap())
            }
        })
        .collect()
}

fn solve(instructions: &Instructions) -> (i32, Image) {
    let mut strength_sum = 0;
    let mut image: Image = vec![vec!["."; 40]; 6];

    let mut cycle = 1;
    let mut x = 1;
    for operation in instructions {
        match operation {
            Operation::Noop => {
                draw_pixel(&mut image, &cycle, &x);
                increase_cycle(&mut strength_sum, &mut cycle, &mut x);
            },
            Operation::Add(n) => {
                draw_pixel(&mut image, &cycle, &x);
                increase_cycle(&mut strength_sum, &mut cycle, &mut x);

                draw_pixel(&mut image, &cycle, &x);
                increase_cycle(&mut strength_sum, &mut cycle, &mut x);

                x += n;
            }
        }
    }

    (strength_sum, image)
}

fn increase_cycle(strength_sum: &mut i32, cycle: &mut i32, x: &i32) {
    if *cycle == 20 || *cycle % 40 == 20 {
        *strength_sum += *cycle * *x;
    }

    *cycle += 1;
}

fn draw_pixel(image: &mut Image, cycle: &i32, x: &i32) {
    let row = ((cycle - 1) / 40) % 6;
    let col = (cycle - 1) % 40;

    if x - 1 <= col && col <= x + 1 {
        image[row as usize][col as usize] = "#";
    }
}
