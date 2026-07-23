use std::collections::HashSet;

pub fn priority(c: char) -> u32 {
    if c.is_ascii_lowercase() {
        c as u32 - 'a' as u32 + 1
    } else {
        c as u32 - 'A' as u32 + 27
    }
}

pub fn common_char(groups: &[&str]) -> char {
    let mut sets = groups.iter().map(|s| s.chars().collect::<HashSet<char>>());
    let first = sets.next().expect("at least one group");
    let common = sets.fold(first, |acc, s| acc.intersection(&s).copied().collect());
    *common.iter().next().expect("a common item exists")
}
