pub struct Icon {
    pub nerdfont: String,
    pub color: String,
    pub emoji: String,
}

impl Icon {
    pub fn new(nerdfont: String, color: String, emoji: String) -> Self {
        Self {
            nerdfont,
            color,
            emoji,
        }
    }
}
