package serializer_test

import (
	"github.com/huang-baoding/protobuf-to-binary/sample"
	"github.com/huang-baoding/protobuf-to-binary/serializer"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jasonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	err = serializer.WriteProtobufToJSONFile(laptop1, jasonFile)
	require.NoError(t, err)

	laptop2 := sample.NewLaptop()
	err = serializer.ReadProtobufFromBinaryFile(laptop2, binaryFile)
	require.NoError(t, err)

}
