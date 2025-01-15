package aireq

import (
	"context"
	conf "derma/detect/config"
	"encoding/json"
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type DocterData struct {
	Answer string `json:"answer"`
}

type PredictionData struct {
	Predictions []struct {
		Class       string  `json:"class"`
		Probability float64 `json:"probability"`
	} `json:"predictions"`
}

func DoctorReq(question string) (string, error) {
	var postArgs protocol.Args
	addr := fmt.Sprintf("http://%s/ask", conf.AISERVER.Addr)
	postArgs.Set("question", question) // Set post args

	_, body, err := AIClinet.Post(context.Background(), nil, addr, &postArgs)
	if err != nil {
		return "", err
	}

	var responseData DocterData

	// 解析JSON响应
	err = sonic.Unmarshal(body, &responseData)
	if err != nil {
		return "", err
	}

	// 提取answer字段的值
	answer := responseData.Answer

	return answer, nil
}

func PictureReq(dir string) (*PredictionData, error) {
	req := &protocol.Request{}
	res := &protocol.Response{}

	req.SetMethod(consts.MethodPost)
	req.SetRequestURI(fmt.Sprintf("http://%s/predict", conf.AISERVER.Addr))
	req.SetFile("file", dir)

	err := AIClinet.Do(context.Background(), req, res)
	if err != nil {
		return nil, err
	}

	// 解析JSON响应
	var responseData PredictionData
	err = json.Unmarshal(res.Body(), &responseData)
	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
