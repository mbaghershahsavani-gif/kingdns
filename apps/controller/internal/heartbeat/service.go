package heartbeat

// Production heartbeat service.
//
// Flow:
// receive heartbeat
// calculate health
// persist status
// update cache
