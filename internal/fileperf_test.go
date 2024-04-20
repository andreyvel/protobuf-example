package internal

import (
	"encoding/binary"
	"errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	impl_v2 "protobuf-dynamic-go/impl.v2"
	"strconv"
	"testing"
	"time"
)

var dataRootPath = "/tmp/proto"

func TestProtoBufToFile(t *testing.T) {
	fileNum := 1
	// create directory if not exists
	if _, err := os.Stat(dataRootPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dataRootPath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	start := time.Now()
	for num := 0; num < fileNum; num++ {
		arr := generateProtos(1000)

		fileName := filepath.Join(dataRootPath, "data"+strconv.Itoa(num)+".dat")
		log.Printf("saveToFile: %s\n", fileName)
		saveToFile(arr, fileName)
	}
	diff := time.Now().Sub(start)
	log.Printf("saveToFile: fileNum=%v, %v seconds\n", fileNum, diff.Seconds())
	//saveToFile:   fileNum=70, 34.807012899 seconds

	start = time.Now()
	for num := 0; num < fileNum; num++ {
		fileName := filepath.Join(dataRootPath, "data"+strconv.Itoa(num)+".dat")
		arr := loadFromFile(fileName)
		log.Printf("loadFromFile: %s, %v rows\n", fileName, len(arr))
	}

	diff = time.Now().Sub(start)
	log.Printf("loadFromFile: fileNum=%v, %v seconds\n", fileNum, diff.Seconds())
	//loadFromFile: fileNum=70, 1.6563136269999998 seconds
}

func generateProtos(msgNum int) []*impl_v2.Message {
	var arr []*impl_v2.Message

	for num := 0; num < msgNum; num++ {
		name := "name" + strconv.Itoa(num)

		var msg = &impl_v2.Message{
			Id:          int32(num),
			Name:        name,
			Email:       name + "@example.com",
			NewField1:   "newField1*" + name,
			NewField2:   generateRandomString(20_000),
			LastUpdated: timestamppb.Now(),
		}
		arr = append(arr, msg)
	}

	return arr
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, length)
	for ind := range result {
		result[ind] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

func saveToFile(arr []*impl_v2.Message, fileName string) {
	// delete file if exists
	_, err := os.Stat(fileName)
	if err == nil {
		err = os.Remove(fileName)
		if err != nil {
			log.Fatal(err)
		}
	}

	fl, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()

	for _, msg := range arr {
		dataBin, err := proto.Marshal(msg)
		if err != nil {
			log.Fatal(err)
		}

		msgLenBin := make([]byte, 4)
		binary.LittleEndian.PutUint32(msgLenBin, uint32(len(dataBin)))

		// write message length
		_, err = fl.Write(msgLenBin)
		if err != nil {
			log.Fatal(err)
		}
		// write message
		_, err = fl.Write(dataBin)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func loadFromFile(fileName string) []*impl_v2.Message {
	fl, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()

	var arr []*impl_v2.Message
	for {
		msgLenBin := make([]byte, 4)
		len, err := fl.Read(msgLenBin)
		if err != nil || len != 4 {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		msgLen := binary.LittleEndian.Uint32(msgLenBin)

		dataBin := make([]byte, msgLen)
		len, err = fl.Read(dataBin)
		if err != nil || len != int(msgLen) {
			log.Fatal(err)
		}

		msg := &impl_v2.Message{}
		err = proto.Unmarshal(dataBin, msg)
		if err != nil {
			log.Fatalln(err)
		}

		arr = append(arr, msg)
	}
	return arr
}
