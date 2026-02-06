use ttt_rust::model::state_model::{State, gen_board};
use ttt_rust::utils::next_states::{next_states};

fn main() {
    // let s1: State = 0b101000000_010000000;
    let s1: State = 0b0;
    for next_move in next_states(&gen_board(&s1)){
        println!("{:032b}", next_move);
    }
}
