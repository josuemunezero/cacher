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

const (
	ARRAY_TEXT = "array"
	STRING_TEXT = "string"
	BULK_TEXT = "bulk"
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

type Writer struct {
	writer io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{writer: w}
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
			fmt.Printf("Invalid type: %v", typ)
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
	v.typ = BULK_TEXT
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
	v.typ = ARRAY_TEXT
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

func (v Value) marshal() []byte {
	switch v.typ {
		case STRING_TEXT :
			return v.marshalString()
		case BULK_TEXT :
			return v.marshalBulk()
		case ARRAY_TEXT :
			return v.marshalArray()
		default :
			return []byte{}
	}
}

func (v Value) marshalString() []byte {
	var response []byte
	response = append(response, STRING)
	response = append(response, v.str...)
	response = append(response, '\r', '\n')
	return response
}

func (v Value) marshalBulk() []byte {
	var response []byte
	response = append(response, BULK)
	response = append(response, strconv.Itoa(len(v.bulk))...)
	response = append(response, '\r', '\n')
	response = append(response, v.bulk...)
	response = append(response, '\r', '\n')
	return response
}

func (v Value) marshalArray() []byte {
	var response []byte
	size := len(v.array)
	response = append(response, ARRAY)
	response = append(response, strconv.Itoa(size)...)
	response = append(response, '\r', '\n')
	for i := 0; i < size; i++ {
		response = append(response, v.marshal()...)
	}
	return response
}

func (w *Writer) write(v Value) (string, error) {
	data := v.marshal()
	fmt.Printf("Marshalled data: %v ", data)
	_, error := w.writer.Write(data)
	if(error != nil) {
		return "", error
	}
	return "Done.", nil
}
