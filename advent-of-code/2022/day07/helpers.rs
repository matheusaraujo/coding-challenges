use std::collections::HashMap;

pub fn directory_sizes(puzzle_input: &[String]) -> HashMap<Vec<String>, u64> {
    let mut dir_sizes: HashMap<Vec<String>, u64> = HashMap::new();
    let mut stack: Vec<String> = Vec::new();

    for line in puzzle_input.iter().filter(|line| !line.is_empty()) {
        if let Some(target) = line.strip_prefix("$ cd ") {
            match target {
                "/" => stack = vec!["/".to_string()],
                ".." => {
                    stack.pop();
                }
                name => stack.push(name.to_string()),
            }
        } else if line == "$ ls" || line.starts_with("dir ") {
            continue;
        } else {
            let size: u64 = line
                .split_whitespace()
                .next()
                .and_then(|token| token.parse().ok())
                .expect("file size");
            for i in 1..=stack.len() {
                *dir_sizes.entry(stack[..i].to_vec()).or_insert(0) += size;
            }
        }
    }

    dir_sizes
}
