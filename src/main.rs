mod cli;
mod color;
mod config;
mod defaults;
mod icon;
mod types;

use crate::cli::Args;
use crate::types::{Format, Source};

use color_eyre::eyre::Result;

fn main() -> Result<()> {
    color_eyre::install()?;

    let args: Args = cli::parse_args();
    let cmd_name = &args.cmd_name;
    let config_path = "~/.config/despell/config.toml";

    let (source, format) =
        get_source_and_format(args.use_custom_mappings, args.use_emoji, args.use_color);

    let icon = match source {
        Source::Custom => config::parse_config_and_get_icon(config_path, cmd_name),
        Source::Default => defaults::get_icon(cmd_name).unwrap_or_default(),
    };

    let output = match format {
        Format::Emoji => icon.emoji,
        Format::Colored => format!("#[fg={}]{}", icon.color, icon.nerdfont),
        Format::Default => icon.nerdfont,
    };

    println!("{}", output);

    Ok(())
}

fn get_source_and_format(
    use_custom_mappings: bool,
    use_emoji: bool,
    use_color: bool,
) -> (Source, Format) {
    let mapping_source = match use_custom_mappings {
        true => Source::Custom,
        false => Source::Default,
    };

    let format = if use_emoji {
        Format::Emoji
    } else if use_color {
        Format::Colored
    } else {
        Format::Default
    };

    (mapping_source, format)
}
