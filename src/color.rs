use std::fmt;

#[derive(Clone)]
pub enum Color {
    Red,
    Green,
    Blue,
    Yellow,
    Magenta,
    Cyan,
    Orange,
    // More colors here if needed
}

impl fmt::Display for Color {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match *self {
            Color::Red => write!(f, "color1"),
            Color::Green => write!(f, "color2"),
            Color::Blue => write!(f, "color3"),
            Color::Yellow => write!(f, "color3"),
            Color::Magenta => write!(f, "color3"),
            Color::Cyan => write!(f, "color3"),
            Color::Orange => write!(f, "color3"),
        }
    }
}
