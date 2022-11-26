use aoc::read_file_input;

fn main() {
    let input = read_file_input(01);

    for (i, line) in input.lines().enumerate() {
        println!("Line {} contains {}", i, line);
    }

    println!("Hello from Day 1");
}
