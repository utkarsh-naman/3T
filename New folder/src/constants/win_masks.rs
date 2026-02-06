use crate::model::state_model::State;

pub static WIN_ROW1_MASK: State = 0b111000000;
pub static WIN_ROW2_MASK: State = 0b000111000;
pub static WIN_ROW3_MASK: State = 0b000000111;
pub static WIN_COL1_MASK: State = 0b100100100;
pub static WIN_COL2_MASK: State = 0b010010010;
pub static WIN_COL3_MASK: State = 0b001001001;
pub static WIN_DIAG1_MASK: State = 0b100010001;
pub static WIN_DIAG2_MASK: State = 0b001010100;

pub static WIN_MASKS: [State; 8] = [WIN_ROW1_MASK, WIN_ROW2_MASK, WIN_ROW3_MASK, WIN_COL1_MASK, WIN_COL2_MASK, WIN_COL3_MASK, WIN_DIAG1_MASK, WIN_DIAG2_MASK];