package types

type Events struct{
    ID string `json:"id" bson:"_id,omitempty"`
    Owner string `json:"owner" bson:"owner"`
    Title string `json:"title" bson:"title"`
    Notes string `json:"notes" bson:"notes"`
    Date string `json:"date" bson:"date"`
    Time_start string `json:"time_start" bson:"time_start"`
    Time_end string `json:"time_end" bson:"time_end"`
    Tags []string `json:"tags" bson:"tags"`
    Dot_color string `json:"dot_color" bson:"dot_color"`
}
