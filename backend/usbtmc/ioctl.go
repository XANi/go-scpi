package usbtmc

import (
	"fmt"
)




// ./include/uapi/asm-generic/ioctl.h
/* ioctl command encoding: 32 bits total, command in lower 16 bits,
 * size of the parameter structure in the lower 14 bits of the
 * upper 16 bits.
*/
// DDSSSSSS SSSSSSSS TTTTTTTT NNNNNNNN


const _IOC_NONE = 0
const _IOC_WRITE = 1
const _IOC_READ = 2



const _IOCTL_DIRBITS = 2
const _IOC_TYPEBITS = 8
const _IOC_NRBITS = 8
const _IOC_SIZEBITS = 14
func genIoctl(dir uint8, ioctl_type uint8, ioctl_nr uint8, size uint16) (out uint32, err error) {
	if size > 2^14 - 1 {
		return 0, fmt.Errorf("Size must be smaller than %d, ", 2^14)
	}
	var word1 uint16
	var word2 uint16
	word1 = size
	word1 |= (uint16(dir)<< 14)

	word2 = uint16(ioctl_nr)
	word2 |= uint16(ioctl_type)<<8

	return uint32(word2) |  (uint32(word1)<<16 ), nil

}
