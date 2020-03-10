package icmpping

import (
	"bytes"
	"encoding/binary"
	"iochen.com/v2gen/app/ping"
	"iochen.com/v2gen/infra/vmess"
	"log"
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

func getICMP() (*ICMP, error) {
	icmp := &ICMP{Type: 8}

	var buffer bytes.Buffer

	if err := binary.Write(&buffer, binary.BigEndian, *icmp); err != nil {
		return nil, err
	}

	icmp.CheckSum = checkSum(buffer.Bytes())
	buffer.Reset()

	return icmp, nil
}

func sendICMPRequest(icmp *ICMP, ip *net.IPAddr, timeout time.Duration) (time.Duration, error) {
	conn, err := net.DialIP("ip4:ICMP", nil, ip)
	if err != nil {
		return -1, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var buffer bytes.Buffer

	if err = binary.Write(&buffer, binary.BigEndian, icmp); err != nil {
		return -1, err
	}

	if _, err := conn.Write(buffer.Bytes()); err != nil {
		return -1, err
	}

	start := time.Now()

	if err = conn.SetReadDeadline(time.Now().Add(time.Second * timeout)); err != nil {
		return -1, err
	}

	recv := make([]byte, 1024)
	_, err = conn.Read(recv)
	if err != nil {
		return -1, err
	}

	end := time.Now()

	return end.Sub(start), nil
}

func Ping(lk *vmess.Link, count int, totalTimeout, eachTimeout time.Duration) (*ping.PingStat, error) {
	ps := &ping.PingStat{}
	icmp, err := getICMP()
	if err != nil {
		return nil, err
	}
	ip, err := net.ResolveIPAddr("ip", lk.Add)
	if err != nil {
		return nil, err
	}

	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(totalTimeout * time.Second)
		timeout <- true
	}()

L:
	for round := 0; round < count; round++ {
		chDelay := make(chan time.Duration)
		go func() {
			delay, err := sendICMPRequest(icmp, ip, eachTimeout)
			if err != nil {
				ps.ErrCounter++
			}
			chDelay <- delay
		}()
		select {
		case delay := <-chDelay:
			if delay > 0 {
				ps.Durations = append(ps.Durations, ping.Duration(delay))
			}
		case <-timeout:
			break L
		}
	}

	return ps, nil
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
