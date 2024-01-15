mod inputs;

use anyhow::Result;
use std::ops::ControlFlow;

fn first_digit<'a, T>(mut iter: T) -> Option<u32>
where
    T: Iterator<Item = (&'a str, u8)>,
{
    const DIGITS: [&str; 9] = [
        "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];
    let result = iter.try_fold((), |_, (s, b)| {
        if let Some(digit) = (b as char).to_digit(10) {
            return ControlFlow::Break(digit);
        }
        if let Some(index) = DIGITS.iter().position(|&w| s.starts_with(w)) {
            return ControlFlow::Break((index as u32) + 1);
        }
        ControlFlow::Continue(())
    });
    match result {
        ControlFlow::Break(digit) => Some(digit),
        ControlFlow::Continue(_) => None,
    }
}

fn number(s: &str) -> u32 {
    let first: u32 = first_digit(s.bytes().enumerate().map(|(i, b)| (&s[i..], b)))
        .expect("must have first digit");
    let last: u32 = first_digit(s.bytes().enumerate().rev().map(|(i, b)| (&s[i..], b)))
        .expect("must have last digit");
    (first * 10) + last
}

fn part2<T: Iterator<Item = String>>(it: T) -> u32 {
    it.map(|line| number(&line)).sum()
}

fn main() -> Result<()>{
    let lines = inputs::read_lines("input01.txt")?;
    let sum = part2(lines.flatten());
    println!("sum: {}", sum);
    Result::Ok(())
}

#[cfg(test)]
mod tests {
    use crate::part2;
    use std::io::{self, BufRead};

    #[test]
    fn short() {
        let cursor = io::Cursor::new(
            vec![
                "two1nine".to_string(),
                "eightwothree".to_string(),
                "abcone2threexyz".to_string(),
                "xtwone3four".to_string(),
                "4nineeightseven2".to_string(),
                "zoneight234".to_string(),
                "7pqrstsixteen".to_string(),
            ]
            .join("\n"),
        );
        let sum = part2(cursor.lines().map(|line| line.unwrap()));
        assert_eq!(sum, 281);
    }
}
