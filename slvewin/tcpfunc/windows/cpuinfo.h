/*
    获取cpu硬件信息
    原理:  cpuid， 由c执行汇编命令获取；
*/
#include<stdio.h>
#include<windows.h>

struct cpuid_result {
    DWORD eax;
    DWORD ebx;
    DWORD ecx;
    DWORD edx;
};

//Generic CPUID function
static inline struct cpuid_result cpuid(unsigned int op)
{
    struct cpuid_result result;
    __asm volatile(
        "mov %%ebx, %%edi;"
        "cpuid;"
        "mov %%ebx, %%esi;"
        "mov %%edi, %%ebx;"
        : "=a" (result.eax),
        "=S" (result.ebx),
        "=c" (result.ecx),
        "=d" (result.edx)
        : "0" (op)
        : "edi");
    return result;
}

static inline unsigned int cpuid_eax(unsigned int op)
{
    //unsigned int eax, ebx, ecx, edx;
    struct cpuid_result regs;

    regs = cpuid(op);

    return regs.eax;
}

void get_cpu_vendor(char* cpu_vendor, unsigned int* cpuid_level)
{
    unsigned int cpuid_op = 0x00000000;
    char vendor_name[16] = {'\0'};
    struct cpuid_result result;
    unsigned int level = 0;

    vendor_name[0] = '\0';
	//eax为0表示读取vendor id，一共12字节，依次在ebx、edx、ecx。
    result = cpuid(cpuid_op);
    level = result.eax;
    vendor_name[0] = (result.ebx >> 0) & 0xff;
    vendor_name[1] = (result.ebx >> 8) & 0xff;
    vendor_name[2] = (result.ebx >> 16) & 0xff;
    vendor_name[3] = (result.ebx >> 24) & 0xff;
    vendor_name[4] = (result.edx >> 0) & 0xff;
    vendor_name[5] = (result.edx >> 8) & 0xff;
    vendor_name[6] = (result.edx >> 16) & 0xff;
    vendor_name[7] = (result.edx >> 24) & 0xff;
    vendor_name[8] = (result.ecx >> 0) & 0xff;
    vendor_name[9] = (result.ecx >> 8) & 0xff;
    vendor_name[10] = (result.ecx >> 16) & 0xff;
    vendor_name[11] = (result.ecx >> 24) & 0xff;
    vendor_name[12] = '\0';

    strcpy(cpu_vendor, vendor_name);
    *cpuid_level = level;
}

void get_cpu_id(char* cpu_id, unsigned int* cpu_sign)
{
    unsigned int cpuid_op = 0x00000001;
    struct cpuid_result result;
    unsigned int sign = 0, id = 0;
    unsigned int tmp = 0;

    result = cpuid(cpuid_op);
    sign = result.eax;
    id = result.edx;

    sprintf(cpu_id, "%02X-%02X-%02X-%02X-%02X-%02X-%02X-%02X", (sign >> 0) & 0xff, (sign >> 8) & 0xff, (sign >> 16) & 0xff, (sign >> 24) & 0xff,
        (id >> 0) & 0xff, (id >> 8) & 0xff, (id >> 16) & 0xff, (id >> 24) & 0xff);
    *cpu_sign = sign;
}

struct cpuinfo_x86 {
    //CPU family
	DWORD    x86;

	//CPU vendor
    DWORD    x86_vendor;

	//CPU model
    DWORD    x86_model;

	//CPU stepping
    DWORD    x86_step;
};

// 参考IA32开发手册第2卷第3章。CPUID exa==0x01的图3-6
static inline void get_fms(struct cpuinfo_x86 *c, DWORD tfms)
{
    c->x86 = (tfms >> 8) & 0xf;
    c->x86_model = (tfms >> 4) & 0xf;
    c->x86_step = tfms & 0xf;
    if (c->x86 == 0xf)
        c->x86 += (tfms >> 20) & 0xff;
    if (c->x86 >= 0x6)
        c->x86_model += ((tfms >> 16) & 0xF) << 4;
}

// 参考IA32开发手册第2卷第3章。CPUID exa==0x01的图3-6
void get_cpu_fms(unsigned int* family, unsigned int* model, unsigned int* stepping)
{
    unsigned int cpuid_op = 0x00000001;
    struct cpuinfo_x86 c;
    unsigned int ver = 0;

    ver = cpuid_eax(cpuid_op);
    get_fms(&c, ver);

    *family = c.x86;
    *model = c.x86_model;
    *stepping = c.x86_step;
}

