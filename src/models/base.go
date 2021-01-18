package models

import (
	"database/sql/driver"
	"time"
)

/* 目前使用自定義會導致時區錯誤 要使用記得先研究怎麼修好 */
type Base struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt LocalTime
	UpdatedAt LocalTime
	DeletedAt *LocalTime
}

const TimeFormat = "2006-01-02 15:04:05"

/* 自定義time類型，使的時間輸出可依自己設定的輸出格式化  */
type LocalTime time.Time

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	// 空值不進行解析
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}

	// 指定解析的格式
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = LocalTime(now)
	return
}

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(*t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t LocalTime) String() string {
	return time.Time(t).Format(TimeFormat)
}

// 写入 mysql 时调用
func (t LocalTime) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return time.Time(t), nil
}

// 检出 mysql 时调用
func (t *LocalTime) Scan(v interface{}) error {
	// mysql 内部日期的格式可能是 2006-01-02 15:04:05 +0800 CST 格式，所以检出的时候还需要进行一次格式化
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	// tTime = v.(time.Time)
	*t = LocalTime(tTime)
	return nil
}
