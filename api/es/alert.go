package es

import (
	"fmt"
	"github.com/prometheus/alertmanager/types"
	"github.com/prometheus/common/model"
	"strconv"
	"time"
)

type LabelValue interface {
}

type LabelSet map[model.LabelName]LabelValue

const (
	labelNameValue model.LabelName = "value"
)

type ESAlert struct {
	Labels LabelSet `json:"labels"`

	Annotations model.LabelSet `json:"annotations"`

	StartsAt     time.Time `json:"startsAt,omitempty"`
	EndsAt       time.Time `json:"endsAt,omitempty"`
	GeneratorURL string    `json:"generatorURL"`
	UpdatedAt    time.Time
	Timeout      bool
}

func Convert(alert *types.Alert) *ESAlert {
	//var result float64
	//var err error
	esAlert := new(ESAlert)
	esAlert.EndsAt = alert.EndsAt
	esAlert.StartsAt = alert.StartsAt
	esAlert.UpdatedAt = alert.UpdatedAt
	esAlert.Annotations = alert.Annotations
	esAlert.GeneratorURL = alert.GeneratorURL
	esAlert.Timeout = alert.Timeout
	//	esAlert.Labels = alert.Labels
	esls := make(LabelSet)
	for k, v := range alert.Labels {
		if k == labelNameValue {
			result, err := strconv.ParseFloat(fmt.Sprintf("%s", v), 32)
			if err != nil {
				esls[k] = v
			} else {
				esls[k] = result
			}
		} else {
			esls[k] = v
		}
	}
	//	esls = alert.Labels
	esAlert.Labels = esls
	return esAlert
}
