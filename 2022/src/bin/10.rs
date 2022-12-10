use aoc::read_file_input;

fn main() {
    let input = read_file_input("10.txt".to_string());

    let mut cycle: i32 = 0;
    let mut signal: i32 = 1;
    let mut strengths = 0;
    let mut sprite = String::new();

    for line in input.lines() {
        let cycles = if line.len() > 4 { 2 } else { 1 };

        for _ in 1..=cycles {
            if cycle % 40 == 0 {
                sprite.push('\n');
            }
            sprite.push(if (signal - cycle % 40).abs() < 2 {
                '█'
            } else {
                ' '
            });
            cycle += 1;
            if (cycle - 20) % 40 == 0 {
                strengths += cycle * signal;
            }
        }
        if line.len() > 4 {
            signal += line.split_once(" ").unwrap().1.parse::<i32>().unwrap();
        }
    }

    println!("{}", strengths);
    println!("{}", sprite);
}
