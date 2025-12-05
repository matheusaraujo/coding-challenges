pub fn string_to_int(s: &str) -> i32 {
    s.parse::<i32>().expect("Failed to parse string to integer")
}

pub fn parse_input(puzzle_input: &[String]) -> (Vec<i32>, Vec<i32>) {
    let mut left: Vec<i32> = Vec::new();
    let mut right: Vec<i32> = Vec::new();

    for line in puzzle_input {
        let parts: Vec<&str> = line.split_whitespace().collect();

        if parts.len() >= 2 {
            left.push(string_to_int(parts[0]));
            right.push(string_to_int(parts[1]));
        }
    }

    left.sort();
    right.sort();

    (left, right)
}
