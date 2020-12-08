/*
    获取windows 系统的相关信息
    原理: GetSystemMetrics
*/

#include <stdio.h>
#include <windows.h>

//设置windows如何排列最小化窗口的一个标志。参考api32.txt中的ARW常数
int GET_ARRANGE(){
	int a;
	a = GetSystemMetrics(SM_CXSCREEN);
	printf("%d",a);
	return 0;
}

//指定启动模式。0=普通模式；1=带网络支持的安全模式
int GET_CLEANBOOT(){
	int a;
	a = GetSystemMetrics(SM_CLEANBOOT);
	printf("%d",a);
	return 0;
}

//可用系统环境的数量
int GET_CMETRICS(){
	int a;
	a = GetSystemMetrics(SM_CMETRICS);
	printf("%d",a);
	return 0;
}

//鼠标按钮（按键）的数量。如没有鼠标，就为零
int GET_CMOUSEBUTTON(){
	int a;
	a = GetSystemMetrics(SM_CMOUSEBUTTONS);
	printf("%d",a);
	return 0;
}

//尺寸不可变边框的大小
int GET_CXBORDER(){
	int x,y;
	x = GetSystemMetrics(SM_CXBORDER);
	y = GetSystemMetrics(SM_CYBORDER);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//标准指针大小
int GET_CXCURSOR(){
	int x,y;
	x = GetSystemMetrics(SM_CXCURSOR);
	y = GetSystemMetrics(SM_CYCURSOR);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//对话框边框的大小
int GET_CXDLGFRAME(){
	int x,y;
	x = GetSystemMetrics(SM_CXDLGFRAME);
	y = GetSystemMetrics(SM_CYDLGFRAME);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//双击区域的大小, 指定屏幕上一个特定的显示区域，只有在这个区域内连续进行两次鼠标单击，才有可能被当作双击事件处理
int GET_CXDOUBLECLK(){
	int x,y;
	x = GetSystemMetrics(SM_CXDOUBLECLK);
	y = GetSystemMetrics(SM_CYDOUBLECLK);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//尺寸可变边框的大小（在win95和nt 4.0中使用SM_C?FIXEDFRAME）
int GET_CXFRAME(){
	int x,y;
	x = GetSystemMetrics(SM_CXFRAME);
	y = GetSystemMetrics(SM_CYFRAME);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}


//最大化窗口客户区的大小
int GET_CXFULLSCREEN(){
	int x,y;
	x = GetSystemMetrics(SM_CXFULLSCREEN);
	y = GetSystemMetrics(SM_CYFULLSCREEN);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//水平滚动条上的箭头大小
int GET_CXHSCROLL(){
	int x,y;
	x = GetSystemMetrics(SM_CXHSCROLL);
	y = GetSystemMetrics(SM_CYHSCROLL);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//滚动块在水平滚动条上的大小
int GET_CXHTHUMB(){
	int x,y;
	x = GetSystemMetrics(SM_CXHTHUMB);
	y = GetSystemMetrics(SM_CYVTHUMB);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//标准图标的大小
int GET_CXICON(){
	int x,y;
	x = GetSystemMetrics(SM_CXICON);
	y = GetSystemMetrics(SM_CYICON);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//桌面图标之间的间隔距离。在win95和nt 4.0中是指大图标的间距
int GET_CXICONSPACING(){
	int x,y;
	x = GetSystemMetrics(SM_CXICONSPACING);
	y = GetSystemMetrics(SM_CYICONSPACING);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//最大化窗口的默认尺寸
int GET_CXMAXIMIZED(){
	int x,y;
	x = GetSystemMetrics(SM_CXMAXIMIZED);
	y = GetSystemMetrics(SM_CYMAXIMIZED);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//改变窗口大小时，最大的轨迹宽度
int GET_CXMAXTRACK(){
	int x,y;
	x = GetSystemMetrics(SM_CXMAXTRACK);
	y = GetSystemMetrics(SM_CYMAXTRACK);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//菜单复选号位图的大小
int GET_CXMENUCHECK(){
	int x,y;
	x = GetSystemMetrics(SM_CXMENUCHECK);
	y = GetSystemMetrics(SM_CYMENUCHECK);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//菜单栏上的按钮大小
int GET_CXMENUSIZE(){
	int x,y;
	x = GetSystemMetrics(SM_CXMENUSIZE);
	y = GetSystemMetrics(SM_CYMENUSIZE);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//窗口的最小尺寸
int GET_CXMIN(){
	int x,y;
	x = GetSystemMetrics(SM_CXMIN);
	y = GetSystemMetrics(SM_CYMIN);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//最小化的窗口必须填充进去的一个矩形小于或等于SM_C?ICONSPACING
int GET_CXMINIMIZED(){
	int x,y;
	x = GetSystemMetrics(SM_CXMINIMIZED);
	y = GetSystemMetrics(SM_CYMINIMIZED);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//窗口的最小轨迹宽度
int GET_CXMINTRACK(){
	int x,y;
	x = GetSystemMetrics(SM_CXMINTRACK);
	y = GetSystemMetrics(SM_CYMINTRACK);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//屏幕大小
char* GET_CXSCREEN(){
    char *xy = (char *) malloc(100);
	int x,y;
	x = GetSystemMetrics(SM_CXSCREEN);
	y = GetSystemMetrics(SM_CYSCREEN);
	printf("%d \n",x);
	printf("%d \n",y);
    sprintf(xy,"%d x %d xp", x, y);
	return xy;
}

//标题栏位图的大小
int GET_CXSIZE(){
	int x,y;
	x = GetSystemMetrics(SM_CXSIZE);
	y = GetSystemMetrics(SM_CYSIZE);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//具有WS_THICKFRAME样式的窗口的大小
int GET_CXSIZEFRAME(){
	int x,y;
	x = GetSystemMetrics(SM_CXSIZEFRAME);
	y = GetSystemMetrics(SM_CYSIZEFRAME);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//小图标的大小
int GET_CXSMICON(){
	int x,y;
	x = GetSystemMetrics(SM_CXSMICON);
	y = GetSystemMetrics(SM_CYSMICON);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//小标题按钮的大小
int GET_CXSMSIZE(){
	int x,y;
	x = GetSystemMetrics(SM_CXSMSIZE);
	y = GetSystemMetrics(SM_CYSMSIZE);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//垂直滚动条中的箭头按钮的大小
int GET_CXVSCROLL(){
	int x,y;
	x = GetSystemMetrics(SM_CXVSCROLL);
	y = GetSystemMetrics(SM_CYVSCROLL);
	printf("%d \n",x);
	printf("%d \n",y);
	return 0;
}

//窗口标题的高度
int GET_CYCAPTION(){
	int a;
	a = GetSystemMetrics(SM_CYCAPTION);
	printf("%d \n",a);
	return 0;
}

//Kanji窗口的大小（Height of Kanji window）
int GET_CYKANJIWINDOW(){
	int a;
	a = GetSystemMetrics(SM_CYKANJIWINDOW);
	printf("%d \n",a);
	return 0;
}

//菜单高度
int GET_CYMENU(){
	int a;
	a = GetSystemMetrics(SM_CYMENU);
	printf("%d \n",a);
	return 0;
}

//小标题的高度
int GET_CYSMCAPTION(){
	int a;
	a = GetSystemMetrics(SM_CYSMCAPTION);
	printf("%d \n",a);
	return 0;
}

//垂直滚动条上滚动块的高度
int GET_CYVTHUMB(){
	int a;
	a = GetSystemMetrics(SM_CYVTHUMB);
	printf("%d \n",a);
	return 0;
}

//如支持双字节则为TRUE
int GET_DBCSENABLED(){
	int a;
	a = GetSystemMetrics(SM_DBCSENABLED);
	printf("%d \n",a);
	return 0;
}

//如windows的调试版正在运行，则为TRUE
int GET_DEBUG(){
	int a;
	a = GetSystemMetrics(SM_DEBUG);
	printf("%d \n",a);
	return 0;
}

//如弹出式菜单对齐菜单栏项目的左侧，则为零
int GET_MENUDROPALIGNMENT(){
	int a;
	a = GetSystemMetrics(SM_MENUDROPALIGNMENT);
	printf("%d \n",a);
	return 0;
}

//允许了希伯来和阿拉伯语
int GET_MIDEASTENABLED(){
	int a;
	a = GetSystemMetrics(SM_MIDEASTENABLED);
	printf("%d \n",a);
	return 0;
}

//如安装了鼠标则为TRUE
int GET_MOUSEPRESENT(){
	int a;
	a = GetSystemMetrics(SM_MOUSEPRESENT);
	printf("%d \n",a);
	return 0;
}

//如安装了带轮鼠标则为TRUE；只适用于nt 4.0
int GET_MOUSEWHEELPRESENT(){
	int a;
	a = GetSystemMetrics(SM_MOUSEWHEELPRESENT);
	printf("%d \n",a);
	return 0;
}

//如安装了网络，则设置位0。其他位保留未用
int GET_NETWORK(){
	int a;
	a = GetSystemMetrics(SM_NETWORK);
	printf("%d \n",a);
	return 0;
}

//如装载了支持笔窗口的DLL，则表示笔窗口的句柄
int GET_PENWINDOWS(){
	int a;
	a = GetSystemMetrics(SM_PENWINDOWS);
	printf("%d \n",a);
	return 0;
}

//如安装了安全（保密）机制，则为TRUE
int GET_SECURE(){
	int a;
	a = GetSystemMetrics(SM_SECURE);
	printf("%d \n",a);
	return 0;
}

//强制视觉提示播放声音
int GET_SHOWSOUNDS(){
	int a;
	a = GetSystemMetrics(SM_SHOWSOUNDS);
	printf("%d \n",a);
	return 0;
}

//系统速度太慢，但仍在运行中（System is too slow for effective use but is being run anyway）
int GET_SLOWMACHINE(){
	int a;
	a = GetSystemMetrics(SM_SLOWMACHINE);
	printf("%d \n",a);
	return 0;
}

//如左右鼠标键已经交换，则为TRUE
int GET_SWAPBUTTON(){
	int a;
	a = GetSystemMetrics(SM_SWAPBUTTON);
	printf("%d \n",a);
	return 0;
}