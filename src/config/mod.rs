use crate::color::Color;
use crate::icon::Icon;
use serde::Deserialize;
use std::collections::HashMap;
use std::path::PathBuf;
use std::{env, fs};

#[derive(Deserialize, Debug, Clone)]
pub struct CustomIcon {
    nerdfont: Option<String>,
    color: Option<String>,
    emoji: Option<String>,
}

#[derive(Deserialize, Debug)]
pub struct Icons {
    default_icon: CustomIcon,
    custom_icons: HashMap<String, CustomIcon>,
}

impl Icons {
    pub fn get_icon(&self, command: &str) -> Icon {
        let icon = self
            .custom_icons
            .get(command)
            .unwrap_or(&self.default_icon)
            .clone();

        Icon {
            nerdfont: icon
                .nerdfont
                .unwrap_or_else(|| self.default_icon.nerdfont.clone().unwrap()),
            color: Color::None,
            emoji: icon
                .emoji
                .unwrap_or_else(|| self.default_icon.emoji.clone().unwrap()),
        }
    }
}

pub fn parse_config_and_get_icon(path: &str, command: &str) -> Icon {
    let expanded_path = if path.starts_with('~') {
        let home = env::var("HOME").expect("The HOME environment variable is not set");
        PathBuf::from(home).join(path.strip_prefix('~').expect("Invalid path"))
    } else {
        PathBuf::from(path)
    };

    let content =
        fs::read_to_string(expanded_path).expect("Failed to read the TOML configuration file");

    let icons: Icons =
        toml::from_str(&content).expect("Failed to parse the TOML configuration file");

    icons.get_icon(command)
}
