package socketcan

/*
#include <stdio.h>
#include <stdlib.h>
#include <sys/ioctl.h>
#include <sys/socket.h>
#include <stdint.h>
#include <ctype.h>
#include <errno.h>
#include <libgen.h>
#include <signal.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <net/if.h>
#include <sys/epoll.h>
#include <sys/ioctl.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <linux/can/raw.h>

#define CAN_FILTER_PASS     0x01    //过滤方式-通过
#define CAN_FILTER_REJECT   0x02    //过滤方式-拒绝

int rcvFiltersSet(int canfd, const uint canId, const uint filterType)
{
    if(canfd <= 0)	//canfd就不用解释了…
        return -1;

    if(0 == canId){
        setsockopt(canfd, SOL_CAN_RAW, CAN_RAW_FILTER, NULL, 0);    //不需要接收任何报文
        return 0;
    }

    struct can_filter rfilter;

    if(filterType & CAN_FILTER_PASS){
        rfilter.can_id = canId;
    } else {
        rfilter.can_id = canId | CAN_INV_FILTER;
    }
	if(canId &0x80000000) {
        rfilter.can_mask = 0x1fffffff; 
	} else {
		rfilter.can_mask = 0x7ff;
	}

    if(filterType & CAN_FILTER_REJECT){
        int join_filter = 1;
        setsockopt(canfd, SOL_CAN_RAW, CAN_RAW_JOIN_FILTERS, &join_filter, sizeof(join_filter));
    }
    setsockopt(canfd, SOL_CAN_RAW, CAN_RAW_FILTER, &rfilter, sizeof(rfilter));
    return 0;
}
*/
import "C"

import (
	"errors"
)


/*
*const uint can_id_reject[] = {0x123, 0x11111678, 0x282};	//除这三种canid之外的报文，都接收
canRcvFiltersSet(canfd, can_id_reject, sizeof(can_id_reject)/sizeof(uint), CAN_FILTER_REJECT);
           
*const uint can_id_pass[] = {0x123, 0x11111678, 0x282};	//只接收这三种canid报文，其它不接收
canRcvFiltersSet(canfd, can_id_pass, sizeof(can_id_pass)/sizeof(uint), CAN_FILTER_PASS);
*/

func (itf *rawInterface) AddfilterPass(canid_pass uint) error {
	succ := C.rcvFiltersSet(C.int(itf.fd), C.uint(canid_pass), C.CAN_FILTER_PASS)
	if succ == 0 {
		return nil
	}

	return errors.New("can filter failed")
}
