use crate::color::Color;
use crate::{emojis, nerdfonts};

pub struct Icon {
    pub nerdfont: String,
    pub color: Color,
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
            nerdfont: nerdfonts::SHELL.to_string(),
            color: Color::None,
            emoji: emojis::TOP_HAT.to_string(),
        }
    }
}
