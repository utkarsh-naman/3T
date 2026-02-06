use std::collections::{HashMap, VecDeque};
use crate::model::map::GameMap;
use crate::model::state_model::{gen_board, State};

pub fn gen_map(start: &State)->GameMap{
    let mut map: GameMap = GameMap::new();
    let mut visited: HashMap<State, bool> = HashMap::new();
    visited.insert(*start, true);

    let mut to_visit: VecDeque<State> = VecDeque::new();
    to_visit.push_back(*start);

    while !to_visit.is_empty() {
        let state_node = to_visit.pop_front();
        
    }

    map
}