mod cli;
mod color;
mod config;
mod defaults;
mod emojis;
mod nerdfonts;
mod types;

use crate::cli::Args;
use crate::types::{Format, Icon, MappingSource};

use color_eyre::eyre::Result;

fn main() -> Result<()> {
    color_eyre::install()?;

    let args: Args = cli::parse_args();
    let cmd_name = &args.cmd_name;
    let config_path = "~/.config/despell/config.toml";

    let mapping_selection = determine_mapping(args.use_custom_mappings);
    let format = determine_format(args.use_emoji, args.use_color);

    let icon = match mapping_selection {
        MappingSource::Custom => config::parse_config_and_get_icon(config_path, cmd_name),
        MappingSource::Default => defaults::get_icon(cmd_name).unwrap_or_default(),
    };

    let output = match format {
        Format::Emoji => icon.emoji,
        Format::Colored => format!("#[fg={}]{}", icon.color, icon.nerdfont),
        Format::Default => icon.nerdfont,
    };

    println!("{}", output);

    Ok(())
}

fn determine_mapping(use_custom_mappings: bool) -> MappingSource {
    match use_custom_mappings {
        true => MappingSource::Custom,
        false => MappingSource::Default,
    }
}

fn determine_format(use_emoji: bool, use_color: bool) -> Format {
    if use_emoji {
        return Format::Emoji;
    }

    if use_color {
        return Format::Colored;
    }

    Format::Default
}
