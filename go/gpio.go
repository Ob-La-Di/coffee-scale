package main

import "github.com/warthog618/gpiod"

type Scale struct {
	chip *gpiod.Chip

	tare  int
	pdSck *gpiod.Line
	dout  *gpiod.Line
}

func NewScale() (Scale, error) {
	chip, err := gpiod.NewChip("gpiochip0", gpiod.WithConsumer("coffeescale"))

	if err != nil {
		return Scale{}, err
	}

	pdsck, err := chip.RequestLine(6)

	if err != nil {
		return Scale{}, err
	}

	dout, err := chip.RequestLine(6)

	if err != nil {
		return Scale{}, err
	}

	return Scale{chip: chip, pdSck: pdsck, dout: dout}, nil
}

func (s Scale) readBit() (int, error) {
	err := s.pdSck.SetValue(1)
	if err != nil {
		return 0, err
	}

	err = s.pdSck.SetValue(0)
	if err != nil {
		return 0, err
	}

	return s.dout.Value()
}

func (s Scale) readByte() (int, error) {
	result := 0

	for i := 0; i < 8; i++ {
		result <<= 1
		readBit, err := s.readBit()

		if err != nil {
			return 0, err
		}

		result |= readBit
	}

	return result, nil
}

func (s Scale) getWeight() (int, error) {
	return s.readByte()
}
