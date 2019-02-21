package transliterations

import (
	"unicode"
)

// Tables is data map
var Tables = map[rune][]string{}

func Transliterate(r rune) (ret string) {
	ret = string(r)
	if r < unicode.MaxASCII {
		return ret
	}
	if r > 0xeffff {
		ret = ""
	}
	section := r >> 8   // Chop off the last two hex digits
	position := r % 256 // Last two hex digits
	if tb, ok := Tables[section]; ok {
		if len(tb) > int(position) {
			ret = tb[position]
		}
	}

	return
}

func init() {
	Tables[0x000] = x000
	Tables[0x001] = x001
	Tables[0x002] = x002
	Tables[0x003] = x003
	Tables[0x004] = x004
	Tables[0x005] = x005
	Tables[0x006] = x006
	Tables[0x007] = x007
	Tables[0x009] = x009
	Tables[0x00a] = x00a
	Tables[0x00b] = x00b
	Tables[0x00c] = x00c
	Tables[0x00d] = x00d
	Tables[0x00e] = x00e
	Tables[0x00f] = x00f
	Tables[0x010] = x010
	Tables[0x011] = x011
	Tables[0x012] = x012
	Tables[0x013] = x013
	Tables[0x014] = x014
	Tables[0x015] = x015
	Tables[0x016] = x016
	Tables[0x017] = x017
	Tables[0x018] = x018
	Tables[0x01d] = x01d
	Tables[0x01e] = x01e
	Tables[0x01f] = x01f
	Tables[0x020] = x020
	Tables[0x021] = x021
	Tables[0x022] = x022
	Tables[0x023] = x023
	Tables[0x024] = x024
	Tables[0x025] = x025
	Tables[0x026] = x026
	Tables[0x027] = x027
	Tables[0x028] = x028
	Tables[0x029] = x029
	Tables[0x02a] = x02a
	Tables[0x02c] = x02c
	Tables[0x02e] = x02e
	Tables[0x02f] = x02f
	Tables[0x030] = x030
	Tables[0x031] = x031
	Tables[0x032] = x032
	Tables[0x033] = x033
	Tables[0x04d] = x04d
	Tables[0x04e] = x04e
	Tables[0x04f] = x04f
	Tables[0x050] = x050
	Tables[0x051] = x051
	Tables[0x052] = x052
	Tables[0x053] = x053
	Tables[0x054] = x054
	Tables[0x055] = x055
	Tables[0x056] = x056
	Tables[0x057] = x057
	Tables[0x058] = x058
	Tables[0x059] = x059
	Tables[0x05a] = x05a
	Tables[0x05b] = x05b
	Tables[0x05c] = x05c
	Tables[0x05d] = x05d
	Tables[0x05e] = x05e
	Tables[0x05f] = x05f
	Tables[0x060] = x060
	Tables[0x061] = x061
	Tables[0x062] = x062
	Tables[0x063] = x063
	Tables[0x064] = x064
	Tables[0x065] = x065
	Tables[0x066] = x066
	Tables[0x067] = x067
	Tables[0x068] = x068
	Tables[0x069] = x069
	Tables[0x06a] = x06a
	Tables[0x06b] = x06b
	Tables[0x06c] = x06c
	Tables[0x06d] = x06d
	Tables[0x06e] = x06e
	Tables[0x06f] = x06f
	Tables[0x070] = x070
	Tables[0x071] = x071
	Tables[0x072] = x072
	Tables[0x073] = x073
	Tables[0x074] = x074
	Tables[0x075] = x075
	Tables[0x076] = x076
	Tables[0x077] = x077
	Tables[0x078] = x078
	Tables[0x079] = x079
	Tables[0x07a] = x07a
	Tables[0x07b] = x07b
	Tables[0x07c] = x07c
	Tables[0x07d] = x07d
	Tables[0x07e] = x07e
	Tables[0x07f] = x07f
	Tables[0x080] = x080
	Tables[0x081] = x081
	Tables[0x082] = x082
	Tables[0x083] = x083
	Tables[0x084] = x084
	Tables[0x085] = x085
	Tables[0x086] = x086
	Tables[0x087] = x087
	Tables[0x088] = x088
	Tables[0x089] = x089
	Tables[0x08a] = x08a
	Tables[0x08b] = x08b
	Tables[0x08c] = x08c
	Tables[0x08d] = x08d
	Tables[0x08e] = x08e
	Tables[0x08f] = x08f
	Tables[0x090] = x090
	Tables[0x091] = x091
	Tables[0x092] = x092
	Tables[0x093] = x093
	Tables[0x094] = x094
	Tables[0x095] = x095
	Tables[0x096] = x096
	Tables[0x097] = x097
	Tables[0x098] = x098
	Tables[0x099] = x099
	Tables[0x09a] = x09a
	Tables[0x09b] = x09b
	Tables[0x09c] = x09c
	Tables[0x09d] = x09d
	Tables[0x09e] = x09e
	Tables[0x09f] = x09f
	Tables[0x0a0] = x0a0
	Tables[0x0a1] = x0a1
	Tables[0x0a2] = x0a2
	Tables[0x0a3] = x0a3
	Tables[0x0a4] = x0a4
	Tables[0x0ac] = x0ac
	Tables[0x0ad] = x0ad
	Tables[0x0ae] = x0ae
	Tables[0x0af] = x0af
	Tables[0x0b0] = x0b0
	Tables[0x0b1] = x0b1
	Tables[0x0b2] = x0b2
	Tables[0x0b3] = x0b3
	Tables[0x0b4] = x0b4
	Tables[0x0b5] = x0b5
	Tables[0x0b6] = x0b6
	Tables[0x0b7] = x0b7
	Tables[0x0b8] = x0b8
	Tables[0x0b9] = x0b9
	Tables[0x0ba] = x0ba
	Tables[0x0bb] = x0bb
	Tables[0x0bc] = x0bc
	Tables[0x0bd] = x0bd
	Tables[0x0be] = x0be
	Tables[0x0bf] = x0bf
	Tables[0x0c0] = x0c0
	Tables[0x0c1] = x0c1
	Tables[0x0c2] = x0c2
	Tables[0x0c3] = x0c3
	Tables[0x0c4] = x0c4
	Tables[0x0c5] = x0c5
	Tables[0x0c6] = x0c6
	Tables[0x0c7] = x0c7
	Tables[0x0c8] = x0c8
	Tables[0x0c9] = x0c9
	Tables[0x0ca] = x0ca
	Tables[0x0cb] = x0cb
	Tables[0x0cc] = x0cc
	Tables[0x0cd] = x0cd
	Tables[0x0ce] = x0ce
	Tables[0x0cf] = x0cf
	Tables[0x0d0] = x0d0
	Tables[0x0d1] = x0d1
	Tables[0x0d2] = x0d2
	Tables[0x0d3] = x0d3
	Tables[0x0d4] = x0d4
	Tables[0x0d5] = x0d5
	Tables[0x0d6] = x0d6
	Tables[0x0d7] = x0d7
	Tables[0x0f9] = x0f9
	Tables[0x0fa] = x0fa
	Tables[0x0fb] = x0fb
	Tables[0x0fc] = x0fc
	Tables[0x0fd] = x0fd
	Tables[0x0fe] = x0fe
	Tables[0x0ff] = x0ff
	Tables[0x1d4] = x1d4
	Tables[0x1d5] = x1d5
	Tables[0x1d6] = x1d6
	Tables[0x1d7] = x1d7
}
