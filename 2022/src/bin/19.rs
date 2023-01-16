use aoc::read_file_input;

fn main() {
    let input = read_file_input("19.txt".to_string());
    let blueprints = parse(input);

    println!("{:?}", blueprints);
}

#[derive(Debug)]
struct Blueprint {
    ore: u8,
    clay: u8,
    obsidian: (u8, u8),
    geode: (u8, u8),
}

fn parse(input: String) -> Vec<Blueprint> {
    let mut blueprints: Vec<Blueprint> = vec![];

    for line in input.lines() {
        println!("{}", line);

        let mut blueprint = Blueprint {
            ore: 0,
            clay: 0,
            obsidian: (0, 0),
            geode: (0, 0),
        };

        let costs: Vec<&str> = line.split("costs ").skip(1).collect();
        for (i, cost) in costs.iter().enumerate() {
            if i == 0 {
                blueprint.ore = cost.split_once(" ").unwrap().0.parse::<u8>().unwrap();
            }
            if i == 1 {
                blueprint.clay = cost.split_once(" ").unwrap().0.parse::<u8>().unwrap();
            }
            if i == 2 {
                let c = cost.split_once("and ").unwrap();
                blueprint.obsidian = (
                    c.0.split_once(" ").unwrap().0.parse::<u8>().unwrap(),
                    c.1.split_once(" ").unwrap().0.parse::<u8>().unwrap(),
                );
            }
            if i == 3 {
                let c = cost.split_once("and ").unwrap();
                blueprint.geode = (
                    c.0.split_once(" ").unwrap().0.parse::<u8>().unwrap(),
                    c.1.split_once(" ").unwrap().0.parse::<u8>().unwrap(),
                );
            }
        }
        blueprints.push(blueprint);
    }

    return blueprints;
}
