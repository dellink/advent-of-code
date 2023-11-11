use std::process::Command;

fn main() {
    for day in 1..=25 {
        let day = format!("{:02}", day);

        let cmd = Command::new("cargo")
            .args(["run", "--release", "--bin", &day])
            .output()
            .unwrap();

        println!("----------");
        println!("| Day {} |", day);
        println!("----------");

        let output = String::from_utf8(cmd.stdout).unwrap();
        let is_empty = output.is_empty();

        println!(
            "{}",
            if is_empty {
                "Not solved."
            } else {
                output.trim()
            }
        )
    }
}
