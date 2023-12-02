use regex::Regex;

use aoc::read_file_input;

fn main() {
    let input = read_file_input("02.txt".to_string());
    let mut part1 = 0;

    for line in input.lines() {
        let parts: Vec<&str> = line.split(':').collect();
        let game_number: Vec<&str> = parts[0].split_whitespace().collect();

        let game_number: i32 = game_number[1].parse().unwrap();

        let sets: Vec<&str> = parts[1].split(';').collect();

        if check_game_possibility(sets) == true {
            part1 += game_number;
        }
    }

    println!("{}", part1);

    let mut part2 = 0;

    for line in input.lines() {
        let parts: Vec<&str> = line.split(':').collect();
        let sets: Vec<&str> = parts[1].split(';').collect();

        let max = get_max_color(sets);

        part2 += max.0 * max.1 * max.2;
    }

    println!("{}", part2);
}

fn check_game_possibility(sets: Vec<&str>) -> bool {
    let re = Regex::new(r"(\d+) (red|blue|green)").unwrap();

    for set in sets {
        for cap in re.captures_iter(set) {
            let mut red = 0;
            let mut blue = 0;
            let mut green = 0;

            let number: i32 = cap[1].parse().unwrap();
            match &cap[2] {
                "red" => red += number,
                "blue" => blue += number,
                "green" => green += number,
                _ => (),
            }

            if red > 12 || green > 13 || blue > 14 {
                return false;
            }
        }
    }

    return true;
}

fn get_max_color(sets: Vec<&str>) -> (i32, i32, i32) {
    let re = Regex::new(r"(\d+) (red|blue|green)").unwrap();

    let mut max_red = 0;
    let mut max_blue = 0;
    let mut max_green = 0;

    for set in sets {
        let mut red = 0;
        let mut blue = 0;
        let mut green = 0;

        for cap in re.captures_iter(set) {
            let number: i32 = cap[1].parse().unwrap();
            match &cap[2] {
                "red" => red += number,
                "blue" => blue += number,
                "green" => green += number,
                _ => (),
            }
        }

        max_red = std::cmp::max(max_red, red);
        max_blue = std::cmp::max(max_blue, blue);
        max_green = std::cmp::max(max_green, green);
    }

    (max_red, max_blue, max_green)
}
