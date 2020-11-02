package algorithm

import "time"

var (
    machineID int64 // 机器id占10位,十进制范围是 [0, 1023]
    sn        int64 // 序列号占12位,十进制范围是 [0, 4095]
    latest    int64 // 上次的时间戳(毫秒 > 微秒 > 纳秒)
)

func init() {
    latest = time.Now().UnixNano() / 1000000
}

func SetMachineId(mid int64) {
    machineID = mid << 12
}

// fixme: 并发安全
func GetSnowflakeId() int64 {
    curTimeStamp := time.Now().UnixNano() / 1000000
    if curTimeStamp == latest {
        sn++
        if sn > 4095 {
            time.Sleep(time.Millisecond)
            curTimeStamp = time.Now().UnixNano() / 1000000
            latest = curTimeStamp
            sn = 0
        }
        rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
        rightBinValue <<= 22
        id := rightBinValue | machineID | sn
        return id
    }
    if curTimeStamp > latest {
        sn = 0
        latest = curTimeStamp
        rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
        rightBinValue <<= 22
        id := rightBinValue | machineID | sn
        return id
    }
    if curTimeStamp < latest {
        return 0
    }
    return 0
}
