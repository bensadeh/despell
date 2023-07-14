use std::env;
use std::process::exit;

mod color;
mod defaults;
mod emojis;
mod icon;
mod nerdfonts;

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() <= 1 {
        println!("No command provided!");
        exit(0);
    }

    let command = &args[1];

    let icon = defaults::get_icon(command).unwrap_or_default();

    println!("Nerdfont: {}", icon.nerdfont);
    println!("Color: {}", icon.color);
    println!("Emoji: {}", icon.emoji);
}
