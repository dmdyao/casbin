package main

import (
    "fmt"
    "github.com/casbin/casbin"
)

func main() {
    // 最基础Demo
    Demo();
    // 简单用户访问权限控制
    Simple();
    // 角色用户访问权限控制
    Role();
}

func Demo() {
    fmt.Println("-------------------- 最基础demo --------------------")
    // 通过策略文件和模型配置穿件一个casbin访问控制实例
    e, _ := casbin.NewEnforcer("./config/perm.conf", "./config/policy.csv")

    // 定义各种sub，obj和act的数组
    subs := []string{"bob", "zeta"}
    objs := []string{"data1", "data2"}
    acts := []string{"read", "write"}

    // 遍历组合sub，obj，act并打印出对应策略匹配结果。
    for _, sub := range subs {
        for _, obj := range objs {
            for _, act := range acts {
                result, _ :=  e.Enforce(sub, obj, act)
                fmt.Println(sub, "/", obj, "/", act, "=", result)
            }
        }
    }
}

func Simple() {
    fmt.Println("-------------------- 简单用户访问权限控制 --------------------")
    //通过策略文件和模型配置穿件一个casbin访问控制实例
    e, _ := casbin.NewEnforcer("./config/simple_perm.conf", "./config/simple_policy.csv")

    //定义各种sub，obj和act的数组
    subs := []string{"cjc", "zym"}
    objs := []string{"/cust/CreateUser", "/cust/GetUser", "/content/CreateContent", "/content/GetContent"}
    acts := []string{"read", "write"}

    //遍历组合sub，obj，act并打印出对应策略匹配结果。
    for _, sub := range subs {
        for _, obj := range objs {
            for _, act := range acts {
                result, _ :=  e.Enforce(sub, obj, act)
                fmt.Println(sub, ",", obj, ",", act, "=", result)
            }
        }
    }
}

func Role() {
    fmt.Println("-------------------- 角色用户访问权限控制 --------------------")
    //通过策略文件和模型配置穿件一个casbin访问控制实例
    e, _ := casbin.NewEnforcer("./config/role_perm.conf", "./config/role_policy.csv")

    //定义各种sub，obj和act的数组
    subs := []string{"cjc", "zym"}
    objs := []string{"/cust/GetUser", "/cust/CreateUser", "/content/CreateContent", "/content/GetContent"}
    acts := []string{"read", "write"}

    //遍历组合sub，obj，act并打印出对应策略匹配结果。
    for _, sub := range subs {
        for _, obj := range objs {
            for _, act := range acts {
                result, _ :=  e.Enforce(sub, obj, act)
                fmt.Println(sub, ",", obj, ",", act, "=", result)
            }
        }
    }
}