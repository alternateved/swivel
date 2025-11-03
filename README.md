# swivel

Cycle focus between windows in the same Sway container.

## Installation

```sh
go install github.com/alternateved/swivel
```

## Usage

```sh
swivel           # focus next window in container
swivel --prev    # focus previous window in container
```

Bind to keys in your Sway config:

```
bindsym $mod+Tab exec swivel
bindsym $mod+Shift+Tab exec swivel --prev
```

## Details

Finds the focused window's parent container and cycles focus to the next/previous sibling window, wrapping around at boundaries. If the focused window isn't in a multi-window container, does nothing.
