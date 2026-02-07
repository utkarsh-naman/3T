use std::cmp::{max, min};
use std::collections::HashMap;
use crate::constants::values::{NEGINF, POSINF};
use crate::mapbuilder::set_terminal_values::get_terminal_values;
use crate::model::map::GameMap;
use crate::model::state_model::State;
use crate::utils::prev_states::parent_of;

pub fn set_values(mut map: GameMap) ->GameMap{
    let win_states = get_terminal_values(&map);
    let mut worked_history: HashMap<State, bool> = HashMap::new();

    wins(&win_states, &mut worked_history, &mut map);
    map
}


fn wins(win_states: &Vec<State>, worked_history : &mut HashMap<State, bool>, map: &mut GameMap) {
    if win_states.len() == 0 {
        return;
    }
    let mut lose_states: Vec<State> = vec![];



    for winkey in win_states{

        let win_wd = if map[winkey].next_state.is_empty() {
                0
            } else {
                max_depth(map, winkey)
            };



        if let Some(win_props) = map.get_mut(winkey) {
            win_props.score = POSINF;
            win_props.win_depth = win_wd;
        }

        for parent_key in parent_of(&map, winkey){
            if worked_history.insert(parent_key, true).is_none() {


                let max_ld = if map[&parent_key].next_state.len() > 0 {
                        min_depth(map, &parent_key) + 1
                    } else {
                        0 // Default or handle as needed, though logic suggests > 0
                    };

                if let Some(parent_props) = map.get_mut(&parent_key) {
                    parent_props.score = NEGINF;
                    if parent_props.next_state.len() > 0 {
                        parent_props.lose_depth = max_ld;
                    }
                    lose_states.push(parent_key);
                }
            }
        }
    }

    loses(&lose_states, worked_history, map);
    return
}



fn loses(lose_states: &Vec<State>, worked_history: &mut HashMap<State, bool>, map: &mut GameMap) {
    if lose_states.len() == 0 {
        return;
    }
    let mut win_states: Vec<State> = vec![];

    for lose_key in lose_states{
        for parent_key in parent_of(&map, lose_key){
            if !worked_history.insert(parent_key, true).is_none() {
                if is_all_neg(&map, &map[&parent_key].next_state) {
                    win_states.push(parent_key);
                } else{
                    if map[&parent_key].score == 0.0{
                        let parent_new_score = zero_score(map, &parent_key);
                        if let Some(parent_props) = map.get_mut(&parent_key){
                            parent_props.score = parent_new_score;
                        }
                    }
                }
            }
        }

    }



    wins(&win_states, worked_history, map);
    return
}

fn min_depth(map: &GameMap, state: &State)->u8{
    let mut mind = 10;
    for &child in &map[state].next_state {
        mind = min(mind, map[&child].win_depth);
    }
    mind
}

fn max_depth(map: &GameMap, state: &State)->u8{
    let mut maxd = 0;
    for &child in &map[state].next_state {
        maxd = max(maxd, map[&child].lose_depth);
    }
    maxd
}

fn is_all_neg(map: &GameMap, children : &Vec<State>) -> bool {
    for child in children {
        if map[child].score != NEGINF{
            return false
        }
    }
    true
}

fn zero_score(map: &GameMap, state: &State)->f32{
    let mut score:f32 = 0.0;
    let children = &map[state].next_state;

    for child in children{
        let child_score = map[child].score;
        if child_score == NEGINF{
            score += 1.0;
        } else if child_score == POSINF{
            score -=1.0;
        } else{
            score -= child_score;
        }
    }
    score = score/(children.len() as f32);
    score
}