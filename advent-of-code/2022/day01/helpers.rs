pub fn elf_calorie_totals(puzzle_input: &[String]) -> Vec<u32> {
    puzzle_input
        .split(|line| line.is_empty())
        .map(|elf| elf.iter().filter_map(|line| line.parse::<u32>().ok()).sum())
        .collect()
}
