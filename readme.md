# WMI接口
> windows的WMI查询接口

## 内置方法
- [vela.wmi.query(sql , function)](#查询方法)  &emsp;查询方法

## 查询方法
> err = vela.wmi.query(sql , pipe) <br />
> sql: 查询语句  pipe: 处理方法 [查询结果](#结果对象)
```lua
    vela.wmi.query("SELECT * FROM Win32_Account" , function(item)
        print(item.s'Name')
        print(item.s'Caption')
    end)
```

## 结果对象
> 封装单个数据对象

内置方法:

- [item.s(string)](#) &emsp;字符内容
- [item.n(string)](#) &emsp;数字内容
- [item.b(string)](#) &emsp;字符数组
- [item.t(string)](#) &emsp;时间

```lua
    print(item.s'Name')
    print(item.n'AccountType' or -1)
    print(item.s'Caption')
    print(item.s'Description')
    print(item.b'Disabled')
    print(item.s'Domain')
    print(item.s'FullName')
    print(item.s'InstallDate')
    print(item.b'LocalAccount')
    print(item.b'Lockout')
    print(item.b'PasswordChangeable')
    print(item.b'PasswordExpires')
    print(item.b'PasswordRequired')
    print(item.s'SID')
    print(item.s'SIDType')
    print(item.s'Status') 

```