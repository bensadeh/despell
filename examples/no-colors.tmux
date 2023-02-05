tmux_active_fg=color255
tmux_inactive_fg=default

tmux_bg=#232235

set-window-option -g window-status-separator ''
set-option -g status-style bg=$tmux_bg
set-option -g status-left ""
set-option -g status-right ""

set -g status-justify centre

# Active
set-window-option -g window-status-current-format "\
#[bg=$tmux_bg] #(despell #W)\
#[fg=$tmux_active_fg,bold bg=$tmux_bg] #W "

# Inactive
set-window-option -g window-status-format "\
#[fg=$tmux_inactive_fg,bg=$tmux_bg] #(despell #W)\
#[fg=$tmux_inactive_fg,dim bg=$tmux_bg] #W "
