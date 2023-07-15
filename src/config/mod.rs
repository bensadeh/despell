use crate::color::Color;
use crate::icon::Icon;
use serde::Deserialize;
use std::collections::HashMap;
use std::path::PathBuf;
use std::{env, fs};

#[derive(Deserialize, Debug, Clone)]
pub struct CustomIcon {
    nerdfont: Option<String>,
    color: Option<Color>,
    emoji: Option<String>,
}

#[derive(Deserialize, Debug)]
pub struct Icons {
    #[serde(rename = "default")]
    default_icon: CustomIcon,

    #[serde(rename = "icons")]
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
            color: icon
                .color
                .unwrap_or_else(|| self.default_icon.color.clone().unwrap()),
            emoji: icon
                .emoji
                .unwrap_or_else(|| self.default_icon.emoji.clone().unwrap()),
        }
    }
}

pub fn parse_config_and_get_icon(path: &str, command: &str) -> Icon {
    let absolute_path = get_path(path);

    let content = match fs::read_to_string(absolute_path) {
        Ok(content) => content,
        Err(_) => exit_with_message("Could not read the TOML configuration file"),
    };

    let icons: Icons = match toml::from_str(&content) {
        Ok(icons) => icons,
        Err(e) => exit_with_message(&format!("{}", e)),
    };

    icons.get_icon(command)
}

fn get_path(path: &str) -> PathBuf {
    if !path.starts_with('~') {
        return PathBuf::from(path);
    }

    let home = match env::var("HOME") {
        Ok(home) => home,
        Err(_) => exit_with_message("The HOME environment variable is not set"),
    };

    let path_without_prefix = match path.strip_prefix("~/") {
        Some(path) => path,
        None => exit_with_message("Could not strip ~ from path"),
    };

    PathBuf::from(home).join(path_without_prefix)
}

fn exit_with_message(message: &str) -> ! {
    eprintln!("{}", message);
    std::process::exit(1);
}
