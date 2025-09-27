use std::io::{self, Read, Write};
use std::fs::File;

use dshuf::drand_api;
use dshuf::shuffle;

use clap::Parser;

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Args {
    #[arg(
        short = 'n',
        long = "head-count",
        value_name = "COUNT",
        help = "output at most this many lines"
    )]
    head_count: Option<usize>,
    #[arg(
        short = 'b',
        long = "beacon",
        help = "round number of beacon to use for randomness"
    )]
    beacon: u64,
    #[arg(short = 'r', long = "repeat", help = "output lines can be repeated")]
    repetitions: bool,
    #[arg(value_name = "FILE")]
    file: Option<String>,
    #[arg(
        short = 'z',
        long = "zero-terminated",
        help = "line delimiter is NUL, not newline"
    )]
    zero_sep: bool,
}

fn main() {
    let args = Args::parse();
    let separator = if args.zero_sep { 0 } else { b'\n' };
    let count = args.head_count;
    let round_number = args.beacon;
    let inputfile: Option<File> = match args.file.as_deref() {
        None | Some("-") => None, // stdin
        Some(path) => Some(File::open(path).unwrap()),
    };

    let randomness = drand_api::get_beacon(round_number).unwrap().randomness();
    let mut buf = Vec::new();
    (match inputfile {
        None => io::stdin().read_to_end(&mut buf),
        Some(mut f) => f.read_to_end(&mut buf),
    })
    .unwrap();
    let mut input = Vec::from_iter(buf.split(|c| *c == separator));
    if input.last().is_some_and(|e| e.is_empty()) {
        input.truncate(input.len() - 1);
    }
    let output = shuffle(
        randomness.as_slice().try_into().unwrap(),
        input,
        args.repetitions,
    );

    let mut stdout = io::stdout();
    if let Some(limit) = count {
        for e in output.take(limit) {
            stdout.write_all(e).unwrap();
            stdout.write_all(&[separator]).unwrap();
        }
    } else {
        for e in output {
            stdout.write_all(e).unwrap();
            stdout.write_all(&[separator]).unwrap();
        }
    }
}
