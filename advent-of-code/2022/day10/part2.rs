use crate::helpers::{parse_ocr, x_during_cycles};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let history = x_during_cycles(puzzle_input);

    let rows: Vec<String> = history
        .chunks(40)
        .map(|row| {
            row.iter()
                .enumerate()
                .map(|(col, &x)| {
                    if (col as i32 - x).abs() <= 1 {
                        '#'
                    } else {
                        '.'
                    }
                })
                .collect()
        })
        .collect();

    any(parse_ocr(&rows))
}
