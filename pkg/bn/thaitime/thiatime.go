package thaitime

import "time"

type ThaiTime struct {
	Location *time.Location
}

func NewThaiTime() (*ThaiTime, error) {
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return nil, err
	}
	return &ThaiTime{Location: location}, nil
}

func (t *ThaiTime) GetBinanceStartOfDayDateTime() time.Time {
	_t := t.GetNow()
	return time.Date(_t.Year(), _t.Month(), _t.Day(), 7, 0, 0, 0, t.Location)
}

func (t *ThaiTime) GetBinanceEndOfDayDateTime() time.Time {
	_t := t.GetNow()
	return time.Date(_t.Year(), _t.Month(), _t.Day(), 24+6, 59, 59, 0, t.Location)
}

func (t *ThaiTime) GetNow() time.Time {
	return time.Now()
}
