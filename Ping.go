package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type ICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func getICMP() ICMP {
	icmp := ICMP{
		Type:       8,
		Code:       0,
		CheckSum:   0,
		Identifier: 0,
	}

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp)
	icmp.CheckSum = CheckSum(buffer.Bytes())
	buffer.Reset()

	return icmp
}

func sendICMPRequest(icmp ICMP, destAddr *net.IPAddr) (error, float64) {
	conn, err := net.DialIP("ip4:icmp", nil, destAddr)
	if err != nil {
		return err, 0
	}
	defer conn.Close()

	var buffer bytes.Buffer
	err = binary.Write(&buffer, binary.BigEndian, icmp)

	if err != nil {
		return err, 0
	}

	if _, err := conn.Write(buffer.Bytes()); err != nil {
		return err, 0
	}

	tStart := time.Now()

	if err = conn.SetReadDeadline((time.Now().Add(time.Second * 2))); err != nil {
		return err, 0
	}

	recv := make([]byte, 1024)
	_, err = conn.Read(recv)

	if err != nil {
		return err, 0
	}

	tEnd := time.Now()
	duration := float64(tEnd.Sub(tStart).Nanoseconds()) / 1e6

	return err, duration
}

func Ping(host string) string {
	raddr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		return "Fail to resolve host." + err.Error()
	}
	err, duration := sendICMPRequest(getICMP(), raddr)
	if err != nil {
		return "Error: " + err.Error()
	}

	return fmt.Sprintf("%.2fms (%s)", duration, host)
}

func CheckSum(data []byte) uint16 {
	var (
		sum    uint32
		length = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)

	return uint16(^sum)
}
