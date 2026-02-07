use crate::symmetry_tables::SYMMETRY_TABLES;
use crate::state_model::{extract, State};

pub fn collapse(state: &State) -> State{
    let head:u32 = *state&0b11111_000000000_000000000_000000000;
    let (_, _, _, _, _, x, o, v )= extract(state);

    let mut canonical:u32 = ((x<<18) | (o<<9) | v)&0b111111111_111111111_111111111;

    for table in SYMMETRY_TABLES{
        let hybrid = (table[x as usize] << 18) | (table[o as usize] << 9) | table[v as usize];
        canonical = canonical.min(hybrid);
    }

    head|canonical
}