func StartMonitor() {

 ticker := time.NewTicker(...)

 for range ticker.C {

     UpdateHeartbeat()

     SyncPeer()

     hb := GetHeartbeat()

     log.Printf(
        "Heartbeat sent: %s %s %s",
        hb.NodeID,
        hb.Region,
        hb.Status,
     )
 }
}