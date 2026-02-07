// // use std::collections::HashMap;
// use std::sync::OnceLock;
// use crate::map::{load_map_from_file, GameMap};
//
// static GMAP: OnceLock<GameMap> = OnceLock::new();
//
// pub fn gmap() -> &'static GameMap {
//     GMAP.get_or_init(|| {
//         load_map_from_file("bin/output/map/valued/map2.ttt")
//             .expect("failed to load file")
//     })
// }


// let map = gmap();
// let state = map.get(&some_state);


//  WEB ASSEMBLY

use std::sync::OnceLock;
use crate::map::GameMap; // Adjusted path based on module structure

static GMAP: OnceLock<GameMap> = OnceLock::new();

// This embeds the file into the binary at compile time!
// constant path assumes map2.ttt is inside the 'src' folder.
static MAP_BYTES: &[u8] = include_bytes!("map2.ttt");

pub fn gmap() -> &'static GameMap {
    GMAP.get_or_init(|| {
        // Deserialize directly from the memory slice, not a file
        bincode::deserialize(MAP_BYTES)
            .expect("failed to deserialize embedded map")
    })
}