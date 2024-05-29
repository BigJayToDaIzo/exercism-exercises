package clock

import "fmt"

type Clock struct {
	Hour   int
	Minute int
}

func New(h, m int) Clock {
	if m >= 60 {
		h += m / 60
		m = m % 60
	} else if m == -60 {
		h--
		m = 0
	} else if m < 0 {
		h -= 1 - m/60
		m = 60 - m%60*-1
	}
	if h == 24 {
		h = 0
	}
	if h > 24 {
		h = h % 24
	}
	if h < 0 {
		h = 24 + h%24
	}
	return Clock{h, m}
}

func (c Clock) Add(m int) Clock {
	c.Minute += m
	if c.Minute >= 60 {
		c.Hour += c.Minute / 60
		c.Minute = c.Minute % 60
	}
	c.Hour = c.Hour % 24
	return c
}

func (c Clock) Subtract(m int) Clock {
	c.Minute -= m
	if c.Minute < 0 {
		c.Hour += c.Minute/60 - 1
		c.Minute = 60 + c.Minute%60
	}
	if c.Hour < 0 {
		c.Hour = 24 + c.Hour%24
	} else if c.Hour > 24 {
		c.Hour = c.Hour % 24
	}
	if c.Hour == 24 {
		c.Hour = 0
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}
