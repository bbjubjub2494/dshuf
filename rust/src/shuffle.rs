use num_bigint::BigUint;
use num_traits::cast::ToPrimitive;

const SAMPLE_LEN: usize = 24;

#[derive(Debug)]
pub struct ShuffleIter<'a> {
    prng: blake3::OutputReader,
    input: Vec<&'a [u8]>,
    i: usize,
    repetitions: bool,
}

impl<'a> ShuffleIter<'a> {
    pub fn new(randomness: &[u8; 32], input: Vec<&'a [u8]>, repetitions: bool) -> ShuffleIter<'a> {
        let mut h = blake3::Hasher::new_keyed(randomness);
        for e in &input {
            h.update(&(e.len() as u64).to_be_bytes());
            h.update(e);
        }
        let prng = h.finalize_xof();
        ShuffleIter {
            prng,
            input,
            i: 0,
            repetitions,
        }
    }
}

impl<'a> Iterator for ShuffleIter<'a> {
    type Item = &'a [u8];

    fn next(&mut self) -> Option<Self::Item> {
        if !self.repetitions && self.i == self.input.len() {
            return None;
        }
        let mut sample = [0u8; SAMPLE_LEN];
        self.prng.fill(&mut sample);
        let r = BigUint::from_bytes_be(&sample);
        if self.repetitions {
            let j = (r % self.input.len()).to_usize().unwrap();
            Some(self.input[j])
        } else {
            let j = (self.i + (r % (self.input.len() - self.i)))
                .to_usize()
                .unwrap();
            self.input.swap(self.i, j);
            let r = self.input[self.i];
            self.i += 1;
            Some(r)
        }
    }
}
pub fn shuffle<'a>(
    randomness: &[u8; 32],
    input: Vec<&'a [u8]>,
    repetitions: bool,
) -> ShuffleIter<'a> {
    ShuffleIter::new(randomness, input, repetitions)
}
