package types

type Color struct{
    Name string
    Code string
    Tw_bg string 
    Tw_border string
}

func NewColor(name string, code string) Color{
    return Color{Name: name,Code: code, Tw_bg: "bg-["+code+"]", Tw_border:"border-["+code+"]"}
} 

