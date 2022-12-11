use aoc::read_file_input;

#[derive(Clone)]
enum Operation {
    Add(usize),
    Multiply(usize),
    SelfMultiply,
}

fn main() {
    let input = read_file_input("11.txt".to_string()).to_string();

    let monkeys: Vec<(Vec<usize>, Operation, usize, usize, usize)> = input
        .split("\n\n")
        .map(|m| {
            let lines: Vec<&str> = m.lines().map(|l| l.trim()).collect();
            (
                lines[1]["Starting items: ".len()..]
                    .split(", ")
                    .map(|x| x.parse().unwrap())
                    .collect(),
                lines[2]["Operation: new = old * ".len()..]
                    .parse()
                    .map(|v| {
                        if lines[2].contains('+') {
                            Operation::Add(v)
                        } else {
                            Operation::Multiply(v)
                        }
                    })
                    .unwrap_or(Operation::SelfMultiply),
                lines[3]["Test: divisible by ".len()..].parse().unwrap(),
                lines[4]["If true: throw to monkey ".len()..]
                    .parse()
                    .unwrap(),
                lines[5]["If false: throw to monkey ".len()..]
                    .parse()
                    .unwrap(),
            )
        })
        .collect::<Vec<_>>();

    let magic = monkeys.iter().map(|m| m.2).product::<usize>();

    println!("{}", process(monkeys.clone(), 20, |x| x / 3));
    println!("{}", process(monkeys, 10000, |x| x % magic));
}

fn process(
    mut monkeys: Vec<(Vec<usize>, Operation, usize, usize, usize)>,
    rounds: usize,
    f: impl Fn(usize) -> usize,
) -> usize {
    let mut inspections = vec![0; monkeys.len()];

    for _ in 0..rounds {
        for i in 0..monkeys.len() {
            let (items, op, test, monkey_true, monkey_false) = monkeys[i].clone();

            for item in items {
                let worry = match op {
                    Operation::Add(v) => f(item + v),
                    Operation::Multiply(v) => f(item * v),
                    Operation::SelfMultiply => f(item * item),
                };
                monkeys[if worry % test == 0 {
                    monkey_true
                } else {
                    monkey_false
                }]
                .0
                .push(worry);
            }

            inspections[i] += monkeys[i].0.len();
            monkeys[i].0.clear();
        }
    }

    inspections.sort();
    inspections[inspections.len() - 1] * inspections[inspections.len() - 2]
}
