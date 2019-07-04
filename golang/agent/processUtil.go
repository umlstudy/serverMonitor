package main

import (
	"strings"

	"github.com/pkg/errors"
	process "github.com/shirou/gopsutil/process"
	"sejong.asia/serverMonitor/common"
)

func FindMatchedPids(procNameParts []string, alarmConditionWithWarningLevelChangeConditionMap map[string]common.AlarmConditionWithWarningLevelChangeCondition) ([]common.ProcessStatus, error) {
	pids, err := process.Pids()
	if err != nil {
		return nil, errors.Wrap(err, "FindMatchedPids #1")
	}
	if len(pids) == 0 {
		return nil, errors.New("could not get pids #2")
	}
	processStatuses := []common.ProcessStatus{}
	for _, pid := range pids {
		proc, err := process.NewProcess(int32(pid))
		if err != nil {
			continue
		}
		procName, err := proc.Name()
		if err != nil {
			continue
		}
		for _, procNamePart := range procNameParts {
			if strings.Contains(procName, procNamePart) {
				acwlcc := alarmConditionWithWarningLevelChangeConditionMap[procNamePart]
				ac := acwlcc.AlarmCondition
				wl := common.NORMAL
				if pid == 0 {
					wl = common.ERROR
				}
				processStatuses = append(processStatuses, common.ProcessStatus{AbstractStatus: common.AbstractStatus{Id: procNamePart, Name: procName, WarningLevel: wl, AlarmCondition: ac}, RealName: procNamePart, ProcId: pid})
				continue
			}
		}
	}
	// 일치하는 프로세스가 없을 경우
	for _, procNamePart := range procNameParts {
		found := false
		for _, ps := range processStatuses {
			if strings.Contains(ps.RealName, procNamePart) {
				found = true
				continue
			}
		}
		if !found {
			acwlcc := alarmConditionWithWarningLevelChangeConditionMap[procNamePart]
			ac := acwlcc.AlarmCondition
			processStatuses = append(processStatuses, common.ProcessStatus{AbstractStatus: common.AbstractStatus{Id: procNamePart, Name: procNamePart, WarningLevel: common.ERROR, AlarmCondition: ac}, RealName: procNamePart, ProcId: 0})
		}
	}

	return processStatuses, nil
}

func CheckAliveProcessStatuses(pss []common.ProcessStatus, procNameParts []string, alarmConditionWithWarningLevelChangeConditionMap map[string]common.AlarmConditionWithWarningLevelChangeCondition) ([]common.ProcessStatus, error) {
	if len(pss) < 1 {
		return nil, errors.New("empty process statuses")
	}

	var newPss []common.ProcessStatus = nil
	for _, ps := range pss {
		if ps.ProcId < 1 {
			newPss_, err := FindMatchedPids(procNameParts, alarmConditionWithWarningLevelChangeConditionMap)
			if err != nil {
				return nil, errors.Wrap(err, "FindMatchedPids failed #5")
			}
			newPss = newPss_
			break
		}
	}

	if newPss == nil {
		for _, ps := range pss {
			proc, err := process.NewProcess(int32(ps.ProcId))
			if err != nil {
				ps.ProcId = 0
			} else {
				procName, err := proc.Name()
				if err != nil {
					ps.ProcId = 0
				} else {
					if strings.Compare(ps.RealName, procName) != 0 {
						ps.ProcId = 0
					}
				}
			}
		}
		newPss = pss
	}

	return newPss, nil
}