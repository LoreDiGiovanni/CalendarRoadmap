package types

type Events struct{
    Title string `json:"title"`
    Notes string `json:"notes"`
    Date string `json:"date"`
    Time_start string `json:"time_start"`
    Time_end string `json:"time_end"`
    Tags []string 
    Dot_color Color
}
