package help

import (
	"github.com/golang/protobuf/proto"
	cb "github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"
	putils "github.com/hyperledger/fabric/protos/utils"
	"github.com/pkg/errors"
	"time"
)

// TransactionDetail获取了交易的具体信息
type TransactionDetail struct {
	TransactionId string
	CreateTime    string
	Args          []string
}

// 从SDK中Block.BlockDara.Data中提取交易具体信息
func GetTransactionInfoFromData(data []byte, needArgs bool) (*TransactionDetail, error) {
	env, err := putils.GetEnvelopeFromBlock(data)
	if err != nil {
		return nil, errors.Wrap(err, "error extracting Envelope from block")
	}
	if env == nil {
		return nil, errors.New("nil envelope")
	}
	payload, err := putils.GetPayload(env)
	if err != nil {
		return nil, errors.Wrap(err, "error extracting Payload from envelope")
	}
	channelHeaderBytes := payload.Header.ChannelHeader
	channelHeader := &cb.ChannelHeader{}
	if err := proto.Unmarshal(channelHeaderBytes, channelHeader); err != nil {
		return nil, errors.Wrap(err, "error extracting ChannelHeader from payload")
	}
	var (
		args []string
	)
	if needArgs {
		tx, err := putils.GetTransaction(payload.Data)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling transaction payload")
		}
		chaincodeActionPayload, err := putils.GetChaincodeActionPayload(tx.Actions[0].Payload)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling chaincode action payload")
		}
		propPayload := &pb.ChaincodeProposalPayload{}
		if err := proto.Unmarshal(chaincodeActionPayload.ChaincodeProposalPayload, propPayload); err != nil {
			return nil, errors.Wrap(err, "error extracting ChannelHeader from payload")
		}
		invokeSpec := &pb.ChaincodeInvocationSpec{}
		err = proto.Unmarshal(propPayload.Input, invokeSpec)
		if err != nil {
			return nil, errors.Wrap(err, "error extracting ChannelHeader from payload")
		}
		for _, v := range invokeSpec.ChaincodeSpec.Input.Args {
			args = append(args, string(v))
		}
	}
	result := &TransactionDetail{
		TransactionId: channelHeader.TxId,
		Args:          args,
		CreateTime:    FormatTime(time.Unix(channelHeader.Timestamp.Seconds, 0)),
	}
	return result, nil
}
