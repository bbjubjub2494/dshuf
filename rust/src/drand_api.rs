use drand_core::beacon::RandomnessBeacon;
use drand_core::chain::{ChainInfo, ChainOptions, ChainVerification};
use drand_core::{DrandError, HttpClient};

pub fn quicknet_chaininfo() -> ChainInfo {
    serde_json::from_str(r#"{"public_key":"83cf0f2896adee7eb8b5f01fcad3912212c437e0073e911fb90022d3e760183c8c4b450b6a0a6c3ac6a5776a2d1064510d1fec758c921cc22b0e17e63aaf4bcb5ed66304de9cf809bd274ca73bab4af5a6e9c76a4bc09e76eae8991ef5ece45a","period":3,"genesis_time":1692803367,"hash":"52db9ba70e0cc0f6eaf7803dd07447a1f5477735fd3f661792ba94600c84e971","groupHash":"f477d5c89f21a17c863a7f937c6a6d15859414d2be09cd448d4279af331c5d3e","schemeID":"bls-unchained-g1-rfc9380","metadata":{"beaconID":"quicknet"}}"#).unwrap()
}

pub fn get_beacon(round_number: u64) -> Result<RandomnessBeacon, DrandError> {
    let endpoint =
        std::env::var("DSHUF_ENDPOINT").unwrap_or("https://drand.cloudflare.com".to_string());
    let api_baseurl = format!("{}/{}", endpoint, hex::encode(quicknet_chaininfo().hash()));
    let options = ChainOptions::new(
        true,
        true,
        Some(ChainVerification::from(quicknet_chaininfo())),
    );
    let client = HttpClient::new(&api_baseurl, Some(options))?;
    client.get(round_number)
}
