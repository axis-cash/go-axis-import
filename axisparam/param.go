package axisparam

func XIP0() uint64 {
	if is_dev {
		return 0
	} else {
		return uint64(0)//lic
	}
}

func XIP1() uint64 {
	if is_dev {
		return 1
	} else {
		return uint64(1)//other
	}
}

func VP0() uint64{
	if is_dev {
		return uint64(20)
	} else {
		return uint64(20) //for init data
	}
}
const MAX_O_INS_LENGTH = int(2500)

const MAX_O_OUT_LENGTH = int(10)

const MAX_Z_OUT_LENGTH_OLD = int(6)

const MAX_Z_OUT_LENGTH_XIP2 = int(500)

const MAX_CONTRACT_OUT_COUNT_LENGTH = int(256)

const LOWEST_STAKING_NODE_FEE_RATE = uint32(2500)

const HIGHEST_STAKING_NODE_FEE_RATE = uint32(7500)

const PKR_PREFIX = byte(88)
