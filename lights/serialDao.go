package lights

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/spf13/viper"
	"github.com/tarm/serial"
	"log"
	"strconv"
	"sync"
	"time"
)

var MUTEX = &sync.Mutex{}

type SerialDao struct {
	Port *serial.Port

	Mutex *sync.Mutex
}

func NewSerialDao() (SerialDao, error) {
	c := &serial.Config{Name: viper.GetString("serial_port"), Baud: 9600}
	MUTEX.Lock()
	port, err := serial.OpenPort(c)
	if err != nil {
		MUTEX.Unlock()
		return SerialDao{}, errors.New("Could not fetch Sserial Port")
	}
	return SerialDao{Port: port, Mutex: MUTEX}, nil
}

func (serial SerialDao) GetChannel(vars map[string]string) (int, error) {
	number, err := strconv.Atoi(vars["channel"])
	if err != nil {
		return 0, errors.New("Device does not exist")
	}
	if number > 0 && number <= viper.GetInt("channels") {
		return number, nil
	}
	return 0, errors.New("Device does not exist")

}
func (serial SerialDao) GetAcion(vars map[string]string) (string, error) {
	switch {
	case vars["action"] == "true":
		{
			return strconv.Itoa(255), nil

		}
	case vars["action"] == "false":
		{
			return strconv.Itoa(0), nil
		}
	default:
		{
			return "", errors.New("Buahhhhh!")
		}
	}
}

func (serial SerialDao) GetReadCommand(vars map[string]string) (string, error) {
	channel, err := serial.GetChannel(vars)
	if err != nil {
		return "", err
	}
	shield := (channel / 8)
	if channel%8 != 0 {
		shield++
	}
	return "@106," + strconv.Itoa(shield) + ":", nil
}

func (serial SerialDao) GetWriteCommand(vars map[string]string) (string, error) {
	channel, err := serial.GetChannel(vars)
	if err != nil {
		return "", err
	}
	action, err := serial.GetAcion(vars)
	if err != nil {
		return "", err
	}

	cmd := "@104," + strconv.Itoa(channel) + "," + action + ":"

	return cmd, nil
}

func (serial SerialDao) ReadFromSerial(vars map[string]string) (bool, error) {

	reader := bufio.NewReader(serial.Port)
	reply, err := reader.ReadBytes('\x0a')
	if err != nil {
		panic(err)
	}
	raw := string(bytes.Trim(reply, " \r\n\t"))
	status, _ := strconv.Atoi(raw)

	var channel_status []bool
	for status > 0 {
		channel_status = append(channel_status, status%10 == 1)
		status /= 10
	}
	for len(channel_status) < 8 {
		channel_status = append(channel_status, false)
	}

	index, _ := strconv.Atoi(vars["channels"])

	return channel_status[index%8], nil

}
func (serial SerialDao) Free() {
	serial.Port.Close()
	serial.Mutex.Unlock()
	log.Println("unlocked")
}

func (serial SerialDao) CurrentState() (map[string]bool, error) {
	var channel []bool
	status := make(map[string]bool)
	numberOfShields := viper.GetInt("channels") / 8

	for i := 1; i <= numberOfShields; i++ {
		cmd := "@106," + strconv.Itoa(i) + ":"

		_, err := serial.Port.Write([]byte(cmd))

		if err != nil {
			return status, err

		}

		reader := bufio.NewReader(serial.Port)
		reply, err := reader.ReadBytes('\x0a')
		if err != nil {
			return status, err
		}
		raw := string(bytes.Trim(reply, " \r\n\t"))
		result, _ := strconv.Atoi(raw)

		for result > 0 {
			channel = append(channel, result%10 == 1)
			result /= 10
		}
		for len(channel) < 8 {
			channel = append(channel, false)
		}

		time.Sleep(1 / 4 * time.Second)
	}

	var buffer bytes.Buffer
	for index, element := range channel {
		buffer.WriteString(strconv.FormatBool(element))
		if (index + 1) != len(channel) {
			buffer.WriteString(" ")
		}
	}

	for index, ele := range channel {
		status[strconv.Itoa(index+1)] = ele
	}

	return status, nil
}
