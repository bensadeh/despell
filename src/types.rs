use crate::color::Color;
use crate::{emojis, nerdfonts};
use serde::Deserialize;

#[derive(Deserialize, Debug, Clone)]
pub struct Icon {
    #[serde(default = "default_nerdfont")]
    pub nerdfont: String,
    #[serde(default = "default_color")]
    pub color: Color,
    #[serde(default = "default_emoji")]
    pub emoji: String,
}

impl Icon {
    pub fn new(nerdfont: String, color: Color, emoji: String) -> Self {
        Self {
            nerdfont,
            color,
            emoji,
        }
    }
}

impl Default for Icon {
    fn default() -> Self {
        Icon {
            nerdfont: default_nerdfont(),
            color: default_color(),
            emoji: default_emoji(),
        }
    }
}

fn default_nerdfont() -> String {
    nerdfonts::SHELL.to_string()
}

fn default_color() -> Color {
    Color::None
}

fn default_emoji() -> String {
    emojis::TOP_HAT.to_string()
}

pub enum MappingSelection {
    Custom,
    Default,
}

pub enum OutputSelection {
    Emoji,
    Colored,
    Default,
}
