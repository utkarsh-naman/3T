use crate::model::state_model::{extract, gen_board, State};
use crate::symmetry_reduction::symmred::collapse;

pub fn next_states(state: &State) ->Vec<State>{
    let mut next_moves:Vec<State> = Vec::new();
    let (continuum, _, _, _, turn, x, o, vacancy) = extract(&state);
    
    if continuum == 0{
        return next_moves;
    }

    let mut i0mask:u32;
    for i in 0..9{
        i0mask = 1<<i;
        if vacancy&i0mask != 0 {
            let next_state: State;

            if turn == 0 {
                let xoboard: State = ((x^i0mask)<<9) | o;
                next_state = collapse(&gen_board(&xoboard));
            } else {
                let xoboard: State = (x << 9) | (o^i0mask);
                next_state = collapse(&gen_board(&xoboard));
            }

            if !next_moves.contains(&next_state){
                next_moves.push(next_state);
            }
        }
    }
    next_moves
}