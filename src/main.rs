mod cli;
mod color;
mod config;
mod defaults;
mod emojis;
mod icon;
mod nerdfonts;
mod types;

use crate::cli::Args;
use crate::icon::Icon;
use crate::types::{MappingSelection, OutputSelection};

use color_eyre::eyre::Result;

fn main() -> Result<()> {
    color_eyre::install()?;

    let args: Args = cli::parse_args();
    let cmd_name = &args.cmd_name;
    let config_path = "~/.config/despell/config.toml";

    let mapping_selection = determine_mapping(args.use_custom_mappings);
    let output_selection = determine_output_selection(args.use_emoji, args.use_color);

    let icon = match mapping_selection {
        MappingSelection::Custom => config::parse_config_and_get_icon(config_path, cmd_name),
        MappingSelection::Default => defaults::get_icon(cmd_name).unwrap_or_default(),
    };

    let output = get_output(output_selection, icon);

    println!("{}", output);

    Ok(())
}

fn determine_mapping(use_custom_mappings: bool) -> MappingSelection {
    match use_custom_mappings {
        true => MappingSelection::Custom,
        false => MappingSelection::Default,
    }
}

fn determine_output_selection(use_emoji: bool, use_color: bool) -> OutputSelection {
    if use_emoji {
        return OutputSelection::Emoji;
    }

    if use_color {
        return OutputSelection::Colored;
    }

    OutputSelection::Default
}

fn get_output(selection: OutputSelection, icon: Icon) -> String {
    match selection {
        OutputSelection::Emoji => icon.emoji,
        OutputSelection::Colored => format!("#[fg={}]{}", icon.color, icon.nerdfont),
        OutputSelection::Default => icon.nerdfont,
    }
}
