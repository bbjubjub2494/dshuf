use std::io::{self, Read};

use dshuf::drand_api;
use dshuf::shuffle;

fn main() {
    // TODO: no hardcoding
    let round_number = 1337;
    let randomness = drand_api::get_beacon(round_number).unwrap().randomness();

    // simulate shuf -n 3
    let mut stdin = io::stdin();
    let mut buf = Vec::new();
    stdin.read_to_end(&mut buf).unwrap();
    let separator = '\n';
    let mut input = Vec::from_iter(buf.split(|c| *c == separator as u8));
    if input.last().map_or(false, |e| e.len() == 0) {
        input.truncate(input.len()-1);
    }
    let output = shuffle(randomness.as_slice().try_into().unwrap(), input, 3);
    println!("{:?}", output);
}
