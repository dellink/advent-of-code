use aoc::read_file_input;

fn main() {
    let input = read_file_input("01.txt".to_string());

    let mut current = 0;
    let mut max: (i32, i32, i32) = (0, 0, 0);

    for (_i, line) in input.lines().enumerate() {
        if line.is_empty() {
            max = calculate(current, max);
            current = 0;
        } else {
            current += line.trim().parse::<i32>().unwrap();
        }
    }

    max = calculate(current, max);

    println!("{}", max.0);
    println!("{}", max.0 + max.1 + max.2);
}

fn calculate(current: i32, max: (i32, i32, i32)) -> (i32, i32, i32) {
    let mut max = max;
    if current > max.0 {
        max.2 = max.1;
        max.1 = max.0;
        max.0 = current;
    } else if current > max.1 {
        max.2 = max.1;
        max.1 = current;
    } else if current > max.2 {
        max.2 = current;
    }
    return max;
}
