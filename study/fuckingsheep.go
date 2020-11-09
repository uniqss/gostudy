package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var testAddr = "111.222.123.234:56789"
var working = true
var count int32
var sentcount int32

var validMsgIds []uint32

var rb = NewRoundBuffer()

type RoundBuffer struct {
	info []byte
	mtx  sync.Mutex
}

func NewRoundBuffer() *RoundBuffer {
	ret := &RoundBuffer{}
	ret.info = append(ret.info, byte(0))
	return ret
}

func (rb *RoundBuffer) PP() []byte {
	rb.mtx.Lock()
	defer rb.mtx.Unlock()

	buf := make([]byte, 40*1024)

	msgid := validMsgIds[rand.Int()%len(validMsgIds)]

	msgNo := 0
	if msgid != 2460596793 {
		msgNo = rand.Int()
	}

	dataLen := rand.Int()

	errCode := 0

	HeadLen := 14

	buf[0] = byte(msgid)
	buf[1] = byte(msgid >> 8)
	buf[2] = byte(msgid >> 16)
	buf[3] = byte(msgid >> 24)

	buf[4] = byte(dataLen)
	buf[5] = byte(dataLen >> 8)
	buf[6] = byte(dataLen >> 16)
	buf[7] = byte(dataLen >> 24)

	buf[8] = byte(errCode)
	buf[9] = byte(errCode >> 8)

	buf[10] = byte(msgNo)
	buf[11] = byte(msgNo >> 8)
	buf[12] = byte(msgNo >> 16)
	buf[13] = byte(msgNo >> 24)

	rbLen := len(rb.info)
	if len(rb.info) == 0 {
		panic("use NewRoundBuffer instead")
	}
	pos := -1
	for i := 0; i < rbLen; i++ {
		if rb.info[i] != byte(255) {
			pos = i
			break
		}
	}
	if pos == -1 {
		for i := 0; i < rbLen; i++ {
			rb.info[i] = byte(0)
		}
		rb.info = append(rb.info, byte(0))
	} else if pos == 0 {
		rb.info[pos]++
	} else {
		rb.info[pos]++
		for i := 0; i < pos; i++ {
			rb.info[i] = byte(0)
		}
	}

	copy(buf[HeadLen:], rb.info)

	return buf
}

func fuckingsheep() {
	rand.Seed(time.Now().UnixNano())
	var conn net.Conn = nil
	var err error
	atomic.AddInt32(&count, 1)

	conn, err = net.Dial("tcp", testAddr)
	if err != nil {
		fmt.Println("connect failed. err:", err)
		atomic.AddInt32(&count, -1)
		//conn.Close()
		return
	}
	defer func() {
		conn.Close()
		atomic.AddInt32(&count, -1)
	}()

	for working {
		if 1 == 0 {
			length := rand.Int() % 65536
			var buffer []byte
			for i := 0; i < length; i++ {
				buffer = append(buffer, byte(rand.Int()%256))
			}
			conn.Write(buffer)
		} else {
			conn.Write(rb.PP())
		}
		atomic.AddInt32(&sentcount, 1)
		//time.Sleep(time.Millisecond * 100)
		//conn.Close()

		time.Sleep(time.Millisecond * 1)
	}
}

