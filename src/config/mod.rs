use crate::defaults;
use crate::icon::Icon;
use serde::Deserialize;
use std::collections::HashMap;
use std::path::PathBuf;
use std::{env, fs};

#[derive(Deserialize, Debug)]
pub struct CustomMappings {
    #[serde(rename = "default")]
    default_icon: Icon,

    #[serde(rename = "icons")]
    icons: HashMap<String, Icon>,
}

pub fn parse_config_and_get_icon(path: &str, command: &str) -> Icon {
    let custom_mappings = parse_config(path);

    if let Some(icon) = custom_mappings.icons.get(command) {
        return icon.clone();
    }

    if let Some(icon) = defaults::get_icon(command) {
        return icon;
    }

    custom_mappings.default_icon
}

pub fn parse_config(path: &str) -> CustomMappings {
    let absolute_path = get_path(path);

    let content = match fs::read_to_string(absolute_path) {
        Ok(content) => content,
        Err(_) => exit_with_message("Could not read the TOML configuration file"),
    };

    let custom_mappings: CustomMappings = match toml::from_str(&content) {
        Ok(icons) => icons,
        Err(e) => exit_with_message(&format!("{}", e)),
    };

    custom_mappings
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
