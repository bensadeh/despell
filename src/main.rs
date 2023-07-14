use clap::Parser;

mod color;
mod defaults;
mod emojis;
mod icon;
mod nerdfonts;

#[derive(Parser)]
#[clap(version = env!("CARGO_PKG_VERSION"))]
struct Args {
    /// Command name
    #[clap(name = "COMMAND")]
    cmd_name: String,

    /// Let despell override the color of the icon
    #[clap(short = 'c', long = "color", conflicts_with = "emoji")]
    color: bool,

    /// Use emojis instead of nerdfonts
    #[clap(short = 'e', long = "emoji", conflicts_with = "color")]
    emoji: bool,

    /// Use custom mappings from ~/.config/despell/config.toml
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
