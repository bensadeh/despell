use clap::Parser;

#[derive(Parser)]
#[clap(version = env!("CARGO_PKG_VERSION"))]
pub struct Args {
    /// Command name
    #[clap(name = "COMMAND")]
    pub cmd_name: String,

    /// Let despell override the color of the icon
    #[clap(short = 'c', long = "color", conflicts_with = "use_emoji")]
    pub use_color: bool,

    /// Use emojis instead of nerdfonts
    #[clap(short = 'e', long = "emoji", conflicts_with = "use_color")]
    pub use_emoji: bool,

    /// Use custom mappings from ~/.config/despell/config.toml
    #[clap(short = 'u', long = "custom")]
    pub use_custom_mappings: bool,
}

pub fn parse_args() -> Args {
    Args::parse()
}
