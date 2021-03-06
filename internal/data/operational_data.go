package data

// Generated by https://quicktype.io

type OperationalData struct {
    OpMode                    string  `json:"OpMode"`                   
    RTime                     string  `json:"RTime"`                    
    RobTimer                  string  `json:"RobTimer"`                 
    A1Pos                     string  `json:"A1Pos"`                    
    A1BrakeStatus             string  `json:"A1BrakeStatus"`            
    A1HoldingTorqueMandStatus string  `json:"A1HoldingTorqueMandStatus"`
    A1HoldingTorque           string  `json:"A1HoldingTorque"`          
    A1Curr                    string  `json:"A1Curr"`                   
    A1Temp                    float64 `json:"A1TEMP"`                   
    A1Torque                  string  `json:"A1Torque"`                 
    A1Speed                   float64 `json:"A1Speed"`                  
    A1Load                    string  `json:"A1load"`                   
    A1DriveLoad               string  `json:"A1DriveLoad"`              
    A2Pos                     string  `json:"A2Pos"`                    
    A2BrakeStatus             string  `json:"A2BrakeStatus"`            
    A2HoldingTorqueMandStatus string  `json:"A2HoldingTorqueMandStatus"`
    A2HoldingTorque           string  `json:"A2HoldingTorque"`          
    A2Curr                    string  `json:"A2Curr"`                   
    A2Temp                    float64 `json:"A2TEMP"`                   
    A2Torque                  string  `json:"A2Torque"`                 
    A2Speed                   float64 `json:"A2Speed"`                  
    A2Load                    string  `json:"A2load"`                   
    A2DriveLoad               string  `json:"A2DriveLoad"`              
    A3Pos                     string  `json:"A3Pos"`                    
    A3BrakeStatus             string  `json:"A3BrakeStatus"`            
    A3HoldingTorqueMandStatus string  `json:"A3HoldingTorqueMandStatus"`
    A3HoldingTorque           string  `json:"A3HoldingTorque"`          
    A3Curr                    string  `json:"A3Curr"`                   
    A3Temp                    float64 `json:"A3TEMP"`                   
    A3Torque                  string  `json:"A3Torque"`                 
    A3Speed                   float64 `json:"A3Speed"`                  
    A3Load                    string  `json:"A3load"`                   
    A3DriveLoad               string  `json:"A3DriveLoad"`              
    A4Pos                     string  `json:"A4Pos"`                    
    A4BrakeStatus             string  `json:"A4BrakeStatus"`            
    A4HoldingTorqueMandStatus string  `json:"A4HoldingTorqueMandStatus"`
    A4HoldingTorque           string  `json:"A4HoldingTorque"`          
    A4Curr                    string  `json:"A4Curr"`                   
    A4Temp                    float64 `json:"A4TEMP"`                   
    A4Torque                  string  `json:"A4Torque"`                 
    A4Speed                   float64 `json:"A4Speed"`                  
    A4Load                    string  `json:"A4load"`                   
    A4DriveLoad               string  `json:"A4DriveLoad"`              
    A5Pos                     string  `json:"A5Pos"`                    
    A5BrakeStatus             string  `json:"A5BrakeStatus"`            
    A5HoldingTorqueMandStatus string  `json:"A5HoldingTorqueMandStatus"`
    A5HoldingTorque           string  `json:"A5HoldingTorque"`          
    A5Curr                    string  `json:"A5Curr"`                   
    A5Temp                    float64 `json:"A5TEMP"`                   
    A5Torque                  string  `json:"A5Torque"`                 
    A5Speed                   float64 `json:"A5Speed"`                  
    A5DriveLoad               string  `json:"A5DriveLoad"`              
    A5Load                    string  `json:"A5load"`                   
    A6Pos                     string  `json:"A6Pos"`                    
    A6BrakeStatus             string  `json:"A6BrakeStatus"`            
    A6HoldingTorqueMandStatus string  `json:"A6HoldingTorqueMandStatus"`
    A6HoldingTorque           string  `json:"A6HoldingTorque"`          
    A6Curr                    string  `json:"A6Curr"`                   
    A6Temp                    float64 `json:"A6TEMP"`                   
    A6Torque                  string  `json:"A6Torque"`                 
    A6Speed                   float64 `json:"A6Speed"`                  
    A6Load                    string  `json:"A6load"`                   
    A6DriveLoad               string  `json:"A6DriveLoad"`              
    StopCause                 int64   `json:"StopCause"`                
    RobStopped                string  `json:"RobStopped"`               
    StopMess                  string  `json:"StopMess"`                 
    SmartPadState             string  `json:"SmartPadState"`            
    SmartPadIsConnected       string  `json:"SmartPadIsConnected"`      
    ProName                   string  `json:"ProName"`                  
    WaitState                 bool    `json:"WaitState"`                
    Base                      Base    `json:"Base"`                     
    Tool                      Base    `json:"Tool"`                     
    BaseNo                    string  `json:"BaseNo"`                   
    ToolNo                    string  `json:"ToolNo"`                   
    BaseName                  string  `json:"BaseName"`                 
    ToolName                  string  `json:"ToolName"`                 
    Overload                  bool    `json:"overload"`                 
    ModeT1                    int64   `json:"modeT1"`                   
    ModeT2                    int64   `json:"modeT2"`                   
    ModeEX                    int64   `json:"modeEX"`                   
    ModeAU                    int64   `json:"modeAU"`                   
    DeviceID                  string  `json:"DeviceId"`                 
    DeviceName                string  `json:"DeviceName"`               
    Ts                        string  `json:"ts"`                       
}

type Base struct {
    X string `json:"X"`
    Y string `json:"Y"`
    Z string `json:"Z"`
    A string `json:"A"`
    B string `json:"B"`
    C string `json:"C"`
}
