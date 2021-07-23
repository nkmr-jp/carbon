package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_DaysInYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-08-05", 366},
		{7, "2021-08-05", 365},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DaysInYear(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DaysInMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-01-05", 31},
		{7, "2020-02-05", 29},
		{8, "2020-03-05", 31},
		{9, "2020-04-05", 30},
		{10, "2020-05-05", 31},
		{11, "2021-06-05", 30},
		{12, "2021-07-05", 31},
		{13, "2021-08-05", 31},
		{14, "2021-09-05", 30},
		{15, "2021-10-05", 31},
		{16, "2021-11-05", 30},
		{17, "2021-12-05", 31},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DaysInMonth(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_MonthOfYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-01-05", 1},
		{7, "2020-02-05", 2},
		{8, "2020-03-05", 3},
		{9, "2020-04-05", 4},
		{10, "2020-05-05", 5},
		{11, "2021-06-05", 6},
		{12, "2021-07-05", 7},
		{13, "2021-08-05", 8},
		{14, "2021-09-05", 9},
		{15, "2021-10-05", 10},
		{16, "2021-11-05", 11},
		{17, "2021-12-05", 12},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.MonthOfYear(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DayOfYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-01-01", 1},
		{7, "2020-01-31", 31},
		{8, "2020-08-05", 218},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DayOfYear(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DayOfMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-01-01", 1},
		{7, "2020-01-31", 31},
		{8, "2020-08-05", 5},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DayOfMonth(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_DayOfWeek(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-08-03", 1},
		{7, "2020-08-04", 2},
		{8, "2020-08-05", 3},
		{9, "2020-08-06", 4},
		{10, "2020-08-07", 5},
		{11, "2020-08-08", 6},
		{12, "2020-08-09", 7},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DayOfWeek(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_WeekOfYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2019-12-31", 1},
		{7, "2020-01-01", 1},
		{8, "2020-01-31", 5},
		{9, "2020-02-28", 9},
		{10, "2020-01-29", 5},
		{11, "2020-08-05", 32},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.WeekOfYear(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_WeekOfMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-01-01", 1},
		{7, "2020-01-31", 4},
		{8, "2020-02-28", 1},
		{9, "2020-01-29", 2},
		{10, "2020-08-05", 1},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.WeekOfMonth(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Century(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-08-05", 21},
		{7, "2020-08-05 13:14:15", 21},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Century(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Decade(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2019-08-05", 10},
		{7, "2019-08-05 13:14:15", 10},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Decade(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Year(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-08-05", 2020},
		{7, "2020-08-05 13:14:15", 2020},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Year(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Quarter(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-01-05", 1},
		{7, "2020-04-05", 2},
		{8, "2020-08-05", 3},
		{9, "2020-11-05", 4},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Quarter(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Month(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-08-05 13:14:15", 8},
		{7, "2020-08-05", 8},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Month(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Week(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", -1},
		{2, "0", -1},
		{3, "0000-00-00", -1},
		{4, "00:00:00", -1},
		{5, "0000-00-00 00:00:00", -1},

		{6, "2020-08-03", 1},
		{7, "2020-08-04", 2},
		{8, "2020-08-05", 3},
		{9, "2020-08-06", 4},
		{10, "2020-08-07", 5},
		{11, "2020-08-08", 6},
		{12, "2020-08-09", 0},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Week(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Day(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-08-05 13:14:15", 5},
		{7, "2020-08-05", 5},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Day(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Hour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-08-05 13:14:15", 13},
		{7, "2020-08-05", 0},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Hour(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Minute(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-08-05 13:14:15", 14},
		{7, "2020-08-05", 0},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Minute(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Second(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-08-05 13:14:15", 15},
		{7, "2020-08-05", 0},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Second(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Millisecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-08-05 13:14:15.999", 999},
		{7, "2020-08-05", 0},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Millisecond(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Microsecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-08-05 13:14:15", 0},
		{7, "2020-08-05", 0},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Microsecond(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Nanosecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, "2020-08-05 13:14:15", 0},
		{7, "2020-08-05", 0},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Nanosecond(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Timezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, PRC, "CST"},
		{2, Tokyo, "JST"},
	}

	for _, test := range tests {
		c := Now(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Timezone(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Location(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, PRC, PRC},
		{2, Tokyo, Tokyo},
	}

	for _, test := range tests {
		c := Now(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Location(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Offset(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, PRC, 28800},
		{2, Tokyo, 32400},
	}

	for _, test := range tests {
		c := Now(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Offset(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Locale(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected string // 期望值
	}{
		{1, "en", "en"},
		{2, "zh-CN", "zh-CN"},
	}

	for _, test := range tests {
		c := SetLocale(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Locale(), "Current test id is "+strconv.Itoa(test.id))
	}
}

func TestCarbon_Age(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		id       int    // 测试id
		input    string // 输入值
		expected int    // 期望值
	}{
		{1, "", 0},
		{2, "0", 0},
		{3, "0000-00-00", 0},
		{4, "00:00:00", 0},
		{5, "0000-00-00 00:00:00", 0},

		{6, Now().AddYears(18).ToDateTimeString(), 0},
		{7, Now().SubYears(18).ToDateTimeString(), 18},
	}

	for _, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Age(), "Current test id is "+strconv.Itoa(test.id))
	}
}
