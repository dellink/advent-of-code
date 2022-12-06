use aoc::read_file_input;

fn main() {
    let input = read_file_input("05.txt".to_string());

    let (boxes, instructions) = input.split_once("\n\n").unwrap();

    let mut stacks1 = vec![vec![]; 9];

    for line in boxes.lines().rev() {
        for (i, char) in line.chars().enumerate() {
            if char.is_alphabetic() {
                stacks1[i / 4].push(char as char);
            }
        }
    }

    let mut stack2 = stacks1.clone();

    for line in instructions.lines() {
        let instruction: Vec<usize> = line
            .split_whitespace()
            .filter_map(|s| s.parse::<usize>().ok())
            .collect();

        for _ in 0..instruction[0] {
            let item = stacks1[instruction[1] - 1].pop().unwrap();
            stacks1[instruction[2] - 1].push(item);
        }
    }

    println!(
        "{}",
        stacks1.iter().filter_map(|s| s.last()).collect::<String>()
    );

    for line in instructions.lines() {
        let instruction: Vec<usize> = line
            .split_whitespace()
            .filter_map(|s| s.parse::<usize>().ok())
            .collect();

        let len = stack2[instruction[2] - 1].len() + instruction[0];
        stack2[instruction[2] - 1].resize(len, 'x');
        for i in 0..instruction[0] {
            let item = stack2[instruction[1] - 1].pop().unwrap();
            stack2[instruction[2] - 1][len - 1 - i] = item;
        }
    }

    println!(
        "{}",
        stack2.iter().filter_map(|s| s.last()).collect::<String>()
    );
}
