package linux

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char LinuxCPUID[50];

char * GetCPUID(){
	char id[50] = {'\0'};
	unsigned   long   s1,s2,s3,s4;
	char   sel;
       asm volatile
       ( "movl $0x01 , %%eax ; \n\t"
       "xorl %%edx , %%edx ;\n\t"
       "cpuid ;\n\t"
       "movl %%edx ,%0 ;\n\t"
       "movl %%eax ,%1 ; \n\t"
       :"=m"(s1),"=m"(s2)
       );
       //printf("%08X-%08X-",s1,s2);
       asm volatile
       ("movl $0x03,%%eax ;\n\t"
       "xorl %%ecx,%%ecx ;\n\t"
       "xorl %%edx,%%edx ;\n\t"
       "cpuid ;\n\t"
       "movl %%edx,%0 ;\n\t"
       "movl %%ecx,%1 ;\n\t"
       :"=m"(s3),"=m"(s4)
       );
       // printf("%08X-%08X \n",s3,s4);
	sprintf(id,"%08X-%08X-%08X-%08X",s1,s2,s3,s4);
	//printf("cpu id: %s\n",id);
	strcpy(LinuxCPUID, id);
	return LinuxCPUID;
}
*/
import "C"

import (
	"log"
)

func GetCPUIDFromLinux() string {
	a := C.GetCPUID()
	log.Println(C.GoString(a))
	return C.GoString(a)
}