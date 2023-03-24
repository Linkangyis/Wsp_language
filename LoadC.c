#include <windows.h>
#include <stdio.h>

typedef void (*MyMessageBox)(void*,char*,char*);
typedef void (*MyMessageBoxConfig)();

void INITCDLL(char* file,char* INITNAME,void* FuncName,void* MapList,void* RpcPort){
	HMODULE hModule=LoadLibrary(file);
	MyMessageBox NewMessageBox=(MyMessageBox)GetProcAddress(hModule,INITNAME);
	return NewMessageBox(MapList,FuncName,RpcPort);
}

void READCONFIG(char* file,char* INITNAME){
	HMODULE hModule=LoadLibrary(file);
	MyMessageBoxConfig NewMessageBox=(MyMessageBox)GetProcAddress(hModule,INITNAME);
	return NewMessageBox();
}

/*
typedef void (*MyMessageBox)(void*);

void Test(char file[],char FuncName[],void *MapList){
	HMODULE hModule=LoadLibrary(file);
	MyMessageBox NewMessageBox=(MyMessageBox)GetProcAddress(hModule,FuncName);
	return NewMessageBox(MapList);
}
*/