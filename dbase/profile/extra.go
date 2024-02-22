package profile

import (
	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/dbase/tables"
)

func Get(wxid, roomid string) *tables.Profile {

	p, err := Fetch(&FetchParam{
		Wxid:   wxid,
		Roomid: roomid,
	})

	if err != nil {
		p = &tables.Profile{Wxid: wxid, Roomid: roomid}
	}

	return p

}

func GetAiModel(wxid, roomid string) *args.LLModel {

	p := Get(wxid, roomid)

	var llmc *args.LLModel

	if p != nil {
		llmc = args.LLM.Models[p.AiModel]
	}

	if llmc == nil {
		llmc = args.LLM.Models[args.LLM.Default]
	}

	if llmc == nil {
		for _, v := range args.LLM.Models {
			return v
		}
	}

	return llmc

}

func SetAiModel(wxid, roomid, argot, model string) error {

	p, err := Fetch(&FetchParam{
		Wxid:   wxid,
		Roomid: roomid,
	})

	if err == nil && p.Rd > 0 {
		err = Update(&UpdateParam{
			Rd:      p.Rd,
			AiArgot: argot,
			AiModel: model,
		})
	} else {
		_, err = Create(&CreateParam{
			Wxid:    wxid,
			Roomid:  roomid,
			Level:   0,
			AiArgot: argot,
			AiModel: model,
		})
	}

	return err

}
