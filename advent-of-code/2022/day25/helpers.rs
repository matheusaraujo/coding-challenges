pub fn snafu_to_decimal(s: &str) -> i64 {
    s.chars().fold(0i64, |acc, c| {
        let digit = match c {
            '2' => 2,
            '1' => 1,
            '0' => 0,
            '-' => -1,
            '=' => -2,
            _ => panic!("invalid SNAFU digit: {c}"),
        };
        acc * 5 + digit
    })
}

pub fn decimal_to_snafu(mut n: i64) -> String {
    let mut digits = Vec::new();
    while n != 0 {
        let rem = n.rem_euclid(5);
        let (digit_char, carry) = match rem {
            0 => ('0', 0),
            1 => ('1', 0),
            2 => ('2', 0),
            3 => ('=', 1),
            4 => ('-', 1),
            _ => unreachable!(),
        };
        digits.push(digit_char);
        n = n / 5 + carry;
    }
    if digits.is_empty() {
        digits.push('0');
    }
    digits.iter().rev().collect()
}
