// use std::collections::HashMap;
use std::sync::OnceLock;
use crate::model::map::{load_map_from_file, GameMap};

static GMAP: OnceLock<GameMap> = OnceLock::new();

pub fn gmap() -> &'static GameMap {
    GMAP.get_or_init(|| {
        load_map_from_file("bin/output/map/valued/map2.ttt")
            .expect("failed to load file")
    })
}


// let map = gmap();
// let state = map.get(&some_state);
