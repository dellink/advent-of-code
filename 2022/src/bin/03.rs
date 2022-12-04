use aoc::read_file_input;

fn main() {
    let input = read_file_input("03.txt".to_string());
    let mut score1 = 0;
    let mut score2 = 0;
    let alphabet = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";

    for line in input.lines() {
        let compartments = line.split_at(line.len() / 2);

        for char in compartments.0.chars() {
            match compartments.1.find(char) {
                Some(_) => match alphabet.find(char) {
                    Some(index) => {
                        score1 += index;
                        break;
                    }
                    None => {}
                },
                None => {}
            };
        }
    }

    println!("{}", score1);

    let mut lines = input.lines();
    while let (Some(line1), Some(line2), Some(line3)) = (lines.next(), lines.next(), lines.next()) {
        for char in line1.chars() {
            match line2.find(char) {
                Some(_) => {
                    match line3.find(char) {
                        Some(_) => match alphabet.find(char) {
                            Some(index) => {
                                score2 += index;
                                break;
                            }
                            None => {}
                        },
                        None => {}
                    };
                }
                None => {}
            };
        }
    }

    println!("{}", score2);
}
