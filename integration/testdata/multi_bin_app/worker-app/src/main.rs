use std::time::Duration;

fn main() {
    println!("worker started");

    loop {
        std::thread::sleep(Duration::from_secs(60));
    }
}
