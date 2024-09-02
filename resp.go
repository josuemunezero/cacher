package main

import (
	"io"
	"fmt"
	"bufio"
	"strconv"
)

const (
	STRING = '+'
	ERROR = '-'
	INTEGER = ':'
	BULK = '$'
	ARRAY = '*'
)

type Value struct {
	typ string
	str string
	num int
	bulk string
	array []Value
}

type Resp struct {
	reader *bufio.Reader
}

func newResp(r io.Reader) *Resp{
	return &Resp{reader: bufio.NewReader(r)}
}

func (r *Resp) read() (Value, error) {
	typ, err := r.reader.ReadByte()
	if err != nil {
		fmt.Println("Error while reading Type")
		return Value{},err
	}
	switch typ {
		case BULK :
			return r.readBulk()
		case ARRAY :
			return r.readArray()
		default :
			fmt.Printf("Invalid type", typ)
			return Value{}, nil
	}
}

func (r *Resp) readLine() (line []byte, err error) {
	for {
		nextVal, err := r.reader.ReadByte()
		if err != nil {
			return nil, err
		}
		if nextVal == '\r' {
			r.reader.ReadByte()
			break
		}
		line = append(line, nextVal)
	}
	return line, nil
}


func (r *Resp) readInteger() (int , error) {
	line, err := r.readLine()
	if err != nil {
		return 0, err
	}

	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, err
	}
	return int(i64), nil
}

func (r *Resp) readBulk() (Value, error)  {
	v := Value{}
	v.typ = "BULK"
	size, err := r.readInteger()
	if err != nil {
		return v, err
	}

	line := make([]byte, size)
	r.reader.Read(line)
	v.bulk = string(line)
	r.readLine()
	return v, nil
}

func (r *Resp) readArray() (Value, error) {
	v := Value{}
	v.typ = "ARRAY" 
	size, err := r.readInteger()
	if err != nil {
		return Value{}, err
	}

	for size > 0 {
		newVal, err := r.read()
		if err != nil {
			return Value{}, err
		}
		v.array = append(v.array, newVal)
		size--
	}
	return v, nil
}
