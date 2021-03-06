package core

import (
	"fmt"
	mkr "github.com/mackerelio/mackerel-client-go"
	"github.com/sacloud/libsacloud/sacloud"
	"reflect"
	"strconv"
	"strings"
)

//-----------------------------------------------------------------------------
// CreateHostPayload
//-----------------------------------------------------------------------------

// CreateHostPayload Mackerel上でのホスト作成/更新を伴うジョブで利用するペイロード
type CreateHostPayload struct {
	// SourcePayload ソースとなるサーバー/ホスト情報
	*SourcePayload

	// MackerelHostParam mackerelへのサーバー登録用パラメータ
	MackerelHostParam *mkr.CreateHostParam
}

// NewCreateHostPayload CreateHostPayloadの新規作成
func NewCreateHostPayload(source interface{}, zone string, resourceID int64, sourceType reflect.Type) *CreateHostPayload {
	return &CreateHostPayload{
		SourcePayload: NewSourcePayload(source, zone, resourceID, sourceType),
	}
}

// IsStatusUpdated ホストステータスの更新が行われているか判定
func (p *CreateHostPayload) IsStatusUpdated() (bool, error) {

	if p.MackerelHost == nil {
		return false, fmt.Errorf("MackerelHost is nil.")
	}
	if string(p.MackerelHostStatus) == "" {
		return false, fmt.Errorf("MackerelHostStatus is empty.")
	}

	return p.MackerelHost.Status != string(p.MackerelHostStatus), nil
}

// IsRoleUpdated ロール情報の更新が行われているか判定
func (p *CreateHostPayload) IsRoleUpdated() (bool, error) {
	if p.MackerelHost == nil {
		return false, fmt.Errorf("MackerelHost is nil.")
	}
	if p.MackerelHostParam == nil {
		return false, fmt.Errorf("MackerelHostParam is nil.")
	}

	// 数が違えばその時点で変更ありと判定
	if len(p.MackerelHost.GetRoleFullnames()) != len(p.MackerelHostParam.RoleFullnames) {
		return true, nil
	}

	for _, source := range p.MackerelHost.GetRoleFullnames() {
		isExists := false
		for _, target := range p.MackerelHostParam.RoleFullnames {
			if source == target {
				isExists = true
				break
			}
		}
		if !isExists {
			return true, nil
		}
	}

	return false, nil
}

// GetFindParam Mackerelからのホスト検索用パラメータを取得する
func (p *CreateHostPayload) GetFindParam() *mkr.FindHostsParam {

	mackerelName := p.GenerateMackerelName()
	if mackerelName == "" {
		return nil
	}

	return &mkr.FindHostsParam{
		CustomIdentifier: mackerelName,
		//Name:             mackerelName,
		Statuses: []string{
			string(MackerelHostStatusWorking),
			string(MackerelHostStatusStandby),
			string(MackerelHostStatusPowerOff),
			string(MackerelHostStatusMaintenance),
		},
	}
}

//-----------------------------------------------------------------------------
// CollectMetricsPayload
//-----------------------------------------------------------------------------

// CollectMetricsPayload メトリクス収集用ジョブで利用するペイロード
type CollectMetricsPayload struct {
	*SourcePayload
	Metrics              *SacloudMetrics
	MackerelMetricsParam []*mkr.HostMetricValue
}

// NewCollectMetricsPayload CollectMetricsPayloadを新規作成する
func NewCollectMetricsPayload(sourcePayload *SourcePayload) *CollectMetricsPayload {
	return &CollectMetricsPayload{
		SourcePayload: sourcePayload,
		Metrics:       &SacloudMetrics{},
	}
}

// SacloudMetrics さくらのクラウドから収集するメトリクスを格納する構造体
type SacloudMetrics struct {
	CPU       []*sacloud.MonitorValues
	Disk      []*sacloud.MonitorValues
	Interface []*sacloud.MonitorValues
	Database  []*sacloud.MonitorValues
}

//-----------------------------------------------------------------------------
// ReconcileHostsPayload
//-----------------------------------------------------------------------------

// ReconcileHostsPayload ホストリコンサイル用ジョブで利用するペイロード
type ReconcileHostsPayload struct {
	// FromSackerelHost sackerelが登録したホスト
	FromSackerelHost *mkr.Host
	// FromAgentHost mackerel-agentが登録したホスト
	FromAgentHost *mkr.Host
}

// GetSacloudServerInfo 対象Mackerelホストからさくらのクラウドリソース情報を取得する
func (p *ReconcileHostsPayload) GetSacloudServerInfo() (string, int64, error) {
	tokens := strings.Split(p.FromSackerelHost.Name, "-") // 想定形式:  SakuraCloud-[ゾーン名]-[リソースID]
	if len(tokens) != 3 {
		return "", -1, fmt.Errorf("TargetHost.Name is invalid : ['%s']", p.FromSackerelHost.Name)
	}

	strZone := strings.ToLower(tokens[1])
	strID := tokens[2]

	// id to int64
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		return "", 01, err
	}

	return strZone, id, nil
}
