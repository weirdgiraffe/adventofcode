mod inputs;

use anyhow::Result;

fn to_digit(iter: &mut dyn Iterator<Item = char>) -> Option<u32> {
    while let Some(v) = iter.next().map(|c| c.to_digit(10)) {
        if let Some(d) = v {
            return Some(d);
        }
    }
    None
}

fn number(s: &str) -> u32 {
    let first = to_digit(&mut s.chars()).expect("must have first digit");
    let last = to_digit(&mut s.chars().rev()).expect("must have last digit");
    first * 10 + last
}

fn part1<T: Iterator<Item = String>>(it: T) -> u32 {
    it.map(|line| number(&line)).sum::<u32>()
}

fn main() -> Result<()>{
    let lines = inputs::read_lines("input01.txt")?;
    let sum = part1(lines.flatten());
    println!("sum: {}", sum);
    Result::Ok(())
}

#[cfg(test)]
mod tests {
    use crate::part1;
    use std::io::{self, BufRead};

    #[test]
    fn short() {
        let cursor = io::Cursor::new(
            vec![
                "1abc2".to_string(),
                "pqr3stu8vwx".to_string(),
                "a1b2c3d4e5f".to_string(),
                "treb7uchet".to_string(),
            ]
            .join("\n"),
        );
        let sum = part1(cursor.lines().map(|line| line.unwrap()));
        assert_eq!(sum, 142);
    }
}
