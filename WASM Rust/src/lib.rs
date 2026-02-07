mod utils;
mod game_map;
mod state_model;
mod symmetry_tables;
mod symmred;
mod ttt_engine;
mod values;
mod win_masks;
mod map;

use wasm_bindgen::prelude::*;

// When the `wee_alloc` feature is enabled, use `wee_alloc` as the global
// allocator.
#[cfg(feature = "wee_alloc")]
#[global_allocator]
static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;

// #[wasm_bindgen]
// extern {
//     fn alert(s: &str);
// }
//
// #[wasm_bindgen]
// pub fn greet() {
//     alert("Hello, tte_wasm!");
// }

use crate::ttt_engine::ttt as engine_ttt;

// This attribute makes the function callable from JavaScript
#[wasm_bindgen]
pub fn utnamttt(xo_board: u32) -> u8 {
    // You might need to handle panic hooks here for better debugging in console
    engine_ttt(xo_board)
}