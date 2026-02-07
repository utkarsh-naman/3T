use rand::seq::SliceRandom;
use rand::thread_rng;
use crate::constants::symmetry_tables::SYMMETRY_TABLES;
use crate::constants::values::{NEGINF, POSINF};
use crate::game_engine::game_map::gmap;
use crate::model::state_model::{extract, gen_board, State};
use crate::symmetry_reduction::symmred::collapse;
// use crate::utils::next_states::next_states;
fn best_next_state(state: &State)->State{
    // let input_state = gen_board(board);
    // let state:State = collapse(&gen_board(board));


    let mut best_score:f32 = NEGINF;
    let mut best_moves: Vec<State> = Vec::new();
    let mut best_move:State = 0b0;

    let map = gmap();

    let next_states = &map.get(state).expect("state not found").next_state;
    if next_states.len() == 0 {
        println!("No next state");
        return best_move;
    }

    let mut max_lose_depth: u8 = 0;
    let mut min_win_depth: u8 = 10;
    for &next_state in next_states {
        let ns_score = map.get(&next_state).expect("no state score").score;
        max_lose_depth = max_lose_depth.max(map.get(&next_state).expect("no state score").lose_depth);
        min_win_depth = min_win_depth.min(map.get(&next_state).expect("no state score").win_depth);
        best_score = best_score.max(ns_score);
    }

    for &next_state in next_states {
        let ns_score: f32 = map.get(&next_state).expect("no state score").score;
        if ns_score == best_score {
            best_moves.push(next_state);
        }
    }

    if best_score == NEGINF{
        let mut moves: Vec<State> = Vec::new();
        for option in best_moves{
            if map.get(&option).expect("").lose_depth == max_lose_depth{
                moves.push(option);
            }
        }
        best_move = *moves.choose(&mut thread_rng()).expect("moves should not be empty");
    } else if best_score == POSINF{
        let mut moves: Vec<State> = Vec::new();
        for option in best_moves{
            if map.get(&option).expect("").win_depth == min_win_depth{
                moves.push(option);
            }
        }
        best_move = *moves.choose(&mut thread_rng()).expect("moves for infinity win should not be empty");
    } else{
        best_move = *best_moves.choose(&mut thread_rng()).expect("moves should not be empty");
    }
    //some more code later
    println!("AI best move:\t\t{:032b}", best_move);
    best_move
}




fn pick_cell(input_state: State, canonical_soln:State)->u8{
    let cell:u8 =0;
    let board: State = (1 << 18)-1;

    for eq in expand_eq(&canonical_soln){
        for i in 0..18{
            let inputboardxo = (input_state >>9) & board;
            let soluboardxo = (eq >>9) & board;

            if inputboardxo^soluboardxo == (1 << i) {
                return 10 - ( (i)%9 + 1)
            }
        }
    }
    cell
}




fn expand_eq(state: &State) -> Vec<State>{
    let mut expansions: Vec<State> = Vec::new();
    expansions.push(*state);

    let head: State = state&0b11111_000000000_000000000_000000000;
    let (_, _, _, _, _, x, o, v ) = extract(state);

    for table in SYMMETRY_TABLES{
        let hybrid = (table[x as usize] << 18) | (table[o as usize] << 9) | table[v as usize];
        expansions.push(head|hybrid);
    }
    expansions
}


pub fn ttt(xo_board:State)->u8{
    let input_full = gen_board(&xo_board);
    println!("input_full:\t\t{:032b}", input_full);

    let canonical_input = collapse(&input_full);
    println!("canonical_input:\t{:032b}", canonical_input);
    return pick_cell(input_full, best_next_state(&canonical_input));
}