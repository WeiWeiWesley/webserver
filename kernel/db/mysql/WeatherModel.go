package mysql

//Weather Weather
type Weather struct {
	City    string `gorm:"column:city;NOT NULL;type:varchar(20);"`
	Content string `gorm:"column:content;type:text;"`
}
