use clap::Parser;
use drand_core::beacon::RandomnessBeaconTime;

use core::option::Option;

use dshuf::drand_api::quicknet_chaininfo;
use std::time::Duration;

#[derive(Parser)]
struct Args {
    #[arg(short = 'd', long)]
    delay: Option<humantime::Duration>,
}

fn main() {
    let args = Args::parse();
    let delay = args.delay.unwrap_or(Duration::ZERO.into());
    let current_round =
        RandomnessBeaconTime::new(&quicknet_chaininfo().into(), &delay.to_string()).round();
    println!("{}", current_round + 1);
}
