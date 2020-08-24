package axisparam

func XIP1() uint64 {
	if is_dev {
		return 0
	} else {
		return uint64(0) //for miner rewards
	}
}

func XIP2() uint64 {
	if is_dev {
		return 0
	} else {
		return uint64(0)
	}
}

func XIP3() uint64 {
	if is_dev {
		return uint64(940410)
	} else {
		return uint64(0)
	}
}

func VP1() uint64 {
	if is_dev {
		return 0
	} else {
		return uint64(0)
	}
}

func VP0() uint64 {
	if is_dev {
		return 0
	} else {
		return uint64(0)
	}
}

func XIP4() uint64 {
	if is_dev {
		return 0
	} else {
		return uint64(0)
	}
}

func XIP5() uint64 {
	if is_dev {
		return 0
	} else {
		return uint64(0)
	}
}

func XIP6() uint64 {
	if is_dev {
		return 50
	} else {
		return uint64(0)
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
