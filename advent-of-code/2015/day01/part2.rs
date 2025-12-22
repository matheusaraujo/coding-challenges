use crate::{Answer, answer};

pub fn part2(puzzle_input: &[String]) -> Answer {
    let mut floor = 1;

    for (i, c) in puzzle_input[0].chars().enumerate() {
        match c {
            '(' => floor += 1,
            ')' => floor -= 1,
            _ => {}
        }

        if floor == -1 {
            return answer(i);
        }
    }

    answer(-1)
}
