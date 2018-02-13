# go-bamboo-spec
Bamboo Spec file in golang

This package can be used to define a Bamboo Spec and dump it to YAML.

Example:
``` golang
//go:generate go run example.go                                                 
package main                                                                    
                                                                                
import (                                                                        
    "fmt"                                                                       
    . "github.com/predmond/go-bamboo-spec"                                      
)                                                                               
                                                                                
func main() {                                                                   
    yaml, err := NewSpec(                                                       
        NewProject("DRAGON",                                                    
            NewPlan("SLAYER", "Dragon Slayer Quest"),                           
        ),                                                                      
        NewStage(                                                               
            NewJob(                                                             
                `echo 'Going to slay the red dragon, watch me'`,                
                `sleep 1`,                                                      
                `echo 'Victory!'`,                                              
            ),                                                                  
        ),                                                                      
    ).Build()                                                                   
    if err != nil {                                                             
        panic(err)                                                              
    }                                                                           
    fmt.Print(yaml)                                                             
}      
```

generates:

``` yaml
project:
  key: DRAGON
  plan:
    key: SLAYER
    name: Dragon Slayer Quest
stages:
- jobs:
  - scripts:
    - echo 'Going to slay the red dragon, watch me'
    - sleep 1
    - echo 'Victory!'
```
