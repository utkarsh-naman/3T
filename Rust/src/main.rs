use ttt_rust::game_engine::ttt_engine::ttt;
// use ttt_rust::mapbuilder::build_initial::gen_map;
// use ttt_rust::mapbuilder::set_terminal_values::set_terminal_values;
// use ttt_rust::mapbuilder::set_values::set_values;
// use ttt_rust::model::map::{load_map_from_file, print_map, save_map_to_file};
// use ttt_rust::model::state_model::{State, gen_board};


fn main() {

    // CREATING THE MAP0

    // let s1: State = 0b101000000_010000000;
    // let mut s1: State = 0b0;
    // s1 = gen_board(&s1);
    // println!(" full board: {:032b}", s1);
    // let gmap = gen_map(&s1);
    // print_map(&gmap);
    // println!("length of map: {}", gmap.len());
    // save_map_to_file(&gmap, "bin/output/map/unvalued/map0.ttt").expect("TODO: panic message");



    // LOADING BACK THE MAP

    // let loaded_gmap = load_map_from_file("bin/output/map/unvalued/map0.ttt").expect("failed to load file.");
    // print_map(&loaded_gmap);
    // println!("length of map: {}", loaded_gmap.len());



    // SETTING TERMINAL STATES

    // let termmap = set_terminal_values(load_map_from_file("bin/output/map/unvalued/map0.ttt").expect("failed to load file."));
    // save_map_to_file(&termmap, "bin/output/map/valued/map1.ttt").expect("TODO: panic message");



    // SETTING ALL VALS

    // let gmap1 = load_map_from_file("bin/output/map/valued/map1.ttt").expect("failed to load file.");
    // // print_map(&loaded_gmap);
    // let gmap2 = set_values(gmap1);
    // print_map(&gmap2);
    // println!("length of map: {}", gmap2.len());
    // save_map_to_file(&gmap2, "bin/output/map/valued/map2.ttt").expect("failed to save file.");


    // TEST THE ENGINE
    let xoboard = 0b000000001_000000010;
    println!("Cell chosen: {}", ttt(xoboard));

}
