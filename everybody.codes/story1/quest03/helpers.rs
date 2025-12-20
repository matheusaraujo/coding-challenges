#[derive(Debug, Clone)]
pub struct Snail {
    pub x: i128,
    pub y: i128,
}

impl Snail {
    fn from_string(input: &str) -> Self {
        let parts: Vec<&str> = input.split_whitespace().collect();
        let x = parts[0].trim_start_matches("x=").parse::<i128>().unwrap();
        let y = parts[1].trim_start_matches("y=").parse::<i128>().unwrap();
        Snail { x, y }
    }

    pub fn disc_size(&self) -> i128 {
        self.x + self.y - 1
    }

    fn days_until_y1(&self) -> i128 {
        let s = self.disc_size();
        (s - self.x).rem_euclid(s)
    }
}

pub struct SnailClock {
    pub snails: Vec<Snail>,
}

impl SnailClock {
    pub fn from_input(input: &[String]) -> Self {
        let snails = input
            .iter()
            .filter(|l| !l.trim().is_empty())
            .map(|l| Snail::from_string(l))
            .collect();
        SnailClock { snails }
    }

    pub fn find_first_alignment(&self) -> i128 {
        let conditions: Vec<(i128, i128)> = self
            .snails
            .iter()
            .map(|s| (s.days_until_y1(), s.disc_size()))
            .collect();

        self.solve_congruences(conditions)
    }

    fn solve_congruences(&self, conditions: Vec<(i128, i128)>) -> i128 {
        let (mut a1, mut n1) = conditions[0];
        for &(a2, n2) in &conditions[1..] {
            let (g, x, _) = self.extended_gcd(n1, n2);
            if (a2 - a1) % g != 0 {
                panic!("No solution");
            }

            let lcm = (n1 * n2) / g;
            let k = ((a2 - a1) / g * x).rem_euclid(n2 / g);
            a1 = (a1 + k * n1).rem_euclid(lcm);
            n1 = lcm;
        }
        a1
    }

    fn extended_gcd(&self, a: i128, b: i128) -> (i128, i128, i128) {
        if a == 0 {
            (b, 0, 1)
        } else {
            let (g, x1, y1) = self.extended_gcd(b % a, a);
            (g, y1 - (b / a) * x1, x1)
        }
    }
}
