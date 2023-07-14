use clap::Parser;

mod color;
mod defaults;
mod emojis;
mod icon;
mod nerdfonts;

#[derive(Parser)]
struct Args {
    /// Command name
    #[clap(name = "COMMAND")]
    cmd_name: String,

    /// Follow (tail) the contents of the file
    #[clap(short = 'c', long = "color")]
    color: bool,

    /// Provide a custom path configuration file
    #[clap(short = 'e', long = "emoji")]
    emoji: bool,

    /// Provide a custom path configuration file
    #[clap(short = 'u', long = "custom")]
    custom: bool,
}
fn main() {
    let args: Args = Args::parse();
    let cmd_name = args.cmd_name;

    let icon = defaults::get_icon(&cmd_name).unwrap_or_default();

    println!("Nerdfont: {}", icon.nerdfont);
    println!("Color: {}", icon.color);
    println!("Emoji: {}", icon.emoji);
}
