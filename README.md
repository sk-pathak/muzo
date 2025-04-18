```
mp3tagger/
├── main.go                 # Entry point initializes and runs app
├── config/
│   ├── config.go           # Configuration handling and persistence
│   └── themes.go           # UI theme definitions
├── tui/
│   ├── model.go            # Data model structure definitions
│   ├── update.go           # Event handling and state updates
│   ├── view.go             # UI rendering functions
│   └── styles.go           # UI style definitions using lipgloss
├── mp3/
│   ├── tags.go             # MP3 tag reading and writing
│   ├── artwork.go          # Artwork extraction and embedding
│   └── lyrics.go           # Lyrics extraction and embedding
├── api/
│   ├── artwork_finder.go   # API client for artwork search
│   └── lyrics_finder.go    # API client for lyrics search
├── util/
│   ├── image.go            # Image processing utilities
│   └── file.go             # File handling utilities
└── assets/
    └── default_artwork.png # Default artwork when none is found

```
## Features
- [ ] manual edit MP3 tags including lyrics, artwork
- [ ] automatically fetch and embed lyrics and artwork from APIs
- [ ] batch processing of MP3 files
- [ ] image preview
- [ ] animated UI with lipgloss