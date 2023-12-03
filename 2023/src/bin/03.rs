use std::collections::HashMap;

use regex::Regex;

use aoc::read_file_input;

fn main() {
    let input = read_file_input("03.txt".to_string());
    let mut part1 = 0;

    let lines: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();
    let re = Regex::new(r"\d+").unwrap();

    for (line_number, line) in lines.iter().enumerate() {
        for mat in re.find_iter(&line.iter().collect::<String>()) {
            let number: i32 = mat.as_str().parse().unwrap();
            let start_column = mat.start();
            let end_column = mat.end();
            let mut found = false;

            for column in start_column..end_column {
                let positions = get_positions(line_number as i32, column as i32);

                for (x, y) in &positions {
                    if *x >= 0 && *y >= 0 && *x < lines.len() as i32 && *y < line.len() as i32 {
                        let adjacent = lines[*x as usize][*y as usize];
                        if adjacent != '.' && !adjacent.is_numeric() {
                            found = true;
                            part1 += number;
                            break;
                        }
                    }
                }

                if found {
                    break;
                }
            }
        }
    }

    println!("{}", part1);

    let mut part2 = 0;
    let mut map: HashMap<(i32, i32), Vec<i32>> = HashMap::new();

    for (line_number, line) in lines.iter().enumerate() {
        for mat in re.find_iter(&line.iter().collect::<String>()) {
            let number: i32 = mat.as_str().parse().unwrap();
            let start_column = mat.start();
            let end_column = mat.end();
            let mut found = false;

            for column in start_column..end_column {
                let positions = get_positions(line_number as i32, column as i32);

                for (x, y) in &positions {
                    if *x >= 0 && *y >= 0 && *x < lines.len() as i32 && *y < line.len() as i32 {
                        let adjacent = lines[*x as usize][*y as usize];
                        if adjacent == '*' && !adjacent.is_numeric() {
                            found = true;

                            map.entry((*x, *y)).or_insert_with(Vec::new).push(number);

                            break;
                        }
                    }
                }

                if found {
                    break;
                }
            }
        }
    }

    for v in map.values() {
        if v.len() == 2 {
            part2 += v[0] * v[1];
        }
    }

    println!("{}", part2);
}

fn get_positions(line_number: i32, column: i32) -> [(i32, i32); 8] {
    return [
        (line_number - 1, column - 1),
        (line_number - 1, column),
        (line_number - 1, column + 1),
        (line_number, column - 1),
        (line_number, column + 1),
        (line_number + 1, column - 1),
        (line_number + 1, column),
        (line_number + 1, column + 1),
    ];
}
