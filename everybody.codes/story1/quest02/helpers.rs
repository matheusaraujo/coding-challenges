use std::cell::RefCell;
use std::collections::{BTreeMap, HashMap};
use std::rc::Rc;

type NodePtr = Rc<RefCell<TreeNode>>;
type SwapFn = fn(&mut TreeNode, &mut TreeNode);

#[derive(Debug)]
pub struct TreeNode {
    pub rank: i32,
    pub symbol: char,
    pub left: Option<NodePtr>,
    pub right: Option<NodePtr>,
}

fn insert_node(root: &NodePtr, rank: i32, symbol: char) -> NodePtr {
    let next_node = {
        let node = root.borrow_mut();
        if rank < node.rank {
            node.left.as_ref().map(Rc::clone)
        } else {
            node.right.as_ref().map(Rc::clone)
        }
    };

    if let Some(next) = next_node {
        insert_node(&next, rank, symbol)
    } else {
        let mut node = root.borrow_mut();
        let new_node = Rc::new(RefCell::new(TreeNode { rank, symbol, left: None, right: None }));
        if rank < node.rank {
            node.left = Some(Rc::clone(&new_node));
        } else {
            node.right = Some(Rc::clone(&new_node));
        }
        new_node
    }
}

fn get_max_level_msg(root_opt: &Option<NodePtr>, deepest_wins: bool) -> String {
    let Some(root) = root_opt else {
        return String::new();
    };

    let mut level_map: BTreeMap<i32, Vec<char>> = BTreeMap::new();

    fn walk(node: &NodePtr, depth: i32, map: &mut BTreeMap<i32, Vec<char>>) {
        let n = node.borrow();
        map.entry(depth).or_default().push(n.symbol);
        if let Some(ref l) = n.left {
            walk(l, depth + 1, map);
        }
        if let Some(ref r) = n.right {
            walk(r, depth + 1, map);
        }
    }

    walk(root, 0, &mut level_map);

    let max_len = level_map.values().map(|v| v.len()).max().unwrap_or(0);

    let mut iter: Box<dyn Iterator<Item = (&i32, &Vec<char>)>> = if deepest_wins {
        Box::new(level_map.iter().rev())
    } else {
        Box::new(level_map.iter())
    };

    iter.find(|(_, v)| v.len() == max_len)
        .map(|(_, v)| v.iter().collect())
        .unwrap_or_default()
}

pub fn solve(puzzle_input: &[String], swap_fn: Option<SwapFn>, deepest_wins: bool) -> String {
    let mut left_root: Option<NodePtr> = None;
    let mut right_root: Option<NodePtr> = None;

    let mut left_map: HashMap<String, NodePtr> = HashMap::new();
    let mut right_map: HashMap<String, NodePtr> = HashMap::new();

    for line in puzzle_input {
        let line = line.trim();
        if line.is_empty() {
            continue;
        }

        if line.starts_with("ADD") {
            let id = line
                .split("id=")
                .nth(1)
                .unwrap()
                .split_whitespace()
                .next()
                .unwrap()
                .to_string();

            let parts: Vec<&str> = line.split(['[', ']']).filter(|s| s.contains(',')).collect();

            let (l_rank, l_sym) = {
                let p: Vec<&str> = parts[0].split(',').collect();
                (
                    p[0].trim().parse().unwrap(),
                    p[1].trim().chars().next().unwrap(),
                )
            };

            let (r_rank, r_sym) = {
                let p: Vec<&str> = parts[1].split(',').collect();
                (
                    p[0].trim().parse().unwrap(),
                    p[1].trim().chars().next().unwrap(),
                )
            };

            let l_ptr = if let Some(ref root) = left_root {
                insert_node(root, l_rank, l_sym)
            } else {
                let root = Rc::new(RefCell::new(TreeNode {
                    rank: l_rank,
                    symbol: l_sym,
                    left: None,
                    right: None,
                }));
                left_root = Some(Rc::clone(&root));
                Rc::clone(&root)
            };

            let r_ptr = if let Some(ref root) = right_root {
                insert_node(root, r_rank, r_sym)
            } else {
                let root = Rc::new(RefCell::new(TreeNode {
                    rank: r_rank,
                    symbol: r_sym,
                    left: None,
                    right: None,
                }));
                right_root = Some(Rc::clone(&root));
                Rc::clone(&root)
            };

            left_map.insert(id.clone(), l_ptr);
            right_map.insert(id, r_ptr);
        } else if line.starts_with("SWAP") && swap_fn.is_some() {
            let swap_fn = swap_fn.unwrap();
            let id = line.split_whitespace().nth(1).unwrap();

            if let (Some(l), Some(r)) = (left_map.get(id), right_map.get(id)) {
                let mut ln = l.borrow_mut();
                let mut rn = r.borrow_mut();
                swap_fn(&mut ln, &mut rn);
            }
        }
    }

    format!(
        "{}{}",
        get_max_level_msg(&left_root, deepest_wins),
        get_max_level_msg(&right_root, deepest_wins)
    )
}
