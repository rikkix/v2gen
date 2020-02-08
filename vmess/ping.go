package vmess

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
)

type icmp struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func getICMP() (icmp, error) {
	icmpInstance := icmp{Type: 8}

	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, icmpInstance)
	if err != nil {
		return icmp{}, err
	}

	icmpInstance.CheckSum = checkSum(buffer.Bytes())
	buffer.Reset()

	return icmpInstance, nil
}

func sendICMPRequest(icmp icmp, destAddr *net.IPAddr) (error, float64) {
	conn, err := net.DialIP("ip4:icmp", nil, destAddr)
	if err != nil {
		return err, 0
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var buffer bytes.Buffer

	if err = binary.Write(&buffer, binary.BigEndian, icmp); err != nil {
		return err, 0
	}

	if _, err := conn.Write(buffer.Bytes()); err != nil {
		return err, 0
	}

	start := time.Now()

	if err = conn.SetReadDeadline(time.Now().Add(time.Second * 2)); err != nil {
		return err, 0
	}

	recv := make([]byte, 1024)
	_, err = conn.Read(recv)

	if err != nil {
		return err, 0
	}

	end := time.Now()
	duration := float64(end.Sub(start).Nanoseconds()) / 1e6

	return err, duration
}

func (node *Link) Ping() string {
	rAddr, err := net.ResolveIPAddr("ip", node.Host)
	if err != nil {
		return "Fail to resolve host." + err.Error()
	}

	icmp, err := getICMP()
	if err != nil {
		log.Println(err)
	}

	err, duration := sendICMPRequest(icmp, rAddr)
	if err != nil {
		return "Error: " + err.Error()
	}

	return fmt.Sprintf("%.2fms (%s)", duration, node.Host)
}

func checkSum(data []byte) uint16 {
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
	sum += sum >> 16

	return uint16(^sum)
}
