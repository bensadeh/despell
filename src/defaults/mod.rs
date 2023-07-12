use crate::icon::Icon;
use crate::nerdfonts;
use phf::phf_map;

static STOCK_MAPPINGS: phf::Map<&'static str, (&'static str, &'static str, &'static str)> = phf_map! {
      "Python" => (nerdfonts::PYTHON, nerdfonts::PYTHON, nerdfonts::PYTHON)
};

pub fn get_icon(command: &str) -> Option<Icon> {
    STOCK_MAPPINGS
        .get(command)
        .map(|&(nerdfont, color, emoji)| {
            Icon::new(nerdfont.to_string(), color.to_string(), emoji.to_string())
        })
}
