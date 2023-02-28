//将proto.Message写入二进制文件、读出文件、写入json文件

package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)
//proto.Message是官方包，指在proto里定义的结构体都可以作为参数
func WriteProtobufToBinaryFile(message proto.Message, fileName string) error {
		data, err := proto.Marshal(message)	//message应该是指针类型
	if err != nil {
		return fmt.Errorf("can not marshal proto message to binary:%w", err)
	}
	err = ioutil.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("can not write to file:%w", err)
	}
	return nil
}

func ReadProtobufFromBinaryFile(message proto.Message, fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("can not read the binary file:%w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("can not unmarshal binary to proto message:%w", err)
	}
	return nil
}

func WriteProtobufToJSONFile(message proto.Message, fileName string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to JSON: %w", err)
	}

	err = ioutil.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON data to file: %w", err)
	}

	return nil
}
