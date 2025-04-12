package tui

type Model struct {
    files         []MP3File       // List of loaded MP3 files
    selectedIndex int             // Index of currently selected file
    batchMode     bool            // Whether batch editing is enabled
    editingField  TagField        // The currently edited tag field
    inputBuffer   string          // Input text buffer for editing
    fetchStatus   FetchStatus     // Status of auto-fetching tags
    showLyrics    bool            // Toggle lyrics view
    showArtwork   bool            // Toggle artwork view
    uiState       UIState         // Controls animations and active views
}

type MP3File struct {
    Path       string            // File path
    Metadata   ID3Tags           // Metadata tags
    Lyrics     string            // Embedded or external lyrics
    Artwork    ImageData         // Cover art
    Selected   bool              // For batch operations
}

type ID3Tags struct {
    Title       string
    Artist      string
    Album       string
    Year        string
    Genre       string
    Track       int
    AlbumArtist string
    Comment     string
}

type ImageData struct {
    Raw      []byte // Image file bytes
    Format   string // "jpeg", "png", etc.
    Preview  string // ASCII/terminal preview
}

type UIState struct {
    ActivePanel string  // "metadata", "artwork", "lyrics"
    ShowHelp    bool    // Toggle help menu
    Animating   bool    // Animation status
}

type FetchStatus struct {
    Fetching    bool
    Success     bool
    ErrorMsg    string
}

type TagField string

const (
    FieldTitle       TagField = "Title"
    FieldArtist      TagField = "Artist"
    FieldAlbum       TagField = "Album"
    FieldYear        TagField = "Year"
    FieldGenre       TagField = "Genre"
    FieldTrack       TagField = "Track"
    FieldAlbumArtist TagField = "Album Artist"
    FieldComment     TagField = "Comment"
)
