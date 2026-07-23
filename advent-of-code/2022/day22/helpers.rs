use std::collections::{HashMap, HashSet, VecDeque};

pub enum Instr {
    Move(u32),
    Turn(char),
}

type Vec3 = (i32, i32, i32);

fn neg(v: Vec3) -> Vec3 {
    (-v.0, -v.1, -v.2)
}

fn dot(a: Vec3, b: Vec3) -> i32 {
    a.0 * b.0 + a.1 * b.1 + a.2 * b.2
}

struct Face {
    block: (usize, usize),
    right: Vec3,
    down: Vec3,
    normal: Vec3,
}

pub struct Board {
    grid: Vec<Vec<u8>>,
    side: usize,
    faces: Vec<Face>,
    face_at: HashMap<(usize, usize), usize>,
}

pub fn parse_board_and_path(puzzle_input: &[String]) -> (Board, Vec<Instr>) {
    let blank_idx = puzzle_input
        .iter()
        .position(|line| line.is_empty())
        .expect("blank separator line");
    let grid_lines = &puzzle_input[..blank_idx];
    let path_line = &puzzle_input[blank_idx + 1];

    let width = grid_lines.iter().map(|line| line.len()).max().unwrap();
    let grid: Vec<Vec<u8>> = grid_lines
        .iter()
        .map(|line| {
            let mut row = line.as_bytes().to_vec();
            row.resize(width, b' ');
            row
        })
        .collect();

    let filled_count = grid.iter().flatten().filter(|&&c| c != b' ').count();
    let side = ((filled_count / 6) as f64).sqrt().round() as usize;

    let block_rows = grid.len().div_ceil(side);
    let block_cols = width.div_ceil(side);
    let filled: HashSet<(usize, usize)> = (0..block_rows)
        .flat_map(|br| (0..block_cols).map(move |bc| (br, bc)))
        .filter(|&(br, bc)| grid[br * side][bc * side] != b' ')
        .collect();

    let (faces, face_at) = build_faces(&filled);

    let instrs = parse_path(path_line);
    (
        Board {
            grid,
            side,
            faces,
            face_at,
        },
        instrs,
    )
}

fn parse_path(line: &str) -> Vec<Instr> {
    let mut instrs = Vec::new();
    let mut num = String::new();
    for c in line.trim().chars() {
        if c.is_ascii_digit() {
            num.push(c);
        } else {
            if !num.is_empty() {
                instrs.push(Instr::Move(num.parse().unwrap()));
                num.clear();
            }
            instrs.push(Instr::Turn(c));
        }
    }
    if !num.is_empty() {
        instrs.push(Instr::Move(num.parse().unwrap()));
    }
    instrs
}

fn build_faces(filled: &HashSet<(usize, usize)>) -> (Vec<Face>, HashMap<(usize, usize), usize>) {
    let start = *filled.iter().min().unwrap();
    let mut faces = vec![Face {
        block: start,
        right: (1, 0, 0),
        down: (0, 1, 0),
        normal: (0, 0, 1),
    }];
    let mut face_at = HashMap::new();
    face_at.insert(start, 0);

    let mut queue = VecDeque::new();
    queue.push_back(start);

    while let Some(block) = queue.pop_front() {
        let idx = face_at[&block];
        let (right, down, normal) = (faces[idx].right, faces[idx].down, faces[idx].normal);
        let (br, bc) = block;

        let neighbors = [
            ((br.wrapping_sub(1), bc), (right, normal, neg(down))),
            ((br + 1, bc), (right, neg(normal), down)),
            ((br, bc.wrapping_sub(1)), (normal, down, neg(right))),
            ((br, bc + 1), (neg(normal), down, right)),
        ];

        for (block_pos, (nright, ndown, nnormal)) in neighbors {
            if filled.contains(&block_pos) && !face_at.contains_key(&block_pos) {
                let nidx = faces.len();
                faces.push(Face {
                    block: block_pos,
                    right: nright,
                    down: ndown,
                    normal: nnormal,
                });
                face_at.insert(block_pos, nidx);
                queue.push_back(block_pos);
            }
        }
    }

    (faces, face_at)
}

impl Board {
    fn cell(&self, face_idx: usize, r: usize, c: usize) -> u8 {
        let block = self.faces[face_idx].block;
        self.grid[block.0 * self.side + r][block.1 * self.side + c]
    }

