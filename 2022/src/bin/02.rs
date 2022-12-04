use aoc::read_file_input;

fn main() {
    let input = read_file_input("02.txt".to_string());
    let mut score1 = 0;
    let mut score2 = 0;

    for line in input.lines() {
        let columns: Vec<&str> = line.split(' ').collect();
        if columns[1] == "X" {
            score1 += 1;
            if columns[0] == "A" {
                score1 += 3;
                score2 += 3;
            } else if columns[0] == "C" {
                score1 += 6;
                score2 += 2;
            } else {
                score2 += 1;
            }
        } else if columns[1] == "Y" {
            score1 += 2;
            if columns[0] == "B" {
                score1 += 3;
                score2 += 2;
            } else if columns[0] == "A" {
                score1 += 6;
                score2 += 1;
            } else {
                score2 += 3;
            }
            score2 += 3;
        } else {
            score1 += 3;
            if columns[0] == "C" {
                score1 += 3;
                score2 += 1;
            } else if columns[0] == "B" {
                score1 += 6;
                score2 += 3;
            } else {
                score2 += 2;
            }
            score2 += 6;
        }
    }

    println!("{}", score1);
    println!("{}", score2);
}
