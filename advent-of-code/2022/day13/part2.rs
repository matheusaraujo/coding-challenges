use crate::helpers::{Packet, parse_packet};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let divider1 = parse_packet("[[2]]");
    let divider2 = parse_packet("[[6]]");

    let mut packets: Vec<Packet> = puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .map(|line| parse_packet(line))
        .collect();
    packets.push(divider1.clone());
    packets.push(divider2.clone());
    packets.sort();

    let pos1 = packets.iter().position(|p| p == &divider1).unwrap() + 1;
    let pos2 = packets.iter().position(|p| p == &divider2).unwrap() + 1;
    any(pos1 * pos2)
}
