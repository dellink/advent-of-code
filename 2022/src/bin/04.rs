use aoc::read_file_input;

fn main() {
    let input = read_file_input("04.txt".to_string());
    let mut score1 = 0;
    let mut score2 = 0;

    for line in input.lines() {
        let elves: Vec<&str> = line.split(',').collect();
        let section1: Vec<i32> = elves[0]
            .split("-")
            .map(|s| s.parse::<i32>().unwrap())
            .collect();
        let section2: Vec<i32> = elves[1]
            .split("-")
            .map(|s| s.parse::<i32>().unwrap())
            .collect();

        if section1[0] >= section2[0] && section1[1] <= section2[1] {
            score1 += 1;
        } else if section2[0] >= section1[0] && section2[1] <= section1[1] {
            score1 += 1;
        }

        if section1[1] >= section2[0] && section1[0] <= section2[1] {
            score2 += 1;
        }
    }

    println!("{}", score1);
    println!("{}", score2);
}
