fn main() {
    let input = include_str!("../inputs/day-02.txt");

    let presents: Vec<Present> = input
        .lines()
        .map(Present::parse)
        .collect();

    println!("{}", part1(&presents));
    println!("{}", part2(&presents));
}

fn part1(presents: &[Present]) -> u32 {
    presents
        .iter()
        .map(|p| p.paper_required())
        .sum()
}

fn part2(presents: &[Present]) -> u32 {
    presents
        .iter()
        .map(|p| p.ribbon_required())
        .sum()
}

#[derive(Debug, PartialEq, Eq, Clone, Copy)]
struct Present {
    l: u32,
    w: u32,
    h: u32,
}

impl Present {
    fn parse(s: &str) -> Present {
        let mut dimensions: Vec<u32> = s
            .split("x")
            .map(|c| c.parse::<u32>().unwrap())
            .collect();

        dimensions.sort();

        Present {
            l: dimensions[0],
            w: dimensions[1],
            h: dimensions[2],
        }
    }

    fn paper_required(&self) -> u32 {
        3 * self.l * self.w + 2 * self.w * self.h + 2 * self.h * self.l
    }

    fn ribbon_required(&self) -> u32 {
        2 * (self.l + self.w) + (self.l * self.w) * self.h
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_paper_required() {
        let present = Present { l: 2, w: 3, h: 4 };
        assert_eq!(58, present.paper_required());

        let present = Present { l: 1, w: 1, h: 10 };
        assert_eq!(43, present.paper_required());
    }

    #[test]
    fn test_ribbon_required() {
        let present = Present { l: 2, w: 3, h: 4 };
        assert_eq!(34, present.ribbon_required());

        let present = Present { l: 1, w: 1, h: 10 };
        assert_eq!(14, present.ribbon_required());
    }
}
