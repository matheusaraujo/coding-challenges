use std::cmp::Ordering;
use std::iter::Peekable;
use std::str::Chars;

#[derive(Clone, Eq, PartialEq)]
pub enum Packet {
    Int(u32),
    List(Vec<Packet>),
}

impl Ord for Packet {
    fn cmp(&self, other: &Self) -> Ordering {
        match (self, other) {
            (Packet::Int(a), Packet::Int(b)) => a.cmp(b),
            (Packet::List(a), Packet::List(b)) => a.cmp(b),
            (Packet::Int(a), Packet::List(b)) => [Packet::Int(*a)][..].cmp(b),
            (Packet::List(a), Packet::Int(b)) => a[..].cmp(&[Packet::Int(*b)]),
        }
    }
}

impl PartialOrd for Packet {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

fn parse(chars: &mut Peekable<Chars>) -> Packet {
    if chars.peek() == Some(&'[') {
        chars.next();
        let mut items = Vec::new();
        loop {
            match chars.peek() {
                Some(']') => {
                    chars.next();
                    break;
                }
                Some(',') => {
                    chars.next();
                }
                Some(_) => items.push(parse(chars)),
                None => break,
            }
        }
        Packet::List(items)
    } else {
        let mut digits = String::new();
        while let Some(&c) = chars.peek() {
            if c.is_ascii_digit() {
                digits.push(c);
                chars.next();
            } else {
                break;
            }
        }
        Packet::Int(digits.parse().expect("packet integer"))
    }
}

pub fn parse_packet(line: &str) -> Packet {
    parse(&mut line.chars().peekable())
}

pub fn parse_pairs(puzzle_input: &[String]) -> Vec<(Packet, Packet)> {
    puzzle_input
        .split(|line| line.is_empty())
        .filter(|block| !block.is_empty())
        .map(|block| (parse_packet(&block[0]), parse_packet(&block[1])))
        .collect()
}
