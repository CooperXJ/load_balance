package model

type SumReport struct {
	CPU    *CPU   `json:"cpu"`
	Memory *Mem   `json:"memory"`
	Disk   *Disk  `json:"disk"`
	NetIO  *NetIO `json:"net_io"`
}

type Mem struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

type CPU struct {
	Used float64 `json:"used"`
}

type Disk struct {
}

type NetIO struct {
}
