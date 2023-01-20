use aoc::read_file_input;

fn main() {
    let input = read_file_input("20.txt".to_string());

    let mut numbers = vec![];

    for line in input.lines() {
        numbers.push(line.parse::<isize>().unwrap());
    }

    println!("{}", mix(&numbers));
}

fn mix(numbers: &[isize]) -> isize {
    let mut res = (0..numbers.len()).collect::<Vec<usize>>();

    for (i, &n) in numbers.iter().enumerate() {
        let pos = res.iter().position(|&y| y == i).unwrap();
        res.remove(pos);
        let new_index = (pos as isize + n).rem_euclid(res.len() as isize) as usize;
        res.insert(new_index, i);
    }

    let zero_index = numbers.iter().position(|&i| i == 0).unwrap();
    let zero_index = res.iter().position(|&i| i == zero_index).unwrap();

    numbers[res[(zero_index + 1000) % res.len()]]
        + numbers[res[(zero_index + 2000) % res.len()]]
        + numbers[res[(zero_index + 3000) % res.len()]]
}
