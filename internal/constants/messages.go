package constants

type NewSessionMsg struct {
    Battle Battle
}

type ReloadSessionMsg struct {
    SessionPath string
}

type BackMsg struct {}