func printcount() {
	for working {
		c := atomic.LoadInt32(&count)
		fmt.Println("count:", c, " now:", time.Now(), " sentcount:", sentcount)
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	initmsgs()

	for i := 0; i < 100; i++ {
		go fuckingsheep()
	}

	go printcount()

	var input string
	for {
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input == "e" || input == "exit" {
			break
		}
	}
	working = false
}

func initmsgs() {
	validMsgIds = append(validMsgIds, 2460596793)
	validMsgIds = append(validMsgIds, 2237866322)
	validMsgIds = append(validMsgIds, 3618208248)
	validMsgIds = append(validMsgIds, 2476496455)
	validMsgIds = append(validMsgIds, 2220281132)
	validMsgIds = append(validMsgIds, 1722598537)
	validMsgIds = append(validMsgIds, 3398631595)
	validMsgIds = append(validMsgIds, 1006001858)
	validMsgIds = append(validMsgIds, 1839217367)
	validMsgIds = append(validMsgIds, 3693113986)
	validMsgIds = append(validMsgIds, 1229969043)
	validMsgIds = append(validMsgIds, 3306165154)
	validMsgIds = append(validMsgIds, 933048411)
	validMsgIds = append(validMsgIds, 288295790)
	validMsgIds = append(validMsgIds, 1758280586)
	validMsgIds = append(validMsgIds, 1505263473)
	validMsgIds = append(validMsgIds, 2131206369)
	validMsgIds = append(validMsgIds, 327785508)
	validMsgIds = append(validMsgIds, 3110179518)
	validMsgIds = append(validMsgIds, 1026835030)
	validMsgIds = append(validMsgIds, 721344829)
	validMsgIds = append(validMsgIds, 2288155311)
	validMsgIds = append(validMsgIds, 593346307)
	validMsgIds = append(validMsgIds, 882320488)
	validMsgIds = append(validMsgIds, 294981960)
	validMsgIds = append(validMsgIds, 3049382912)
	validMsgIds = append(validMsgIds, 549570321)
	validMsgIds = append(validMsgIds, 855750706)
	validMsgIds = append(validMsgIds, 931322545)
	validMsgIds = append(validMsgIds, 608347538)
	validMsgIds = append(validMsgIds, 1581077056)
	validMsgIds = append(validMsgIds, 1308437859)
	validMsgIds = append(validMsgIds, 1490042769)
	validMsgIds = append(validMsgIds, 1259356338)
	validMsgIds = append(validMsgIds, 567957880)
	validMsgIds = append(validMsgIds, 840574555)
	validMsgIds = append(validMsgIds, 1439048276)
	validMsgIds = append(validMsgIds, 2553242803)
	validMsgIds = append(validMsgIds, 2414215128)
	validMsgIds = append(validMsgIds, 3204791785)
	validMsgIds = append(validMsgIds, 2832180866)
	validMsgIds = append(validMsgIds, 3843863954)
	validMsgIds = append(validMsgIds, 4074123001)
	validMsgIds = append(validMsgIds, 2601227002)
	validMsgIds = append(validMsgIds, 2295066073)
	validMsgIds = append(validMsgIds, 1545551230)
	validMsgIds = append(validMsgIds, 1340031581)
	validMsgIds = append(validMsgIds, 3577798685)
	validMsgIds = append(validMsgIds, 3330318142)
	validMsgIds = append(validMsgIds, 3123182279)
	validMsgIds = append(validMsgIds, 1880428781)
	validMsgIds = append(validMsgIds, 855789369)
	validMsgIds = append(validMsgIds, 2955968476)
	validMsgIds = append(validMsgIds, 3858590696)
	validMsgIds = append(validMsgIds, 2936280460)
	validMsgIds = append(validMsgIds, 945446300)
	validMsgIds = append(validMsgIds, 2393522589)
	validMsgIds = append(validMsgIds, 1966250362)
	validMsgIds = append(validMsgIds, 766500948)
	validMsgIds = append(validMsgIds, 3540795184)
	validMsgIds = append(validMsgIds, 1936781183)
	validMsgIds = append(validMsgIds, 896183179)
	validMsgIds = append(validMsgIds, 1466224056)
	validMsgIds = append(validMsgIds, 2955248888)
	validMsgIds = append(validMsgIds, 2460858751)
	validMsgIds = append(validMsgIds, 2220099579)
	validMsgIds = append(validMsgIds, 4157421398)
	validMsgIds = append(validMsgIds, 2932110564)
	validMsgIds = append(validMsgIds, 1346025614)
	validMsgIds = append(validMsgIds, 889546577)
	validMsgIds = append(validMsgIds, 1546089627)
	validMsgIds = append(validMsgIds, 3083062478)
	validMsgIds = append(validMsgIds, 2684966821)
	validMsgIds = append(validMsgIds, 1414487208)
	validMsgIds = append(validMsgIds, 1132840899)
	validMsgIds = append(validMsgIds, 3407586265)
	validMsgIds = append(validMsgIds, 2330783825)
	validMsgIds = append(validMsgIds, 1212327502)
	validMsgIds = append(validMsgIds, 1602764069)
	validMsgIds = append(validMsgIds, 1705211124)
	validMsgIds = append(validMsgIds, 1919545247)
	validMsgIds = append(validMsgIds, 2416176450)
	validMsgIds = append(validMsgIds, 2278127145)
	validMsgIds = append(validMsgIds, 3524825974)
	validMsgIds = append(validMsgIds, 3318876189)
	validMsgIds = append(validMsgIds, 507846297)
	validMsgIds = append(validMsgIds, 3031240707)
	validMsgIds = append(validMsgIds, 1983279634)
	validMsgIds = append(validMsgIds, 1643964793)
	validMsgIds = append(validMsgIds, 2671241060)
	validMsgIds = append(validMsgIds, 2298313743)
	validMsgIds = append(validMsgIds, 3476983357)
	validMsgIds = append(validMsgIds, 2805514155)
}
