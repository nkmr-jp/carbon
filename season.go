package carbon

import (
	"strings"
	"time"
)

// Season 获取当前季节(以气象划分)，支持i18n
func (c Carbon) Season() string {
	if c.IsZero() {
		return ""
	}
	if len(c.Lang.resources) == 0 {
		c.Lang.SetLocale(defaultLocale)
	}
	index := -1
	switch {
	case c.Month() == 3 || c.Month() == 4 || c.Month() == 5:
		index = 0
	case c.Month() == 6 || c.Month() == 7 || c.Month() == 8:
		index = 1
	case c.Month() == 9 || c.Month() == 10 || c.Month() == 11:
		index = 2
	case c.Month() == 12 || c.Month() == 1 || c.Month() == 2:
		index = 3
	}
	if seasons, ok := c.Lang.resources["seasons"]; ok {
		slice := strings.Split(seasons, "|")
		return slice[index]
	}
	return ""
}

// StartOfSeason 本季节开始时间
func (c Carbon) StartOfSeason() Carbon {
	if c.IsZero() {
		return c
	}
	if c.Month() == 1 || c.Month() == 2 {
		c.Time = time.Date(c.Year()-1, time.Month(12), 1, 0, 0, 0, 0, c.Loc)
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()/3*3), 1, 0, 0, 0, 0, c.Loc)
	return c
}

// EndOfSeason 本季节结束时间
func (c Carbon) EndOfSeason() Carbon {
	if c.IsZero() {
		return c
	}
	if c.Month() == 1 || c.Month() == 2 {
		c.Time = time.Date(c.Year(), time.Month(2), 1, 23, 59, 59, 999999999, c.Loc).AddDate(0, 1, -1)
		return c
	}
	if c.Month() == 12 {
		c.Time = time.Date(c.Year()+1, time.Month(2), 1, 23, 59, 59, 999999999, c.Loc).AddDate(0, 1, -1)
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()/3*3+2), 1, 23, 59, 59, 999999999, c.Loc).AddDate(0, 1, -1)
	return c
}

// IsSpring 是否是春季
func (c Carbon) IsSpring() bool {
	if c.IsZero() {
		return false
	}
	if c.Month() == 3 || c.Month() == 4 || c.Month() == 5 {
		return true
	}
	return false
}

// IsSummer 是否是夏季
func (c Carbon) IsSummer() bool {
	if c.IsZero() {
		return false
	}
	if c.Month() == 6 || c.Month() == 7 || c.Month() == 8 {
		return true
	}
	return false
}

// IsAutumn 是否是秋季
func (c Carbon) IsAutumn() bool {
	if c.IsZero() {
		return false
	}
	if c.Month() == 9 || c.Month() == 10 || c.Month() == 11 {
		return true
	}
	return false
}

// IsWinter 是否是冬季
func (c Carbon) IsWinter() bool {
	if c.IsZero() {
		return false
	}
	if c.Month() == 12 || c.Month() == 1 || c.Month() == 2 {
		return true
	}
	return false
}
