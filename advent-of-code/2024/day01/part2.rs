use super::helpers::parse_input;
use std::collections::HashMap;

pub fn part2(puzzle_input: &[String]) -> String {
    let (left, right) = parse_input(puzzle_input);

    let mut count: HashMap<i32, i32> = HashMap::new();

    for x in left {
        *count.entry(x).or_insert(0) += 1;
    }

    let mut result: i32 = 0;

    for item in right {
        if let Some(count_value) = count.get(&item) {
            result += item * count_value;
        }
    }

    result.to_string()
}
