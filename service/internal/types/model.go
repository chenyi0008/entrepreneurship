package types

type User struct {
	Id       int64
	Account  string
	Password string
}

type AlarmEquipment struct {
	Id   int64
	Name string
}

type InspectEquipment struct {
	Id       int64
	Status   string
	Name     string
	Location string
	UserId   int64
}

type History struct {
	Id        int64
	Position  string
	Notes     string
	VideoId   string
	PicId     string
	InspectId int64
}

type SmsNum struct {
	Id      int64
	IsAlarm bool
	Num     int64
}

type InspectAlarm struct {
	InspectId    int64
	AlarmId      int64
	IsBuzzing    bool
	IsFlashing   bool
	ShowLocation bool
}

type AlarmNum struct {
	Id      int64
	IsAlarm bool
	Num     int64
}
