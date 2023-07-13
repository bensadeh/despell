use crate::color::Color;
use crate::emojis;
use crate::icon::Icon;
use crate::nerdfonts;
use phf::phf_map;

static DEFAULT_MAPPINGS: phf::Map<&'static str, (&'static str, Color, &'static str)> = phf_map! {
    "Python" => (nerdfonts::PYTHON, Color::Yellow, emojis::SNAKE),
    "[tmux]" => (nerdfonts::TMUX, Color::Green, emojis::NUT_AND_BOLD),
    "ack" => (nerdfonts::SEARCH, Color::Cyan, emojis::MAGNIFYING_GLASS),
    "atop" => (nerdfonts::GRAPH, Color::Yellow, emojis::MICROSCOPE),
    "bat" => (nerdfonts::ANIMAL, Color::Magenta, emojis::BAT),
    "cargo" => (nerdfonts::RUST, Color::Red, emojis::CRAB),
    "cat" => (nerdfonts::ANIMAL, Color::Red, emojis::CAT),
    "clx" => (nerdfonts::YC, Color::Orange, emojis::NEWSPAPER),
    "curl" => (nerdfonts::GLOBE, Color::Blue, emojis::GLOBE),
    "docker" => (nerdfonts::DOCKER, Color::Blue, emojis::WHALE),
    "duplicate" => (nerdfonts::DUPLICATE, Color::Yellow, emojis::LEAVES),
    "emacs" => (nerdfonts::EMACS, Color::Magenta, emojis::PEN),
    "exa" => (nerdfonts::DIRECTORIES, Color::Yellow, emojis::FOLDER),
    "fd" => (nerdfonts::SEARCH, Color::Cyan, emojis::MAGNIFYING_GLASS),
    "find" => (nerdfonts::SEARCH, Color::Cyan, emojis::MAGNIFYING_GLASS),
    "fzf" => (nerdfonts::SEARCH, Color::Cyan, emojis::MAGNIFYING_GLASS),
    "gh" => (nerdfonts::GITHUB, Color::Blue, emojis::TANABATA_TREE),
    "git" => (nerdfonts::GIT, Color::Red, emojis::WOOD),
    "glow" => (nerdfonts::MARKDOWN, Color::Magenta, emojis::STAR),
    "go" => (nerdfonts::GOLANG, Color::Cyan, emojis::HAMSTER_FACE),
    "grep" => (nerdfonts::SEARCH, Color::Cyan, emojis::MAGNIFYING_GLASS),
    "htop" => (nerdfonts::GRAPH, Color::Yellow, emojis::MICROSCOPE),
    "http" => (nerdfonts::GLOBE, Color::Blue, emojis::GLOBE),
    "java" => (nerdfonts::JAVA, Color::Red, emojis::COFFEE),
    "lazygit" => (nerdfonts::GIT, Color::Red, emojis::TANABATA_TREE),
    "less" => (nerdfonts::TEXT, Color::Magenta, emojis::BOOK_RED),
    "lf" => (nerdfonts::DIRECTORIES, Color::Yellow, emojis::FOLDER),
    "ls" => (nerdfonts::DIRECTORIES, Color::Yellow, emojis::FOLDER),
    "lua" => (nerdfonts::LUA, Color::Blue, emojis::MOON),
    "lynx" => (nerdfonts::GLOBE, Color::Blue, emojis::GLOBE),
    "man" => (nerdfonts::BOOK, Color::Magenta, emojis::BOOK_BLUE),
    "more" => (nerdfonts::TEXT, Color::Magenta, emojis::BOOK_RED),
    "mv" => (nerdfonts::MOVE, Color::Yellow, emojis::MOVING_TRUCK),
    "nano" => (nerdfonts::PENCIL, Color::Magenta, emojis::PEN),
    "nnn" => (nerdfonts::DIRECTORIES, Color::Yellow, emojis::FOLDER),
    "node" => (nerdfonts::NODE, Color::Green, emojis::GLOBE),
    "npm" => (nerdfonts::PACKAGE, Color::Red, emojis::PACKAGE),
    "nvim" => (nerdfonts::VI, Color::Green, emojis::PEN),
    "pico" => (nerdfonts::PENCIL, Color::Magenta, emojis::PEN),
    "ping" => (nerdfonts::GLOBE, Color::Blue, emojis::GLOBE),
    "ranger" => (nerdfonts::DIRECTORIES, Color::Yellow, emojis::FOLDER),
    "rg" => (nerdfonts::SEARCH, Color::Cyan, emojis::MAGNIFYING_GLASS),
    "rm" => (nerdfonts::TRASH, Color::Red, emojis::CROSS_MARK),
    "rsync" => (nerdfonts::SYNC, Color::Red, emojis::ANTICLOCKWISE_ARROWS),
    "ruby" => (nerdfonts::RUBY, Color::Red, emojis::DIAMONDS_SUIT),
    "scp" => (nerdfonts::COMPUTERS, Color::Cyan, emojis::COMPUTER),
    "sed" => (nerdfonts::TEXT, Color::Magenta, emojis::BOOK_RED),
    "sk" => (nerdfonts::SEARCH, Color::Cyan, emojis::MAGNIFYING_GLASS),
    "sleep" => (nerdfonts::HOURGLASS, Color::Cyan, emojis::ZZZ),
    "spin" => (nerdfonts::DOWN, Color::Cyan, emojis::SPIN),
    "ssh" => (nerdfonts::COMPUTERS, Color::Yellow, emojis::KEY),
    "sudo" => (nerdfonts::WARNING, Color::Red, emojis::RED_EXCLAMATION_MARK),
    "tail" => (nerdfonts::DOWN, Color::Green, emojis::DOWN_POINTING_TRIANGLE),
    "tar" => (nerdfonts::ZIP, Color::Orange, emojis::COMPRESSED),
    "top" => (nerdfonts::GRAPH, Color::Yellow, emojis::MICROSCOPE),
    "vi" => (nerdfonts::VI, Color::Green, emojis::PEN),
    "vim" => (nerdfonts::VI, Color::Green, emojis::PEN),
    "w3m" => (nerdfonts::GLOBE, Color::Blue, emojis::GLOBE),
    "wget" => (nerdfonts::GLOBE, Color::Blue, emojis::GLOBE),
};

pub fn get_icon(command: &str) -> Option<Icon> {
    DEFAULT_MAPPINGS
        .get(command)
        .map(|(nerdfont, color, emojis)| {
            Icon::new(nerdfont.to_string(), color.clone(), emojis.to_string())
        })
}
