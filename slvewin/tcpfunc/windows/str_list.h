#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct{
    unsigned int size;        //子字符串数量
    char **list;            //用字符串数组来存放字符串列表
}StringList;
 
/*
 * 初始化一个字符串列表
 * 注意：strlist_malloc() 和 strlist_free(StringList*) 要配对使用
 */
static StringList* strlist_malloc()
{
    StringList *strlist = (StringList*)malloc(sizeof(StringList));
    memset(strlist, 0, sizeof(StringList));
    return strlist;
}
 
/*
 * 在strlist列表的末端增加一个字符串
 */
static int strlist_add(StringList *strlist, char *str)
{
    int len = strlen(str);
    strlist->size++;
    strlist->list = (char**)realloc(strlist->list, sizeof(char*) * strlist->size);
    if(strlist->list == NULL){
        /* out of memory! */
        printf("error: not enough memory (realloc returned NULL)\n");
        return -1;
    }
    strlist->list[strlist->size-1]=(char*)malloc(len+1);
    memcpy(strlist->list[strlist->size-1], str, len);
    (strlist->list[strlist->size-1])[len]=0;
    return strlist->size;
}
 
/*
 * 在字符串列表的index序号后插入str字符串
 */
static int strlist_insertAfter(StringList *strlist, char* str, unsigned int index)
{
    int i, len;
    if(index >= strlist->size){
        return -1;
    }
    len = strlen(str);
    strlist->size++;
    strlist->list = (char**)realloc(strlist->list, sizeof(char*) * strlist->size);
    if(strlist->list == NULL){
        /* out of memory! */
        printf("error: not enough memory (realloc returned NULL)\n");
        return -1;
    }
    for(i= strlist->size-1; i>index+1; i--){
        strlist->list[i] = strlist->list[i-1];
    }
    strlist->list[index+1] = (char*)malloc(len+1);
    memcpy(strlist->list[index+1], str, len);
    strlist->list[index+1][len]=0;
    return 0;
}
 
/*
 * 删除字符串列表strlist末端的字符串
 */
static int strlist_removeLast(StringList *strlist){
    if(strlist->size <=0){
        return -1;
    }
    free(strlist->list[strlist->size-1]);
    strlist->size--;
    return strlist->size;
}
 
/*
 * 删除字符串列表index位置处的字符串
 */
static int strlist_removeAt(StringList *strlist, unsigned int index)
{
    int i;
    if(index >= strlist->size){
        return -1;
    }
    free(strlist->list[index]);
    for(i=index; i< strlist->size-1; i++){
        strlist->list[i] = strlist->list[i+1];
    }
    strlist->size--;
    return strlist->size;
}
 
/*
 * 获取strlist中index处的字符串。
 */
static char* strlist_getStrAt(StringList *strlist, unsigned int index)
{
    if(strlist==NULL || index>=strlist->size){
        return NULL;
    }
    return strlist->list[index];
}
 
/*
 * 判断字符串列表strlist内是否包含了字符串str。
 */
static int strlist_contains(StringList *strlist, char* str)
{
    int i;
    for(i=0; i<strlist->size; i++){
        if(0 == (strcmp(strlist->list[i], str))){
            return 1;
        }
    }
    return 0;
}
 
/*
 * 释放字符串列表占用的内存空间
 */
static int strlist_free(StringList *strlist)
{
    int i;
    if(!strlist){
        return -1;
    }
    for(i=0; i<strlist->size; i++){
        free(strlist->list[i]);
    }
    free(strlist->list);
    free(strlist);
    return 0;
}
 
/*
 * 打印字符串列表的相关信息
 */
static void strlist_infolog(StringList *strlist){
    int i;
    printf("\n>>>\n");
    printf("strlist info:\n");
    printf("strlist size: %d\n", strlist->size);
    for(i=0; i<strlist->size; i++){
        printf("%d, %s\n",i, strlist->list[i]);
    }
}