void get_cpu_name(char* processor_name)
{
    unsigned int cpuid_op = 0x80000002;
    struct cpuid_result regs;
    char temp_processor_name[49];
    char* processor_name_start;
    unsigned int *name_as_ints = (unsigned int *)temp_processor_name;
    unsigned int i;

	//用cpuid指令，eax传入0x80000002/0x80000003/0x80000004，
    //共3个，每个4个寄存器，每个寄存器4字节，故一共48字节。
    //参考IA32开发手册第2卷第3章。
    for (i = 0; i < 3; i++) {
        regs = cpuid(cpuid_op + i);
        name_as_ints[i * 4 + 0] = regs.eax;
        name_as_ints[i * 4 + 1] = regs.ebx;
        name_as_ints[i * 4 + 2] = regs.ecx;
        name_as_ints[i * 4 + 3] = regs.edx;
    }

    temp_processor_name[49] = '\0'; // 最后的字节为0，结束

    processor_name_start = temp_processor_name;
    while (*processor_name_start == ' ')
        processor_name_start++;

    memset(processor_name, 0, 49);
    strcpy(processor_name, processor_name_start);
}

void get_address_bits(unsigned int* linear, unsigned int* physical)
{
    unsigned int cpuid_op = 0x80000008;
    unsigned int tmp = 0;
    tmp = cpuid_eax(cpuid_op);
    *linear = (tmp >> 8) & 0xff;
    *physical = (tmp >> 0) & 0xff;

}

typedef struct {
	char* vendor_id;
	unsigned int cpuid_level;
	char* cpuid_serial;
	unsigned int cpuid_sign;
	char* cpu_name;
	unsigned int cpu_family;
	unsigned int cpu_model;
	unsigned int cpu_stepping;
	unsigned int phy_bits;
	unsigned int vir_bits;
}RetuenCpuInfo;

RetuenCpuInfo _Test()
{
	RetuenCpuInfo cpuinfo;
	printf("go");
    char buffer[49] = { '\0' };
    unsigned int num = 0;
    unsigned int f = 0, m = 0, s = 0;
    unsigned int phy_bits = 0, vir_bits = 0;
	//char* vendor_id;
	//vendor_id = (char *)malloc(100);
	memset((void *)buffer, '\0', sizeof(buffer));
	get_cpu_vendor(buffer,&num);
	strcpy(cpuinfo.vendor_id, buffer);
	//get_cpu_vendor(cpuinfo.vendor_id,&num);
	printf("cpuinfo.vendor_id \t: %s\n", cpuinfo.vendor_id);
	cpuinfo.cpuid_level = num;
	printf("cpuinfo.cpuid_level \t: %u\n", cpuinfo.cpuid_level);
    num = 0;
    memset((void *)buffer, '\0', sizeof(buffer));
    get_cpu_id(buffer, &num);
    strcpy(cpuinfo.cpuid_serial,buffer);
    printf("cpuinfo.cpuid_serial \t: %s\n", cpuinfo.cpuid_serial);
    cpuinfo.cpuid_sign = num;
    printf("cpuinfo.cpuid_sign \t: %u\n", cpuinfo.cpuid_sign);
	memset((void *)buffer, '\0', sizeof(buffer));
	get_cpu_name(buffer);
	strcpy(cpuinfo.cpu_name,buffer);
	printf("cpuinfo.cpu_name \t: %s\n", cpuinfo.cpu_name);
    get_cpu_fms(&f, &m, &s);
    cpuinfo.cpu_family = f;
    cpuinfo.cpu_model = m;
    cpuinfo.cpu_stepping = s;
    printf("cpuinfo.cpu_family \t: %u  (0x%0X)\n", cpuinfo.cpu_family, cpuinfo.cpu_family);
    printf("cpuinfo.cpu_model \t: %u  (0x%0X)\n", cpuinfo.cpu_model, cpuinfo.cpu_model);
    printf("cpuinfo.cpu_stepping \t: %u  (0x%0X)\n", cpuinfo.cpu_stepping, cpuinfo.cpu_stepping);
    get_address_bits(&vir_bits, &phy_bits);
    cpuinfo.vir_bits,cpuinfo.phy_bits = vir_bits,phy_bits;
    printf("address sizes \t: %u bits physical - %u bits virtual\n", cpuinfo.vir_bits, cpuinfo.phy_bits);
	return cpuinfo;
}
char* WindowsGetCpuVendorId(){
	char* vendor_id;
	vendor_id = (char *)malloc(49);
	char buffer[49] = { '\0' };
    unsigned int num = 0;
	memset((void *)buffer, '\0', sizeof(buffer));
	get_cpu_vendor(buffer,&num);
	strcpy(vendor_id, buffer);
	return vendor_id;
}

char* WindowsGetCpuId(){
	char* id;
	id = (char *)malloc(49);
	char buffer[49] = { '\0' };
    unsigned int num = 0;
	memset((void *)buffer, '\0', sizeof(buffer));
	get_cpu_id(buffer,&num);
	strcpy(id, buffer);
	return id;
}

char* WindowsGetCpuName(){
	char* name;
	name = (char *)malloc(100);
	get_cpu_name(name);
	return name;
}
