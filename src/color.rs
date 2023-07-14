use serde::{Deserialize, Deserializer};
use std::fmt;

#[derive(Clone, Debug)]
pub enum Color {
    None,
    Red,
    Green,
    Blue,
    Yellow,
    Magenta,
    Cyan,
    Orange,
}

impl fmt::Display for Color {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match *self {
            Color::None => write!(f, "color0"),
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

impl<'de> Deserialize<'de> for Color {
    fn deserialize<D>(deserializer: D) -> Result<Self, D::Error>
    where
        D: Deserializer<'de>,
    {
        let color_str = String::deserialize(deserializer)?;
        match color_str.as_str() {
            "none" => Ok(Color::None),
            "red" => Ok(Color::Red),
            "green" => Ok(Color::Green),
            "blue" => Ok(Color::Blue),
            "yellow" => Ok(Color::Yellow),
            "magenta" => Ok(Color::Magenta),
            "cyan" => Ok(Color::Cyan),
            "orange" => Ok(Color::Orange),
            _ => Err(serde::de::Error::custom(format!(
                "Invalid color: {}",
                color_str
            ))),
        }
    }
}
