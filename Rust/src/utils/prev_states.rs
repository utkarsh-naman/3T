use crate::model::map::GameMap;
use crate::model::state_model::State;


pub fn parent_of(map: &GameMap, state: &State)-> Vec<State> {
    let mut parents: Vec<State> = Vec::new();

    for (key, props) in map.iter(){
        if props.next_state.contains(state){
            parents.push(*key);
        }
    }
    parents
}