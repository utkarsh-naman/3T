use std::collections::{HashMap, VecDeque};
use crate::model::map::{GameMap, StateProps};
use crate::model::state_model::{State};
use crate::utils::next_states::next_states;

pub fn gen_map(start: &State) ->GameMap{
    let mut map: GameMap = GameMap::new();
    let mut visited: HashMap<State, bool> = HashMap::new();
    visited.insert(*start, true);

    let mut to_visit: VecDeque<State> = VecDeque::new();
    to_visit.push_back(*start);

    while let Some(state_node) = to_visit.pop_front(){
        // do something
        let children = next_states(&state_node);

        for &child in &children{
            if visited.insert(child, true).is_none(){
                to_visit.push_back(child);
            }
        }

        let node_props = StateProps {
            score: 0.0,
            win_depth: 10,
            lose_depth: 10,
            next_state: children,
        };

        map.insert(state_node, node_props);


    }
    map
}
