package handlers


type Day struct {
    Number       int    `json:"number"`
    NameoftheDay string `json:"nameoftheday"`
    Lessons      string `json:"lessons"`
    HomeWork     string `json:"homework"`
    Leisure      string `json:"leisure"`
}

var Days = []Day{
	{NameoftheDay: "Monday", Lessons: "Golang\nElectronics&DD\nKazakh", HomeWork: "Optional(Which lesson has to be done immediately)", Leisure: "None"},
	{NameoftheDay: "Tuesday", Lessons: "OOP", HomeWork: "Web Development", Leisure: "Drawing Hobby"},
	{NameoftheDay: "Wednesday", Lessons: "WebDev\nOOP\nKazakh", HomeWork: "Optional(Which lesson has to be done immediately)", Leisure: "None"},
	{NameoftheDay: "Thursday", Lessons: "WebDev\nElectronics&DD", HomeWork: "ICT", Leisure: "None"},
	{NameoftheDay: "Friday", Lessons: "Kazakh", HomeWork: "Electronics&DD", Leisure: "Drawing Hobby"},
	{NameoftheDay: "Saturday", Lessons: "ICT\nGolang", HomeWork: "OOP", Leisure: "None"},
	{NameoftheDay: "Sunday", Lessons: "ICT", HomeWork: "Golang", Leisure: "Football: Liverpool v Arsenal"},
}

func AllWeeklyTasks() []Day{
	return Days
}

func BytheDay(n int) (Day, bool) {
	for i := 0; i < len(Days); i++ {
		if Days[i].Number == n {
			return Days[i], true
		}
	}
	return Day{}, false
}