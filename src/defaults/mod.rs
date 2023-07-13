use crate::color::Color;
use crate::icon::Icon;
use crate::nerdfonts;
use phf::phf_map;

static STOCK_MAPPINGS: phf::Map<&'static str, (&'static str, Color, &'static str)> = phf_map! {
    "Python" => (nerdfonts::PYTHON, Color::Red, nerdfonts::PYTHON),
};

pub fn get_icon(command: &str) -> Option<Icon> {
    STOCK_MAPPINGS.get(command).map(|(nerdfont, color, emoji)| {
        Icon::new(nerdfont.to_string(), color.clone(), emoji.to_string())
    })
}
