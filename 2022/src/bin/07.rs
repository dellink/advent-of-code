use std::collections::{HashMap, HashSet};
use std::path::PathBuf;

use aoc::read_file_input;

fn main() {
    let input = read_file_input("07.txt".to_string());

    let mut fs = HashMap::new();
    let mut pwd = PathBuf::new();

    for instructions in input.split('$').skip(1) {
        match instructions.trim().lines().next().unwrap() {
            "ls" => {
                let sizes = instructions
                    .lines()
                    .skip(1)
                    .map(|instruction| {
                        let (size, f) = instruction.split_once(' ').unwrap();
                        (size.parse::<i32>().unwrap_or(-1), f)
                    })
                    .collect();

                fs.insert(pwd.clone(), sizes);
            }
            "cd .." => {
                pwd.pop();
            }
            dir => {
                pwd.push(dir.split_once(' ').unwrap().1);
            }
        }
    }

    let mut sizes = HashMap::new();

    for path in fs.keys() {
        compute_size(&fs, &mut sizes, path);
    }

    println!("{}", sizes.values().filter(|&&s| s <= 100000).sum::<i32>());
    println!(
        "{}",
        sizes
            .values()
            .filter(|&&s| 40000000 + s >= sizes[&PathBuf::from("/")])
            .min()
            .unwrap()
    );
}

fn compute_size(
    fs: &HashMap<PathBuf, HashSet<(i32, &str)>>,
    sizes: &mut HashMap<PathBuf, i32>,
    path: &PathBuf,
) {
    if sizes.contains_key(path) {
        return;
    }

    let size = fs[path]
        .iter()
        .map(|&(s, d)| match s {
            -1 => {
                let path = path.join(d);
                compute_size(fs, sizes, &path);
                sizes[&path]
            }
            s => s,
        })
        .sum();

    sizes.insert(path.clone(), size);
}
