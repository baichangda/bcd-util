package immotors_json

import (
	"bcd-util/support_parse/immotors"
	"encoding/json"
	"github.com/pkg/errors"
	"sort"
)

type Json struct {
	SAIC_FileVersion string    `json:"SAIC_FileVersion"`
	FileCreationTime int64     `json:"FileCreationTime"`
	Tboxinfo         Tboxinfo  `json:"tboxinfo"`
	Journey          Journey   `json:"journey"`
	Channels         []Channel `json:"channels"`
}

func (e *Json) ToBytes(ts int64, count int) ([]byte, error) {
	newJson := *e
	tss := ts / 1000
	startTss := tss - int64(count) + 1
	newJson.FileCreationTime = startTss
	channels := newJson.Channels
	newChannels := make([]Channel, len(channels))
	for j, channel := range channels {
		data0 := channel.Data[0]
		newData := make([]map[string]any, count)
		if channel.ID == 1 {
			for i := range count {
				temp := make(map[string]any)
				for k, v := range data0 {
					if k == "TBOXSysTim" {
						temp[k] = startTss + int64(i)
					} else {
						temp[k] = v
					}
				}
				newData[i] = temp
			}
		} else {
			for i := range count {
				newData[i] = data0
			}
		}
		channel.Starttime = startTss
		channel.Data = newData
		newChannels[j] = channel
	}

	newJson.Channels = newChannels

	marshal, err := json.Marshal(newJson)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return marshal, nil
}

type Tboxinfo struct {
	ID         string `json:"ID"`
	VIN        string `json:"VIN"`
	PartNumber string `json:"PartNumber"`
}

type Journey struct {
	JourneyID int `json:"JourneyID"`
}

type Channel struct {
	ID                  int              `json:"ID"`
	Starttime           int64            `json:"starttime"`
	CollectiofrequecyHz int              `json:"collectiofrequecyHz"`
	Data                []map[string]any `json:"data"`
}

var GroupId_groupName = make(map[int]string)
var GroupName_groupId = make(map[string]int)

func init() {
	GroupId_groupName[1] = "IAM"
	GroupId_groupName[2] = "ICC"
	GroupId_groupName[3] = "IMCU"
	GroupId_groupName[4] = "IBS"
	GroupId_groupName[5] = "ESS"
	GroupId_groupName[6] = "CCU"
	GroupId_groupName[7] = "ISC"
	GroupId_groupName[8] = "SAC"
	GroupId_groupName[9] = "ICM"
	GroupId_groupName[10] = "EHBS"
	GroupId_groupName[11] = "ECM"
	GroupId_groupName[12] = "PLCM"
	GroupId_groupName[13] = "TCM"
	GroupId_groupName[14] = "BCM"
	GroupId_groupName[15] = "CCP"
	GroupId_groupName[16] = "VCU"
	GroupId_groupName[17] = "TC"
	GroupId_groupName[18] = "AMP"
	GroupId_groupName[19] = "AMR"
	GroupId_groupName[20] = "ATC"
	GroupId_groupName[21] = "BPEPS"
	GroupId_groupName[22] = "CARLog"
	GroupId_groupName[23] = "DLP"
	GroupId_groupName[24] = "EOPC"
	GroupId_groupName[25] = "EPMCU"
	GroupId_groupName[26] = "EPS"
	GroupId_groupName[27] = "EPS_SFCANFD"
	GroupId_groupName[28] = "FLDCM"
	GroupId_groupName[29] = "FLIDAR"
	GroupId_groupName[30] = "FLSM"
	GroupId_groupName[31] = "FRDCM"
	GroupId_groupName[32] = "FRSM"
	GroupId_groupName[33] = "FVCM"
	GroupId_groupName[34] = "HUD"
	GroupId_groupName[35] = "IECU"
	GroupId_groupName[36] = "IMATE"
	GroupId_groupName[37] = "IMU"
	GroupId_groupName[38] = "IPD"
	GroupId_groupName[39] = "IPS"
	GroupId_groupName[40] = "LFSDA"
	GroupId_groupName[41] = "LHCMS"
	GroupId_groupName[42] = "LHRDA"
	GroupId_groupName[43] = "LVBM"
	GroupId_groupName[44] = "PGM"
	GroupId_groupName[45] = "PMA"
	GroupId_groupName[46] = "RBM"
	GroupId_groupName[47] = "RDR"
	GroupId_groupName[48] = "RFSDA"
	GroupId_groupName[49] = "RHCMS"
	GroupId_groupName[50] = "RHRDA"
	GroupId_groupName[51] = "RLDCM"
	GroupId_groupName[52] = "RLSM"
	GroupId_groupName[53] = "RRDCM"
	GroupId_groupName[54] = "RRSM"
	GroupId_groupName[55] = "RWSGW"
	GroupId_groupName[56] = "WLC"
	GroupId_groupName[57] = "FDR"
	GroupId_groupName[58] = "SAS"
	GroupId_groupName[59] = "SCM"
	GroupId_groupName[60] = "SCU"
	GroupId_groupName[61] = "SDM"
	GroupId_groupName[62] = "SPD"

	for k, v := range GroupId_groupName {
		GroupName_groupId[v] = k
	}
}

