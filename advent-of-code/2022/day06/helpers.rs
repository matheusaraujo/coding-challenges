use std::collections::HashSet;

pub fn find_marker(signal: &str, marker_len: usize) -> usize {
    let chars: Vec<char> = signal.chars().collect();
    chars
        .windows(marker_len)
        .position(|window| window.iter().collect::<HashSet<_>>().len() == marker_len)
        .map(|i| i + marker_len)
        .expect("no marker found")
}