    fn cross_edge(
        &self,
        face_idx: usize,
        local_r: usize,
        local_c: usize,
        facing: usize,
    ) -> (usize, usize, usize, usize) {
        let n = self.side;
        let f = &self.faces[face_idx];

        let v_exit = match facing {
            0 => f.right,
            1 => f.down,
            2 => neg(f.right),
            3 => neg(f.down),
            _ => unreachable!(),
        };

        let dest_idx = self
            .faces
            .iter()
            .position(|ff| ff.normal == v_exit)
            .unwrap();
        let dest = &self.faces[dest_idx];
        let v_enter = neg(f.normal);

        let new_facing = if v_enter == dest.right {
            0
        } else if v_enter == dest.down {
            1
        } else if v_enter == neg(dest.right) {
            2
        } else if v_enter == neg(dest.down) {
            3
        } else {
            unreachable!()
        };

        let (tangent_f, along_f) = if facing == 0 || facing == 2 {
            (f.down, local_r)
        } else {
            (f.right, local_c)
        };
        let tangent_dest = if new_facing == 0 || new_facing == 2 {
            dest.down
        } else {
            dest.right
        };

        let sign = dot(tangent_f, tangent_dest);
        let along_dest = if sign == 1 {
            along_f
        } else {
            (n - 1) - along_f
        };

        let (new_r, new_c) = match new_facing {
            0 => (along_dest, 0),
            1 => (0, along_dest),
            2 => (along_dest, n - 1),
            3 => (n - 1, along_dest),
            _ => unreachable!(),
        };

        (dest_idx, new_r, new_c, new_facing)
    }
}

const DIRS: [(i32, i32); 4] = [(0, 1), (1, 0), (0, -1), (-1, 0)];

fn turn(facing: usize, t: char) -> usize {
    if t == 'R' {
        (facing + 1) % 4
    } else {
        (facing + 3) % 4
    }
}

pub fn walk_flat(board: &Board, path: &[Instr]) -> (usize, usize, usize) {
    let height = board.grid.len() as i32;
    let width = board.grid[0].len() as i32;
    let mut row = 0usize;
    let mut col = board.grid[0].iter().position(|&c| c == b'.').unwrap();
    let mut facing = 0usize;

    for instr in path {
        match instr {
            Instr::Turn(t) => facing = turn(facing, *t),
            Instr::Move(steps) => {
                for _ in 0..*steps {
                    let (dr, dc) = DIRS[facing];
                    let mut r = row as i32;
                    let mut c = col as i32;
                    loop {
                        r = (r + dr).rem_euclid(height);
                        c = (c + dc).rem_euclid(width);
                        if board.grid[r as usize][c as usize] != b' ' {
                            break;
                        }
                    }
                    if board.grid[r as usize][c as usize] == b'#' {
                        break;
                    }
                    row = r as usize;
                    col = c as usize;
                }
            }
        }
    }

    (row, col, facing)
}

pub fn walk_cube(board: &Board, path: &[Instr]) -> (usize, usize, usize) {
    let n = board.side;
    let start_col = board.grid[0].iter().position(|&c| c == b'.').unwrap();
    let mut face_idx = board.face_at[&(0, start_col / n)];
    let mut local_r = 0usize;
    let mut local_c = start_col % n;
    let mut facing = 0usize;

    for instr in path {
        match instr {
            Instr::Turn(t) => facing = turn(facing, *t),
            Instr::Move(steps) => {
                for _ in 0..*steps {
                    let (dr, dc) = DIRS[facing];
                    let nr = local_r as i32 + dr;
                    let nc = local_c as i32 + dc;

                    let (nf, nr2, nc2, nfacing) =
                        if (0..n as i32).contains(&nr) && (0..n as i32).contains(&nc) {
                            (face_idx, nr as usize, nc as usize, facing)
                        } else {
                            board.cross_edge(face_idx, local_r, local_c, facing)
                        };

                    if board.cell(nf, nr2, nc2) == b'#' {
                        break;
                    }
                    face_idx = nf;
                    local_r = nr2;
                    local_c = nc2;
                    facing = nfacing;
                }
            }
        }
    }

    let block = board.faces[face_idx].block;
    (block.0 * n + local_r, block.1 * n + local_c, facing)
}

pub fn password(row: usize, col: usize, facing: usize) -> i64 {
    1000 * (row as i64 + 1) + 4 * (col as i64 + 1) + facing as i64
}
