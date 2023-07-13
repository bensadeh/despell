use std::env;

mod color;
mod defaults;
mod icon;
mod nerdfonts;

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() > 1 {
        let command = &args[1];
        if let Some(icon) = defaults::get_icon(command) {
            println!("Nerdfont: {}", icon.nerdfont);
            println!("Color: {}", icon.color);
            println!("Emoji: {}", icon.emoji);
        } else {
            // the command was not found in `STOCK_MAPPINGS`
            println!("No icon found for command: {}", command);
        }
    } else {
        println!("No command provided!");
    }
}
