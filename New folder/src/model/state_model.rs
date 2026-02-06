pub type State = u32;

// use std::env::var;
// use crate::constants::win_masks;
use crate::constants::win_masks::WIN_MASKS;

static FILTER9: u32 = 0b111111111;

pub fn extract(state: &State) -> (State, State, State, State, State, State, State, State){
    let continuum = (*state >> 31) & 1;
    let tie = (*state >> 30) & 1;
    let won = (*state >> 29) & 1;
    let lose = (*state >> 28) & 1;
    let turn = (*state >> 27) & 1;
    let x = (*state >> 18) & FILTER9;
    let o = (*state >> 9) & FILTER9;
    let vacancy = (*state) & FILTER9;

    (continuum, tie, won, lose, turn, x, o, vacancy)
}

pub fn gen_board(state: &State) -> State{
    let o = *state&FILTER9;
    let x = (*state>>9)&FILTER9;
    let vacancy = !(x|o)&FILTER9;
    let continuum = if_continue(&x, &o, &vacancy);
    let tie:State = if game_tie(&x, &o, &vacancy){1} else{0};

    let mut who_won:State = 0;
    let mut who_lose:State = 0;

    if continuum == 0 && tie == 0{
        if player_won(&x){
            who_won = 0;
            who_lose = 1;
        } else{
            who_won = 1;
            who_lose = 0;
        }
    }

    let turn = whose_turn(&vacancy);


    let full_state = (continuum<<31) | (tie<<30) | (who_won<<29) | (who_lose<<28) | (turn<<27) | (x<<18) | (o<<9) | vacancy;

    full_state
}

fn if_continue(x: &State, o: &State, vacancy: &State)-> State{
    if *vacancy == 0 || player_won(x) || player_won(o){
        return 0
    }
    1
}

fn player_won(player: &State) -> bool{
    for mask in WIN_MASKS{
        if *player&mask == mask{
            return true
        }
    }
    false
}

fn game_tie(x: &State, o: &State, vacancy: &State) -> bool{
    if *vacancy == 0 && !player_won(x) && !player_won(o){
        return true
    }
    false
}

fn whose_turn (vcacancy: &State) -> State{
    if *vcacancy == 0{
        return 1
    }
    if (*vcacancy&(*vcacancy-1))^1 == 0{
        return 1
    }
    0
}