func BinToJson(p *immotors.Packet) Json {
	ts := p.F_evt_0001.F_TBOXSysTim * 1000
	vin := p.F_evt_D00A.F_VIN
	var channels []Channel
	//group RHCMS
	data_RHCMS := make(map[string]any)
	if p.F_evt_D016 != nil {
		data_RHCMS["DTCInfomationRHCMS"] = p.F_evt_D016.F_DTCInfomationRHCMS
	}
	if len(data_RHCMS) > 0 {
		channels = append(channels, Channel{
			ID:                  49,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_RHCMS},
		})
	}

	//group WLC
	data_WLC := make(map[string]any)
	if p.F_evt_D011 != nil {
		data_WLC["DTCInfomationWLC"] = p.F_evt_D011.F_DTCInfomationWLC
	}
	if len(data_WLC) > 0 {
		channels = append(channels, Channel{
			ID:                  56,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_WLC},
		})
	}

	//group SDM
	data_SDM := make(map[string]any)
	if p.F_evt_D012 != nil {
		data_SDM["DTCInfomationSDM"] = p.F_evt_D012.F_DTCInfomationSDM
	}
	if len(data_SDM) > 0 {
		channels = append(channels, Channel{
			ID:                  61,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_SDM},
		})
	}

	//group BPEPS
	data_BPEPS := make(map[string]any)
	if p.F_evt_D013 != nil {
		data_BPEPS["DTCInfomationBPEPS"] = p.F_evt_D013.F_DTCInfomationBPEPS
	}
	if len(data_BPEPS) > 0 {
		channels = append(channels, Channel{
			ID:                  21,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_BPEPS},
		})
	}

	//group IBS
	data_IBS := make(map[string]any)
	if p.F_evt_0801 != nil {
		data_IBS["BrkPdlPos"] = p.F_evt_0801.F_BrkPdlPos
	}
	if p.F_evt_0802 != nil {
		data_IBS["VehSpdAvgDrvn"] = p.F_evt_0802.F_VehSpdAvgDrvn
	}
	if p.F_evt_0802 != nil {
		data_IBS["VehSpdAvgDrvnV"] = p.F_evt_0802.F_VehSpdAvgDrvnV
	}
	if p.F_evt_0803 != nil {
		data_IBS["BrkPdlPosV"] = p.F_evt_0803.F_BrkPdlPosV
	}
	if p.F_evt_D006 != nil {
		data_IBS["BrkSysBrkLghtsReqd"] = p.F_evt_D006.F_BrkSysBrkLghtsReqd
	}
	if p.F_evt_D006 != nil {
		data_IBS["EPBSysBrkLghtsReqd"] = p.F_evt_D006.F_EPBSysBrkLghtsReqd
	}
	if p.F_evt_D006 != nil {
		data_IBS["EPBSysBrkLghtsReqdA"] = p.F_evt_D006.F_EPBSysBrkLghtsReqdA
	}
	if p.F_evt_D006 != nil {
		data_IBS["BrkFludLvlLow"] = p.F_evt_D006.F_BrkFludLvlLow
	}
	if p.F_evt_D006 != nil {
		data_IBS["BrkSysRedBrkTlltReq"] = p.F_evt_D006.F_BrkSysRedBrkTlltReq
	}
	if p.F_evt_D006 != nil {
		data_IBS["ABSF"] = p.F_evt_D006.F_ABSF
	}
	if p.F_evt_D006 != nil {
		data_IBS["VSESts"] = p.F_evt_D006.F_VSESts
	}
	if p.F_evt_D006 != nil {
		data_IBS["BrkFludLvlLowV"] = p.F_evt_D006.F_BrkFludLvlLowV
	}
	if p.F_evt_D008 != nil {
		data_IBS["DTCInfomationEPB"] = p.F_evt_D008.F_DTCInfomationEPB
	}
	if p.F_evt_D012 != nil {
		data_IBS["DTCInfomationIBS"] = p.F_evt_D012.F_DTCInfomationIBS
	}
	if len(data_IBS) > 0 {
		channels = append(channels, Channel{
			ID:                  4,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_IBS},
		})
	}

	//group FLDCM
	data_FLDCM := make(map[string]any)
	if p.F_evt_D013 != nil {
		data_FLDCM["DTCInfomationDCM_FL"] = p.F_evt_D013.F_DTCInfomationDCM_FL
	}
	if len(data_FLDCM) > 0 {
		channels = append(channels, Channel{
			ID:                  28,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_FLDCM},
		})
	}

	//group AMR
	data_AMR := make(map[string]any)
	if p.F_evt_D013 != nil {
		data_AMR["DTCInfomationAMR"] = p.F_evt_D013.F_DTCInfomationAMR
	}
	if len(data_AMR) > 0 {
		channels = append(channels, Channel{
			ID:                  19,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_AMR},
		})
	}

	//group FRSM
	data_FRSM := make(map[string]any)
	if p.F_evt_D013 != nil {
		data_FRSM["DTCInfomationMSM_Psng"] = p.F_evt_D013.F_DTCInfomationMSM_Psng
	}
	if len(data_FRSM) > 0 {
		channels = append(channels, Channel{
			ID:                  32,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_FRSM},
		})
	}

	//group PMA
	data_PMA := make(map[string]any)
	if p.F_evt_D016 != nil {
		data_PMA["DTCInfomationPMA"] = p.F_evt_D016.F_DTCInfomationPMA
	}
	if len(data_PMA) > 0 {
		channels = append(channels, Channel{
			ID:                  45,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_PMA},
		})
	}

	//group IECU
	data_IECU := make(map[string]any)
	if p.F_evt_D010 != nil {
		data_IECU["DTCInfomationIECU"] = p.F_evt_D010.F_DTCInfomationIECU
	}
	if len(data_IECU) > 0 {
		channels = append(channels, Channel{
			ID:                  35,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_IECU},
		})
	}

	//group RBM
	data_RBM := make(map[string]any)
	if p.F_evt_D012 != nil {
		data_RBM["DTCInfomationRBM"] = p.F_evt_D012.F_DTCInfomationRBM
	}
	if len(data_RBM) > 0 {
		channels = append(channels, Channel{
			ID:                  46,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_RBM},
		})
	}

	//group FLSM
	data_FLSM := make(map[string]any)
	if p.F_evt_D013 != nil {
		data_FLSM["DTCInfomationMSM_Drv"] = p.F_evt_D013.F_DTCInfomationMSM_Drv
	}
	if len(data_FLSM) > 0 {
		channels = append(channels, Channel{
			ID:                  30,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_FLSM},
		})
	}

	//group SPD
	data_SPD := make(map[string]any)
	if p.F_evt_D017 != nil {
		data_SPD["DTCInfomationSPD"] = p.F_evt_D017.F_DTCInfomationSPD
	}
	if len(data_SPD) > 0 {
		channels = append(channels, Channel{
			ID:                  62,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_SPD},
		})
	}

	//group ICC
	data_ICC := make(map[string]any)
	if p.F_evt_0800 != nil {
		data_ICC["SysPwrMd"] = p.F_evt_0800.F_SysPwrMd
	}
	if p.F_evt_0800 != nil {
		data_ICC["SysPwrMdV"] = p.F_evt_0800.F_SysPwrMdV
	}
	if p.F_evt_0800 != nil {
		data_ICC["SysVolV"] = p.F_evt_0800.F_SysVolV
	}
	if p.F_evt_0800 != nil {
		data_ICC["SysVol"] = p.F_evt_0800.F_SysVol
	}
	if p.F_evt_0803 != nil {
		data_ICC["VehOdo"] = p.F_evt_0803.F_VehOdo
	}
	if p.F_evt_0803 != nil {
		data_ICC["VehOdoV"] = p.F_evt_0803.F_VehOdoV
	}
	if p.F_evt_D015 != nil {
		data_ICC["DTCInfomationICC"] = p.F_evt_D015.F_DTCInfomationICC
	}
	if p.F_evt_D019 != nil {
		data_ICC["BCMAvlbly"] = p.F_evt_D019.F_BCMAvlbly
	}
	if p.F_evt_D019 != nil {
		data_ICC["CCUAvlbly"] = p.F_evt_D019.F_CCUAvlbly
	}
	if p.F_evt_D019 != nil {
		data_ICC["BatCrnt"] = p.F_evt_D019.F_BatCrnt
	}
	if p.F_evt_D019 != nil {
		data_ICC["BatSOC"] = p.F_evt_D019.F_BatSOC
	}
	if p.F_evt_D019 != nil {
		data_ICC["BatSOCSts"] = p.F_evt_D019.F_BatSOCSts
	}
	if p.F_evt_D019 != nil {
		data_ICC["BatVol"] = p.F_evt_D019.F_BatVol
	}
	if p.F_evt_D019 != nil {
		data_ICC["EnrgSplReq"] = p.F_evt_D019.F_EnrgSplReq
	}
	if p.F_evt_D019 != nil {
		data_ICC["EnrgSplReqScene"] = p.F_evt_D019.F_EnrgSplReqScene
	}
	if p.F_evt_D019 != nil {
		data_ICC["VehEnrgRdyLvl"] = p.F_evt_D019.F_VehEnrgRdyLvl
	}
	if p.F_evt_D019 != nil {
		data_ICC["VehEnrgRdyLvlV"] = p.F_evt_D019.F_VehEnrgRdyLvlV
	}
	if len(data_ICC) > 0 {
		channels = append(channels, Channel{
			ID:                  2,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_ICC},
		})
	}

	//group SAS
	data_SAS := make(map[string]any)
	if p.F_evt_D012 != nil {
		data_SAS["DTCInfomationSAS"] = p.F_evt_D012.F_DTCInfomationSAS
	}
	if len(data_SAS) > 0 {
		channels = append(channels, Channel{
			ID:                  58,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_SAS},
		})
	}

	//group RWSGW
	data_RWSGW := make(map[string]any)
	if p.F_evt_D012 != nil {
		data_RWSGW["DTCInfomationRWSGW"] = p.F_evt_D012.F_DTCInfomationRWSGW
	}
	if p.F_evt_D012 != nil {
		data_RWSGW["DTCInfomationRWS"] = p.F_evt_D012.F_DTCInfomationRWS
	}
	if len(data_RWSGW) > 0 {
		channels = append(channels, Channel{
			ID:                  55,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_RWSGW},
		})
	}

	//group IMU
	data_IMU := make(map[string]any)
	if p.F_evt_D016 != nil {
		data_IMU["DTCInfomationIMU"] = p.F_evt_D016.F_DTCInfomationIMU
	}
	if len(data_IMU) > 0 {
		channels = append(channels, Channel{
			ID:                  37,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_IMU},
		})
	}

	//group HUD
	data_HUD := make(map[string]any)
	if p.F_evt_D017 != nil {
		data_HUD["DTCInfomationHUD"] = p.F_evt_D017.F_DTCInfomationHUD
	}
	if len(data_HUD) > 0 {
		channels = append(channels, Channel{
			ID:                  34,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_HUD},
		})
	}

	//group FDR
	data_FDR := make(map[string]any)
	if p.F_evt_D010 != nil {
		data_FDR["DTCInfomationFDR"] = p.F_evt_D010.F_DTCInfomationFDR
	}
	if len(data_FDR) > 0 {
		channels = append(channels, Channel{
			ID:                  57,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_FDR},
		})
	}

	//group SCM
	data_SCM := make(map[string]any)
	if p.F_evt_D012 != nil {
		data_SCM["DTCInfomationSCM"] = p.F_evt_D012.F_DTCInfomationSCM
	}
	if len(data_SCM) > 0 {
		channels = append(channels, Channel{
			ID:                  59,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_SCM},
		})
	}

	//group RLDCM
	data_RLDCM := make(map[string]any)
	if p.F_evt_D013 != nil {
		data_RLDCM["DTCInfomationDCM_RL"] = p.F_evt_D013.F_DTCInfomationDCM_RL
	}
	if len(data_RLDCM) > 0 {
		channels = append(channels, Channel{
			ID:                  51,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_RLDCM},
		})
	}

	//group AMP
	data_AMP := make(map[string]any)
	if p.F_evt_D014 != nil {
		data_AMP["DTCInfomationAMP"] = p.F_evt_D014.F_DTCInfomationAMP
	}
	if len(data_AMP) > 0 {
		channels = append(channels, Channel{
			ID:                  18,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_AMP},
		})
	}

	//group TC
	data_TC := make(map[string]any)
	if p.F_evt_000F != nil {
		data_TC["TMActuToqHiPre"] = p.F_evt_000F.F_TMActuToqHiPre
	}
	if p.F_evt_000F != nil {
		data_TC["TMInvtrCrntHiPre"] = p.F_evt_000F.F_TMInvtrCrntHiPre
	}

	if len(data_TC) > 0 {
		channels = append(channels, Channel{
			ID:                  17,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_TC},
		})
	}

	//group PLCM
	data_PLCM := make(map[string]any)
	if p.F_evt_D008 != nil {
		data_PLCM["DTCInfomationPLCM"] = p.F_evt_D008.F_DTCInfomationPLCM
	}
	if len(data_PLCM) > 0 {
		channels = append(channels, Channel{
			ID:                  12,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_PLCM},
		})
	}

	//group VCU
	data_VCU := make(map[string]any)
	if p.F_evt_D00F != nil {
		data_VCU["VCUBatPrsAlrm"] = p.F_evt_D00F.F_VCUBatPrsAlrm
	}
	if p.F_evt_D00F != nil {
		data_VCU["VCUBatPrsAlrmV"] = p.F_evt_D00F.F_VCUBatPrsAlrmV
	}
	if len(data_VCU) > 0 {
		channels = append(channels, Channel{
			ID:                  16,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_VCU},
		})
	}

	//group RDR
	data_RDR := make(map[string]any)
	if p.F_evt_D017 != nil {
		data_RDR["DTCInfomationRrDetnRdr"] = p.F_evt_D017.F_DTCInfomationRrDetnRdr
	}
	if len(data_RDR) > 0 {
		channels = append(channels, Channel{
			ID:                  47,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_RDR},
		})
	}

	//group EHBS
	data_EHBS := make(map[string]any)
	if p.F_evt_D006 != nil {
		data_EHBS["IbstrWrnngIO"] = p.F_evt_D006.F_IbstrWrnngIO
	}
	if len(data_EHBS) > 0 {
		channels = append(channels, Channel{
			ID:                  10,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_EHBS},
		})
	}

	//group EOPC
	data_EOPC := make(map[string]any)
	if p.F_evt_D011 != nil {
		data_EOPC["DTCInfomationEOPC"] = p.F_evt_D011.F_DTCInfomationEOPC
	}
	if len(data_EOPC) > 0 {
		channels = append(channels, Channel{
			ID:                  24,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_EOPC},
		})
	}

	//group RLSM
	data_RLSM := make(map[string]any)
	if p.F_evt_D016 != nil {
		data_RLSM["DTCInfomationRLSM"] = p.F_evt_D016.F_DTCInfomationRLSM
	}
	if len(data_RLSM) > 0 {
		channels = append(channels, Channel{
			ID:                  52,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_RLSM},
		})
	}

	//group LHCMS
	data_LHCMS := make(map[string]any)
	if p.F_evt_D016 != nil {
		data_LHCMS["DTCInfomationLHCMS"] = p.F_evt_D016.F_DTCInfomationLHCMS
	}
	if len(data_LHCMS) > 0 {
		channels = append(channels, Channel{
			ID:                  41,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_LHCMS},
		})
	}

	//group DLP
	data_DLP := make(map[string]any)
	if p.F_evt_D013 != nil {
		data_DLP["DTCInfomationDLP"] = p.F_evt_D013.F_DTCInfomationDLP
	}
	if len(data_DLP) > 0 {
		channels = append(channels, Channel{
			ID:                  23,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_DLP},
		})
	}

	//group FLIDAR
	data_FLIDAR := make(map[string]any)
	if p.F_evt_D017 != nil {
		data_FLIDAR["DTCInfomationFLIDAR"] = p.F_evt_D017.F_DTCInfomationFLIDAR
	}
	if len(data_FLIDAR) > 0 {
		channels = append(channels, Channel{
			ID:                  29,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_FLIDAR},
		})
	}

	//group FVCM
	data_FVCM := make(map[string]any)
	if p.F_evt_D017 != nil {
		data_FVCM["DTCInfomationFVCM"] = p.F_evt_D017.F_DTCInfomationFVCM
	}
	if len(data_FVCM) > 0 {
		channels = append(channels, Channel{
			ID:                  33,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_FVCM},
		})
	}

	//group CCU
	data_CCU := make(map[string]any)
	if p.F_evt_D006 != nil {
		data_CCU["HVDCDCSta"] = p.F_evt_D006.F_HVDCDCSta
	}
	if p.F_evt_D006 != nil {
		data_CCU["HVDCDCHVSideVol"] = p.F_evt_D006.F_HVDCDCHVSideVol
	}
	if p.F_evt_D006 != nil {
		data_CCU["HVDCDCHVSideVolV"] = p.F_evt_D006.F_HVDCDCHVSideVolV
	}
	if p.F_evt_D006 != nil {
		data_CCU["HVDCDCTem"] = p.F_evt_D006.F_HVDCDCTem
	}
	if p.F_evt_D011 != nil {
		data_CCU["DTCInfomationCCU"] = p.F_evt_D011.F_DTCInfomationCCU
	}
	if p.F_evt_D019 != nil {
		data_CCU["HVDCDCLVSideVol"] = p.F_evt_D019.F_HVDCDCLVSideVol
	}
	if len(data_CCU) > 0 {
		channels = append(channels, Channel{
			ID:                  6,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_CCU},
		})
	}

	//group EPMCU
	data_EPMCU := make(map[string]any)
	if p.F_evt_D011 != nil {
		data_EPMCU["DTCInfomationEPMCU"] = p.F_evt_D011.F_DTCInfomationEPMCU
	}
	if len(data_EPMCU) > 0 {
		channels = append(channels, Channel{
			ID:                  25,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_EPMCU},
		})
	}

	//group SCU
	data_SCU := make(map[string]any)
	if p.F_evt_D011 != nil {
		data_SCU["DTCInfomationSCU"] = p.F_evt_D011.F_DTCInfomationSCU
	}
	if len(data_SCU) > 0 {
		channels = append(channels, Channel{
			ID:                  60,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_SCU},
		})
	}

	//group LFSDA
	data_LFSDA := make(map[string]any)
	if p.F_evt_D010 != nil {
		data_LFSDA["DTCInfomationLFSDA"] = p.F_evt_D010.F_DTCInfomationLFSDA
	}
	if len(data_LFSDA) > 0 {
		channels = append(channels, Channel{
			ID:                  40,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_LFSDA},
		})
	}

	//group RRSM
	data_RRSM := make(map[string]any)
	if p.F_evt_D016 != nil {
		data_RRSM["DTCInfomationRRSM"] = p.F_evt_D016.F_DTCInfomationRRSM
	}
	if len(data_RRSM) > 0 {
		channels = append(channels, Channel{
			ID:                  54,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_RRSM},
		})
	}

	//group IPD
	data_IPD := make(map[string]any)
	if p.F_evt_D010 != nil {
		data_IPD["DTCInfomationIPD"] = p.F_evt_D010.F_DTCInfomationIPD
	}
	if len(data_IPD) > 0 {
		channels = append(channels, Channel{
			ID:                  38,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_IPD},
		})
	}

	//group PGM
	data_PGM := make(map[string]any)
	if p.F_evt_D014 != nil {
		data_PGM["DTCInfomationPGM"] = p.F_evt_D014.F_DTCInfomationPGM
	}
	if len(data_PGM) > 0 {
		channels = append(channels, Channel{
			ID:                  44,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_PGM},
		})
	}

	//group IMCU
	data_IMCU := make(map[string]any)
	if p.F_evt_0800 != nil {
		data_IMCU["TrShftLvrPos"] = p.F_evt_0800.F_TrShftLvrPos
	}
	if p.F_evt_0800 != nil {
		data_IMCU["TrShftLvrPosV"] = p.F_evt_0800.F_TrShftLvrPosV
	}
	if p.F_evt_D006 != nil {
		data_IMCU["EPTRdy"] = p.F_evt_D006.F_EPTRdy
	}
	if p.F_evt_D006 != nil {
		data_IMCU["ElecVehSysMd"] = p.F_evt_D006.F_ElecVehSysMd
	}
	if p.F_evt_D006 != nil {
		data_IMCU["EPTTrInptShaftToq"] = p.F_evt_D006.F_EPTTrInptShaftToq
	}
	if p.F_evt_D006 != nil {
		data_IMCU["EPTTrInptShaftToqV"] = p.F_evt_D006.F_EPTTrInptShaftToqV
	}
	if p.F_evt_D006 != nil {
		data_IMCU["EPTTrOtptShaftToq"] = p.F_evt_D006.F_EPTTrOtptShaftToq
	}
	if p.F_evt_D006 != nil {
		data_IMCU["EPTTrOtptShaftToqV"] = p.F_evt_D006.F_EPTTrOtptShaftToqV
	}
	if p.F_evt_D006 != nil {
		data_IMCU["EPTBrkPdlDscrtInptSts"] = p.F_evt_D006.F_EPTBrkPdlDscrtInptSts
	}
	if p.F_evt_D006 != nil {
		data_IMCU["EPTBrkPdlDscrtInptStsV"] = p.F_evt_D006.F_EPTBrkPdlDscrtInptStsV
	}
	if p.F_evt_D006 != nil {
		data_IMCU["EPTAccelActuPos"] = p.F_evt_D006.F_EPTAccelActuPos
	}
	if p.F_evt_D006 != nil {
		data_IMCU["EPTAccelActuPosV"] = p.F_evt_D006.F_EPTAccelActuPosV
	}
	if p.F_evt_D006 != nil {
		data_IMCU["TMInvtrCrntV"] = p.F_evt_D006.F_TMInvtrCrntV
	}
	if p.F_evt_D006 != nil {
		data_IMCU["TMInvtrCrnt"] = p.F_evt_D006.F_TMInvtrCrnt
	}
	if p.F_evt_D006 != nil {
		data_IMCU["TMSta"] = p.F_evt_D006.F_TMSta
	}
	if p.F_evt_D006 != nil {
		data_IMCU["TMInvtrTem"] = p.F_evt_D006.F_TMInvtrTem
	}
	if p.F_evt_D006 != nil {
		data_IMCU["TMSpd"] = p.F_evt_D006.F_TMSpd
	}
	if p.F_evt_D006 != nil {
		data_IMCU["TMSpdV"] = p.F_evt_D006.F_TMSpdV
	}
	if p.F_evt_D006 != nil {
		data_IMCU["TMActuToq"] = p.F_evt_D006.F_TMActuToq
	}
	if p.F_evt_D006 != nil {
		data_IMCU["TMActuToqV"] = p.F_evt_D006.F_TMActuToqV
	}
	if p.F_evt_D006 != nil {
		data_IMCU["TMSttrTem"] = p.F_evt_D006.F_TMSttrTem
	}
	if p.F_evt_D006 != nil {
		data_IMCU["TMInvtrVolV"] = p.F_evt_D006.F_TMInvtrVolV
	}
	if p.F_evt_D006 != nil {
		data_IMCU["TMInvtrVol"] = p.F_evt_D006.F_TMInvtrVol
	}
	if p.F_evt_D006 != nil {
		data_IMCU["EPTTrOtptShaftTotToq"] = p.F_evt_D006.F_EPTTrOtptShaftTotToq
	}
	if p.F_evt_D006 != nil {
		data_IMCU["EPTTrOtptShaftTotToqV"] = p.F_evt_D006.F_EPTTrOtptShaftTotToqV
	}
	if p.F_evt_D008 != nil {
		data_IMCU["DTCInfomationIMCU"] = p.F_evt_D008.F_DTCInfomationIMCU
	}
	if p.F_evt_D009 != nil {
		data_IMCU["TMRtrTem"] = p.F_evt_D009.F_TMRtrTem
	}
	if p.F_evt_D009 != nil {
		data_IMCU["TMStrOvTempAlrm"] = p.F_evt_D009.F_TMStrOvTempAlrm
	}
	if p.F_evt_D009 != nil {
		data_IMCU["TMInvtrOvTempAlrm"] = p.F_evt_D009.F_TMInvtrOvTempAlrm
	}
	if p.F_evt_D009 != nil {
		data_IMCU["EPTHVDCDCMdReq"] = p.F_evt_D009.F_EPTHVDCDCMdReq
	}
	if p.F_evt_D009 != nil {
		data_IMCU["VCUSecyWrnngInfo"] = p.F_evt_D009.F_VCUSecyWrnngInfo
	}
	if p.F_evt_D009 != nil {
		data_IMCU["VCUSecyWrnngInfoPV"] = p.F_evt_D009.F_VCUSecyWrnngInfoPV
	}
	if p.F_evt_D009 != nil {
		data_IMCU["VCUSecyWrnngInfoRC"] = p.F_evt_D009.F_VCUSecyWrnngInfoRC
	}
	if p.F_evt_D009 != nil {
		data_IMCU["VCUSecyWrnngInfoCRC"] = p.F_evt_D009.F_VCUSecyWrnngInfoCRC
	}
	if p.F_evt_D019 != nil {
		data_IMCU["EnrgSplReqEPTRdy"] = p.F_evt_D019.F_EnrgSplReqEPTRdy
	}
	if p.F_evt_D019 != nil {
		data_IMCU["HVEstbCond"] = p.F_evt_D019.F_HVEstbCond
	}
	if p.F_evt_D008 != nil {
		data_IMCU["DTCInfomationTC"] = p.F_evt_D008.F_DTCInfomationTC
	}
	if len(data_IMCU) > 0 {
		channels = append(channels, Channel{
			ID:                  3,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_IMCU},
		})
	}

	//group ICM
	data_ICM := make(map[string]any)
	if p.F_evt_D006 != nil {
		data_ICM["AvgFuelCsump"] = p.F_evt_D006.F_AvgFuelCsump
	}
	if p.F_evt_D014 != nil {
		data_ICM["DTCInfomationICM"] = p.F_evt_D014.F_DTCInfomationICM
	}
	if len(data_ICM) > 0 {
		channels = append(channels, Channel{
			ID:                  9,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_ICM},
		})
	}

	//group ECM
	data_ECM := make(map[string]any)
	if p.F_evt_D006 != nil {
		data_ECM["EnSpd"] = p.F_evt_D006.F_EnSpd
	}
	if p.F_evt_D006 != nil {
		data_ECM["EnSpdSts"] = p.F_evt_D006.F_EnSpdSts
	}
	if p.F_evt_D006 != nil {
		data_ECM["FuelCsump"] = p.F_evt_D006.F_FuelCsump
	}
	if p.F_evt_D008 != nil {
		data_ECM["DTCInfomationECM"] = p.F_evt_D008.F_DTCInfomationECM
	}
	if len(data_ECM) > 0 {
		channels = append(channels, Channel{
			ID:                  11,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_ECM},
		})
	}

	//group RHRDA
	data_RHRDA := make(map[string]any)
	if p.F_evt_D010 != nil {
		data_RHRDA["DTCInfomationRHRDA"] = p.F_evt_D010.F_DTCInfomationRHRDA
	}
	if len(data_RHRDA) > 0 {
		channels = append(channels, Channel{
			ID:                  50,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_RHRDA},
		})
	}

	//group EPS
	data_EPS := make(map[string]any)
	if p.F_evt_D012 != nil {
		data_EPS["DTCInfomationEPS"] = p.F_evt_D012.F_DTCInfomationEPS
	}
	if len(data_EPS) > 0 {
		channels = append(channels, Channel{
			ID:                  26,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_EPS},
		})
	}

	//group FRDCM
	data_FRDCM := make(map[string]any)
	if p.F_evt_D013 != nil {
		data_FRDCM["DTCInfomationDCM_FR"] = p.F_evt_D013.F_DTCInfomationDCM_FR
	}
	if len(data_FRDCM) > 0 {
		channels = append(channels, Channel{
			ID:                  31,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_FRDCM},
		})
	}

	//group ATC
	data_ATC := make(map[string]any)
	if p.F_evt_D013 != nil {
		data_ATC["DTCInfomationATC"] = p.F_evt_D013.F_DTCInfomationATC
	}
	if len(data_ATC) > 0 {
		channels = append(channels, Channel{
			ID:                  20,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_ATC},
		})
	}

	//group IMATE
	data_IMATE := make(map[string]any)
	if p.F_evt_D014 != nil {
		data_IMATE["DTCInfomationIMATE"] = p.F_evt_D014.F_DTCInfomationIMATE
	}
	if len(data_IMATE) > 0 {
		channels = append(channels, Channel{
			ID:                  36,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_IMATE},
		})
	}

	//group ESS
	data_ESS := make(map[string]any)
	if p.F_evt_000E != nil {
		data_ESS["BMSChrgSts"] = p.F_evt_000E.F_BMSChrgSts
	}
	if p.F_evt_000E != nil {
		data_ESS["BMSPackSOCBkup"] = p.F_evt_000E.F_BMSPackSOCBkup
	}
	if p.F_evt_000E != nil {
		data_ESS["BMSPackSOCVBkup"] = p.F_evt_000E.F_BMSPackSOCVBkup
	}
	if p.F_evt_000E != nil {
		data_ESS["BMSOfbdChrgSpRsn"] = p.F_evt_000E.F_BMSOfbdChrgSpRsn
	}
	if p.F_evt_000E != nil {
		data_ESS["BMSWrlsChrgSpRsn"] = p.F_evt_000E.F_BMSWrlsChrgSpRsn
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSBscSta"] = p.F_evt_D006.F_BMSBscSta
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSPackCrnt"] = p.F_evt_D006.F_BMSPackCrnt
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSPackCrntV"] = p.F_evt_D006.F_BMSPackCrntV

	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSPackSOC"] = p.F_evt_D006.F_BMSPackSOC
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSPackSOCV"] = p.F_evt_D006.F_BMSPackSOCV
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSPackSOCDsp"] = p.F_evt_D006.F_BMSPackSOCDsp
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSPackSOCDspV"] = p.F_evt_D006.F_BMSPackSOCDspV
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSPackVol"] = p.F_evt_D006.F_BMSPackVol
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSPackVolV"] = p.F_evt_D006.F_BMSPackVolV
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSPtIsltnRstc"] = p.F_evt_D006.F_BMSPtIsltnRstc
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSCellMaxTemIndx"] = p.F_evt_D006.F_BMSCellMaxTemIndx
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSCellMaxTem"] = p.F_evt_D006.F_BMSCellMaxTem
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSCellMaxTemV"] = p.F_evt_D006.F_BMSCellMaxTemV
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSCellMinTemIndx"] = p.F_evt_D006.F_BMSCellMinTemIndx
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSCellMinTem"] = p.F_evt_D006.F_BMSCellMinTem
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSCellMinTemV"] = p.F_evt_D006.F_BMSCellMinTemV
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSCellMaxVolIndx"] = p.F_evt_D006.F_BMSCellMaxVolIndx
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSCellMaxVol"] = p.F_evt_D006.F_BMSCellMaxVol
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSCellMaxVolV"] = p.F_evt_D006.F_BMSCellMaxVolV
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSCellMinVolIndx"] = p.F_evt_D006.F_BMSCellMinVolIndx
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSCellMinVol"] = p.F_evt_D006.F_BMSCellMinVol
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSCellMinVolV"] = p.F_evt_D006.F_BMSCellMinVolV
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSPtIsltnRstcV"] = p.F_evt_D006.F_BMSPtIsltnRstcV
	}
	if p.F_evt_D006 != nil {
		data_ESS["BMSHVILClsd"] = p.F_evt_D006.F_BMSHVILClsd
	}
	if p.F_evt_D008 != nil {
		data_ESS["DTCInfomationBMS"] = p.F_evt_D008.F_DTCInfomationBMS
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSCMUFlt"] = p.F_evt_D009.F_BMSCMUFlt
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSCellVoltFlt"] = p.F_evt_D009.F_BMSCellVoltFlt
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSPackTemFlt"] = p.F_evt_D009.F_BMSPackTemFlt
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSPackVoltFlt"] = p.F_evt_D009.F_BMSPackVoltFlt
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSWrnngInfo"] = p.F_evt_D009.F_BMSWrnngInfo

	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSWrnngInfoPV"] = p.F_evt_D009.F_BMSWrnngInfoPV
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSWrnngInfoRC"] = p.F_evt_D009.F_BMSWrnngInfoRC
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSPreThrmFltInd"] = p.F_evt_D009.F_BMSPreThrmFltInd
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSKeepSysAwkScene"] = p.F_evt_D009.F_BMSKeepSysAwkScene
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSTemOverDifAlrm"] = p.F_evt_D009.F_BMSTemOverDifAlrm
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSOverTemAlrm"] = p.F_evt_D009.F_BMSOverTemAlrm
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSOverPackVolAlrm"] = p.F_evt_D009.F_BMSOverPackVolAlrm
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSUnderPackVolAlrm"] = p.F_evt_D009.F_BMSUnderPackVolAlrm
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSHVILAlrm"] = p.F_evt_D009.F_BMSHVILAlrm
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSOverCellVolAlrm"] = p.F_evt_D009.F_BMSOverCellVolAlrm
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSUnderCellVolAlrm"] = p.F_evt_D009.F_BMSUnderCellVolAlrm
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSLowSOCAlrm"] = p.F_evt_D009.F_BMSLowSOCAlrm
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSJumpngSOCAlrm"] = p.F_evt_D009.F_BMSJumpngSOCAlrm
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSHiSOCAlrm"] = p.F_evt_D009.F_BMSHiSOCAlrm

	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSPackVolMsmchAlrm"] = p.F_evt_D009.F_BMSPackVolMsmchAlrm
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSPoorCellCnstncyAlrm"] = p.F_evt_D009.F_BMSPoorCellCnstncyAlrm
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSCellOverChrgdAlrm"] = p.F_evt_D009.F_BMSCellOverChrgdAlrm
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSLowPtIsltnRstcAlrm"] = p.F_evt_D009.F_BMSLowPtIsltnRstcAlrm
	}
	if p.F_evt_D009 != nil {
		data_ESS["BMSOnbdChrgSpRsn"] = p.F_evt_D009.F_BMSOnbdChrgSpRsn
	}
	if p.F_evt_D00B != nil {
		data_ESS["BMSCellVolSumNum"] = p.F_evt_D00B.F_BMSCellVolSumNum
	}
	if p.F_evt_D00B != nil {
		data_ESS["BMSCellVol"] = p.F_evt_D00B.F_BMSCellVol
	}
	if p.F_evt_D00B != nil {
		data_ESS["BMSCellVolV"] = p.F_evt_D00B.F_BMSCellVolV
	}
	if p.F_evt_D00C != nil {
		data_ESS["BMSCellTemSumNum"] = p.F_evt_D00C.F_BMSCellTemSumNum
	}
	if p.F_evt_D00C != nil {
		data_ESS["BMSCellTem"] = p.F_evt_D00C.F_BMSCellTem
	}
	if p.F_evt_D00C != nil {
		data_ESS["BMSCellTemV"] = p.F_evt_D00C.F_BMSCellTemV
	}
	if p.F_evt_D00D != nil {
		data_ESS["BMSBusbarTemSumNum"] = p.F_evt_D00D.F_BMSBusbarTemSumNum
	}
	if p.F_evt_D00D != nil {
		data_ESS["BMSBusbarTem"] = p.F_evt_D00D.F_BMSBusbarTem

	}
	if p.F_evt_D00D != nil {
		data_ESS["BMSBusbarTemV"] = p.F_evt_D00D.F_BMSBusbarTemV
	}
	if p.F_evt_D00E != nil {
		data_ESS["BMSRptBatCodeNum"] = p.F_evt_D00E.F_BMSRptBatCodeNum
	}
	if p.F_evt_D00E != nil {
		data_ESS["BMSRptBatCodeAsc"] = p.F_evt_D00E.F_BMSRptBatCodeAsc
	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSWrnngInfoCRC"] = p.F_evt_D00F.F_BMSWrnngInfoCRC
	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSBusbarTempMax"] = p.F_evt_D00F.F_BMSBusbarTempMax
	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSPreThrmFltIndBkup"] = p.F_evt_D00F.F_BMSPreThrmFltIndBkup
	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSWrnngInfoRCBkup"] = p.F_evt_D00F.F_BMSWrnngInfoRCBkup
	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSBatPrsFlt"] = p.F_evt_D00F.F_BMSBatPrsFlt

	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSWrnngInfoBkup"] = p.F_evt_D00F.F_BMSWrnngInfoBkup
	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSBatPrsAlrm"] = p.F_evt_D00F.F_BMSBatPrsAlrm
	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSBatPrsAlrmV"] = p.F_evt_D00F.F_BMSBatPrsAlrmV
	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSBatPrsSnsrV"] = p.F_evt_D00F.F_BMSBatPrsSnsrV
	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSBatPrsSnsrValBkup"] = p.F_evt_D00F.F_BMSBatPrsSnsrValBkup
	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSBatPrsSnsrVBkup"] = p.F_evt_D00F.F_BMSBatPrsSnsrVBkup
	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSBatPrsSnsrVal"] = p.F_evt_D00F.F_BMSBatPrsSnsrVal
	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSClntPumpPWMReq"] = p.F_evt_D00F.F_BMSClntPumpPWMReq
	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSPumpPwrOnReq"] = p.F_evt_D00F.F_BMSPumpPwrOnReq
	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSBatPrsAlrmVBkup"] = p.F_evt_D00F.F_BMSBatPrsAlrmVBkup
	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSBatPrsAlrmBkup"] = p.F_evt_D00F.F_BMSBatPrsAlrmBkup
	}
	if p.F_evt_D00F != nil {
		data_ESS["BMSWrnngInfoCRCBkup"] = p.F_evt_D00F.F_BMSWrnngInfoCRCBkup
	}
	if len(data_ESS) > 0 {
		channels = append(channels, Channel{
			ID:                  5,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_ESS},
		})
	}

	//group BCM
	data_BCM := make(map[string]any)
	if p.F_evt_D008 != nil {
		data_BCM["DTCInfomationTPMS"] = p.F_evt_D008.F_DTCInfomationTPMS
	}
	if p.F_evt_D013 != nil {
		data_BCM["DTCInfomationBCM"] = p.F_evt_D013.F_DTCInfomationBCM
	}
	if len(data_BCM) > 0 {
		channels = append(channels, Channel{
			ID:                  14,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_BCM},
		})
	}

	//group CCP
	data_CCP := make(map[string]any)
	if p.F_evt_D00F != nil {
		data_CCP["OtsdAirTemCrVal"] = p.F_evt_D00F.F_OtsdAirTemCrVal
	}
	if p.F_evt_D00F != nil {
		data_CCP["OtsdAirTemCrValV"] = p.F_evt_D00F.F_OtsdAirTemCrValV
	}
	if len(data_CCP) > 0 {
		channels = append(channels, Channel{
			ID:                  15,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_CCP},
		})
	}

	//group LVBM
	data_LVBM := make(map[string]any)
	if p.F_evt_D016 != nil {
		data_LVBM["DTCInfomationLVBM"] = p.F_evt_D016.F_DTCInfomationLVBM
	}
	if len(data_LVBM) > 0 {
		channels = append(channels, Channel{
			ID:                  43,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_LVBM},
		})
	}

	//group SAC
	data_SAC := make(map[string]any)
	if p.F_evt_D006 != nil {
		data_SAC["SAMInvtrCrnt"] = p.F_evt_D006.F_SAMInvtrCrnt

	}
	if p.F_evt_D006 != nil {
		data_SAC["SAMInvtrCrntV"] = p.F_evt_D006.F_SAMInvtrCrntV
	}
	if p.F_evt_D006 != nil {
		data_SAC["SAMSta"] = p.F_evt_D006.F_SAMSta
	}
	if p.F_evt_D006 != nil {
		data_SAC["SAMInvtrTem"] = p.F_evt_D006.F_SAMInvtrTem
	}
	if p.F_evt_D006 != nil {
		data_SAC["SAMSpdV"] = p.F_evt_D006.F_SAMSpdV
	}
	if p.F_evt_D006 != nil {
		data_SAC["SAMSpd"] = p.F_evt_D006.F_SAMSpd
	}
	if p.F_evt_D006 != nil {
		data_SAC["SAMActuToqV"] = p.F_evt_D006.F_SAMActuToqV
	}
	if p.F_evt_D006 != nil {
		data_SAC["SAMActuToq"] = p.F_evt_D006.F_SAMActuToq
	}
	if p.F_evt_D006 != nil {
		data_SAC["SAMSttrTem"] = p.F_evt_D006.F_SAMSttrTem
	}
	if p.F_evt_D006 != nil {
		data_SAC["SAMInvtrVolV"] = p.F_evt_D006.F_SAMInvtrVolV

	}
	if p.F_evt_D006 != nil {
		data_SAC["SAMInvtrVol"] = p.F_evt_D006.F_SAMInvtrVol
	}
	if p.F_evt_D008 != nil {
		data_SAC["DTCInfomationSAC"] = p.F_evt_D008.F_DTCInfomationSAC
	}
	if p.F_evt_D009 != nil {
		data_SAC["SAMStrOvTempAlrm"] = p.F_evt_D009.F_SAMStrOvTempAlrm
	}
	if p.F_evt_D009 != nil {
		data_SAC["SAMInvtrOvTempAlrm"] = p.F_evt_D009.F_SAMInvtrOvTempAlrm
	}
	if len(data_SAC) > 0 {
		channels = append(channels, Channel{
			ID:                  8,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_SAC},
		})
	}

	//group LHRDA
	data_LHRDA := make(map[string]any)
	if p.F_evt_D010 != nil {
		data_LHRDA["DTCInfomationLHRDA"] = p.F_evt_D010.F_DTCInfomationLHRDA
	}
	if len(data_LHRDA) > 0 {
		channels = append(channels, Channel{
			ID:                  42,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_LHRDA},
		})
	}

	//group TCM
	data_TCM := make(map[string]any)
	if p.F_evt_D008 != nil {
		data_TCM["DTCInfomationTCM"] = p.F_evt_D008.F_DTCInfomationTCM
	}
	if len(data_TCM) > 0 {
		channels = append(channels, Channel{
			ID:                  13,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_TCM},
		})
	}

	//group EPS_SFCANFD
	data_EPS_SFCANFD := make(map[string]any)
	if p.F_evt_D012 != nil {
		data_EPS_SFCANFD["DTCInfomationEPS_S"] = p.F_evt_D012.F_DTCInfomationEPS_S
	}
	if len(data_EPS_SFCANFD) > 0 {
		channels = append(channels, Channel{
			ID:                  27,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_EPS_SFCANFD},
		})
	}

	//group IPS
	data_IPS := make(map[string]any)
	if p.F_evt_D017 != nil {
		data_IPS["DTCInfomationIPS"] = p.F_evt_D017.F_DTCInfomationIPS
	}
	if len(data_IPS) > 0 {
		channels = append(channels, Channel{
			ID:                  39,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_IPS},
		})
	}

	//group IAM
	data_IAM := make(map[string]any)
	if p.F_evt_0001 != nil {
		data_IAM["TBOXSysTim"] = p.F_evt_0001.F_TBOXSysTim
	}
	if p.F_evt_0003 != nil {
		data_IAM["RelwakeupTim"] = p.F_evt_0003.F_RelwakeupTim
	}
	if p.F_evt_0004 != nil {
		data_IAM["GnssAlt"] = p.F_evt_0004.F_GnssAlt
	}
	if p.F_evt_0004 != nil {
		data_IAM["Longitude"] = p.F_evt_0004.F_Longitude
	}
	if p.F_evt_0004 != nil {
		data_IAM["GPSSts"] = p.F_evt_0004.F_GPSSts
	}
	if p.F_evt_0005 != nil {
		data_IAM["Latitude"] = p.F_evt_0005.F_Latitude
	}
	if p.F_evt_0005 != nil {
		data_IAM["VehTyp"] = p.F_evt_0005.F_VehTyp
	}
	if p.F_evt_0005 != nil {
		data_IAM["GNSSDirection"] = p.F_evt_0005.F_GNSSDirection
	}
	if p.F_evt_0006 != nil {
		data_IAM["HDop"] = p.F_evt_0006.F_HDop
	}
	if p.F_evt_0006 != nil {
		data_IAM["VDop"] = p.F_evt_0006.F_VDop
	}
	if p.F_evt_0007 != nil {
		data_IAM["AcceX"] = p.F_evt_0007.F_AcceX
	}
	if p.F_evt_0007 != nil {
		data_IAM["AcceY"] = p.F_evt_0007.F_AcceY
	}
	if p.F_evt_0007 != nil {
		data_IAM["AcceZ"] = p.F_evt_0007.F_AcceZ
	}
	if p.F_evt_0008 != nil {
		data_IAM["cellMCC"] = p.F_evt_0008.F_cellMCC
	}
	if p.F_evt_0008 != nil {
		data_IAM["cellMNC"] = p.F_evt_0008.F_cellMNC
	}
	if p.F_evt_0008 != nil {
		data_IAM["millisecond"] = p.F_evt_0008.F_millisecond
	}
	if p.F_evt_0008 != nil {
		data_IAM["spistatus"] = p.F_evt_0008.F_spistatus
	}
	if p.F_evt_0009 != nil {
		data_IAM["cellLAC"] = p.F_evt_0009.F_cellLAC
	}
	if p.F_evt_0009 != nil {
		data_IAM["CellID"] = p.F_evt_0009.F_CellID
	}
	if p.F_evt_000A != nil {
		data_IAM["cellSignalStrength"] = p.F_evt_000A.F_cellSignalStrength
	}
	if p.F_evt_000A != nil {
		data_IAM["cellRAT"] = p.F_evt_000A.F_cellRAT
	}
	if p.F_evt_000A != nil {
		data_IAM["cellRATadd"] = p.F_evt_000A.F_cellRATadd
	}
	if p.F_evt_000A != nil {
		data_IAM["cellChanID"] = p.F_evt_000A.F_cellChanID
	}
	if p.F_evt_000A != nil {
		data_IAM["GNSSSATS"] = p.F_evt_000A.F_GNSSSATS
	}
	if p.F_evt_000B != nil {
		data_IAM["ModemStates"] = p.F_evt_000B.F_ModemStates
	}
	if p.F_evt_000B != nil {
		data_IAM["iNetworkSts"] = p.F_evt_000B.F_iNetworkSts
	}
	if p.F_evt_000B != nil {
		data_IAM["iNetworkSts_ErrCode"] = p.F_evt_000B.F_iNetworkSts_ErrCode
	}
	if p.F_evt_000C != nil {
		data_IAM["PotclVer"] = p.F_evt_000C.F_PotclVer
	}
	if p.F_evt_000C != nil {
		data_IAM["PotclSecyVer"] = p.F_evt_000C.F_PotclSecyVer
	}
	if p.F_evt_000C != nil {
		data_IAM["CalendarYear"] = p.F_evt_000C.F_CalendarYear
	}
	if p.F_evt_000C != nil {
		data_IAM["CalendarDay"] = p.F_evt_000C.F_CalendarDay
	}
	if p.F_evt_000C != nil {
		data_IAM["CalendarMonth"] = p.F_evt_000C.F_CalendarMonth
	}
	if p.F_evt_000D != nil {
		data_IAM["CellFrequency"] = p.F_evt_000D.F_CellFrequency
	}
	if p.F_evt_D00A != nil {
		data_IAM["VIN"] = p.F_evt_D00A.F_VIN
	}
	if p.F_evt_D00A != nil {
		data_IAM["IAMSN"] = p.F_evt_D00A.F_IAMSN
	}
	if p.F_evt_D00A != nil {
		data_IAM["EsimIccid"] = p.F_evt_D00A.F_EsimIccid
	}
	if p.F_evt_D00A != nil {
		data_IAM["EsimID"] = p.F_evt_D00A.F_EsimID
	}
	if p.F_evt_D010 != nil {
		data_IAM["DTCInfomationIAM"] = p.F_evt_D010.F_DTCInfomationIAM
	}
	if p.F_evt_D018 != nil {
		data_IAM["APN1ConnSts"] = p.F_evt_D018.F_APN1ConnSts
	}
	if p.F_evt_D018 != nil {
		data_IAM["APN2ConnSts"] = p.F_evt_D018.F_APN2ConnSts
	}
	if p.F_evt_D018 != nil {
		data_IAM["ECallSts"] = p.F_evt_D018.F_ECallSts
	}
	if p.F_evt_D018 != nil {
		data_IAM["MqttConnFailRsn"] = p.F_evt_D018.F_MqttConnFailRsn
	}
	if p.F_evt_D018 != nil {
		data_IAM["LocDRSts"] = p.F_evt_D018.F_LocDRSts
	}
	if p.F_evt_D018 != nil {
		data_IAM["LongitudeDR"] = p.F_evt_D018.F_LongitudeDR
	}
	if p.F_evt_D018 != nil {
		data_IAM["LatitudeDR"] = p.F_evt_D018.F_LatitudeDR
	}
	if p.F_evt_D018 != nil {
		data_IAM["LocGnns1Sts"] = p.F_evt_D018.F_LocGnns1Sts
	}
	if p.F_evt_D018 != nil {
		data_IAM["TBOXGPSTime"] = p.F_evt_D018.F_TBOXGPSTime
	}
	if p.F_evt_D018 != nil {
		data_IAM["LocGnns2Sts"] = p.F_evt_D018.F_LocGnns2Sts
	}
	if p.F_evt_D018 != nil {
		data_IAM["LocRTKSts"] = p.F_evt_D018.F_LocRTKSts
	}
	if p.F_evt_D018 != nil {
		data_IAM["LocGnns1SatNum"] = p.F_evt_D018.F_LocGnns1SatNum
	}
	if p.F_evt_D018 != nil {
		data_IAM["LocGnns2SatNum"] = p.F_evt_D018.F_LocGnns2SatNum
	}
	if p.F_evt_D01A != nil {
		data_IAM["iEcuSts"] = p.F_evt_D01A.F_iEcuSts
	}
	if p.F_evt_D01A != nil {
		data_IAM["iIAMInterSts"] = p.F_evt_D01A.F_iIAMInterSts
	}
	if p.F_evt_D01A != nil {
		data_IAM["iMpuIPTableRuleSts"] = p.F_evt_D01A.F_iMpuIPTableRuleSts
	}
	if p.F_evt_D01A != nil {
		data_IAM["iModemIPTableRuleSts"] = p.F_evt_D01A.F_iModemIPTableRuleSts
	}
	if p.F_evt_D01A != nil {
		data_IAM["iARPRuleSts"] = p.F_evt_D01A.F_iARPRuleSts
	}
	if p.F_evt_D01A != nil {
		data_IAM["iICC2PHYSGMIISts"] = p.F_evt_D01A.F_iICC2PHYSGMIISts
	}
	if p.F_evt_D01A != nil {
		data_IAM["iMpuRGMIISts"] = p.F_evt_D01A.F_iMpuRGMIISts
	}
	if p.F_evt_D01A != nil {
		data_IAM["iModemRGMIISts"] = p.F_evt_D01A.F_iModemRGMIISts
	}
	if p.F_evt_D01A != nil {
		data_IAM["iSwitchSGMIISts"] = p.F_evt_D01A.F_iSwitchSGMIISts
	}
	if p.F_evt_D01A != nil {
		data_IAM["iUSBConnSts"] = p.F_evt_D01A.F_iUSBConnSts
	}
	if p.F_evt_D01A != nil {
		data_IAM["iIPASts"] = p.F_evt_D01A.F_iIPASts
	}
	if p.F_evt_D01A != nil {
		data_IAM["iAPSts"] = p.F_evt_D01A.F_iAPSts
	}
	if p.F_evt_D01A != nil {
		data_IAM["networkbackupinfo"] = p.F_evt_D01A.F_networkbackupinfo
	}
	if p.F_evt_D01B != nil {
		data_IAM["WANStatus"] = p.F_evt_D01B.F_WANStatus
	}
	if p.F_evt_D01B != nil {
		data_IAM["ChannelType1"] = p.F_evt_D01B.F_ChannelType1
	}
	if p.F_evt_D01B != nil {
		data_IAM["ChannelStates1"] = p.F_evt_D01B.F_ChannelStates1
	}
	if p.F_evt_D01B != nil {
		data_IAM["IPAddress1"] = p.F_evt_D01B.F_IPAddress1
	}
	if p.F_evt_D01B != nil {
		data_IAM["ChannelType2"] = p.F_evt_D01B.F_ChannelType2
	}
	if p.F_evt_D01B != nil {
		data_IAM["ChannelStates2"] = p.F_evt_D01B.F_ChannelStates2
	}
	if p.F_evt_D01B != nil {
		data_IAM["IPAddress2"] = p.F_evt_D01B.F_IPAddress2
	}
	if p.F_evt_D01B != nil {
		data_IAM["ChannelType3"] = p.F_evt_D01B.F_ChannelType3
	}
	if p.F_evt_D01B != nil {
		data_IAM["ChannelStates3"] = p.F_evt_D01B.F_ChannelStates3
	}
	if p.F_evt_D01B != nil {
		data_IAM["IPAddress3"] = p.F_evt_D01B.F_IPAddress3
	}
	if p.F_evt_D01B != nil {
		data_IAM["ChannelType4"] = p.F_evt_D01B.F_ChannelType4
	}
	if p.F_evt_D01B != nil {
		data_IAM["ChannelStates4"] = p.F_evt_D01B.F_ChannelStates4
	}
	if p.F_evt_D01C != nil {
		data_IAM["IPAddress4"] = p.F_evt_D01C.F_IPAddress4
	}
	if p.F_evt_D01C != nil {
		data_IAM["CurIMSI"] = p.F_evt_D01C.F_CurIMSI
	}
	if p.F_evt_D01C != nil {
		data_IAM["NetType"] = p.F_evt_D01C.F_NetType
	}
	if p.F_evt_D01C != nil {
		data_IAM["rssi"] = p.F_evt_D01C.F_rssi
	}
	if p.F_evt_D01C != nil {
		data_IAM["rsrp"] = p.F_evt_D01C.F_rsrp
	}
	if p.F_evt_D01C != nil {
		data_IAM["rscp"] = p.F_evt_D01C.F_rscp
	}
	if p.F_evt_D01C != nil {
		data_IAM["sinr"] = p.F_evt_D01C.F_sinr
	}
	if p.F_evt_D01C != nil {
		data_IAM["ecio"] = p.F_evt_D01C.F_ecio
	}
	if p.F_evt_D01D != nil {
		data_IAM["cellLAC5G"] = p.F_evt_D01D.F_cellLAC5G
	}
	if p.F_evt_D01D != nil {
		data_IAM["CellID5G"] = p.F_evt_D01D.F_CellID5G
	}
	if p.F_evt_D01F != nil {
		data_IAM["NetRecvRsn"] = p.F_evt_D01F.F_NetRecvRsn
	}
	if p.F_evt_D01F != nil {
		data_IAM["NetRecvActn"] = p.F_evt_D01F.F_NetRecvActn
	}
	if p.F_evt_D01F != nil {
		data_IAM["NetRecvActnTimstmp"] = p.F_evt_D01F.F_NetRecvActnTimstmp
	}
	if p.F_evt_D01F != nil {
		data_IAM["NetRecvActnCnt"] = p.F_evt_D01F.F_NetRecvActnCnt
	}
	if p.F_evt_D01F != nil {
		data_IAM["NetRecvActnRst"] = p.F_evt_D01F.F_NetRecvActnRst
	}
	if p.F_evt_D01F != nil {
		data_IAM["NetRecvtime"] = p.F_evt_D01F.F_NetRecvtime
	}
	if p.F_evt_FFFF != nil {
		data_IAM["CRC32"] = p.F_evt_FFFF.F_CRC32
	}
	if len(data_IAM) > 0 {
		channels = append(channels, Channel{
			ID:                  1,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_IAM},
		})
	}

	//group RFSDA
	data_RFSDA := make(map[string]any)
	if p.F_evt_D010 != nil {
		data_RFSDA["DTCInfomationRFSDA"] = p.F_evt_D010.F_DTCInfomationRFSDA
	}
	if len(data_RFSDA) > 0 {
		channels = append(channels, Channel{
			ID:                  48,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_RFSDA},
		})
	}

	//group RRDCM
	data_RRDCM := make(map[string]any)
	if p.F_evt_D013 != nil {
		data_RRDCM["DTCInfomationDCM_RR"] = p.F_evt_D013.F_DTCInfomationDCM_RR
	}
	if len(data_RRDCM) > 0 {
		channels = append(channels, Channel{
			ID:                  53,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_RRDCM},
		})
	}

	//group ISC
	data_ISC := make(map[string]any)
	if p.F_evt_D006 != nil {
		data_ISC["ISGInvtrCrntV"] = p.F_evt_D006.F_ISGInvtrCrntV
	}
	if p.F_evt_D006 != nil {
		data_ISC["ISGInvtrCrnt"] = p.F_evt_D006.F_ISGInvtrCrnt
	}
	if p.F_evt_D006 != nil {
		data_ISC["ISGSta"] = p.F_evt_D006.F_ISGSta
	}
	if p.F_evt_D006 != nil {
		data_ISC["ISGInvtrTem"] = p.F_evt_D006.F_ISGInvtrTem
	}
	if p.F_evt_D006 != nil {
		data_ISC["ISGSpd"] = p.F_evt_D006.F_ISGSpd
	}
	if p.F_evt_D006 != nil {
		data_ISC["ISGSpdV"] = p.F_evt_D006.F_ISGSpdV
	}
	if p.F_evt_D006 != nil {
		data_ISC["ISGActuToq"] = p.F_evt_D006.F_ISGActuToq
	}
	if p.F_evt_D006 != nil {
		data_ISC["ISGActuToqV"] = p.F_evt_D006.F_ISGActuToqV
	}
	if p.F_evt_D006 != nil {
		data_ISC["ISGSttrTem"] = p.F_evt_D006.F_ISGSttrTem
	}
	if p.F_evt_D006 != nil {
		data_ISC["ISGInvtrVolV"] = p.F_evt_D006.F_ISGInvtrVolV
	}
	if p.F_evt_D006 != nil {
		data_ISC["ISGInvtrVol"] = p.F_evt_D006.F_ISGInvtrVol
	}
	if p.F_evt_D008 != nil {
		data_ISC["DTCInfomationISC"] = p.F_evt_D008.F_DTCInfomationISC
	}
	if p.F_evt_D009 != nil {
		data_ISC["ISCStrOvTempAlrm"] = p.F_evt_D009.F_ISCStrOvTempAlrm
	}
	if p.F_evt_D009 != nil {
		data_ISC["ISCInvtrOvTempAlrm"] = p.F_evt_D009.F_ISCInvtrOvTempAlrm
	}
	if len(data_ISC) > 0 {
		channels = append(channels, Channel{
			ID:                  7,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_ISC},
		})
	}

	//group CARLog
	data_CARLog := make(map[string]any)
	if p.F_evt_D014 != nil {
		data_CARLog["DTCInfomationCARLog"] = p.F_evt_D014.F_DTCInfomationCARLog
	}
	if len(data_CARLog) > 0 {
		channels = append(channels, Channel{
			ID:                  22,
			Starttime:           ts / 1000,
			CollectiofrequecyHz: 1,
			Data:                []map[string]any{data_CARLog},
		})
	}

	sort.Slice(channels, func(i, j int) bool {
		return channels[i].ID < channels[j].ID
	})

	return Json{
		SAIC_FileVersion: "0.0.1.2",
		FileCreationTime: ts / 1000,
		Tboxinfo: Tboxinfo{
			ID:         "026HJ125J022",
			VIN:        vin,
			PartNumber: "1080533001",
		},
		Journey: Journey{
			JourneyID: 3681,
		},
		Channels: channels,
	}
}
