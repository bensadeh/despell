use crate::color::Color;
use serde::Deserialize;

pub mod emojis;
pub mod nerdfonts;

#[derive(Deserialize, Debug, Clone)]
pub struct Icon {
    #[serde(default = "Icon::default_nerdfont")]
    pub nerdfont: String,
    #[serde(default = "Icon::default_color")]
    pub color: Color,
    #[serde(default = "Icon::default_emoji")]
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

    fn default_nerdfont() -> String {
        nerdfonts::SHELL.to_string()
    }

    fn default_color() -> Color {
        Color::None
    }

    fn default_emoji() -> String {
        emojis::TOP_HAT.to_string()
    }
}

impl Default for Icon {
    fn default() -> Self {
        Icon {
            nerdfont: Icon::default_nerdfont(),
            color: Icon::default_color(),
            emoji: Icon::default_emoji(),
        }
    }
}
