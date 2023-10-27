use crypto::md5::Md5;
use crypto::digest::Digest;

fn main() {
    let input =  "ckczppom";

    println!("{}", part1(&input));
    println!("{}", part2(&input));
}

fn part1(input: &str) -> u64 {
    try_md5_hash_func(input, |output| {
        output[0] as u32 + output[1] as u32 + (output[2] >> 4) as u32 == 0
    })
}

fn part2(input: &str) -> u64 {
    try_md5_hash_func(input, |output| {
        output[0] as u32 + output[1] as u32 + output[2] as u32 == 0
    })
}

fn try_md5_hash_func(input: &str, func: fn(&[u8; 16]) -> bool) -> u64 {
    let mut hasher = Md5::new();
    let mut output = [0; 16];

    for i in 0..u64::MAX {
        hasher.input(input.as_bytes());
        hasher.input(i.to_string().as_bytes());

        hasher.result(&mut output);

        if func(&output) {
            return i;
        }

        hasher.reset();
    }

    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("abcdef"), 609043);

        assert_eq!(part1("pqrstuv"), 1048970);
    }
}
