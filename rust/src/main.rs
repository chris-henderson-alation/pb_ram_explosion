use std::io::BufRead;

mod example {
    tonic::include_proto!("main");
}

fn main() {
    println!("htop --pid {}", std::process::id());
    let m = example::MySmallMessage{tiny: "just tricking dead code analysis".to_string()};
    std::io::BufReader::new(std::io::stdin()).read_line(&mut String::new()).unwrap();
    println!("{:?}", m);
}
