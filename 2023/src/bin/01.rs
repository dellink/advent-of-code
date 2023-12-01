use aoc::read_file_input;

fn main() {
    let input = read_file_input("01.txt".to_string());
    let mut part1 = 0;
    let mut part2 = 0;

    for line in input.lines() {
        part1 += calculate(line.to_string());

        let replaced = line
            .replace("nine", "n9e")
            .replace("eight", "e8t")
            .replace("seven", "s7n")
            .replace("six", "s6x")
            .replace("five", "f5e")
            .replace("four", "f4r")
            .replace("three", "t3e")
            .replace("two", "t2o")
            .replace("one", "o1e");

        part2 += calculate(replaced);
    }

    println!("{}", part1);
    println!("{}", part2);
}

fn calculate(str: String) -> i32 {
    let mut digits = str.chars().filter(|c| c.is_numeric());

    let first = digits.next();
    let last = digits.last();

    let mut s = String::new();

    if first.is_some() && last.is_some() {
        s.push(first.unwrap());
        s.push(last.unwrap());
    } else {
        s.push(first.unwrap());
        s.push(first.unwrap());
    }

    return s.parse::<i32>().unwrap();
}
