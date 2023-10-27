use anyhow::Result;

fn main() -> Result<()> {
    let input = std::fs::read_to_string("src/inputs/day-11.txt")?;

    let monkeys = MonkeyFactory::create_from_string(&input);

    println!("{:?}", part1(monkeys.clone()));
    println!("{:?}", part2(monkeys.clone()));

    Ok(())
}

fn part1(monkeys: Vec<Monkey>) -> Option<usize> {
    Some(simulate(monkeys, 20, |x| x / 3))
}

fn part2(monkeys: Vec<Monkey>) -> Option<usize> {
    let base: u64 = monkeys.iter().map(|m| m.divide_by).product();

    Some(simulate(monkeys, 10000, |x| x % base))
}

fn simulate(mut monkeys: Vec<Monkey>, rounds: u64, relieve_func: impl Fn(u64) -> u64) -> usize {
    let mut inspections: Vec<usize> = vec![0; monkeys.len()];

    for _ in 0..rounds {
        for i in 0..(monkeys.len()) {
            let monkey = &mut monkeys[i];

            let items: Vec<u64> = monkey
                .items
                .drain(..)
                .map(|item| relieve_func(monkey.operation.apply(item)))
                .collect();

            inspections[i] += items.len();

            let divide_by = monkey.divide_by;
            let next_if_truthy = monkey.next_if_truthy;
            let next_if_falsy = monkey.next_if_falsy;

            for item in items {
                let next;

                if item % divide_by == 0 {
                    next = next_if_truthy;
                } else {
                    next = next_if_falsy;
                }

                monkeys[next].items.push(item);
            }
        }
    }

    inspections.sort();

    inspections
        .iter()
        .rev()
        .take(2)
        .product()
}

struct MonkeyFactory {}

impl MonkeyFactory {
    fn create_from_string(s: &str) -> Vec<Monkey> {
        s
            .split("\n\n")
            .into_iter()
            .filter_map(|l| MonkeyFactory{}.create_monkey(l))
            .collect()
    }

    fn create_monkey(&self, s: &str) -> Option<Monkey> {
        let mut lines = s.lines().skip(1);

        Some(Monkey {
            items: self.parse_items(lines.next()?)?,
            operation: self.parse_operation(lines.next()?)?,
            divide_by: self.get_last_number_string(lines.next()?).parse().unwrap(),
            next_if_truthy: self.get_last_number_string(lines.next()?).parse().unwrap(),
            next_if_falsy: self.get_last_number_string(lines.next()?).parse().unwrap(),
        })
    }

    fn parse_items(&self, s: &str) -> Option<Vec<u64>> {
        Some(s
            .split_once(":")?
            .1
            .split(',')
            .filter_map(|x| x.trim().parse().ok())
            .collect())
    }

    fn parse_operation(&self, s: &str) -> Option<Operation> {
        let mut operation_line = s
            .split_once("= ")?
            .1
            .split_ascii_whitespace();

        let operand_a = operation_line.next().map(Operand::from_str)??;

        let operator = operation_line.next()?;

        let operand_b = operation_line.next().map(Operand::from_str)??;

        Some(match operator {
            "+" => Operation::Add(operand_a, operand_b),
            "*" => Operation::Multiply(operand_a, operand_b),
            _ => unreachable!(),
        })
    }

    fn get_last_number_string<'a>(&'a self, s: &'a str) -> &'a str {
        s.split_ascii_whitespace().last().unwrap()
    }
}

#[derive(Debug, PartialEq, Clone)]
enum Operand {
    Old,
    Num(u64),
}

impl Operand {
    fn from_str(s: &str) -> Option<Self> {
        match s {
            "old" => Some(Operand::Old),
            x => Some(Operand::Num(x.parse().ok()?)),
        }
    }

    fn apply(&self, x: u64) -> u64 {
        match self {
            Operand::Old => x,
            Operand::Num(y) => *y,
        }
    }
}

#[derive(Debug, PartialEq, Clone)]
enum Operation {
    Add(Operand, Operand),
    Multiply(Operand, Operand),
}

impl Operation {
    fn apply(&self, x: u64) -> u64 {
        match self {
            Self::Add(a, b) => a.apply(x) + b.apply(x),
            Self::Multiply(a, b) => a.apply(x) * b.apply(x),
        }
    }
}

#[derive(Debug, Clone)]
struct Monkey {
    items: Vec<u64>,
    operation: Operation,
    divide_by: u64,
    next_if_truthy: usize,
    next_if_falsy: usize,
}

#[test]
fn test_monkey_from_str() {
    let input = concat!("Monkey 0:\n",
        "  Starting items: 79, 98\n",
        "  Operation: new = old * 19\n",
        "Test: divisible by 23\n",
        "If true: throw to monkey 2\n",
        "If false: throw to monkey 3"
    );

    let result = &MonkeyFactory::create_from_string(&input)[0];

    assert_eq!(vec![79, 98], result.items);
    assert_eq!(
        Operation::Multiply(Operand::Old, Operand::Num(19)),
        result.operation
    );
    assert_eq!(23, result.divide_by);
    assert_eq!(2, result.next_if_truthy);
    assert_eq!(3, result.next_if_falsy);
}

#[test]
fn test_simulate_20() {
    let input = std::fs::read_to_string("src/examples/day-11.txt").unwrap();

    let monkeys = MonkeyFactory::create_from_string(&input);
    let result = part1(monkeys).unwrap();

    assert_eq!(10605, result);
}

#[test]
fn test_simulate_10_000() {
    let input = std::fs::read_to_string("src/examples/day-11.txt").unwrap();

    let monkeys = MonkeyFactory::create_from_string(&input);
    let result = part2(monkeys).unwrap();

    assert_eq!(2713310158, result);
}
