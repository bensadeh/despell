use crate::color::Color;

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
