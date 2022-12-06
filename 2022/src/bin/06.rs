use std::collections::HashSet;

use aoc::read_file_input;

fn main() {
    let input: Vec<char> = read_file_input("06.txt".to_string()).chars().collect();

    for i in 0..input.len() {
        let part = Vec::from_iter(input[i..(i + 4)].iter().cloned());

        if part.iter().collect::<HashSet<_>>().len() == 4 {
            println!("{}", i + 4);
            break;
        }
    }

    for i in 0..input.len() {
        let part = Vec::from_iter(input[i..(i + 14)].iter().cloned());

        if part.iter().collect::<HashSet<_>>().len() == 14 {
            println!("{}", i + 14);
            break;
        }
    }
}
