package models

type RoutingRule struct {
    ID uint `gorm:"primaryKey"`
    ServiceID uint
    Condition string
    TargetNode uint
    Priority int
    Enabled bool
}
