use std::collections::HashMap;
use std::fs::File;
use std::io::{BufReader, BufWriter};
use std::path::Path;
use serde::{Serialize, Deserialize}; // Import Serde traits

use crate::state_model::State;

#[derive(Serialize, Deserialize, Debug)]
pub struct StateProps{
    pub score: f32,
    pub win_depth: u8,
    pub lose_depth: u8,
    pub next_state: Vec<State>,
}

pub type GameMap = HashMap<State, StateProps>;

// pub fn print_map(map: &GameMap) {
//     if map.is_empty() {
//         println!("The map is empty.");
//         return;
//     }
//
//     println!("--- Game Map Dump (Total States: {}) ---", map.len());
//
//     for (state, props) in map {
//         println!("State [binary]: {:032b}", state);
//         println!("\tState Properties {{");
//         println!("\t\tscore: {}", props.score);
//         println!("\t\twin_depth: {}", props.win_depth);
//         println!("\t\tlose_depth: {}", props.lose_depth);
//
//         println!("\t\tnext_state: [");
//         for next in &props.next_state {
//             println!("\t\t\t{:032b}", next);
//         }
//         println!("\t\t]");
//
//         println!("\t}}\n");
//     }
//
//     println!("------------------------------------------");
// }


// pub fn save_map_to_file(map: &GameMap, filename: &str) -> bincode::Result<()> {
//     let path = Path::new(filename);
//     let file = File::create(path)?;
//     let writer = BufWriter::new(file);
//
//     // Serialize the map directly into the file
//     bincode::serialize_into(writer, map)?;
//
//     println!("Successfully saved map to: {}", filename);
//     Ok(())
// }

pub fn load_map_from_file(filename: &str) -> bincode::Result<GameMap> {
    let path = Path::new(filename);
    let file = File::open(path)?;
    let reader = BufReader::new(file);

    // Deserialize the data back into a HashMap
    let map: GameMap = bincode::deserialize_from(reader)?;

    // println!("Successfully loaded map from: {}", filename);
    Ok(map)
}