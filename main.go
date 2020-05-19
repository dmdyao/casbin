package main

import (
    "fmt"

    "github.com/casbin/casbin"
)

func main() {
    //通过策略文件和模型配置穿件一个casbin访问控制实例
    e, _ := casbin.NewEnforcer("./perm.conf", "./policy.csv")

    //定义各种sub，obj和act的数组
    subs := []string{"bob", "zeta"}
    objs := []string{"data1", "data2"}
    acts := []string{"read", "write"}

    //遍历组合sub，obj，act并打印出对应策略匹配结果。
    for _, sub := range subs {
        for _, obj := range objs {
            for _, act := range acts {
                result, _ :=  e.Enforce(sub, obj, act)
                fmt.Println(sub, "/", obj, "/", act, "=", result)
            }
        }
    }

}