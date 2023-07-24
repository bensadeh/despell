use crate::icon::Icon;
use clap::Parser;

mod color;
mod config;
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
    #[clap(short = 'c', long = "color", conflicts_with = "use_emoji")]
    use_color: bool,

    /// Use emojis instead of nerdfonts
    #[clap(short = 'e', long = "emoji", conflicts_with = "use_color")]
    use_emoji: bool,

    /// Use custom mappings from ~/.config/despell/config.toml
    #[clap(short = 'u', long = "custom")]
    use_custom_mappings: bool,
}

fn main() {
    let args: Args = Args::parse();
    let cmd_name = &args.cmd_name;
    let config_path = "~/.config/despell/config.toml";

    let icon = if args.use_custom_mappings {
        config::parse_config_and_get_icon(config_path, cmd_name)
    } else {
        defaults::get_icon(cmd_name).unwrap_or_default()
    };

    let output = get_output(args.use_emoji, args.use_color, icon);

    println!("{}", output);
}

fn get_output(use_emoji: bool, use_color: bool, icon: Icon) -> String {
    if use_emoji {
        return icon.emoji;
    }

    if use_color {
        return format!("#[fg={}]{}", icon.color, icon.nerdfont);
    }

    icon.nerdfont
}
