package msgpack

// Auto-generated by internal/cmd/gendecoder-numeric/gendecoder-numeric.go. DO NOT EDIT!

import (
	"math"

	"github.com/pkg/errors"
)

func (d *Decoder) DecodeInt16(v *int16) error {
	code, err := d.ReadCode()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code`)
	}

	if code != Int16 {
		return errors.Errorf(`msgpack: expected Int16, got %s`, code)
	}

	x, err := d.src.ReadUint16()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read int16`)
	}

	*v = int16(x)
	return nil
}

func (d *Decoder) DecodeInt32(v *int32) error {
	code, err := d.ReadCode()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code`)
	}

	if code != Int32 {
		return errors.Errorf(`msgpack: expected Int32, got %s`, code)
	}

	x, err := d.src.ReadUint32()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read int32`)
	}

	*v = int32(x)
	return nil
}

func (d *Decoder) DecodeUint(v *uint) error {
	code, err := d.ReadCode()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code`)
	}

	if code != Uint64 {
		return errors.Errorf(`msgpack: expected Uint64, got %s`, code)
	}

	x, err := d.src.ReadUint64()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read uint`)
	}

	*v = uint(x)
	return nil
}

func (d *Decoder) DecodeUint8(v *uint8) error {
	code, err := d.ReadCode()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code`)
	}

	if code != Uint8 {
		return errors.Errorf(`msgpack: expected Uint8, got %s`, code)
	}

	x, err := d.src.ReadUint8()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read uint8`)
	}

	*v = x
	return nil
}

func (d *Decoder) DecodeUint32(v *uint32) error {
	code, err := d.ReadCode()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code`)
	}

	if code != Uint32 {
		return errors.Errorf(`msgpack: expected Uint32, got %s`, code)
	}

	x, err := d.src.ReadUint32()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read uint32`)
	}

	*v = x
	return nil
}

func (d *Decoder) DecodeInt(v *int) error {
	code, err := d.ReadCode()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code`)
	}

	if code != Int64 {
		return errors.Errorf(`msgpack: expected Int64, got %s`, code)
	}

	x, err := d.src.ReadUint64()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read int`)
	}

	*v = int(x)
	return nil
}

func (d *Decoder) DecodeInt64(v *int64) error {
	code, err := d.ReadCode()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code`)
	}

	if code != Int64 {
		return errors.Errorf(`msgpack: expected Int64, got %s`, code)
	}

	x, err := d.src.ReadUint64()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read int64`)
	}

	*v = int64(x)
	return nil
}

func (d *Decoder) DecodeUint16(v *uint16) error {
	code, err := d.ReadCode()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code`)
	}

	if code != Uint16 {
		return errors.Errorf(`msgpack: expected Uint16, got %s`, code)
	}

	x, err := d.src.ReadUint16()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read uint16`)
	}

	*v = x
	return nil
}

func (d *Decoder) DecodeUint64(v *uint64) error {
	code, err := d.ReadCode()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code`)
	}

	if code != Uint64 {
		return errors.Errorf(`msgpack: expected Uint64, got %s`, code)
	}

	x, err := d.src.ReadUint64()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read uint64`)
	}

	*v = x
	return nil
}

func (d *Decoder) DecodeInt8(v *int8) error {
	code, err := d.ReadCode()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code`)
	}

	if code != Int8 {
		return errors.Errorf(`msgpack: expected Int8, got %s`, code)
	}

	x, err := d.src.ReadUint8()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read int8`)
	}

	*v = int8(x)
	return nil
}

func (d *Decoder) DecodeFloat32(v *float32) error {
	code, err := d.ReadCode()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code`)
	}

	if code != Float {
		return errors.Errorf(`msgpack: expected Float, got %s`, code)
	}

	x, err := d.src.ReadUint32()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read float32`)
	}

	*v = math.Float32frombits(x)
	return nil
}

func (d *Decoder) DecodeFloat64(v *float64) error {
	code, err := d.ReadCode()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code`)
	}

	if code != Double {
		return errors.Errorf(`msgpack: expected Double, got %s`, code)
	}

	x, err := d.src.ReadUint64()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read float64`)
	}

	*v = math.Float64frombits(x)
	return nil
}
