use crate::constants::values::POSINF;
use crate::model::map::{GameMap};
use crate::model::state_model::State;

pub fn get_terminal_values(game_map: &GameMap) -> Vec<State> {
    let mut terminal_wins: Vec<State> = Vec::new();

    for key in game_map.keys() {
        if (*key>>31)&1 == 0{ // do not continue means terminal state
            if (*key>>30)&1 == 0 { // no tie so someone won
                terminal_wins.push(*key);
            }
        }
    }
    terminal_wins
}

pub fn set_terminal_values(mut map: GameMap) -> GameMap{
    let terminal_wins = get_terminal_values(&map);

    for state in terminal_wins{
        if let Some(state_props) = map.get_mut(&state) {
            state_props.score = POSINF;
            state_props.win_depth = 0;
        }
    }
    map
}